package binance

import (
	"github.com/segmentio/encoding/json"
	"github.com/valyala/fasthttp"
)

// Tickers24h returns 24 hour price change statistics
func (c *Client) Tickers24h(req *Tickers24hReq) ([]*TickerStatFull, error) {
	res, err := c.Do(fasthttp.MethodGet, EndpointTicker24h, req, false, false)
	if err != nil {
		return nil, err
	}
	var tickerStats []*TickerStatFull
	err = json.Unmarshal(res, &tickerStats)

	return tickerStats, err
}

// Tickers24hMini returns 24 hour price change statistics
func (c *Client) Tickers24hMini(req *Tickers24hReq) ([]*TickerStatMini, error) {
	if req == nil {
		req = &Tickers24hReq{}
	}
	req.RespType = TickerRespTypeMini

	res, err := c.Do(fasthttp.MethodGet, EndpointTicker24h, req, false, false)
	if err != nil {
		return nil, err
	}
	var tickerStats []*TickerStatMini
	err = json.Unmarshal(res, &tickerStats)

	return tickerStats, err
}

// Ticker24h returns 24 hour price change statistics
func (c *Client) Ticker24h(req *Ticker24hReq) (*TickerStatFull, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointTicker24h, req, false, false)
	if err != nil {
		return nil, err
	}
	tickerStats := &TickerStatFull{}
	err = json.Unmarshal(res, tickerStats)

	return tickerStats, err
}

// Ticker24hMini returns 24 hour price change statistics
func (c *Client) Ticker24hMini(req *Ticker24hReq) (*TickerStatMini, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	req.RespType = TickerRespTypeMini
	res, err := c.Do(fasthttp.MethodGet, EndpointTicker24h, req, false, false)
	if err != nil {
		return nil, err
	}
	tickerStats := &TickerStatMini{}
	err = json.Unmarshal(res, tickerStats)

	return tickerStats, err
}

// BookTickers returns best price/qty on the order book for all symbols
func (c *Client) BookTickers(req *BookTickersReq) ([]*BookTicker, error) {
	res, err := c.Do(fasthttp.MethodGet, EndpointTickerBook, req, false, false)
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
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointTickerBook, req, false, false)
	if err != nil {
		return nil, err
	}
	resp := &BookTicker{}
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// Tickers returns rolling window price change statistics
func (c *Client) Tickers(req *TickersReq) ([]*TickerStat, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if len(req.Symbols) == 0 {
		return nil, ErrEmptySymbol
	}
	if req.WindowSize != "" && !req.WindowSize.IsValid() {
		return nil, ErrInvalidTickerWindow
	}
	req.RespType = TickerRespTypeMini
	res, err := c.Do(fasthttp.MethodGet, EndpointTicker, req, false, false)
	if err != nil {
		return nil, err
	}
	var resp []*TickerStat
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// TickersMini returns rolling window price change statistics
func (c *Client) TickersMini(req *TickersReq) ([]*TickerStatMini, error) {
	if req != nil && req.WindowSize != "" && !req.WindowSize.IsValid() {
		return nil, ErrInvalidTickerWindow
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointTicker, req, false, false)
	if err != nil {
		return nil, err
	}
	var resp []*TickerStatMini
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// Ticker returns rolling window price change statistics
func (c *Client) Ticker(req *TickerReq) (*TickerStat, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.WindowSize != "" && !req.WindowSize.IsValid() {
		return nil, ErrInvalidTickerWindow
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointTicker, req, false, false)
	if err != nil {
		return nil, err
	}
	resp := &TickerStat{}
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// TickerMini returns rolling window price change statistics
func (c *Client) TickerMini(req *TickerReq) (*TickerStatMini, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.WindowSize != "" && !req.WindowSize.IsValid() {
		return nil, ErrInvalidTickerWindow
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointTicker, req, false, false)
	if err != nil {
		return nil, err
	}
	resp := &TickerStatMini{}
	err = json.Unmarshal(res, &resp)

	return resp, err
}
