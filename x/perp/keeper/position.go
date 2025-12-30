// x/perp/keeper/position.go
package keeper

import (
    "fmt"
    "math"
    
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cry2133/perpdex/x/perp/types"
)

func (k Keeper) OpenPosition(ctx sdk.Context, msg *types.MsgOpenPosition) (*types.Position, error) {
    // 1. 验证参数
    if err := k.ValidateOpenPosition(msg); err != nil {
        return nil, err
    }
    
    // 2. 获取市场价格
    oraclePrice, err := k.oracleKeeper.GetPrice(ctx, msg.Pair)
    if err != nil {
        return nil, types.ErrPriceUnavailable
    }
    
    // 3. 获取市场配置
    market, err := k.marketKeeper.GetMarket(ctx, msg.Pair)
    if err != nil {
        return nil, types.ErrMarketNotFound
    }
    
    // 4. 计算保证金和仓位大小
    marginDec, ok := sdk.NewDecFromStr(msg.Margin)
    if !ok {
        return nil, types.ErrInvalidMargin
    }
    
    leverageDec, ok := sdk.NewDecFromStr(msg.Leverage)
    if !ok {
        return nil, types.ErrInvalidLeverage
    }
    
    // 检查杠杆限制
    maxLeverage, _ := sdk.NewDecFromStr(market.MaxLeverage)
    if leverageDec.GT(maxLeverage) {
        return nil, types.ErrLeverageTooHigh
    }
    
    // 5. 计算仓位大小
    priceDec := sdk.NewDecFromInt(oraclePrice.Price)
    positionSize := marginDec.Mul(leverageDec).Quo(priceDec)
    
    // 6. 检查初始保证金率
    initialMarginRatio, _ := sdk.NewDecFromStr(market.InitialMarginRatio)
    requiredMargin := positionSize.Mul(priceDec).Quo(leverageDec)
    if marginDec.LT(requiredMargin) {
        return nil, types.ErrInsufficientMargin
    }
    
    // 7. 转移保证金
    traderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
    if err != nil {
        return nil, err
    }
    
    marginCoins := sdk.NewCoins(sdk.NewCoin(
        market.QuoteDenom,
        marginDec.TruncateInt(),
    ))
    
    // 从交易者到模块转账
    if err := k.bankKeeper.SendCoinsFromAccountToModule(
        ctx, traderAddr, types.ModuleName, marginCoins,
    ); err != nil {
        return nil, err
    }
    
    // 8. 创建仓位
    position := types.Position{
        Trader:         msg.Sender,
        Pair:           msg.Pair,
        Side:           msg.Side,
        Size:           positionSize.String(),
        EntryPrice:     priceDec.String(),
        Margin:         msg.Margin,
        Leverage:       msg.Leverage,
        OpenBlock:      ctx.BlockHeight(),
        OpenTime:       uint64(ctx.BlockTime().Unix()),
        LastFundingTime: uint64(ctx.BlockTime().Unix()),
    }
    
    // 生成仓位ID
    positionId := k.getNextPositionId(ctx)
    
    // 保存仓位
    k.SetPosition(ctx, positionId, position)
    
    // 9. 更新交易者仓位映射
    k.SetTraderPosition(ctx, msg.Sender, msg.Pair, positionId)
    
    // 10. 发射事件
    ctx.EventManager().EmitEvent(
        sdk.NewEvent(types.EventTypePositionOpened,
            sdk.NewAttribute(types.AttributeKeyTrader, msg.Sender),
            sdk.NewAttribute(types.AttributeKeyPair, msg.Pair),
            sdk.NewAttribute(types.AttributeKeyPositionId, fmt.Sprintf("%d", positionId)),
            sdk.NewAttribute(types.AttributeKeySize, positionSize.String()),
            sdk.NewAttribute(types.AttributeKeyEntryPrice, priceDec.String()),
        ),
    )
    
    return &position, nil
}


func (k Keeper) ClosePosition(ctx sdk.Context, msg *types.MsgClosePosition) (*types.MsgClosePositionResponse, error) {
    // 1. 获取仓位
    position, found := k.GetPosition(ctx, msg.PositionId)
    if !found {
        return nil, types.ErrPositionNotFound
    }
    
    // 2. 权限检查
    if msg.Sender != position.Trader {
        return nil, types.ErrNotPositionOwner
    }
    
    // 3. 获取当前价格
    oraclePrice, err := k.oracleKeeper.GetPrice(ctx, position.Pair)
    if err != nil {
        return nil, types.ErrPriceUnavailable
    }
    
    currentPrice := sdk.NewDecFromInt(oraclePrice.Price)
    entryPrice, _ := sdk.NewDecFromStr(position.EntryPrice)
    size, _ := sdk.NewDecFromStr(position.Size)
    margin, _ := sdk.NewDecFromStr(position.Margin)
    
    // 4. 计算盈亏
    var pnl sdk.Dec
    if position.Side == types.PositionSide_LONG {
        // 多头：盈利 = (当前价 - 入场价) * 仓位大小
        pnl = currentPrice.Sub(entryPrice).Mul(size)
    } else {
        // 空头：盈利 = (入场价 - 当前价) * 仓位大小
        pnl = entryPrice.Sub(currentPrice).Mul(size)
    }
    
    // 5. 计算总资金
    totalAmount := margin.Add(pnl)
    
    // 6. 获取市场配置
    market, err := k.marketKeeper.GetMarket(ctx, position.Pair)
    if err != nil {
        return nil, err
    }
    
    // 7. 结算资金
    traderAddr := sdk.MustAccAddressFromBech32(position.Trader)
    
    if totalAmount.IsPositive() {
        // 盈利或收回保证金
        payout := sdk.NewCoins(sdk.NewCoin(
            market.QuoteDenom,
            totalAmount.TruncateInt(),
        ))
        
        if err := k.bankKeeper.SendCoinsFromModuleToAccount(
            ctx, types.ModuleName, traderAddr, payout,
        ); err != nil {
            return nil, err
        }
    } else if totalAmount.IsNegative() {
        // 亏损，需要补充保证金
        loss := totalAmount.Abs()
        required := sdk.NewCoins(sdk.NewCoin(
            market.QuoteDenom,
            loss.TruncateInt(),
        ))
        
        // 检查交易者余额
        balance := k.bankKeeper.GetBalance(ctx, traderAddr, market.QuoteDenom)
        if balance.Amount.LT(loss.TruncateInt()) {
            // 触发部分清算
            return k.partialLiquidation(ctx, position, msg.PositionId, loss)
        }
        
        if err := k.bankKeeper.SendCoinsFromAccountToModule(
            ctx, traderAddr, types.ModuleName, required,
        ); err != nil {
            return nil, err
        }
    }
    
    // 8. 删除仓位记录
    k.RemovePosition(ctx, msg.PositionId)
    k.RemoveTraderPosition(ctx, position.Trader, position.Pair)
    
    // 9. 发射事件
    ctx.EventManager().EmitEvent(
        sdk.NewEvent(types.EventTypePositionClosed,
            sdk.NewAttribute(types.AttributeKeyTrader, position.Trader),
            sdk.NewAttribute(types.AttributeKeyPair, position.Pair),
            sdk.NewAttribute(types.AttributeKeyPositionId, fmt.Sprintf("%d", msg.PositionId)),
            sdk.NewAttribute(types.AttributeKeyPNL, pnl.String()),
        ),
    )
    
    return &types.MsgClosePositionResponse{
        RealizedPnl: pnl.String(),
    }, nil
}
