package ws

import (
	"strings"

	"github.com/xenking/binance-api"
	"github.com/xenking/fastws"
)

type Client struct {
	conn *fastws.Conn
}

func NewClient() *Client {
	return &Client{}
}

// Depth opens websocket with depth updates for the given symbol (eg @100ms frequency)
func (c *Client) Depth(symbol string, frequency FrequencyType) (*Depth, error) {
	var b strings.Builder
	b.WriteString(baseWS)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@depth")
	b.WriteString(string(frequency))
	conn, err := fastws.Dial(b.String())
	if err != nil {
		return nil, err
	}
	return &Depth{wsClient{conn: conn}}, nil
}

// DepthLevel opens websocket with depth updates for the given symbol (eg @100ms frequency)
func (c *Client) DepthLevel(symbol, level string, frequency FrequencyType) (*DepthLevel, error) {
	var b strings.Builder
	b.WriteString(baseWS)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@depth")
	b.WriteString(level)
	b.WriteString(string(frequency))
	conn, err := fastws.Dial(b.String())
	if err != nil {
		return nil, err
	}
	return &DepthLevel{wsClient{conn: conn}}, nil
}

// AllMarketTickers opens websocket with with single depth summary for all tickers
func (c *Client) AllMarketTickers() (*AllMarketTicker, error) {
	var b strings.Builder
	b.WriteString(baseWS)
	b.WriteString("!ticker@arr")
	conn, err := fastws.Dial(b.String())
	if err != nil {
		return nil, err
	}
	return &AllMarketTicker{wsClient{conn: conn}}, nil
}

// IndivTicker opens websocket with with single depth summary for all tickers
func (c *Client) IndivTicker(symbol string) (*IndivTicker, error) {
	var b strings.Builder
	b.WriteString(baseWS)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@ticker")
	conn, err := fastws.Dial(b.String())
	if err != nil {
		return nil, err
	}
	return &IndivTicker{wsClient{conn: conn}}, nil
}

// AllBookTickers opens websocket with with single depth summary for all tickers
func (c *Client) AllBookTickers() (*AllMarketTicker, error) {
	var b strings.Builder
	b.WriteString(baseWS)
	b.WriteString("!bookTicker")
	conn, err := fastws.Dial(b.String())
	if err != nil {
		return nil, err
	}
	return &AllMarketTicker{wsClient{conn: conn}}, nil
}

// IndivBookTicker opens websocket with book ticker best bid or ask updates for the given symbol
func (c *Client) IndivBookTicker(symbol string) (*IndivBookTicker, error) {
	var b strings.Builder
	b.WriteString(baseWS)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@bookTicker")
	conn, err := fastws.Dial(b.String())
	if err != nil {
		return nil, err
	}
	return &IndivBookTicker{wsClient{conn: conn}}, nil
}

// Klines opens websocket with klines updates for the given symbol with the given interval
func (c *Client) Klines(symbol string, interval binance.KlineInterval) (*Klines, error) {
	var b strings.Builder
	b.WriteString(baseWS)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@kline_")
	b.WriteString(string(interval))
	conn, err := fastws.Dial(b.String())
	if err != nil {
		return nil, err
	}
	return &Klines{wsClient{conn: conn}}, nil
}

// AggTrades opens websocket with aggregated trades updates for the given symbol
func (c *Client) AggTrades(symbol string) (*AggTrades, error) {
	var b strings.Builder
	b.WriteString(baseWS)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@aggTrade")
	conn, err := fastws.Dial(b.String())
	if err != nil {
		return nil, err
	}
	return &AggTrades{wsClient{conn: conn}}, nil
}

// Trades opens websocket with trades updates for the given symbol
func (c *Client) Trades(symbol string) (*Trades, error) {
	var b strings.Builder
	b.WriteString(baseWS)
	b.WriteString(strings.ToLower(symbol))
	b.WriteString("@trade")
	conn, err := fastws.Dial(b.String())
	if err != nil {
		return nil, err
	}
	return &Trades{wsClient{conn: conn}}, nil
}

// AccountInfo opens websocket with account info updates
func (c *Client) AccountInfo(listenKey string) (*AccountInfo, error) {
	var b strings.Builder
	b.WriteString(baseWS)
	b.WriteString(listenKey)
	conn, err := fastws.Dial(b.String())
	if err != nil {
		return nil, err
	}
	conn.ReadTimeout = 0
	return &AccountInfo{wsClient{conn: conn}}, nil
}
