package ws

import (
	"context"
	"net"
	"strings"

	"github.com/gobwas/ws"

	"github.com/xenking/binance-api"
)

const DefaultStreamPath = "wss://stream.binance.com:9443/ws/"

type Client struct {
	conn       net.Conn
	StreamPath string
}

func NewClient() *Client {
	return &Client{
		StreamPath: DefaultStreamPath,
	}
}

func NewCustomClient(prefix string, conn net.Conn) *Client {
	return &Client{StreamPath: prefix, conn: conn}
}

// DiffDepth opens websocket with depth updates for the given symbol to locally manage an order book
func (c *Client) DiffDepth(ctx context.Context, symbol string, frequency FrequencyType) (*Depth, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, strings.ToLower(symbol), EndpointDepthStream, string(frequency))
	if err != nil {
		return nil, err
	}

	return &Depth{NewConn(wsc)}, nil
}

// DepthLevel opens websocket with depth updates for the given symbol (eg @100ms frequency)
func (c *Client) DepthLevel(ctx context.Context, symbol string, level DepthLevelType, frequency FrequencyType) (*DepthLevel, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, strings.ToLower(symbol), EndpointDepthStream, string(level), string(frequency))
	if err != nil {
		return nil, err
	}

	return &DepthLevel{NewConn(wsc)}, nil
}

// IndividualTicker opens websocket with individual ticker updates for the given symbol
func (c *Client) IndividualTicker(ctx context.Context, symbol string) (*IndividualTicker, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, strings.ToLower(symbol), EndpointTickerStream)
	if err != nil {
		return nil, err
	}

	return &IndividualTicker{NewConn(wsc)}, nil
}

// IndividualRollingWindowTicker opens websocket with individual ticker updates for the given symbol with a custom rolling window
func (c *Client) IndividualRollingWindowTicker(ctx context.Context, symbol string, window WindowSizeType) (*IndividualTicker, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, strings.ToLower(symbol), EndpointWindowTickerStream, string(window))
	if err != nil {
		return nil, err
	}

	return &IndividualTicker{NewConn(wsc)}, nil
}

// AllMarketTickers opens websocket with ticker updates for all symbols
func (c *Client) AllMarketTickers(ctx context.Context) (*AllMarketTicker, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, EndpointAllMarketTickersStream)
	if err != nil {
		return nil, err
	}

	return &AllMarketTicker{NewConn(wsc)}, nil
}

// AllMarketRollingWindowTickers opens websocket with ticker updates for all symbols with a custom rolling window
func (c *Client) AllMarketRollingWindowTickers(ctx context.Context, window WindowSizeType) (*AllMarketTicker, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, EndpointAllMarketWindowTickersStream, string(window), "@arr")
	if err != nil {
		return nil, err
	}

	return &AllMarketTicker{NewConn(wsc)}, nil
}

// AllMarketMiniTickers opens websocket with
func (c *Client) AllMarketMiniTickers(ctx context.Context) (*AllMarketMiniTicker, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, EndpointAllMarketMiniTickersStream)
	if err != nil {
		return nil, err
	}

	return &AllMarketMiniTicker{NewConn(wsc)}, nil
}

// IndividualMiniTicker opens websocket with
func (c *Client) IndividualMiniTicker(ctx context.Context, symbol string) (*IndividualMiniTicker, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, strings.ToLower(symbol), EndpointMiniTickerStream)
	if err != nil {
		return nil, err
	}

	return &IndividualMiniTicker{NewConn(wsc)}, nil
}

// IndividualBookTicker opens websocket with book ticker best bid or ask updates for the given symbol
func (c *Client) IndividualBookTicker(ctx context.Context, symbol string) (*IndividualBookTicker, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, strings.ToLower(symbol), EndpointBookTickerStream)
	if err != nil {
		return nil, err
	}

	return &IndividualBookTicker{NewConn(wsc)}, nil
}

// Klines opens websocket with klines updates for the given symbol with the given interval
func (c *Client) Klines(ctx context.Context, symbol string, interval binance.KlineInterval) (*Klines, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, strings.ToLower(symbol), EndpointKlineStream, string(interval))
	if err != nil {
		return nil, err
	}

	return &Klines{NewConn(wsc)}, nil
}

// AggTrades opens websocket with aggregated trades updates for the given symbol
func (c *Client) AggTrades(ctx context.Context, symbol string) (*AggTrades, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, strings.ToLower(symbol), EndpointAggregatedTradeStream)
	if err != nil {
		return nil, err
	}

	return &AggTrades{NewConn(wsc)}, nil
}

// Trades opens websocket with trades updates for the given symbol
func (c *Client) Trades(ctx context.Context, symbol string) (*Trades, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, strings.ToLower(symbol), EndpointTradeStream)
	if err != nil {
		return nil, err
	}

	return &Trades{NewConn(wsc)}, nil
}

// AccountInfo opens websocket with account info updates
func (c *Client) AccountInfo(ctx context.Context, listenKey string) (*AccountInfo, error) {
	wsc, err := newWSClient(ctx, c.conn, c.StreamPath, listenKey)
	if err != nil {
		return nil, err
	}

	return &AccountInfo{NewConn(wsc)}, nil
}

func newWSClient(ctx context.Context, conn net.Conn, paths ...string) (net.Conn, error) {
	path := strings.Join(paths, "")
	if conn != nil {
		_, err := ws.Upgrade(conn)
		return conn, err
	}
	newConn, _, _, err := ws.Dial(ctx, path)
	return newConn, err
}
