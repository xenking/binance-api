package binance

import (
	"math/rand"
	"testing"

	"github.com/segmentio/encoding/json"
	"github.com/stretchr/testify/suite"
)

func TestClient(t *testing.T) {
	suite.Run(t, new(clientTestSuite))
}

func TestClientHTTP2(t *testing.T) {
	suite.Run(t, new(clientHTTP2TestSuite))
}

func TestMockClient(t *testing.T) {
	suite.Run(t, new(mockedTestSuite))
}

type baseTestSuite struct {
	suite.Suite
	api *Client
}

func (s *baseTestSuite) SetupSuite() {
	s.api = NewClient("", "")
}

type clientTestSuite struct {
	baseTestSuite
}

func (s *clientTestSuite) TestPing() {
	s.Require().NoError(s.api.Ping())
}

func (s *clientTestSuite) TestTime() {
	_, e := s.api.Time()
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestTicker() {
	_, e := s.api.Ticker(&TickerReq{"LTCBTC"})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestDepth() {
	_, e := s.api.Depth(&DepthReq{Symbol: "SNMBTC", Limit: 5})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestTrades() {
	_, e := s.api.Trades(&TradeReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestAggregatedTrades() {
	_, e := s.api.AggregatedTrades(&AggregatedTradeReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestKlines() {
	resp, e := s.api.Klines(&KlinesReq{Symbol: "SNMBTC", Interval: KlineInterval1hour, Limit: 5})
	s.Require().NoError(e)
	s.Require().Len(resp, 5)
}

func (s *clientTestSuite) TestAllBookTickers() {
	_, e := s.api.BookTickers()
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestPrices() {
	_, e := s.api.Prices()
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestExchangeInfo() {
	info, err := s.api.ExchangeInfo()
	s.Require().NoError(err)
	s.Require().NotNil(info)
	s.Require().NotEmpty(info.Symbols)
}

type clientHTTP2TestSuite struct {
	clientTestSuite
}

func (s *clientHTTP2TestSuite) SetupSuite() {
	var err error
	s.api, err = NewClientHTTP2("", "")
	s.Require().NoError(err)
}

type mockedClient struct {
	Response func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error)
	window   int
}

func (m *mockedClient) UsedWeight() map[string]int {
	panic("not used")
}

func (m *mockedClient) OrderCount() map[string]int {
	panic("not used")
}

func (m *mockedClient) RetryAfter() int {
	panic("not used")
}

func (m *mockedClient) SetWindow(w int) {
	m.window = w
}

func (m *mockedClient) Do(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
	return m.Response(method, endpoint, data, sign, stream)
}

type mockedTestSuite struct {
	baseTestSuite
	mock *mockedClient
}

func (s *mockedTestSuite) SetupSuite() {
	s.mock = &mockedClient{}
	s.api = NewCustomClient(s.mock)
}

func (s *mockedTestSuite) SetupTest() {
	s.mock.Response = nil
}

func (s *mockedTestSuite) TestNewOrder() {
	var expected *OrderRespAck
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&OrderReq{}, data)
		req := data.(*OrderReq)
		expected = &OrderRespAck{
			Symbol:       req.Symbol,
			OrderID:      rand.Uint64(),
			TransactTime: rand.Uint64(),
		}
		return json.Marshal(expected)
	}

	actual, e := s.api.NewOrder(&OrderReq{
		Symbol:      "SNMBTC",
		Side:        OrderSideSell,
		Type:        OrderTypeLimit,
		TimeInForce: TimeInForceGTC,
		Quantity:    "1",
		Price:       "0.1",
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedTestSuite) TestNewOrderResult() {
	var expected *OrderRespResult
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&OrderReq{}, data)
		req := data.(*OrderReq)
		expected = &OrderRespResult{
			Symbol:              req.Symbol,
			OrderID:             rand.Uint64(),
			TransactTime:        rand.Uint64(),
			Price:               req.Price,
			OrigQty:             req.Quantity,
			ExecutedQty:         "0",
			CummulativeQuoteQty: req.QuoteQuantity,
			Status:              OrderStatusNew,
			TimeInForce:         string(req.TimeInForce),
			Type:                req.Type,
			Side:                req.Side,
		}
		return json.Marshal(expected)
	}

	actual, e := s.api.NewOrderResult(&OrderReq{
		Symbol:      "SNMBTC",
		Side:        OrderSideSell,
		Type:        OrderTypeLimit,
		TimeInForce: TimeInForceGTC,
		Quantity:    "1",
		Price:       "0.1",
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedTestSuite) TestNewOrderFull() {
	var expected *OrderRespFull
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&OrderReq{}, data)
		req := data.(*OrderReq)
		expected = &OrderRespFull{
			Symbol:              req.Symbol,
			OrderID:             rand.Uint64(),
			TransactTime:        rand.Uint64(),
			Price:               req.Price,
			OrigQty:             req.Quantity,
			ExecutedQty:         "0",
			CummulativeQuoteQty: req.QuoteQuantity,
			Status:              OrderStatusNew,
			TimeInForce:         string(req.TimeInForce),
			Type:                req.Type,
			Side:                req.Side,
		}
		return json.Marshal(expected)
	}

	_, e := s.api.NewOrderFull(&OrderReq{
		Symbol:      "SNMBTC",
		Side:        OrderSideSell,
		Type:        OrderTypeLimit,
		TimeInForce: TimeInForceGTC,
		Quantity:    "1",
		Price:       "0.1",
	})
	s.Require().NoError(e)
}

func (s *mockedTestSuite) TestQueryCancelOrder() {
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&OrderReq{}, data)
		req := data.(*OrderReq)
		return json.Marshal(&OrderRespAck{
			Symbol:       req.Symbol,
			OrderID:      rand.Uint64(),
			TransactTime: rand.Uint64(),
		})
	}
	createReq := &OrderReq{
		Symbol:      "SNMBTC",
		Side:        OrderSideSell,
		Type:        OrderTypeLimit,
		TimeInForce: TimeInForceGTC,
		Quantity:    "1",
		Price:       "0.1",
	}
	resp, e := s.api.NewOrder(createReq)
	s.Require().NoError(e)

	var expectedQuery *QueryOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&QueryOrderReq{}, data)
		req := data.(*QueryOrderReq)
		expectedQuery = &QueryOrder{
			Symbol:              req.Symbol,
			OrderID:             req.OrderID,
			Price:               createReq.Price,
			OrigQty:             createReq.Quantity,
			ExecutedQty:         "0",
			CummulativeQuoteQty: createReq.Quantity,
			Status:              OrderStatusNew,
			TimeInForce:         createReq.TimeInForce,
			Type:                createReq.Type,
			Side:                createReq.Side,
			Time:                rand.Uint64(),
			UpdateTime:          rand.Uint64(),
			OrigQuoteOrderQty:   createReq.QuoteQuantity,
		}
		return json.Marshal(expectedQuery)
	}
	actualQuery, e := s.api.QueryOrder(&QueryOrderReq{
		Symbol:  "SNMBTC",
		OrderID: resp.OrderID,
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expectedQuery, actualQuery)

	var expectedCancel *CancelOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&CancelOrderReq{}, data)
		expectedCancel = &CancelOrder{
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
	actualCancel, e := s.api.CancelOrder(&CancelOrderReq{
		Symbol:  "SNMBTC",
		OrderID: resp.OrderID,
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expectedCancel, actualCancel)
}

func (s *mockedTestSuite) TestDataStream() {
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().Nil(data)
		return json.Marshal(&DatastreamReq{
			ListenKey: "stream-key",
		})
	}
	key, err := s.api.DataStream()
	s.Require().NoError(err)
	s.Require().Equal("stream-key", key)
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(DatastreamReq{}, data)
		return nil, nil
	}
	s.Require().NoError(s.api.DataStreamKeepAlive(key))
	s.Require().NoError(s.api.DataStreamClose(key))
}

func (s *mockedTestSuite) TestAllOrders() {
	var expected []*QueryOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&AllOrdersReq{}, data)
		req := data.(*AllOrdersReq)
		expected = append(expected, &QueryOrder{
			Symbol:              req.Symbol,
			OrderID:             req.OrderID,
			Price:               "0.1",
			OrigQty:             "1",
			ExecutedQty:         "0",
			CummulativeQuoteQty: "1",
			Status:              OrderStatusNew,
			TimeInForce:         TimeInForceGTC,
			Type:                OrderTypeLimit,
			Side:                OrderSideSell,
			Time:                rand.Uint64(),
			UpdateTime:          rand.Uint64(),
			OrigQuoteOrderQty:   "1",
		})
		return json.Marshal(expected)
	}

	actual, e := s.api.AllOrders(&AllOrdersReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedTestSuite) TestOpenOrders() {
	var expected []*QueryOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&OpenOrdersReq{}, data)
		req := data.(*OpenOrdersReq)
		expected = append(expected, &QueryOrder{
			Symbol:              req.Symbol,
			OrderID:             rand.Uint64(),
			Price:               "0.1",
			OrigQty:             "1",
			ExecutedQty:         "0",
			CummulativeQuoteQty: "1",
			Status:              OrderStatusNew,
			TimeInForce:         TimeInForceGTC,
			Type:                OrderTypeLimit,
			Side:                OrderSideSell,
			Time:                rand.Uint64(),
			UpdateTime:          rand.Uint64(),
			OrigQuoteOrderQty:   "1",
		})
		return json.Marshal(expected)
	}

	actual, e := s.api.OpenOrders(&OpenOrdersReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedTestSuite) TestCancelOpenOrders() {
	var expected []*CancelOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&CancelOpenOrdersReq{}, data)
		req := data.(*CancelOpenOrdersReq)
		expected = append(expected, &CancelOrder{
			Symbol:              req.Symbol,
			OrderID:             rand.Uint64(),
			Price:               "0.1",
			OrigQty:             "1",
			ExecutedQty:         "0",
			CummulativeQuoteQty: "1",
			Status:              OrderStatusNew,
			TimeInForce:         TimeInForceGTC,
			Type:                OrderTypeLimit,
			Side:                OrderSideSell,
		})
		return json.Marshal(expected)
	}

	actual, e := s.api.CancelOpenOrders(&CancelOpenOrdersReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedTestSuite) TestAccount() {
	var expected *AccountInfo
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().Nil(data)
		expected = &AccountInfo{
			MakerCommission:  15,
			TakerCommission:  15,
			BuyerCommission:  0,
			SellerCommission: 0,
			CanTrade:         true,
			CanWithdraw:      true,
			CanDeposit:       true,
			AccountType:      AccountTypeSpot,
			Balances: []*Balance{{
				Asset:  "SNM",
				Free:   "1",
				Locked: "",
			}},
		}
		return json.Marshal(expected)
	}

	actual, e := s.api.Account()
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}
