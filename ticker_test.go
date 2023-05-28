package binance_test

import "github.com/xenking/binance-api"

type tickerTestSuite struct {
	baseTestSuite
}

func (s *tickerTestSuite) TestTickers24h() {
	_, e := s.client.Tickers24h(nil)
	s.Require().NoError(e)

	resp, e := s.client.Tickers24h(&binance.Tickers24hReq{Symbols: []string{"LTCBTC", "ETHBTC"}})
	s.Require().NoError(e)
	s.Require().Len(resp, 2)
}

func (s *tickerTestSuite) TestTickers24hMini() {
	_, e := s.client.Tickers24h(nil)
	s.Require().NoError(e)

	resp, e := s.client.Tickers24hMini(&binance.Tickers24hReq{Symbols: []string{"LTCBTC", "ETHBTC"}})
	s.Require().NoError(e)
	s.Require().Len(resp, 2)
}

func (s *tickerTestSuite) TestTicker24h() {
	_, e := s.client.Ticker24h(&binance.Ticker24hReq{})
	s.Require().ErrorIs(e, binance.ErrEmptySymbol)

	_, e = s.client.Ticker24h(&binance.Ticker24hReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
}

func (s *tickerTestSuite) TestTicker24hMini() {
	_, e := s.client.Ticker24hMini(&binance.Ticker24hReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
}

func (s *tickerTestSuite) TestBookTickers() {
	_, e := s.client.BookTickers(nil)
	s.Require().NoError(e)

	resp, e := s.client.BookTickers(&binance.BookTickersReq{Symbols: []string{"LTCBTC", "ETHBTC"}})
	s.Require().NoError(e)
	s.Require().Len(resp, 2)
}

func (s *tickerTestSuite) TestBookTicker() {
	_, e := s.client.BookTicker(&binance.BookTickerReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
}

func (s *tickerTestSuite) TestTickers() {
	resp, e := s.client.Tickers(&binance.TickersReq{Symbols: []string{"LTCBTC", "ETHBTC"}})
	s.Require().NoError(e)
	s.Require().Len(resp, 2)
}

func (s *tickerTestSuite) TestTickersMini() {
	resp, e := s.client.TickersMini(&binance.TickersReq{Symbols: []string{"LTCBTC", "ETHBTC"}})
	s.Require().NoError(e)
	s.Require().Len(resp, 2)
}

func (s *tickerTestSuite) TestTicker() {
	_, e := s.client.Ticker(&binance.TickerReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
}

func (s *tickerTestSuite) TestTickerMini() {
	_, e := s.client.TickerMini(&binance.TickerReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
}
