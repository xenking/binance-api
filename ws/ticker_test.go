package ws_test

import (
	"context"
	"time"

	"github.com/xenking/binance-api"
	"github.com/xenking/binance-api/ws"
)

type tickerTestSuite struct {
	baseTestSuite
}

func (s *tickerTestSuite) TestDepth_Read() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	depth, err := s.ws.DiffDepth(ctx, symbol, ws.Frequency1000ms)
	s.Require().NoError(err)
	defer depth.Close()

	u, err := depth.Read()

	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *tickerTestSuite) TestDepth_Stream() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	depth, err := s.ws.DiffDepth(ctx, symbol, ws.Frequency1000ms)
	s.Require().NoError(err)

	for u := range depth.Stream() {
		s.Require().Equal(symbol, u.Symbol)
		break
	}
	s.Require().NoError(depth.Close())
}

func (s *tickerTestSuite) TestDepthLevel_Read() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	depthLevel, err := s.ws.DepthLevel(ctx, symbol, ws.DepthLevel10, ws.Frequency1000ms)
	s.Require().NoError(err)
	defer depthLevel.Close()

	u, err := depthLevel.Read()

	s.Require().NoError(err)
	s.Require().Equal(len(u.Asks), 10)
	s.Require().Equal(len(u.Bids), 10)
}

func (s *tickerTestSuite) TestDepthLevel_Stream() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	depthLevel, err := s.ws.DepthLevel(ctx, symbol, ws.DepthLevel5, ws.Frequency1000ms)
	s.Require().NoError(err)

	for u := range depthLevel.Stream() {
		s.Require().Equal(len(u.Asks), 5)
		s.Require().Equal(len(u.Bids), 5)

		break
	}

	s.Require().NoError(depthLevel.Close())
}

func (s *tickerTestSuite) TestKlines_Read() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	klines, err := s.ws.Klines(ctx, symbol, binance.KlineInterval1min)
	s.Require().NoError(err)
	defer klines.Close()

	u, err := klines.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *tickerTestSuite) TestKlines_Stream() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	klines, err := s.ws.Klines(ctx, symbol, binance.KlineInterval1min)
	s.Require().NoError(err)

	for u := range klines.Stream() {
		s.Require().Equal(symbol, u.Symbol)
		break
	}
	s.Require().NoError(klines.Close())
}

func (s *tickerTestSuite) TestAllMarketTickers_Read() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tickers, err := s.ws.AllMarketTickers(ctx)
	s.Require().NoError(err)
	defer tickers.Close()

	u, err := tickers.Read()
	s.Require().NoError(err)
	s.Require().NotEmpty(u)
}

func (s *tickerTestSuite) TestAllMarketTickers_Stream() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tickers, err := s.ws.AllMarketTickers(ctx)
	s.Require().NoError(err)

	for u := range tickers.Stream() {
		s.Require().NotEmpty(u)
		break
	}
	s.Require().NoError(tickers.Close())
}

func (s *tickerTestSuite) TestAllMarketMiniTickers_Read() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tickers, err := s.ws.AllMarketMiniTickers(ctx)
	s.Require().NoError(err)
	defer tickers.Close()

	u, err := tickers.Read()
	s.Require().NoError(err)
	s.Require().NotEmpty(u)
}

func (s *tickerTestSuite) TestAllMarketMiniTickers_Stream() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tickers, err := s.ws.AllMarketMiniTickers(ctx)
	s.Require().NoError(err)

	for u := range tickers.Stream() {
		s.Require().NotEmpty(u)
		break
	}
	s.Require().NoError(tickers.Close())
}

func (s *tickerTestSuite) TestIndividualTicker_Read() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticker, err := s.ws.IndividualTicker(ctx, symbol)
	s.Require().NoError(err)
	defer ticker.Close()

	u, err := ticker.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *tickerTestSuite) TestIndividualTickers_Stream() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticker, err := s.ws.IndividualTicker(ctx, symbol)
	s.Require().NoError(err)

	for u := range ticker.Stream() {
		s.Require().Equal(symbol, u.Symbol)
		break
	}
	s.Require().NoError(ticker.Close())

}

func (s *tickerTestSuite) TestIndividualMiniTicker_Read() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticker, err := s.ws.IndividualMiniTicker(ctx, symbol)
	s.Require().NoError(err)
	defer ticker.Close()

	u, err := ticker.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *tickerTestSuite) TestIndividualMiniTickers_Stream() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	miniTicker, err := s.ws.IndividualMiniTicker(ctx, symbol)
	s.Require().NoError(err)

	for u := range miniTicker.Stream() {
		s.Require().Equal(symbol, u.Symbol)
		break
	}
	s.Require().NoError(miniTicker.Close())

}

func (s *tickerTestSuite) TestIndividualBookTicker_Read() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bookTicker, err := s.ws.IndividualBookTicker(ctx, symbol)
	s.Require().NoError(err)
	defer bookTicker.Close()

	u, err := bookTicker.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *tickerTestSuite) TestIndividualBookTickers_Stream() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bookTicker, err := s.ws.IndividualBookTicker(ctx, symbol)
	s.Require().NoError(err)

	for u := range bookTicker.Stream() {
		s.Require().Equal(symbol, u.Symbol)
		break
	}
	s.Require().NoError(bookTicker.Close())
}

func (s *tickerTestSuite) TestAggTrades_Read() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	trades, err := s.ws.AggTrades(ctx, symbol)
	s.Require().NoError(err)
	defer trades.Close()

	u, err := trades.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *tickerTestSuite) TestAggTrades_Stream() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	trades, err := s.ws.AggTrades(ctx, symbol)
	s.Require().NoError(err)

	for u := range trades.Stream() {
		s.Require().Equal(symbol, u.Symbol)
		break
	}
	s.Require().NoError(trades.Close())

}

func (s *tickerTestSuite) TestTrades_Read() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	trades, err := s.ws.Trades(ctx, symbol)
	s.Require().NoError(err)
	defer trades.Close()

	u, err := trades.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *tickerTestSuite) TestTrades_Stream() {
	const symbol = "ETHBTC"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	trades, err := s.ws.Trades(ctx, symbol)
	s.Require().NoError(err)

	for u := range trades.Stream() {
		s.Require().Equal(symbol, u.Symbol)
		break
	}
	s.Require().NoError(trades.Close())
}
