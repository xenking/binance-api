package ws

import (
	"strings"

	"github.com/xenking/websocket"

	"github.com/xenking/binance-api"
)

const (
	BaseWS = "wss://stream.binance.com:9443/ws/"
)

type Client struct {
	prefix string
}

func NewClient() *Client {
	return &Client{
		prefix: BaseWS,
	}
}

func NewCustomClient(baseWS string) *Client {
	return &Client{prefix: baseWS}
}

// Depth opens websocket with depth updates for the given symbol (eg @100ms frequency)
func (c *Client) Depth(symbol string, frequency FrequencyType) (*Depth, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@depth")
	b.WriteString(string(frequency))
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &Depth{client{Client: conn}}, nil
}

// DepthLevel opens websocket with depth updates for the given symbol (eg @100ms frequency)
func (c *Client) DepthLevel(symbol, level string, frequency FrequencyType) (*DepthLevel, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@depth")
	b.WriteString(level)
	b.WriteString(string(frequency))
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &DepthLevel{client{Client: conn}}, nil
}

// AllMarketTickers opens websocket with with single depth summary for all tickers
func (c *Client) AllMarketTickers() (*AllMarketTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString("!ticker@arr")
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &AllMarketTicker{client{Client: conn}}, nil
}

// IndivTicker opens websocket with with single depth summary for all tickers
func (c *Client) IndivTicker(symbol string) (*IndivTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@ticker")
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &IndivTicker{client{Client: conn}}, nil
}

// AllMarketMiniTickers opens websocket with with single depth summary for all mini-tickers
func (c *Client) AllMarketMiniTickers() (*AllMarketMiniTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString("!miniTicker@arr")
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &AllMarketMiniTicker{client{Client: conn}}, nil
}

// IndivMiniTicker opens websocket with with single depth summary for all mini-tickers
func (c *Client) IndivMiniTicker(symbol string) (*IndivMiniTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@miniTicker")
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &IndivMiniTicker{client{Client: conn}}, nil
}

// AllBookTickers opens websocket with with single depth summary for all tickers
func (c *Client) AllBookTickers() (*AllBookTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString("!bookTicker")
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &AllBookTicker{client{Client: conn}}, nil
}

// IndivBookTicker opens websocket with book ticker best bid or ask updates for the given symbol
func (c *Client) IndivBookTicker(symbol string) (*IndivBookTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@bookTicker")
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &IndivBookTicker{client{Client: conn}}, nil
}

// Klines opens websocket with klines updates for the given symbol with the given interval
func (c *Client) Klines(symbol string, interval binance.KlineInterval) (*Klines, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@kline_")
	b.WriteString(string(interval))
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &Klines{client{Client: conn}}, nil
}

// AggTrades opens websocket with aggregated trades updates for the given symbol
func (c *Client) AggTrades(symbol string) (*AggTrades, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@aggTrade")
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &AggTrades{client{Client: conn}}, nil
}

// Trades opens websocket with trades updates for the given symbol
func (c *Client) Trades(symbol string) (*Trades, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@trade")
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &Trades{client{Client: conn}}, nil
}

// AccountInfo opens websocket with account info updates
func (c *Client) AccountInfo(listenKey string) (*AccountInfo, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(listenKey)
	conn, err := websocket.Dial(b.String())
	if err != nil {
		return nil, err
	}

	return &AccountInfo{client{Client: conn}}, nil
}
