package ws

import (
	"net"
	"strings"

	"github.com/xenking/websocket"

	"github.com/xenking/binance-api"
)

const DefaultPrefix = "wss://stream.binance.com:9443/ws/"

type Client struct {
	conn   net.Conn
	Prefix string
}

func NewClient() *Client {
	return &Client{
		Prefix: DefaultPrefix,
	}
}

func NewCustomClient(prefix string, conn net.Conn) *Client {
	return &Client{Prefix: prefix, conn: conn}
}

// Depth opens websocket with depth updates for the given symbol (eg @100ms frequency)
func (c *Client) Depth(symbol string, frequency FrequencyType) (*Depth, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, strings.ToLower(symbol), "@depth", string(frequency))
	if err != nil {
		return nil, err
	}

	return &Depth{NewConn(wsc)}, nil
}

// DepthLevel opens websocket with depth updates for the given symbol (eg @100ms frequency)
func (c *Client) DepthLevel(symbol string, level DepthLevelType, frequency FrequencyType) (*DepthLevel, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, strings.ToLower(symbol), "@depth", string(level), string(frequency))
	if err != nil {
		return nil, err
	}

	return &DepthLevel{NewConn(wsc)}, nil
}

// AllMarketTickers opens websocket with with single depth summary for all tickers
func (c *Client) AllMarketTickers() (*AllMarketTicker, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, "!ticker@arr")
	if err != nil {
		return nil, err
	}

	return &AllMarketTicker{NewConn(wsc)}, nil
}

// IndivTicker opens websocket with with single depth summary for all tickers
func (c *Client) IndivTicker(symbol string) (*IndivTicker, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, strings.ToLower(symbol), "@ticker")
	if err != nil {
		return nil, err
	}

	return &IndivTicker{NewConn(wsc)}, nil
}

// AllMarketMiniTickers opens websocket with with single depth summary for all mini-tickers
func (c *Client) AllMarketMiniTickers() (*AllMarketMiniTicker, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, "!miniTicker@arr")
	if err != nil {
		return nil, err
	}

	return &AllMarketMiniTicker{NewConn(wsc)}, nil
}

// IndivMiniTicker opens websocket with with single depth summary for all mini-tickers
func (c *Client) IndivMiniTicker(symbol string) (*IndivMiniTicker, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, strings.ToLower(symbol), "@miniTicker")
	if err != nil {
		return nil, err
	}

	return &IndivMiniTicker{NewConn(wsc)}, nil
}

// AllBookTickers opens websocket with with single depth summary for all tickers
func (c *Client) AllBookTickers() (*AllBookTicker, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, "!bookTicker")
	if err != nil {
		return nil, err
	}

	return &AllBookTicker{NewConn(wsc)}, nil
}

// IndivBookTicker opens websocket with book ticker best bid or ask updates for the given symbol
func (c *Client) IndivBookTicker(symbol string) (*IndivBookTicker, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, strings.ToLower(symbol), "@bookTicker")
	if err != nil {
		return nil, err
	}

	return &IndivBookTicker{NewConn(wsc)}, nil
}

// Klines opens websocket with klines updates for the given symbol with the given interval
func (c *Client) Klines(symbol string, interval binance.KlineInterval) (*Klines, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, strings.ToLower(symbol), "@kline_", string(interval))
	if err != nil {
		return nil, err
	}

	return &Klines{NewConn(wsc)}, nil
}

// AggTrades opens websocket with aggregated trades updates for the given symbol
func (c *Client) AggTrades(symbol string) (*AggTrades, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, strings.ToLower(symbol), "@aggTrade")
	if err != nil {
		return nil, err
	}

	return &AggTrades{NewConn(wsc)}, nil
}

// Trades opens websocket with trades updates for the given symbol
func (c *Client) Trades(symbol string) (*Trades, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, strings.ToLower(symbol), "@trade")
	if err != nil {
		return nil, err
	}

	return &Trades{NewConn(wsc)}, nil
}

// AccountInfo opens websocket with account info updates
func (c *Client) AccountInfo(listenKey string) (*AccountInfo, error) {
	wsc, err := newWSClient(c.conn, c.Prefix, listenKey)
	if err != nil {
		return nil, err
	}

	return &AccountInfo{NewConn(wsc)}, nil
}

func newWSClient(conn net.Conn, paths ...string) (*websocket.Client, error) {
	path := strings.Join(paths, "")
	if conn != nil {
		return websocket.MakeClient(conn, path)
	}

	return websocket.Dial(path)
}
