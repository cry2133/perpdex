// x/perp/abci.go
package perp

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cry2133/perpdex/x/perp/keeper"
)

// EndBlocker 在每个区块结束时执行
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
    // 每10个区块检查一次清算
    if ctx.BlockHeight()%10 == 0 {
        k.CheckLiquidation(ctx)
    }
    
    // 每天计算一次资金费率
    if ctx.BlockHeight()%5760 == 0 { // 假设5秒一个区块
        k.CalculateFundingRate(ctx)
    }
}
