package binance

import (
	"github.com/segmentio/encoding/json"
	"github.com/valyala/fasthttp"
)

const (
	// BaseHostPort for binance addresses
	BaseHostPort     = "api.binance.com:443"
	BaseHost         = "api.binance.com"
	DefaultUserAgent = "Binance/client"
)

type Client struct {
	RestClient
}

// NewClient creates a new binance client with key and secret
func NewClient(apikey, secret string) *Client {
	return &Client{
		RestClient: NewRestClient(apikey, secret),
	}
}

// NewClientHTTP2 creates a new binance client using HTTP/2 protocol with key and secret
func NewClientHTTP2(apikey, secret string) (*Client, error) {
	c, err := NewRestClientHTTP2(apikey, secret)

	return &Client{
		RestClient: c,
	}, err
}

func NewCustomClient(restClient RestClient) *Client {
	return &Client{
		RestClient: restClient,
	}
}

func (c *Client) ReqWindow(window int) *Client {
	c.RestClient.SetWindow(window)

	return c
}

// General endpoints

// Time tests connectivity to the Rest API and get the current server time
func (c *Client) Time() (*ServerTime, error) {
	res, err := c.Do(fasthttp.MethodGet, EndpointTime, nil, false, false)
	if err != nil {
		return nil, err
	}
	serverTime := &ServerTime{}
	err = json.Unmarshal(res, serverTime)

	return serverTime, err
}

// Ping tests connectivity to the Rest API
func (c *Client) Ping() error {
	_, err := c.Do(fasthttp.MethodGet, EndpointPing, nil, false, false)

	return err
}

// ExchangeInfo get current exchange trading rules and symbols information
func (c *Client) ExchangeInfo(req *ExchangeInfoReq) (*ExchangeInfo, error) {
	if req == nil {
		req = &ExchangeInfoReq{}
	}
	if len(req.Symbols) > 0 && req.Symbol != "" {
		req.Symbol = ""
	}

	res, err := c.Do(fasthttp.MethodGet, EndpointExchangeInfo, req, false, false)
	if err != nil {
		return nil, err
	}
	resp := &ExchangeInfo{}
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// Depth retrieves the order book for the given symbol
func (c *Client) Depth(req *DepthReq) (*Depth, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.Limit < 0 || req.Limit > MaxDepthLimit {
		req.Limit = DefaultDepthLimit
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointDepth, req, false, false)
	if err != nil {
		return nil, err
	}
	depth := &Depth{}
	err = json.Unmarshal(res, &depth)

	return depth, err
}

// Klines returns kline/candlestick bars for a symbol. Kline are uniquely identified by their open time
func (c *Client) Klines(req *KlinesReq) ([]*Kline, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.Interval == "" {
		req.Interval = KlineInterval5min
	}
	if req.Limit < 0 || req.Limit > MaxKlinesLimit {
		req.Limit = DefaultKlinesLimit
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointKlines, req, false, false)
	if err != nil {
		return nil, err
	}
	var klines []*Kline
	err = json.Unmarshal(res, &klines)

	return klines, err
}

// UIKlines returns kline/candlestick bars for a symbol. UIKlines is optimized for presentation of candlestick charts.
func (c *Client) UIKlines(req *KlinesReq) ([]*Kline, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.Interval == "" {
		req.Interval = KlineInterval5min
	}
	if req.Limit < 0 || req.Limit > MaxKlinesLimit {
		req.Limit = DefaultKlinesLimit
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointUIKlines, req, false, false)
	if err != nil {
		return nil, err
	}
	var klines []*Kline
	err = json.Unmarshal(res, &klines)

	return klines, err
}

// AvgPrice returns current average price for a symbol.
func (c *Client) AvgPrice(req *AvgPriceReq) (*AvgPrice, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointAvgPrice, req, false, false)
	if err != nil {
		return nil, err
	}
	price := &AvgPrice{}
	err = json.Unmarshal(res, price)

	return price, err
}

// Prices calculates the latest price for all symbols
func (c *Client) Prices(req *TickerPricesReq) ([]*SymbolPrice, error) {
	res, err := c.Do(fasthttp.MethodGet, EndpointTickerPrice, req, false, false)
	if err != nil {
		return nil, err
	}
	var prices []*SymbolPrice
	err = json.Unmarshal(res, &prices)

	return prices, err
}

// Price calculates the latest price for a symbol
func (c *Client) Price(req *TickerPriceReq) (*SymbolPrice, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointTickerPrice, req, false, false)
	if err != nil {
		return nil, err
	}
	price := &SymbolPrice{}
	err = json.Unmarshal(res, price)

	return price, err
}
