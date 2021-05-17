package binance

import (
	"github.com/segmentio/encoding/json"
	"github.com/valyala/fasthttp"
)

const (
	BaseHost         = "api.binance.com"
	DefaultUserAgent = "Binance/client"
)

type Client struct {
	c *restClient
}

func NewClient(apikey, secret string) *Client {
	return &Client{
		c: newRestClient(apikey, secret),
	}
}

func (c *Client) ReqWindow(window int) *Client {
	c.c.window = window
	return c
}

// General endpoints

// Ping tests connectivity to the Rest API
func (c *Client) Ping() error {
	_, err := c.c.do(fasthttp.MethodGet, EndpointPing, nil, false, false)
	return err
}

// Time tests connectivity to the Rest API and get the current server time
func (c *Client) Time() (*ServerTime, error) {
	res, err := c.c.do(fasthttp.MethodGet, EndpointTime, nil, false, false)
	if err != nil {
		return nil, err
	}
	serverTime := &ServerTime{}
	return serverTime, json.Unmarshal(res, serverTime)
}

// Market Data endpoints

// Depth retrieves the order book for the given symbol
func (c *Client) Depth(req *DepthReq) (*Depth, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Limit == 0 || req.Limit > 100 {
		req.Limit = 100
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointDepth, req, false, false)
	if err != nil {
		return nil, err
	}
	depth := &Depth{}
	return depth, json.Unmarshal(res, &depth)
}

// Trades get for a specific account and symbol
func (c *Client) Trades(req *TradeReq) ([]*Trade, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.Limit < 500 || req.Limit > 1000 {
		req.Limit = 500
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointTrades, req, false, false)
	if err != nil {
		return nil, err
	}
	var trades []*Trade
	return trades, json.Unmarshal(res, &trades)
}

// AggregatedTrades gets compressed, aggregate trades.
// AccountTrades that fill at the time, from the same order, with the same price will have the quantity aggregated
// Remark: If both startTime and endTime are sent, limit should not be sent AND the distance between startTime and endTime must be less than 24 hours.
// Remark: If frondId, startTime, and endTime are not sent, the most recent aggregate trades will be returned.
func (c *Client) AggregatedTrades(req *AggregatedTradeReq) ([]*AggregatedTrade, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Limit < 500 || req.Limit > 1000 {
		req.Limit = 500
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointAggTrades, req, false, false)
	if err != nil {
		return nil, err
	}
	var trades []*AggregatedTrade
	return trades, json.Unmarshal(res, &trades)
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
	if req.Limit == 0 || req.Limit > 500 {
		req.Limit = 500
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointKlines, req, false, false)
	if err != nil {
		return nil, err
	}
	var klines []*Klines
	return klines, json.Unmarshal(res, &klines)

}

// Tickers returns 24 hour price change statistics
func (c *Client) Tickers() ([]*TickerStats, error) {
	res, err := c.c.do(fasthttp.MethodGet, EndpointTicker24h, nil, false, false)
	if err != nil {
		return nil, err
	}
	var tickerStats []*TickerStats
	return tickerStats, json.Unmarshal(res, &tickerStats)
}

// Ticker returns 24 hour price change statistics
func (c *Client) Ticker(req *TickerReq) (*TickerStats, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointTicker24h, req, false, false)
	if err != nil {
		return nil, err
	}
	tickerStats := &TickerStats{}
	return tickerStats, json.Unmarshal(res, tickerStats)
}

// AvgPrice returns 24 hour price change statistics
func (c *Client) AvgPrice(req *AvgPriceReq) (*AvgPrice, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointAvgPrice, req, false, false)
	if err != nil {
		return nil, err
	}
	price := &AvgPrice{}
	return price, json.Unmarshal(res, price)
}

// Prices calculates the latest price for all symbols
func (c *Client) Prices() ([]*SymbolPrice, error) {
	res, err := c.c.do(fasthttp.MethodGet, EndpointTickerPrice, nil, false, false)
	if err != nil {
		return nil, err
	}
	var prices []*SymbolPrice
	return prices, json.Unmarshal(res, &prices)
}

// Price calculates the latest price for a symbol
func (c *Client) Price(req *TickerPriceReq) (*SymbolPrice, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointTickerPrice, req, false, false)
	if err != nil {
		return nil, err
	}
	price := &SymbolPrice{}
	return price, json.Unmarshal(res, price)
}

// BookTickers returns best price/qty on the order book for all symbols
func (c *Client) BookTickers() ([]*BookTicker, error) {
	res, err := c.c.do(fasthttp.MethodGet, EndpointTickerBook, nil, false, false)
	if err != nil {
		return nil, err
	}
	var resp []*BookTicker
	return resp, json.Unmarshal(res, &resp)
}

// BookTicker returns best price/qty on the order book for all symbols
func (c *Client) BookTicker(req *BookTickerReq) (*BookTicker, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointTickerBook, req, false, false)
	if err != nil {
		return nil, err
	}
	resp := &BookTicker{}
	return resp, json.Unmarshal(res, &resp)
}

// Signed endpoints, associated with an account

// NewOrder sends in a new order
func (c *Client) NewOrder(req *OrderReq) (*OrderRespAck, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	switch req.Type {
	case OrderTypeLimit:
		if len(req.Price) == 0 || len(req.Quantity) == 0 {
			return nil, ErrEmptyLimit
		}
		if req.TimeInForce == "" {
			req.TimeInForce = TimeInForceGTC
		}
	case OrderTypeMarket:
		if len(req.Quantity) == 0 && len(req.QuoteQuantity) == 0 {
			return nil, ErrEmptyMarket
		}
	}
	res, err := c.c.do(fasthttp.MethodPost, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}

	resp := &OrderRespAck{}
	return resp, json.Unmarshal(res, resp)
}

// NewOrderResult sends in a new order and return created order
func (c *Client) NewOrderResult(req *OrderReq) (*OrderRespResult, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	switch req.Type {
	case OrderTypeLimit:
		if len(req.Price) == 0 || len(req.Quantity) == 0 {
			return nil, ErrEmptyLimit
		}
		if req.TimeInForce == "" {
			req.TimeInForce = TimeInForceGTC
		}
	case OrderTypeMarket:
		if len(req.Quantity) == 0 && len(req.QuoteQuantity) == 0 {
			return nil, ErrEmptyMarket
		}
	}
	req.OrderRespType = OrderRespTypeResult
	res, err := c.c.do(fasthttp.MethodPost, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}

	resp := &OrderRespResult{}
	return resp, json.Unmarshal(res, resp)
}

// NewOrderFull sends in a new order and return created full order info
func (c *Client) NewOrderFull(req *OrderReq) (*OrderRespFull, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	switch req.Type {
	case OrderTypeLimit:
		if len(req.Price) == 0 || len(req.Quantity) == 0 {
			return nil, ErrEmptyLimit
		}
		if req.TimeInForce == "" {
			req.TimeInForce = TimeInForceGTC
		}
	case OrderTypeMarket:
		if len(req.Quantity) == 0 && len(req.QuoteQuantity) == 0 {
			return nil, ErrEmptyMarket
		}
	}
	req.OrderRespType = OrderRespTypeFull
	res, err := c.c.do(fasthttp.MethodPost, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}

	resp := &OrderRespFull{}
	return resp, json.Unmarshal(res, resp)
}

// NewOrderTest tests new order creation and signature/recvWindow long. Creates and validates a new order but does not send it into the matching engine
func (c *Client) NewOrderTest(req *OrderReq) error {
	if req == nil {
		return ErrNilRequest
	}
	_, err := c.c.do(fasthttp.MethodPost, EndpointOrderTest, req, true, false)
	return err
}

// QueryOrder checks an order's status
func (c *Client) QueryOrder(req *QueryOrderReq) (*QueryOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.OrderID < 0 && req.OrigClientOrderId == "" {
		return nil, ErrEmptyOrderID
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}
	resp := &QueryOrder{}
	return resp, json.Unmarshal(res, resp)
}

// CancelOrder cancel an active order
func (c *Client) CancelOrder(req *CancelOrderReq) (*CancelOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.OrderID < 0 || (req.OrigClientOrderId == "" && req.NewClientOrderId == "") {
		return nil, ErrEmptyOrderID
	}
	res, err := c.c.do(fasthttp.MethodDelete, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}
	resp := &CancelOrder{}
	return resp, json.Unmarshal(res, resp)
}

// OpenOrders get all open orders on a symbol
func (c *Client) OpenOrders(req *OpenOrdersReq) ([]*QueryOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointOpenOrders, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*QueryOrder
	return resp, json.Unmarshal(res, &resp)
}

// CancelOpenOrders cancel all open orders on a symbol
func (c *Client) CancelOpenOrders(req *CancelOpenOrdersReq) ([]*CancelOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	res, err := c.c.do(fasthttp.MethodDelete, EndpointOpenOrders, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*CancelOrder
	return resp, json.Unmarshal(res, &resp)
}

// AllOrders get all account orders; active, canceled, or filled
func (c *Client) AllOrders(req *AllOrdersReq) ([]*QueryOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Limit == 0 {
		req.Limit = 500
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointOrdersAll, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*QueryOrder
	return resp, json.Unmarshal(res, &resp)
}

// Account get current account information
func (c *Client) Account() (*AccountInfo, error) {
	res, err := c.c.do(fasthttp.MethodGet, EndpointAccount, nil, true, false)
	if err != nil {
		return nil, err
	}
	resp := &AccountInfo{}
	return resp, json.Unmarshal(res, &resp)
}

// AccountTrades get trades for a specific account and symbol
func (c *Client) AccountTrades(req *AccountTradesReq) (*AccountTrades, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Limit == 0 || req.Limit > 500 {
		req.Limit = 500
	}
	res, err := c.c.do(fasthttp.MethodGet, EndpointAccountTrades, req, true, false)
	if err != nil {
		return nil, err
	}
	resp := &AccountTrades{}
	return resp, json.Unmarshal(res, &resp)
}

func (c *Client) ExchangeInfo() (*ExchangeInfo, error) {
	res, err := c.c.do(fasthttp.MethodGet, EndpointExchangeInfo, nil, false, false)
	if err != nil {
		return nil, err
	}
	resp := &ExchangeInfo{}
	return resp, json.Unmarshal(res, &resp)
}

// User stream endpoint

// DataStream starts a new user datastream
func (c *Client) DataStream() (string, error) {
	res, err := c.c.do(fasthttp.MethodPost, EndpointDataStream, nil, false, true)
	if err != nil {
		return "", err
	}

	resp := &DatastreamReq{}
	return resp.ListenKey, json.Unmarshal(res, &resp)
}

// DataStreamKeepAlive pings the datastream key to prevent timeout
func (c *Client) DataStreamKeepAlive(listenKey string) error {
	_, err := c.c.do(fasthttp.MethodPut, EndpointDataStream, DatastreamReq{ListenKey: listenKey}, false, true)
	return err
}

// DataStreamClose closes the datastream key
func (c *Client) DataStreamClose(listenKey string) error {
	_, err := c.c.do(fasthttp.MethodDelete, EndpointDataStream, DatastreamReq{ListenKey: listenKey}, false, true)
	return err
}
