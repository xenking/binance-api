package ws

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xenking/binance-api"
)

type binanceCtx struct {
	api *binance.Client
	ws  *Client
}

func newBinanceCtx() *binanceCtx {
	key := os.Getenv("BINANCE_API_KEY")
	secret := os.Getenv("BINANCE_API_SECRET")
	return &binanceCtx{
		ws:  NewClient(),
		api: binance.NewClient(key, secret),
	}
}

func TestClient_Depth_Read(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.Depth(symbol, Frequency1000ms)
	defer ws.Close()
	require.NoError(t, err)
	u, err := ws.Read()
	require.NoError(t, err)
	require.Equal(t, symbol, u.Symbol)
}

func TestClient_Depth_Stream(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.Depth(symbol, Frequency1000ms)
	defer ws.Close()
	require.NoError(t, err)
	for u := range ws.Stream() {
		require.NoError(t, ws.err)
		require.Equal(t, symbol, u.Symbol)
		break
	}
}

func TestClient_Klines_Read(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.Klines(symbol, binance.KlineInterval1min)
	defer ws.Close()
	require.NoError(t, err)
	u, err := ws.Read()
	require.NoError(t, err)
	require.Equal(t, symbol, u.Symbol)
}

func TestClient_Klines_Stream(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.Klines(symbol, binance.KlineInterval1min)
	defer ws.Close()
	require.NoError(t, err)
	for u := range ws.Stream() {
		require.NoError(t, ws.err)
		require.Equal(t, symbol, u.Symbol)
		break
	}
}

func TestClient_IndivTicker_Read(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.IndivTicker(symbol)
	defer ws.Close()
	require.NoError(t, err)
	u, err := ws.Read()
	require.NoError(t, err)
	require.Equal(t, symbol, u.Symbol)
}

func TestClient_IndivTickers_Stream(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.IndivTicker(symbol)
	defer ws.Close()
	require.NoError(t, err)
	for u := range ws.Stream() {
		require.NoError(t, ws.err)
		require.Equal(t, symbol, u.Symbol)
		break
	}
}

func TestClient_IndivBookTicker_Read(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.IndivBookTicker(symbol)
	defer ws.Close()
	require.NoError(t, err)
	u, err := ws.Read()
	require.NoError(t, err)
	require.Equal(t, symbol, u.Symbol)
}

func TestClient_IndivBookTickers_Stream(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.IndivBookTicker(symbol)
	defer ws.Close()
	require.NoError(t, err)
	for u := range ws.Stream() {
		require.NoError(t, ws.err)
		require.Equal(t, symbol, u.Symbol)
		break
	}
}

func TestClient_AggTrades_Read(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.AggTrades(symbol)
	defer ws.Close()
	require.NoError(t, err)
	u, err := ws.Read()
	require.NoError(t, err)
	require.Equal(t, symbol, u.Symbol)
}

func TestClient_AggTrades_Stream(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.AggTrades(symbol)
	defer ws.Close()
	require.NoError(t, err)
	for u := range ws.Stream() {
		require.NoError(t, ws.err)
		require.Equal(t, symbol, u.Symbol)
		break
	}
}

func TestClient_Trades_Read(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.Trades(symbol)
	defer ws.Close()
	require.NoError(t, err)
	u, err := ws.Read()
	require.NoError(t, err)
	require.Equal(t, symbol, u.Symbol)
}

func TestClient_Trades_Stream(t *testing.T) {
	const symbol = "ETHBTC"
	ctx := newBinanceCtx()
	ws, err := ctx.ws.Trades(symbol)
	defer ws.Close()
	require.NoError(t, err)
	for u := range ws.Stream() {
		require.NoError(t, ws.err)
		require.Equal(t, symbol, u.Symbol)
		break
	}
}

func TestClient_AccountInfo_Read(t *testing.T) {
	ctx := newBinanceCtx()
	key, err := ctx.api.DataStream()
	require.NoError(t, err)
	defer ctx.api.DataStreamClose(key)
	ws, err := ctx.ws.AccountInfo(key)
	require.NoError(t, err)
	defer ws.Close()
	u1, u2, err := ws.Read()
	require.NoError(t, err)
	if u1 == UpdateTypeUnknown && u2 == nil {
		require.FailNow(t, "Expected to receive an update")
	}
}
