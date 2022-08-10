package binance_test

import (
	"github.com/xenking/binance-api"
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
	api *binance.Client
}

func (s *baseTestSuite) SetupSuite() {
	s.api = binance.NewClient("", "")
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
	_, e := s.api.Ticker(&binance.TickerReq{"LTCBTC"})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestTickers() {
	_, e := s.api.Tickers()
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestDepth() {
	_, e := s.api.Depth(&binance.DepthReq{Symbol: "SNMBTC", Limit: 5})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestTrades() {
	_, e := s.api.Trades(&binance.TradeReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestAggregatedTrades() {
	_, e := s.api.AggregatedTrades(&binance.AggregatedTradeReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestKlines() {
	resp, e := s.api.Klines(&binance.KlinesReq{Symbol: "SNMBTC", Interval: binance.KlineInterval1hour, Limit: 5})
	s.Require().NoError(e)
	s.Require().Len(resp, 5)
}

func (s *clientTestSuite) TestAllBookTickers() {
	_, e := s.api.BookTickers()
	s.Require().NoError(e)
}
func (s *clientTestSuite) TestBookTicker() {
	_, e := s.api.BookTicker(&binance.BookTickerReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestAvgPrice() {
	_, e := s.api.AvgPrice(&binance.AvgPriceReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestPrices() {
	_, e := s.api.Prices()
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestPrice() {
	_, e := s.api.Price(&binance.TickerPriceReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestExchangeInfo() {
	info, err := s.api.ExchangeInfo()
	s.Require().NoError(err)
	s.Require().NotNil(info)
	s.Require().NotEmpty(info.Symbols)
}

func (s *clientTestSuite) TestExchangeInfoSymbol() {
	info, err := s.api.ExchangeInfoSymbol(&binance.ExchangeInfoReq{Symbol: "SNMBTC"})
	s.Require().NoError(err)
	s.Require().NotNil(info)
	s.Require().NotEmpty(info.Symbols)
}

type clientHTTP2TestSuite struct {
	clientTestSuite
}

func (s *clientHTTP2TestSuite) SetupSuite() {
	var err error
	s.api, err = binance.NewClientHTTP2("", "")
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
	s.api = binance.NewCustomClient(s.mock)
}

func (s *mockedTestSuite) SetupTest() {
	s.mock.Response = nil
}

func (s *mockedTestSuite) TestHistoricalTrades() {
	var expected []*binance.Trade

	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.HistoricalTradeReq{}, data)
		expected = []*binance.Trade{
			{
				ID:       rand.Int63(),
				Price:    "0.1",
				Qty:      "1",
				QuoteQty: "1",
				Time:     rand.Int63(),
			},
		}
		return json.Marshal(expected)
	}
	_, e := s.api.HistoricalTrades(&binance.HistoricalTradeReq{Symbol: "SNMBTC", Limit: 5})
	s.Require().NoError(e)
}

func (s *mockedTestSuite) TestNewOrder() {
	var expected *binance.OrderRespAck
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		expected = &binance.OrderRespAck{
			Symbol:       req.Symbol,
			OrderID:      rand.Uint64(),
			TransactTime: rand.Uint64(),
		}
		return json.Marshal(expected)
	}

	actual, e := s.api.NewOrder(&binance.OrderReq{
		Symbol:   "SNMBTC",
		Side:     binance.OrderSideSell,
		Type:     binance.OrderTypeLimit,
		Quantity: "1",
		Price:    "0.1",
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedTestSuite) TestNewMarketOrder() {
	var expected *binance.OrderRespAck
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		expected = &binance.OrderRespAck{
			Symbol:       req.Symbol,
			OrderID:      rand.Uint64(),
			TransactTime: rand.Uint64(),
		}
		return json.Marshal(expected)
	}

	actual, e := s.api.NewOrder(&binance.OrderReq{
		Symbol:   "SNMBTC",
		Side:     binance.OrderSideSell,
		Type:     binance.OrderTypeMarket,
		Quantity: "1",
		Price:    "0.1",
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedTestSuite) TestNewOrderTest() {
	var expected *binance.OrderRespAck
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		expected = &binance.OrderRespAck{
			Symbol:       req.Symbol,
			OrderID:      rand.Uint64(),
			TransactTime: rand.Uint64(),
		}
		return json.Marshal(expected)
	}

	e := s.api.NewOrderTest(&binance.OrderReq{
		Symbol:   "SNMBTC",
		Side:     binance.OrderSideSell,
		Type:     binance.OrderTypeLimit,
		Quantity: "1",
		Price:    "0.1",
	})
	s.Require().NoError(e)
}

func (s *mockedTestSuite) TestNewOrderResult() {
	var expected *binance.OrderRespResult
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		expected = &binance.OrderRespResult{
			Symbol:              req.Symbol,
			OrderID:             rand.Uint64(),
			TransactTime:        rand.Uint64(),
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

	actual, e := s.api.NewOrderResult(&binance.OrderReq{
		Symbol:      "SNMBTC",
		Side:        binance.OrderSideSell,
		Type:        binance.OrderTypeLimit,
		TimeInForce: binance.TimeInForceGTC,
		Quantity:    "1",
		Price:       "0.1",
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedTestSuite) TestNewOrderFull() {
	var expected *binance.OrderRespFull
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		expected = &binance.OrderRespFull{
			Symbol:              req.Symbol,
			OrderID:             rand.Uint64(),
			TransactTime:        rand.Uint64(),
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

	_, e := s.api.NewOrderFull(&binance.OrderReq{
		Symbol:      "SNMBTC",
		Side:        binance.OrderSideSell,
		Type:        binance.OrderTypeLimit,
		TimeInForce: binance.TimeInForceGTC,
		Quantity:    "1",
		Price:       "0.1",
	})
	s.Require().NoError(e)
}

func (s *mockedTestSuite) TestQueryCancelOrder() {
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OrderReq{}, data)
		req := data.(*binance.OrderReq)
		return json.Marshal(&binance.OrderRespAck{
			Symbol:       req.Symbol,
			OrderID:      rand.Uint64(),
			TransactTime: rand.Uint64(),
		})
	}
	createReq := &binance.OrderReq{
		Symbol:      "SNMBTC",
		Side:        binance.OrderSideSell,
		Type:        binance.OrderTypeLimit,
		TimeInForce: binance.TimeInForceGTC,
		Quantity:    "1",
		Price:       "0.1",
	}
	resp, e := s.api.NewOrder(createReq)
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
			Time:                rand.Uint64(),
			UpdateTime:          rand.Uint64(),
			OrigQuoteOrderQty:   createReq.QuoteQuantity,
		}
		return json.Marshal(expectedQuery)
	}
	actualQuery, e := s.api.QueryOrder(&binance.QueryOrderReq{
		Symbol:  "SNMBTC",
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
	actualCancel, e := s.api.CancelOrder(&binance.CancelOrderReq{
		Symbol:  "SNMBTC",
		OrderID: resp.OrderID,
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expectedCancel, actualCancel)
}

func (s *mockedTestSuite) TestDataStream() {
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().Nil(data)
		return json.Marshal(&binance.DatastreamReq{
			ListenKey: "stream-key",
		})
	}
	key, err := s.api.DataStream()
	s.Require().NoError(err)
	s.Require().Equal("stream-key", key)
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DatastreamReq{}, data)
		return nil, nil
	}
	s.Require().NoError(s.api.DataStreamKeepAlive(key))
	s.Require().NoError(s.api.DataStreamClose(key))
}

func (s *mockedTestSuite) TestAllOrders() {
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
			Time:                rand.Uint64(),
			UpdateTime:          rand.Uint64(),
			OrigQuoteOrderQty:   "1",
		})
		return json.Marshal(expected)
	}

	actual, e := s.api.AllOrders(&binance.AllOrdersReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedTestSuite) TestOpenOrders() {
	var expected []*binance.QueryOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.OpenOrdersReq{}, data)
		req := data.(*binance.OpenOrdersReq)
		expected = append(expected, &binance.QueryOrder{
			Symbol:              req.Symbol,
			OrderID:             rand.Uint64(),
			Price:               "0.1",
			OrigQty:             "1",
			ExecutedQty:         "0",
			CummulativeQuoteQty: "1",
			Status:              binance.OrderStatusNew,
			TimeInForce:         binance.TimeInForceGTC,
			Type:                binance.OrderTypeLimit,
			Side:                binance.OrderSideSell,
			Time:                rand.Uint64(),
			UpdateTime:          rand.Uint64(),
			OrigQuoteOrderQty:   "1",
		})
		return json.Marshal(expected)
	}

	actual, e := s.api.OpenOrders(&binance.OpenOrdersReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedTestSuite) TestCancelOpenOrders() {
	var expected []*binance.CancelOrder
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.CancelOpenOrdersReq{}, data)
		req := data.(*binance.CancelOpenOrdersReq)
		expected = append(expected, &binance.CancelOrder{
			Symbol:              req.Symbol,
			OrderID:             rand.Uint64(),
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

	actual, e := s.api.CancelOpenOrders(&binance.CancelOpenOrdersReq{Symbol: "SNMBTC"})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedTestSuite) TestAccount() {
	var expected *binance.AccountInfo
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().Nil(data)
		expected = &binance.AccountInfo{
			MakerCommission:  15,
			TakerCommission:  15,
			BuyerCommission:  0,
			SellerCommission: 0,
			CanTrade:         true,
			CanWithdraw:      true,
			CanDeposit:       true,
			AccountType:      binance.AccountTypeSpot,
			Balances: []*binance.Balance{{
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

func (s *mockedTestSuite) TestAccountTrades() {
	var expected *binance.AccountTrades
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.AccountTradesReq{}, data)
		req := data.(*binance.AccountTradesReq)
		expected = &binance.AccountTrades{
			Symbol:   req.Symbol,
			OrderID:  rand.Uint64(),
			QuoteQty: "1",
			Price:    "0.1",
			Qty:      "1",
			Time:     rand.Uint64(),
		}
		return json.Marshal(expected)
	}

	actual, e := s.api.AccountTrades(&binance.AccountTradesReq{
		Symbol: "SNMBTC",
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}
