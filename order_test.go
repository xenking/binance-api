package binance_test

import (
	"math/rand"

	"github.com/segmentio/encoding/json"
	"github.com/xenking/binance-api"
)

type mockedOrderTestSuite struct {
	mockedTestSuite
}

func (s *mockedOrderTestSuite) TestNewOrder() {
	var expected *binance.OrderRespAck
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		expected = &binance.OrderRespAck{
			Symbol:       req.Symbol,
			OrderID:      int64(rand.Uint32()),
			TransactTime: int64(rand.Uint32()),
		}
		return json.Marshal(expected)
	}

	actual, e := s.client.NewOrder(&binance.OrderReq{
		Symbol:   "LTCBTC",
		Side:     binance.OrderSideSell,
		Type:     binance.OrderTypeLimit,
		Quantity: "1",
		Price:    "0.1",
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedOrderTestSuite) TestNewMarketOrder() {
	var expected *binance.OrderRespAck
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		expected = &binance.OrderRespAck{
			Symbol:       req.Symbol,
			OrderID:      int64(rand.Uint32()),
			TransactTime: int64(rand.Uint32()),
		}
		return json.Marshal(expected)
	}

	actual, e := s.client.NewOrder(&binance.OrderReq{
		Symbol:   "LTCBTC",
		Side:     binance.OrderSideSell,
		Type:     binance.OrderTypeMarket,
		Quantity: "1",
		Price:    "0.1",
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedOrderTestSuite) TestNewOrderTest() {
	var expected *binance.OrderRespAck
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		expected = &binance.OrderRespAck{
			Symbol:       req.Symbol,
			OrderID:      int64(rand.Uint32()),
			TransactTime: int64(rand.Uint32()),
		}
		return json.Marshal(expected)
	}

	e := s.client.NewOrderTest(&binance.OrderReq{
		Symbol:   "LTCBTC",
		Side:     binance.OrderSideSell,
		Type:     binance.OrderTypeLimit,
		Quantity: "1",
		Price:    "0.1",
	})
	s.Require().NoError(e)
}

func (s *mockedOrderTestSuite) TestNewOrderResult() {
	var expected *binance.OrderRespResult
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		expected = &binance.OrderRespResult{
			Symbol:              req.Symbol,
			OrderID:             int64(rand.Uint32()),
			TransactTime:        int64(rand.Uint32()),
			Price:               req.Price,
			OrigQty:             req.Quantity,
			ExecutedQty:         "0",
			CummulativeQuoteQty: req.QuoteQuantity,
			Status:              binance.OrderStatusNew,
			TimeInForce:         string(req.TimeInForce),
			Type:                req.Type,
			Side:                req.Side,
		}
		return json.Marshal(expected)
	}

	actual, e := s.client.NewOrderResult(&binance.OrderReq{
		Symbol:      "LTCBTC",
		Side:        binance.OrderSideSell,
		Type:        binance.OrderTypeLimit,
		TimeInForce: binance.TimeInForceGTC,
		Quantity:    "1",
		Price:       "0.1",
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedOrderTestSuite) TestNewOrderFull() {
	var expected *binance.OrderRespFull
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		expected = &binance.OrderRespFull{
			Symbol:              req.Symbol,
			OrderID:             int64(rand.Uint32()),
			TransactTime:        int64(rand.Uint32()),
			Price:               req.Price,
			OrigQty:             req.Quantity,
			ExecutedQty:         "0",
			CummulativeQuoteQty: req.QuoteQuantity,
			Status:              binance.OrderStatusNew,
			TimeInForce:         string(req.TimeInForce),
			Type:                req.Type,
			Side:                req.Side,
		}
		return json.Marshal(expected)
	}

	_, e := s.client.NewOrderFull(&binance.OrderReq{
		Symbol:      "LTCBTC",
		Side:        binance.OrderSideSell,
		Type:        binance.OrderTypeLimit,
		TimeInForce: binance.TimeInForceGTC,
		Quantity:    "1",
		Price:       "0.1",
	})
	s.Require().NoError(e)
}

func (s *mockedOrderTestSuite) TestQueryCancelOrder() {
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		return json.Marshal(&binance.OrderRespAck{
			Symbol:       req.Symbol,
			OrderID:      int64(rand.Uint32()),
			TransactTime: int64(rand.Uint32()),
		})
	}
	createReq := &binance.OrderReq{
		Symbol:      "LTCBTC",
		Side:        binance.OrderSideSell,
		Type:        binance.OrderTypeLimit,
		TimeInForce: binance.TimeInForceGTC,
		Quantity:    "1",
		Price:       "0.1",
	}
	resp, e := s.client.NewOrder(createReq)
	s.Require().NoError(e)

	var expectedQuery *binance.QueryOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.QueryOrderReq{}, data)
		req := data.(*binance.QueryOrderReq)
		expectedQuery = &binance.QueryOrder{
			Symbol:              req.Symbol,
			OrderID:             req.OrderID,
			Price:               createReq.Price,
			OrigQty:             createReq.Quantity,
			ExecutedQty:         "0",
			CummulativeQuoteQty: createReq.Quantity,
			Status:              binance.OrderStatusNew,
			TimeInForce:         createReq.TimeInForce,
			Type:                createReq.Type,
			Side:                createReq.Side,
			Time:                int64(rand.Uint32()),
			UpdateTime:          int64(rand.Uint32()),
			OrigQuoteOrderQty:   createReq.QuoteQuantity,
		}
		return json.Marshal(expectedQuery)
	}
	actualQuery, e := s.client.QueryOrder(&binance.QueryOrderReq{
		Symbol:  "LTCBTC",
		OrderID: resp.OrderID,
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expectedQuery, actualQuery)

	var expectedCancel *binance.CancelOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.CancelOrderReq{}, data)
		expectedCancel = &binance.CancelOrder{
			Symbol:              actualQuery.Symbol,
			OrderID:             actualQuery.OrderID,
			Price:               actualQuery.Price,
			OrigQty:             actualQuery.OrigQty,
			ExecutedQty:         actualQuery.ExecutedQty,
			CummulativeQuoteQty: actualQuery.CummulativeQuoteQty,
			Status:              actualQuery.Status,
			TimeInForce:         actualQuery.TimeInForce,
			Type:                actualQuery.Type,
			Side:                actualQuery.Side,
		}
		return json.Marshal(expectedCancel)
	}
	actualCancel, e := s.client.CancelOrder(&binance.CancelOrderReq{
		Symbol:  "LTCBTC",
		OrderID: resp.OrderID,
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expectedCancel, actualCancel)
}

func (s *mockedOrderTestSuite) TestCancelReplaceOrder() {
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		return json.Marshal(&binance.OrderRespAck{
			Symbol:       req.Symbol,
			OrderID:      int64(rand.Uint32()),
			TransactTime: int64(rand.Uint32()),
		})
	}
	createReq := &binance.OrderReq{
		Symbol:      "LTCBTC",
		Side:        binance.OrderSideSell,
		Type:        binance.OrderTypeLimit,
		TimeInForce: binance.TimeInForceGTC,
		Quantity:    "1",
		Price:       "0.1",
	}
	resp, e := s.client.NewOrder(createReq)
	s.Require().NoError(e)

	var expectedQuery *binance.QueryOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.QueryOrderReq{}, data)
		req := data.(*binance.QueryOrderReq)
		expectedQuery = &binance.QueryOrder{
			Symbol:              req.Symbol,
			OrderID:             req.OrderID,
			Price:               createReq.Price,
			OrigQty:             createReq.Quantity,
			ExecutedQty:         "0",
			CummulativeQuoteQty: createReq.Quantity,
			Status:              binance.OrderStatusNew,
			TimeInForce:         createReq.TimeInForce,
			Type:                createReq.Type,
			Side:                createReq.Side,
			Time:                int64(rand.Uint32()),
			UpdateTime:          int64(rand.Uint32()),
			OrigQuoteOrderQty:   createReq.QuoteQuantity,
		}
		return json.Marshal(expectedQuery)
	}
	actualQuery, e := s.client.QueryOrder(&binance.QueryOrderReq{
		Symbol:  "LTCBTC",
		OrderID: resp.OrderID,
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expectedQuery, actualQuery)

	req := &binance.CancelReplaceOrderReq{
		OrderReq:      *createReq,
		CancelOrderID: resp.OrderID,
	}

	var expectedCancel *binance.CancelReplaceOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.CancelReplaceOrderReq{}, data)
		expectedCancel = &binance.CancelReplaceOrder{
			CancelResult:   binance.CancelReplaceResultSuccess,
			NewOrderResult: binance.CancelReplaceResultSuccess,
			CancelResponse: binance.CancelOrder{
				Symbol:              actualQuery.Symbol,
				OrderID:             actualQuery.OrderID,
				Price:               actualQuery.Price,
				OrigQty:             actualQuery.OrigQty,
				ExecutedQty:         actualQuery.ExecutedQty,
				CummulativeQuoteQty: actualQuery.CummulativeQuoteQty,
				Status:              actualQuery.Status,
				TimeInForce:         actualQuery.TimeInForce,
				Type:                actualQuery.Type,
				Side:                actualQuery.Side,
			},
			NewOrderResponse: &binance.OrderRespFull{
				Symbol:              req.Symbol,
				OrderID:             int64(rand.Uint32()),
				TransactTime:        int64(rand.Uint32()),
				Price:               req.Price,
				OrigQty:             req.Quantity,
				ExecutedQty:         "0",
				CummulativeQuoteQty: req.QuoteQuantity,
				Status:              binance.OrderStatusNew,
				TimeInForce:         string(req.TimeInForce),
				Type:                req.Type,
				Side:                req.Side,
			},
		}
		return json.Marshal(expectedCancel)
	}
	actualCancel, e := s.client.CancelReplaceOrder(req)
	s.Require().NoError(e)
	s.Require().EqualValues(expectedCancel, actualCancel)
}

func (s *mockedOrderTestSuite) TestOpenOrders() {
	var expected []*binance.QueryOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OpenOrdersReq{}, data)
		req := data.(*binance.OpenOrdersReq)
		expected = append(expected, &binance.QueryOrder{
			Symbol:              req.Symbol,
			OrderID:             int64(rand.Uint32()),
			Price:               "0.1",
			OrigQty:             "1",
			ExecutedQty:         "0",
			CummulativeQuoteQty: "1",
			Status:              binance.OrderStatusNew,
			TimeInForce:         binance.TimeInForceGTC,
			Type:                binance.OrderTypeLimit,
			Side:                binance.OrderSideSell,
			Time:                int64(rand.Uint32()),
			UpdateTime:          int64(rand.Uint32()),
			OrigQuoteOrderQty:   "1",
		})
		return json.Marshal(expected)
	}

	actual, e := s.client.OpenOrders(&binance.OpenOrdersReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedOrderTestSuite) TestAllOrders() {
	var expected []*binance.QueryOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.AllOrdersReq{}, data)
		req := data.(*binance.AllOrdersReq)
		expected = append(expected, &binance.QueryOrder{
			Symbol:              req.Symbol,
			OrderID:             req.OrderID,
			Price:               "0.1",
			OrigQty:             "1",
			ExecutedQty:         "0",
			CummulativeQuoteQty: "1",
			Status:              binance.OrderStatusNew,
			TimeInForce:         binance.TimeInForceGTC,
			Type:                binance.OrderTypeLimit,
			Side:                binance.OrderSideSell,
			Time:                int64(rand.Uint32()),
			UpdateTime:          int64(rand.Uint32()),
			OrigQuoteOrderQty:   "1",
		})
		return json.Marshal(expected)
	}

	actual, e := s.client.AllOrders(&binance.AllOrdersReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedOrderTestSuite) TestCancelOpenOrders() {
	var expected []*binance.CancelOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.CancelOpenOrdersReq{}, data)
		req := data.(*binance.CancelOpenOrdersReq)
		expected = append(expected, &binance.CancelOrder{
			Symbol:              req.Symbol,
			OrderID:             int64(rand.Uint32()),
			Price:               "0.1",
			OrigQty:             "1",
			ExecutedQty:         "0",
			CummulativeQuoteQty: "1",
			Status:              binance.OrderStatusNew,
			TimeInForce:         binance.TimeInForceGTC,
			Type:                binance.OrderTypeLimit,
			Side:                binance.OrderSideSell,
		})
		return json.Marshal(expected)
	}

	actual, e := s.client.CancelOpenOrders(&binance.CancelOpenOrdersReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}
