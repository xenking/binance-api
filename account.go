package binance

import (
	"github.com/segmentio/encoding/json"
	"github.com/valyala/fasthttp"
)

// Trades get for a specific account and symbol
func (c *Client) Trades(req *TradeReq) ([]*Trade, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.Limit < 0 || req.Limit > MaxTradesLimit {
		req.Limit = DefaultTradesLimit
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointTrades, req, false, false)
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
	if req.Limit < 0 || req.Limit > MaxTradesLimit {
		req.Limit = DefaultTradesLimit
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointHistoricalTrades, req, false, false)
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
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.Limit < 0 || req.Limit > MaxTradesLimit {
		req.Limit = DefaultTradesLimit
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointAggTrades, req, false, false)
	if err != nil {
		return nil, err
	}
	var trades []*AggregatedTrade
	err = json.Unmarshal(res, &trades)

	return trades, err
}

// AccountTrades get trades for a specific account and symbol
func (c *Client) AccountTrades(req *AccountTradesReq) ([]*AccountTrade, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Limit < 0 || req.Limit > MaxAccountTradesLimit {
		req.Limit = MaxAccountTradesLimit
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointAccountTrades, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*AccountTrade
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// Account get current account information
func (c *Client) Account() (*AccountInfo, error) {
	res, err := c.Do(fasthttp.MethodGet, EndpointAccount, nil, true, false)
	if err != nil {
		return nil, err
	}
	resp := &AccountInfo{}
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// OrderRateLimit get the user's current order count usage for all intervals.
func (c *Client) OrderRateLimit() ([]*RateLimit, error) {
	res, err := c.Do(fasthttp.MethodGet, EndpointRateLimit, nil, true, false)
	if err != nil {
		return nil, err
	}

	var resp []*RateLimit
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// MyPreventedMatches get orders that were expired due to STP
func (c *Client) MyPreventedMatches(req *AccountTradesReq) ([]*AccountTrade, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Limit < 0 || req.Limit > MaxAccountTradesLimit {
		req.Limit = MaxAccountTradesLimit
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointMyPreventedMatches, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*AccountTrade
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// User stream endpoint

// DataStream starts a new user data stream
func (c *Client) DataStream() (string, error) {
	res, err := c.Do(fasthttp.MethodPost, EndpointDataStream, nil, false, true)
	if err != nil {
		return "", err
	}

	resp := &DataStream{}
	err = json.Unmarshal(res, &resp)

	return resp.ListenKey, err
}

// DataStreamKeepAlive pings the data stream key to prevent timeout
func (c *Client) DataStreamKeepAlive(listenKey string) error {
	_, err := c.Do(fasthttp.MethodPut, EndpointDataStream, DataStream{ListenKey: listenKey}, false, true)

	return err
}

// DataStreamClose closes the data stream key
func (c *Client) DataStreamClose(listenKey string) error {
	_, err := c.Do(fasthttp.MethodDelete, EndpointDataStream, DataStream{ListenKey: listenKey}, false, true)

	return err
}
