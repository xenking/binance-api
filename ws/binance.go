package ws

import (
	"net"
	"strings"

	"github.com/xenking/websocket"

	"github.com/xenking/binance-api"
)

const (
	BaseWS = "wss://stream.binance.com:9443/ws/"
)

type Client struct {
	conn   net.Conn
	prefix string
}

func NewClient() *Client {
	return &Client{
		prefix: BaseWS,
	}
}

func NewCustomClient(prefix string, conn net.Conn) *Client {
	return &Client{prefix: prefix, conn: conn}
}

// Depth opens websocket with depth updates for the given symbol (eg @100ms frequency)
func (c *Client) Depth(symbol string, frequency FrequencyType) (*Depth, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@depth")
	b.WriteString(string(frequency))

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &Depth{client{Client: wsc}}, nil
}

// DepthLevel opens websocket with depth updates for the given symbol (eg @100ms frequency)
func (c *Client) DepthLevel(symbol, level string, frequency FrequencyType) (*DepthLevel, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@depth")
	b.WriteString(level)
	b.WriteString(string(frequency))

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &DepthLevel{client{Client: wsc}}, nil
}

// AllMarketTickers opens websocket with with single depth summary for all tickers
func (c *Client) AllMarketTickers() (*AllMarketTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString("!ticker@arr")

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &AllMarketTicker{client{Client: wsc}}, nil
}

// IndivTicker opens websocket with with single depth summary for all tickers
func (c *Client) IndivTicker(symbol string) (*IndivTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@ticker")

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &IndivTicker{client{Client: wsc}}, nil
}

// AllMarketMiniTickers opens websocket with with single depth summary for all mini-tickers
func (c *Client) AllMarketMiniTickers() (*AllMarketMiniTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString("!miniTicker@arr")

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &AllMarketMiniTicker{client{Client: wsc}}, nil
}

// IndivMiniTicker opens websocket with with single depth summary for all mini-tickers
func (c *Client) IndivMiniTicker(symbol string) (*IndivMiniTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@miniTicker")

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &IndivMiniTicker{client{Client: wsc}}, nil
}

// AllBookTickers opens websocket with with single depth summary for all tickers
func (c *Client) AllBookTickers() (*AllBookTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString("!bookTicker")

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &AllBookTicker{client{Client: wsc}}, nil
}

// IndivBookTicker opens websocket with book ticker best bid or ask updates for the given symbol
func (c *Client) IndivBookTicker(symbol string) (*IndivBookTicker, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@bookTicker")

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &IndivBookTicker{client{Client: wsc}}, nil
}

// Klines opens websocket with klines updates for the given symbol with the given interval
func (c *Client) Klines(symbol string, interval binance.KlineInterval) (*Klines, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@kline_")
	b.WriteString(string(interval))

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &Klines{client{Client: wsc}}, nil
}

// AggTrades opens websocket with aggregated trades updates for the given symbol
func (c *Client) AggTrades(symbol string) (*AggTrades, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@aggTrade")

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &AggTrades{client{Client: wsc}}, nil
}

// Trades opens websocket with trades updates for the given symbol
func (c *Client) Trades(symbol string) (*Trades, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@trade")

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &Trades{client{Client: wsc}}, nil
}

// AccountInfo opens websocket with account info updates
func (c *Client) AccountInfo(listenKey string) (*AccountInfo, error) {
	var b strings.Builder
	b.WriteString(c.prefix)
	b.WriteString(listenKey)

	var wsc *websocket.Client
	var err error
	if c.conn != nil {
		wsc, err = websocket.MakeClient(c.conn, b.String())
	} else {
		wsc, err = websocket.Dial(b.String())
	}
	if err != nil {
		return nil, err
	}

	return &AccountInfo{client{Client: wsc}}, nil
}
