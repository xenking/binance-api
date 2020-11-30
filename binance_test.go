package binance

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_Ping(t *testing.T) {
	ctx := newBinanceCtx()
	require.NoError(t, ctx.api.Ping())
}

func TestClient_Time(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.Time()
	require.NoError(t, e)
}

func TestClient_Ticker(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.Ticker(&TickerReq{"LTCBTC"})
	require.NoError(t, e)
}

func TestClient_Depth(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.Depth(&DepthReq{Symbol: "NEOBTC", Limit: 5})
	require.NoError(t, e)
}

func TestClient_Trades(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.Trades(&TradeReq{Symbol: "NEOBTC"})
	require.NoError(t, e)
}

func TestClient_AggregatedTrades(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.AggregatedTrades(&AggregatedTradeReq{Symbol: "NEOBTC"})
	require.NoError(t, e)
}

func TestClient_Klines(t *testing.T) {
	ctx := newBinanceCtx()
	s, e := ctx.api.Klines(&KlinesReq{Symbol: "NEOBTC", Interval: KlineInterval1h, Limit: 5})
	require.NoError(t, e)
	require.Len(t, s, 5)
}

func TestClient_AllBookTickers(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.BookTickers()
	require.NoError(t, e)
}
func TestClient_Prices(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.Prices()
	require.NoError(t, e)
}

func TestClient_Order(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.NewOrder(&OrderReq{
		Symbol:      "NEOBTC",
		Side:        OrderSideSell,
		Type:        OrderTypeLimit,
		TimeInForce: TimeInForceGTC,
		Quantity:    1,
		Price:       0.1,
	})
	require.NoError(t, e)
}

func TestClient_QueryCancelOrder(t *testing.T) {
	ctx := newBinanceCtx()
	s, e := ctx.api.NewOrder(&OrderReq{
		Symbol:      "NEOBTC",
		Side:        OrderSideSell,
		Type:        OrderTypeLimit,
		TimeInForce: TimeInForceGTC,
		Quantity:    1,
		Price:       0.1,
	})
	require.NoError(t, e)
	q, e := ctx.api.QueryOrder(&QueryOrderReq{
		Symbol:  "NEOBTC",
		OrderID: s.OrderID,
	})
	require.NoError(t, e)
	require.Equal(t, "NEOBTC", q.Symbol)
	c, e := ctx.api.CancelOrder(&CancelOrderReq{
		Symbol:  "NEOBTC",
		OrderID: s.OrderID,
	})
	require.NoError(t, e)
	require.Equal(t, "NEOBTC", c.Symbol)
}

func TestClient_DataStream(t *testing.T) {
	ctx := newBinanceCtx()
	key, err := ctx.api.DataStream()
	require.NoError(t, err)
	require.NotEmpty(t, key)
	require.NoError(t, ctx.api.DataStreamKeepAlive(key))
	require.NoError(t, ctx.api.DataStreamClose(key))
}

func TestClient_AllOrders(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.AllOrders(&AllOrdersReq{Symbol: "SNMBTC"})
	require.NoError(t, e)
}

func TestClient_OpenOrders(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.OpenOrders(&OpenOrdersReq{Symbol: "SNMBTC"})
	require.NoError(t, e)
}

func TestClient_CancelOpenOrders(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.CancelOpenOrders(&CancelOpenOrdersReq{Symbol: "SNMBTC"})
	require.NoError(t, e)
}

func TestClient_Account(t *testing.T) {
	ctx := newBinanceCtx()
	_, e := ctx.api.Account()
	require.NoError(t, e)
}

func TestClient_ExchangeInfo(t *testing.T) {
	ctx := newBinanceCtx()
	info, err := ctx.api.ExchangeInfo()
	require.NoError(t, err)
	require.NotNil(t, info)
	require.NotEmpty(t, info.Symbols)
}

type binanceCtx struct {
	api *Client
}

func newBinanceCtx() *binanceCtx {
	return &binanceCtx{
		api: NewClient("", ""),
	}
}
