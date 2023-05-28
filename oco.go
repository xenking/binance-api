package binance

import (
	"github.com/segmentio/encoding/json"
	"github.com/valyala/fasthttp"
)

// NewOCO get all account orders; active, canceled, or filled
func (c *Client) NewOCO(req *OCOReq) (*OCOOrder, error) {
	switch {
	case req == nil:
		return nil, ErrNilRequest
	case req.Symbol == "":
		return nil, ErrEmptySymbol
	case req.Quantity == "":
		return nil, ErrEmptyQuantity
	case req.Price == "":
		return nil, ErrEmptyPrice
	case req.StopPrice == "":
		return nil, ErrEmptyStopPrice
	}

	res, err := c.Do(fasthttp.MethodGet, EndpointOCOOrder, req, true, false)
	if err != nil {
		return nil, err
	}
	resp := &OCOOrder{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// CancelOCO cancel an active OCO order
func (c *Client) CancelOCO(req *CancelOCOReq) (*OCOOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.OrderListID == 0 && req.ListClientOrderID == "" {
		return nil, ErrEmptyOrderID
	}
	res, err := c.Do(fasthttp.MethodDelete, EndpointOCOOrders, req, true, false)
	if err != nil {
		return nil, err
	}
	resp := &OCOOrder{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// QueryOCO get an OCO order
func (c *Client) QueryOCO(req *QueryOCOReq) (*OCOOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.OrderListID == 0 && req.ListClientOrderID == "" {
		return nil, ErrEmptyOrderID
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointOCOOrders, req, true, false)
	if err != nil {
		return nil, err
	}
	resp := &OCOOrder{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// AllOCO get all account orders; active, canceled, or filled
func (c *Client) AllOCO(req *AllOCOReq) ([]*OCOOrder, error) {
	if req != nil && (req.Limit < 0 || req.Limit > MaxOrderLimit) {
		req.Limit = DefaultOrderLimit
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointOCOOrdersAll, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*OCOOrder
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// OpenOCO get all open orders on a symbol
func (c *Client) OpenOCO() ([]*OCOOrder, error) {
	res, err := c.Do(fasthttp.MethodGet, EndpointOpenOCOOrders, nil, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*OCOOrder
	err = json.Unmarshal(res, &resp)

	return resp, err
}
