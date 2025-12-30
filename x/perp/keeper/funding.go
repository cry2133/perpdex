// x/perp/keeper/funding.go
package keeper

import (
    "math"
    
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// CalculateFundingRate 计算资金费率
func (k Keeper) CalculateFundingRate(ctx sdk.Context) {
    // 获取所有市场
    markets := k.marketKeeper.GetAllMarkets(ctx)
    
    for _, market := range markets {
        if !market.Enabled {
            continue
        }
        
        // 计算资金费率
        fundingRate := k.calculateMarketFundingRate(ctx, market.Pair)
        
        // 更新资金费率
        k.SetFundingRate(ctx, market.Pair, fundingRate)
        
        // 应用资金费用
        k.applyFunding(ctx, market.Pair, fundingRate)
    }
}

func (k Keeper) calculateMarketFundingRate(ctx sdk.Context, pair string) sdk.Dec {
    // 获取标记价格（Oracle价格）
    oraclePrice, err := k.oracleKeeper.GetPrice(ctx, pair)
    if err != nil {
        return sdk.ZeroDec()
    }
    markPrice := sdk.NewDecFromInt(oraclePrice.Price)
    
    // 计算指数价格（这里简化处理，实际需要从多个交易所获取）
    indexPrice := markPrice // 假设相同
    
    // 计算溢价率
    premium := markPrice.Sub(indexPrice).Quo(indexPrice)
    
    // 资金费率 = 溢价率 * 资金费率系数 + 基础利率
    fundingRateCoefficient := sdk.NewDecWithPrec(1, 3) // 0.001
    baseRate := sdk.NewDecWithPrec(1, 4) // 0.0001
    
    fundingRate := premium.Mul(fundingRateCoefficient).Add(baseRate)
    
    // 限制资金费率在合理范围内
    maxRate := sdk.NewDecWithPrec(5, 3) // 0.5%
    minRate := sdk.NewDecWithPrec(-5, 3) // -0.5%
    
    if fundingRate.GT(maxRate) {
        fundingRate = maxRate
    } else if fundingRate.LT(minRate) {
        fundingRate = minRate
    }
    
    return fundingRate
}

func (k Keeper) applyFunding(ctx sdk.Context, pair string, fundingRate sdk.Dec) {
    // 获取该交易对的所有仓位
    positions := k.GetPositionsByPair(ctx, pair)
    
    for _, position := range positions {
        size, _ := sdk.NewDecFromStr(position.Size)
        entryPrice, _ := sdk.NewDecFromStr(position.EntryPrice)
        
        // 计算仓位价值
        positionValue := size.Mul(entryPrice)
        
        // 计算资金费用（多头支付空头，如果资金费率为正）
        fundingPayment := positionValue.Mul(fundingRate)
        
        // 根据仓位方向决定支付或收取
        if position.Side == types.PositionSide_LONG {
            // 多头支付
            if fundingPayment.IsPositive() {
                k.chargeFunding(ctx, position.Trader, pair, fundingPayment)
            } else {
                k.payFunding(ctx, position.Trader, pair, fundingPayment.Abs())
            }
        } else {
            // 空头收取（如果资金费率为正）
            if fundingPayment.IsPositive() {
                k.payFunding(ctx, position.Trader, pair, fundingPayment)
            } else {
                k.chargeFunding(ctx, position.Trader, pair, fundingPayment.Abs())
            }
        }
        
        // 更新最后资金费率时间
        position.LastFundingTime = uint64(ctx.BlockTime().Unix())
        k.UpdatePosition(ctx, position)
    }
}
