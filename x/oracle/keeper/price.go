// x/oracle/keeper/price.go
package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cry2133/perpdex/x/oracle/types"
)

// SetPrice 设置价格
func (k Keeper) SetPrice(ctx sdk.Context, pair string, price sdk.Int, source string) error {
	// 验证价格
	if price.IsNegative() {
		return types.ErrInvalidPrice
	}

	// 创建价格记录
	priceRecord := types.Price{
		Pair:      pair,
		Price:     price,
		Timestamp: ctx.BlockTime().Unix(),
		Source:    source,
	}

	// 存储价格
	store := ctx.KVStore(k.storeKey)
	key := types.PriceKey(pair)
	bz := k.cdc.MustMarshal(&priceRecord)
	store.Set(key, bz)

	// 发射事件
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypePriceUpdate,
			sdk.NewAttribute(types.AttributeKeyPair, pair),
			sdk.NewAttribute(types.AttributeKeyPrice, price.String()),
			sdk.NewAttribute(types.AttributeKeySource, source),
		),
	)

	return nil
}

// GetPrice 获取最新价格
func (k Keeper) GetPrice(ctx sdk.Context, pair string) (types.Price, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.PriceKey(pair)
	bz := store.Get(key)

	if bz == nil {
		return types.Price{}, types.ErrPriceNotFound
	}

	var price types.Price
	k.cdc.MustUnmarshal(bz, &price)

	// 检查价格是否过期（超过30秒）
	if ctx.BlockTime().Unix()-price.Timestamp > 30 {
		return types.Price{}, types.ErrPriceStale
	}

	return price, nil
}
