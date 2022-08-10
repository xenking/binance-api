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
	c RestClient
}

// NewClient creates a new binance client with key and secret
func NewClient(apikey, secret string) *Client {
	return &Client{
		c: NewRestClient(apikey, secret),
	}
}

// NewClientHTTP2 creates a new binance client using HTTP/2 protocol with key and secret
func NewClientHTTP2(apikey, secret string) (*Client, error) {
	c, err := NewRestClientHTTP2(apikey, secret)

	return &Client{
		c: c,
	}, err
}

func NewCustomClient(restClient RestClient) *Client {
	return &Client{
		c: restClient,
	}
}

func (c *Client) ReqWindow(window int) *Client {
	c.c.SetWindow(window)

	return c
}

// General endpoints

// Ping tests connectivity to the Rest API
func (c *Client) Ping() error {
	_, err := c.c.Do(fasthttp.MethodGet, EndpointPing, nil, false, false)

	return err
}

// Time tests connectivity to the Rest API and get the current server time
func (c *Client) Time() (*ServerTime, error) {
	res, err := c.c.Do(fasthttp.MethodGet, EndpointTime, nil, false, false)
	if err != nil {
		return nil, err
	}
	serverTime := &ServerTime{}
	err = json.Unmarshal(res, serverTime)

	return serverTime, err
}

// Market Data endpoints

// Depth retrieves the order book for the given symbol
func (c *Client) Depth(req *DepthReq) (*Depth, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Limit <= 0 || req.Limit > MaxDepthLimit {
		req.Limit = DefaultDepthLimit
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointDepth, req, false, false)
	if err != nil {
		return nil, err
	}
	depth := &Depth{}
	err = json.Unmarshal(res, &depth)

	return depth, err
}

// Trades get for a specific account and symbol
func (c *Client) Trades(req *TradeReq) ([]*Trade, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.Limit <= 0 || req.Limit > MaxTradesLimit {
		req.Limit = DefaultTradesLimit
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointTrades, req, false, false)
	if err != nil {
		return nil, err
	}
	var trades []*Trade
	err = json.Unmarshal(res, &trades)

	return trades, err
}

// HistoricalTrades get for a specific symbol started from order id
func (c *Client) HistoricalTrades(req *HistoricalTradeReq) ([]*Trade, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.Limit <= 0 || req.Limit > MaxTradesLimit {
		req.Limit = DefaultTradesLimit
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointHistoricalTrades, req, false, false)
	if err != nil {
		return nil, err
	}
	var trades []*Trade
	err = json.Unmarshal(res, &trades)

	return trades, err
}

// AggregatedTrades gets compressed, aggregate trades.
// AccountTrades that fill at the time, from the same order, with the same price will have the quantity aggregated
// Remark: If both startTime and endTime are sent, limit should not be sent AND the distance between startTime and endTime must be less than 24 hours.
// Remark: If frondId, startTime, and endTime are not sent, the most recent aggregate trades will be returned.
func (c *Client) AggregatedTrades(req *AggregatedTradeReq) ([]*AggregatedTrade, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Limit <= 0 || req.Limit > MaxTradesLimit {
		req.Limit = DefaultTradesLimit
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointAggTrades, req, false, false)
	if err != nil {
		return nil, err
	}
	var trades []*AggregatedTrade
	err = json.Unmarshal(res, &trades)

	return trades, err
}

// Klines returns kline/candlestick bars for a symbol. Klines are uniquely identified by their open time
func (c *Client) Klines(req *KlinesReq) ([]*Klines, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.Interval == "" {
		req.Interval = KlineInterval5min
	}
	if req.Limit <= 0 || req.Limit > MaxKlinesLimit {
		req.Limit = DefaultKlinesLimit
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointKlines, req, false, false)
	if err != nil {
		return nil, err
	}
	var klines []*Klines
	err = json.Unmarshal(res, &klines)

	return klines, err
}

// Tickers returns 24 hour price change statistics
func (c *Client) Tickers() ([]*TickerStats, error) {
	res, err := c.c.Do(fasthttp.MethodGet, EndpointTicker24h, nil, false, false)
	if err != nil {
		return nil, err
	}
	var tickerStats []*TickerStats
	err = json.Unmarshal(res, &tickerStats)

	return tickerStats, err
}

// Ticker returns 24 hour price change statistics
func (c *Client) Ticker(req *TickerReq) (*TickerStats, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointTicker24h, req, false, false)
	if err != nil {
		return nil, err
	}
	tickerStats := &TickerStats{}
	err = json.Unmarshal(res, tickerStats)

	return tickerStats, err
}

// AvgPrice returns 24 hour price change statistics
func (c *Client) AvgPrice(req *AvgPriceReq) (*AvgPrice, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointAvgPrice, req, false, false)
	if err != nil {
		return nil, err
	}
	price := &AvgPrice{}
	err = json.Unmarshal(res, price)

	return price, err
}

// Prices calculates the latest price for all symbols
func (c *Client) Prices() ([]*SymbolPrice, error) {
	res, err := c.c.Do(fasthttp.MethodGet, EndpointTickerPrice, nil, false, false)
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
	res, err := c.c.Do(fasthttp.MethodGet, EndpointTickerPrice, req, false, false)
	if err != nil {
		return nil, err
	}
	price := &SymbolPrice{}
	err = json.Unmarshal(res, price)

	return price, err
}

// BookTickers returns best price/qty on the order book for all symbols
func (c *Client) BookTickers() ([]*BookTicker, error) {
	res, err := c.c.Do(fasthttp.MethodGet, EndpointTickerBook, nil, false, false)
	if err != nil {
		return nil, err
	}
	var resp []*BookTicker
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// BookTicker returns best price/qty on the order book for all symbols
func (c *Client) BookTicker(req *BookTickerReq) (*BookTicker, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointTickerBook, req, false, false)
	if err != nil {
		return nil, err
	}
	resp := &BookTicker{}
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// Signed endpoints, associated with an account

// NewOrder sends in a new order
func (c *Client) NewOrder(req *OrderReq) (*OrderRespAck, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	switch req.Type { //nolint:exhaustive
	case OrderTypeLimit:
		if req.Price == "" || req.Quantity == "" {
			return nil, ErrEmptyLimit
		}
		if req.TimeInForce == "" {
			req.TimeInForce = TimeInForceGTC
		}
	case OrderTypeMarket:
		if req.Quantity == "" && req.QuoteQuantity == "" {
			return nil, ErrEmptyMarket
		}
	}
	res, err := c.c.Do(fasthttp.MethodPost, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}

	resp := &OrderRespAck{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// NewOrderResult sends in a new order and return created order
func (c *Client) NewOrderResult(req *OrderReq) (*OrderRespResult, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	switch req.Type { //nolint:exhaustive
	case OrderTypeLimit:
		if req.Price == "" || req.Quantity == "" {
			return nil, ErrEmptyLimit
		}
		if req.TimeInForce == "" {
			req.TimeInForce = TimeInForceGTC
		}
	case OrderTypeMarket:
		if req.Quantity == "" && req.QuoteQuantity == "" {
			return nil, ErrEmptyMarket
		}
	}
	req.OrderRespType = OrderRespTypeResult
	res, err := c.c.Do(fasthttp.MethodPost, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}

	resp := &OrderRespResult{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// NewOrderFull sends in a new order and return created full order info
func (c *Client) NewOrderFull(req *OrderReq) (*OrderRespFull, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	switch req.Type { //nolint:exhaustive
	case OrderTypeLimit:
		if req.Price == "" || req.Quantity == "" {
			return nil, ErrEmptyLimit
		}
		if req.TimeInForce == "" {
			req.TimeInForce = TimeInForceGTC
		}
	case OrderTypeMarket:
		if req.Quantity == "" && req.QuoteQuantity == "" {
			return nil, ErrEmptyMarket
		}
	}
	req.OrderRespType = OrderRespTypeFull
	res, err := c.c.Do(fasthttp.MethodPost, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}

	resp := &OrderRespFull{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// NewOrderTest tests new order creation and signature/recvWindow long. Creates and validates a new order but does not send it into the matching engine
func (c *Client) NewOrderTest(req *OrderReq) error {
	if req == nil {
		return ErrNilRequest
	}
	_, err := c.c.Do(fasthttp.MethodPost, EndpointOrderTest, req, true, false)

	return err
}

// QueryOrder checks an order's status
func (c *Client) QueryOrder(req *QueryOrderReq) (*QueryOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.OrderID == 0 && req.OrigClientOrderID == "" {
		return nil, ErrEmptyOrderID
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}
	resp := &QueryOrder{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// CancelOrder cancel an active order
func (c *Client) CancelOrder(req *CancelOrderReq) (*CancelOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.OrderID == 0 && req.OrigClientOrderID == "" {
		return nil, ErrEmptyOrderID
	}
	res, err := c.c.Do(fasthttp.MethodDelete, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}
	resp := &CancelOrder{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// OpenOrders get all open orders on a symbol
func (c *Client) OpenOrders(req *OpenOrdersReq) ([]*QueryOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointOpenOrders, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*QueryOrder
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// CancelOpenOrders cancel all open orders on a symbol
func (c *Client) CancelOpenOrders(req *CancelOpenOrdersReq) ([]*CancelOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	res, err := c.c.Do(fasthttp.MethodDelete, EndpointOpenOrders, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*CancelOrder
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// AllOrders get all account orders; active, canceled, or filled
func (c *Client) AllOrders(req *AllOrdersReq) ([]*QueryOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Limit <= 0 || req.Limit > MaxOrderLimit {
		req.Limit = DefaultOrderLimit
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointOrdersAll, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*QueryOrder
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// Account get current account information
func (c *Client) Account() (*AccountInfo, error) {
	res, err := c.c.Do(fasthttp.MethodGet, EndpointAccount, nil, true, false)
	if err != nil {
		return nil, err
	}
	resp := &AccountInfo{}
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// AccountTrades get trades for a specific account and symbol
func (c *Client) AccountTrades(req *AccountTradesReq) (*AccountTrades, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Limit <= 0 || req.Limit > MaxAccountTradesLimit {
		req.Limit = MaxAccountTradesLimit
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointAccountTrades, req, true, false)
	if err != nil {
		return nil, err
	}
	resp := &AccountTrades{}
	err = json.Unmarshal(res, &resp)

	return resp, err
}

func (c *Client) ExchangeInfo() (*ExchangeInfo, error) {
	res, err := c.c.Do(fasthttp.MethodGet, EndpointExchangeInfo, nil, false, false)
	if err != nil {
		return nil, err
	}
	resp := &ExchangeInfo{}
	err = json.Unmarshal(res, &resp)

	return resp, err
}

func (c *Client) ExchangeInfoSymbol(req *ExchangeInfoReq) (*ExchangeInfo, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	res, err := c.c.Do(fasthttp.MethodGet, EndpointExchangeInfo, req, false, false)
	if err != nil {
		return nil, err
	}
	resp := &ExchangeInfo{}
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// User stream endpoint

// DataStream starts a new user datastream
func (c *Client) DataStream() (string, error) {
	res, err := c.c.Do(fasthttp.MethodPost, EndpointDataStream, nil, false, true)
	if err != nil {
		return "", err
	}

	resp := &DatastreamReq{}
	err = json.Unmarshal(res, &resp)

	return resp.ListenKey, err
}

// DataStreamKeepAlive pings the datastream key to prevent timeout
func (c *Client) DataStreamKeepAlive(listenKey string) error {
	_, err := c.c.Do(fasthttp.MethodPut, EndpointDataStream, DatastreamReq{ListenKey: listenKey}, false, true)

	return err
}

// DataStreamClose closes the datastream key
func (c *Client) DataStreamClose(listenKey string) error {
	_, err := c.c.Do(fasthttp.MethodDelete, EndpointDataStream, DatastreamReq{ListenKey: listenKey}, false, true)

	return err
}
