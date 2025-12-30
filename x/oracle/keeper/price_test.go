// x/oracle/keeper/price_test.go
package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cry2133/perpdex/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestPriceFlow(t *testing.T) {
	tApp, ctx := setupTestApp(t)
	oracleKeeper := tApp.OracleKeeper

	// 测试设置价格
	err := oracleKeeper.SetPrice(ctx, "BTC/USD", sdk.NewInt(50000), "binance")
	require.NoError(t, err)

	// 测试获取价格
	price, err := oracleKeeper.GetPrice(ctx, "BTC/USD")
	require.NoError(t, err)
	require.Equal(t, "BTC/USD", price.Pair)
	require.Equal(t, sdk.NewInt(50000), price.Price)

	// 测试价格不存在
	_, err = oracleKeeper.GetPrice(ctx, "ETH/USD")
	require.Error(t, err)
	require.Equal(t, types.ErrPriceNotFound, err)
}
