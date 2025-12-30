// x/perp/keeper/liquidation.go
package keeper

import (
    "fmt"
    
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cry2133/perpdex/x/perp/types"
)

// CheckLiquidation 定期检查清算
func (k Keeper) CheckLiquidation(ctx sdk.Context) {
    // 获取所有活跃仓位
    positions := k.GetAllPositions(ctx)
    
    for _, position := range positions {
        // 检查是否需要清算
        if k.shouldLiquidate(ctx, position) {
            // 执行清算
            if err := k.executeLiquidation(ctx, position); err != nil {
                // 记录清算失败，但不中断处理
                ctx.Logger().Error("liquidation failed", "error", err, "position", position.Trader)
            }
        }
    }
}

// shouldLiquidate 判断是否应该清算
func (k Keeper) shouldLiquidate(ctx sdk.Context, position types.Position) bool {
    // 获取市场价格
    oraclePrice, err := k.oracleKeeper.GetPrice(ctx, position.Pair)
    if err != nil {
        return false
    }
    
    currentPrice := sdk.NewDecFromInt(oraclePrice.Price)
    entryPrice, _ := sdk.NewDecFromStr(position.EntryPrice)
    size, _ := sdk.NewDecFromStr(position.Size)
    margin, _ := sdk.NewDecFromStr(position.Margin)
    
    // 计算当前保证金率
    var positionValue sdk.Dec
    if position.Side == types.PositionSide_LONG {
        positionValue = currentPrice.Mul(size)
    } else {
        positionValue = entryPrice.Mul(size) // 空头使用入场价计算价值
    }
    
    // 计算盈亏
    var pnl sdk.Dec
    if position.Side == types.PositionSide_LONG {
        pnl = currentPrice.Sub(entryPrice).Mul(size)
    } else {
        pnl = entryPrice.Sub(currentPrice).Mul(size)
    }
    
    // 当前保证金 = 初始保证金 + 盈亏
    currentMargin := margin.Add(pnl)
    
    // 保证金率 = 当前保证金 / 仓位价值
    marginRatio := currentMargin.Quo(positionValue)
    
    // 获取维持保证金率
    market, err := k.marketKeeper.GetMarket(ctx, position.Pair)
    if err != nil {
        return false
    }
    
    maintenanceMarginRatio, _ := sdk.NewDecFromStr(market.MaintenanceMarginRatio)
    
    // 如果保证金率低于维持保证金率，需要清算
    return marginRatio.LT(maintenanceMarginRatio)
}

// executeLiquidation 执行清算
func (k Keeper) executeLiquidation(ctx sdk.Context, position types.Position) error {
    // 获取仓位ID
    positionId := k.GetPositionId(ctx, position.Trader, position.Pair)
    if positionId == 0 {
        return types.ErrPositionNotFound
    }
    
    // 获取当前价格
    oraclePrice, err := k.oracleKeeper.GetPrice(ctx, position.Pair)
    if err != nil {
        return err
    }
    
    currentPrice := sdk.NewDecFromInt(oraclePrice.Price)
    
    // 计算清算价格（以当前市价平仓）
    var pnl sdk.Dec
    entryPrice, _ := sdk.NewDecFromStr(position.EntryPrice)
    size, _ := sdk.NewDecFromStr(position.Size)
    margin, _ := sdk.NewDecFromStr(position.Margin)
    
    if position.Side == types.PositionSide_LONG {
        pnl = currentPrice.Sub(entryPrice).Mul(size)
    } else {
        pnl = entryPrice.Sub(currentPrice).Mul(size)
    }
    
    // 计算剩余资金
    remaining := margin.Add(pnl)
    
    // 清算惩罚（收取一定比例作为罚金）
    penaltyRatio := sdk.NewDecWithPrec(5, 3) // 0.5%
    penalty := remaining.Mul(penaltyRatio)
    finalPayout := remaining.Sub(penalty)
    
    // 如果有剩余资金，返还给交易者
    if finalPayout.IsPositive() {
        market, _ := k.marketKeeper.GetMarket(ctx, position.Pair)
        traderAddr := sdk.MustAccAddressFromBech32(position.Trader)
        
        payout := sdk.NewCoins(sdk.NewCoin(
            market.QuoteDenom,
            finalPayout.TruncateInt(),
        ))
        
        if err := k.bankKeeper.SendCoinsFromModuleToAccount(
            ctx, types.ModuleName, traderAddr, payout,
        ); err != nil {
            return err
        }
    }
    
    // 删除仓位
    k.RemovePosition(ctx, positionId)
    k.RemoveTraderPosition(ctx, position.Trader, position.Pair)
    
    // 发射清算事件
    ctx.EventManager().EmitEvent(
        sdk.NewEvent(types.EventTypePositionLiquidated,
            sdk.NewAttribute(types.AttributeKeyTrader, position.Trader),
            sdk.NewAttribute(types.AttributeKeyPair, position.Pair),
            sdk.NewAttribute(types.AttributeKeyLiquidationPrice, currentPrice.String()),
            sdk.NewAttribute(types.AttributeKeyPenalty, penalty.String()),
        ),
    )
    
    return nil
}
