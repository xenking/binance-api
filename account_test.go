package binance_test

import (
	"math/rand"

	"github.com/segmentio/encoding/json"
	"github.com/xenking/binance-api"
)

type mockedAccountTestSuite struct {
	mockedTestSuite
}

func (s *mockedAccountTestSuite) TestHistoricalTrades() {
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
	_, e := s.client.HistoricalTrades(&binance.HistoricalTradeReq{Symbol: "LTCBTC", Limit: 5})
	s.Require().NoError(e)
}

func (s *mockedAccountTestSuite) TestAggregatedTrades() {
	var expected []*binance.AggregatedTrade

	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.AggregatedTradeReq{}, data)
		expected = []*binance.AggregatedTrade{
			{
				TradeID:   rand.Int63(),
				Price:     "0.1",
				Quantity:  "1",
				Timestamp: rand.Int63(),
			},
		}
		return json.Marshal(expected)
	}
	_, e := s.client.AggregatedTrades(&binance.AggregatedTradeReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
}

func (s *mockedAccountTestSuite) TestAccountTrades() {
	var expected []*binance.AccountTrade
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(&binance.AccountTradesReq{}, data)
		req := data.(*binance.AccountTradesReq)
		expected = []*binance.AccountTrade{
			{
				Symbol:   req.Symbol,
				OrderID:  int64(rand.Uint32()),
				QuoteQty: "1",
				Price:    "0.1",
				Qty:      "1",
				Time:     int64(rand.Uint32()),
			},
		}
		return json.Marshal(expected)
	}

	actual, e := s.client.AccountTrades(&binance.AccountTradesReq{
		Symbol: "LTCBTC",
	})
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedAccountTestSuite) TestAccount() {
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
			AccountType:      binance.PermissionTypeSpot,
			Balances: []*binance.Balance{{
				Asset:  "SNM",
				Free:   "1",
				Locked: "",
			}},
		}
		return json.Marshal(expected)
	}

	actual, e := s.client.Account()
	s.Require().NoError(e)
	s.Require().EqualValues(expected, actual)
}

func (s *mockedAccountTestSuite) TestDataStream() {
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().Nil(data)
		return json.Marshal(&binance.DataStream{
			ListenKey: "stream-key",
		})
	}
	key, err := s.client.DataStream()
	s.Require().NoError(err)
	s.Require().Equal("stream-key", key)
	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DataStream{}, data)
		return nil, nil
	}
	s.Require().NoError(s.client.DataStreamKeepAlive(key))
	s.Require().NoError(s.client.DataStreamClose(key))
}
