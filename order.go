package binance

import (
	"github.com/segmentio/encoding/json"
	"github.com/valyala/fasthttp"
)

// NewOrder sends in a new order
func (c *Client) NewOrder(req *OrderReq) (*OrderRespAck, error) {
	if err := c.validateNewOrderReq(req); err != nil {
		return nil, err
	}
	req.OrderRespType = OrderRespTypeAsk
	res, err := c.Do(fasthttp.MethodPost, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}

	resp := &OrderRespAck{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// NewOrderResult sends in a new order and return created order
func (c *Client) NewOrderResult(req *OrderReq) (*OrderRespResult, error) {
	if err := c.validateNewOrderReq(req); err != nil {
		return nil, err
	}
	req.OrderRespType = OrderRespTypeResult
	res, err := c.Do(fasthttp.MethodPost, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}

	resp := &OrderRespResult{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// NewOrderFull sends in a new order and return created full order info
func (c *Client) NewOrderFull(req *OrderReq) (*OrderRespFull, error) {
	if err := c.validateNewOrderReq(req); err != nil {
		return nil, err
	}
	req.OrderRespType = OrderRespTypeFull
	res, err := c.Do(fasthttp.MethodPost, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}

	resp := &OrderRespFull{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// NewOrderTest tests new order creation and signature/recvWindow long. Creates and validates a new order but does not send it into the matching engine
func (c *Client) NewOrderTest(req *OrderReq) error {
	if err := c.validateNewOrderReq(req); err != nil {
		return err
	}
	_, err := c.Do(fasthttp.MethodPost, EndpointOrderTest, req, true, false)

	return err
}

// QueryOrder checks an order's status
func (c *Client) QueryOrder(req *QueryOrderReq) (*QueryOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.OrderID == 0 && req.OrigClientOrderID == "" {
		return nil, ErrEmptyOrderID
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointOrder, req, true, false)
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
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.OrderID == 0 && req.OrigClientOrderID == "" {
		return nil, ErrEmptyOrderID
	}
	res, err := c.Do(fasthttp.MethodDelete, EndpointOrder, req, true, false)
	if err != nil {
		return nil, err
	}
	resp := &CancelOrder{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// CancelOpenOrders cancel all open orders on a symbol
func (c *Client) CancelOpenOrders(req *CancelOpenOrdersReq) ([]*CancelOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	res, err := c.Do(fasthttp.MethodDelete, EndpointOpenOrders, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*CancelOrder
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// CancelReplaceOrder cancels an existing order and places a new order on the same symbol
func (c *Client) CancelReplaceOrder(req *CancelReplaceOrderReq) (*CancelReplaceOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.CancelOrderID == 0 && req.CancelOrigClientOrderID == "" {
		return nil, ErrEmptyOrderID
	}
	if err := c.validateNewOrderReq(&req.OrderReq); err != nil {
		return nil, err
	}
	if req.CancelReplaceMode == "" {
		req.CancelReplaceMode = CancelReplaceModeStopOnFailure
	}
	if req.OrderRespType == "" {
		req.OrderRespType = OrderRespTypeAsk
	}

	res, err := c.Do(fasthttp.MethodPost, EndpointCancelReplaceOrder, req, true, false)
	if err != nil {
		return nil, err
	}
	resp := &CancelReplaceOrder{}
	err = json.Unmarshal(res, resp)

	return resp, err
}

// OpenOrders get all open orders on a symbol
func (c *Client) OpenOrders(req *OpenOrdersReq) ([]*QueryOrder, error) {
	res, err := c.Do(fasthttp.MethodGet, EndpointOpenOrders, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*QueryOrder
	err = json.Unmarshal(res, &resp)

	return resp, err
}

// AllOrders get all account orders; active, canceled, or filled
func (c *Client) AllOrders(req *AllOrdersReq) ([]*QueryOrder, error) {
	if req == nil {
		return nil, ErrNilRequest
	}
	if req.Symbol == "" {
		return nil, ErrEmptySymbol
	}
	if req.Limit < 0 || req.Limit > MaxOrderLimit {
		req.Limit = DefaultOrderLimit
	}
	res, err := c.Do(fasthttp.MethodGet, EndpointOrdersAll, req, true, false)
	if err != nil {
		return nil, err
	}
	var resp []*QueryOrder
	err = json.Unmarshal(res, &resp)

	return resp, err
}

func (c *Client) validateNewOrderReq(req *OrderReq) error {
	switch {
	case req == nil:
		return ErrNilRequest
	case req.Symbol == "":
		return ErrEmptySymbol
	case req.Side == "":
		return ErrEmptySide
	case req.StrategyType > 0 && req.StrategyType < MinStrategyType:
		return ErrMinStrategyType
	}

	switch req.Type {
	case OrderTypeLimit, OrderTypeLimitMaker:
		switch {
		case req.Price == "":
			return ErrEmptyPrice
		case req.Quantity == "":
			return ErrEmptyQuantity
		case req.TimeInForce == "":
			req.TimeInForce = TimeInForceGTC
		}
	case OrderTypeMarket:
		if req.Quantity == "" && req.QuoteQuantity == "" {
			return ErrEmptyQuantity
		}
	case OrderTypeStopLoss, OrderTypeTakeProfit:
		switch {
		case req.Quantity == "":
			return ErrEmptyQuantity
		case req.StopPrice == "" && req.TrailingDelta == 0:
			return ErrEmptyStopPrice
		}
	case OrderTypeStopLossLimit, OrderTypeTakeProfitLimit:
		switch {
		case req.Quantity == "":
			return ErrEmptyQuantity
		case req.Price == "":
			return ErrEmptyPrice
		case req.StopPrice == "" && req.TrailingDelta == 0:
			return ErrEmptyStopPrice
		case req.TimeInForce == "":
			req.TimeInForce = TimeInForceGTC
		}
	default:
		return ErrInvalidOrderType
	}

	return nil
}
