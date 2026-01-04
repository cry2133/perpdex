package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"perpdex/x/perpdex/types"
)

func (k Keeper) AppendPrice(ctx context.Context, post types.Price) uint64 {
	count := k.GetPriceCount(ctx)
	post.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.PriceKey))
	appendedValue := k.cdc.MustMarshal(&post)
	store.Set(GetPriceIDBytes(post.Id), appendedValue)

	// Add symbol index
	symbolStore := prefix.NewStore(storeAdapter, []byte(types.PriceSymbolKey))
	symbolStore.Set([]byte(post.Symbol), GetPriceIDBytes(post.Id))

	k.SetPriceCount(ctx, count+1)
	return count
}

func (k Keeper) GetPriceCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := []byte(types.PriceCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetPriceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetPriceCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := []byte(types.PriceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetPrice(ctx context.Context, id uint64) (val types.Price, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.PriceKey))
	b := store.Get(GetPriceIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetPrice(ctx context.Context, post types.Price) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.PriceKey))
	b := k.cdc.MustMarshal(&post)
	store.Set(GetPriceIDBytes(post.Id), b)

	// Update symbol index
	symbolStore := prefix.NewStore(storeAdapter, []byte(types.PriceSymbolKey))
	symbolStore.Set([]byte(post.Symbol), GetPriceIDBytes(post.Id))
}

// GetPriceBySymbol retrieves a price by its symbol using index
func (k Keeper) GetPriceBySymbol(ctx context.Context, symbol string) (val types.Price, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	// Get the price ID from the symbol index
	symbolStore := prefix.NewStore(storeAdapter, []byte(types.PriceSymbolKey))
	idBytes := symbolStore.Get([]byte(symbol))
	if idBytes == nil {
		return val, false
	}

	id := binary.BigEndian.Uint64(idBytes)

	// Get the price using the ID
	return k.GetPrice(ctx, id)
}

// RemovePrice removes a price and its symbol index
func (k Keeper) RemovePrice(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	// First get the price to retrieve its symbol
	price, found := k.GetPrice(ctx, id)
	if !found {
		return
	}

	// Remove from main store
	store := prefix.NewStore(storeAdapter, []byte(types.PriceKey))
	store.Delete(GetPriceIDBytes(id))

	// Remove from symbol index
	symbolStore := prefix.NewStore(storeAdapter, []byte(types.PriceSymbolKey))
	symbolStore.Delete([]byte(price.Symbol))
}
