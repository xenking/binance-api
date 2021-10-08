package ws

import (
	"math/rand"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/segmentio/encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/xenking/fastws"

	"github.com/xenking/binance-api"
)

func TestWSClient(t *testing.T) {
	suite.Run(t, new(clientTestSuite))
}

func TestDataStream(t *testing.T) {
	suite.Run(t, new(mockedTestSuite))
}

type baseTestSuite struct {
	suite.Suite
	ws *Client
}

func (s *baseTestSuite) SetupTest() {
	s.ws = NewClient()
}

type clientTestSuite struct {
	baseTestSuite
}

func (s *clientTestSuite) TestDepth_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Depth(symbol, Frequency1000ms)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()

	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestDepth_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Depth(symbol, Frequency1000ms)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestKlines_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Klines(symbol, binance.KlineInterval1min)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestKlines_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Klines(symbol, binance.KlineInterval1min)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestAllMarketTickers_Read() {
	ws, err := s.ws.AllMarketTickers()
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().NotEmpty(u)
}

func (s *clientTestSuite) TestAllMarketTickers_Stream() {
	ws, err := s.ws.AllMarketTickers()
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().NotEmpty(u)
		break
	}
}

func (s *clientTestSuite) TestAllMarketMiniTickers_Read() {
	ws, err := s.ws.AllMarketMiniTickers()
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().NotEmpty(u)
}

func (s *clientTestSuite) TestAllMarketMiniTickers_Stream() {
	ws, err := s.ws.AllMarketMiniTickers()
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().NotEmpty(u)
		break
	}
}

func (s *clientTestSuite) TestAllBookTickers_Read() {
	ws, err := s.ws.AllBookTickers()
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().NotEmpty(u)
}

func (s *clientTestSuite) TestAllBookTickers_Stream() {
	ws, err := s.ws.AllBookTickers()
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().NotEmpty(u)
		break
	}
}

func (s *clientTestSuite) TestIndivTicker_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestIndivTickers_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestIndivMiniTicker_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivMiniTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestIndivMiniTickers_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivMiniTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestIndivBookTicker_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivBookTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestIndivBookTickers_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.IndivBookTicker(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestAggTrades_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.AggTrades(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestAggTrades_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.AggTrades(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

func (s *clientTestSuite) TestTrades_Read() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Trades(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	u, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(symbol, u.Symbol)
}

func (s *clientTestSuite) TestTrades_Stream() {
	const symbol = "ETHBTC"

	ws, err := s.ws.Trades(symbol)
	s.Require().NoError(err)
	defer ws.Close()

	for u := range ws.Stream() {
		s.Require().NoError(ws.err)
		s.Require().Equal(symbol, u.Symbol)
		break
	}
}

type mockedClient struct {
	Response func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error)
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

func (m *mockedClient) SetWindow(_ int) {
	panic("not used")
}

func (m *mockedClient) Do(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
	return m.Response(method, endpoint, data, sign, stream)
}

type mockedTestSuite struct {
	baseTestSuite
	api  *binance.Client
	mock *mockedClient
}

func (s *mockedTestSuite) SetupSuite() {
	s.mock = &mockedClient{}
	s.api = binance.NewCustomClient(s.mock)
}

func (s *mockedTestSuite) SetupTest() {
}

func (s *mockedTestSuite) TestAccountInfo_Read() {
	s.ws = NewCustomClient("ws://localhost:9844/")
	ln, err := net.Listen("tcp", ":9844")
	s.Require().NoError(err)

	expected := &BalanceUpdate{
		EventType:    UpdateTypeBalanceUpdate,
		Time:         rand.Uint64(),
		Asset:        "BTC",
		BalanceDelta: "1",
	}
	wsHandler := func(conn *fastws.Conn) {
		b, err := json.Marshal(expected)
		s.Require().NoError(err)

		written, err := conn.Write(b)
		s.Require().NoError(err)
		s.Require().NotZero(written)
	}
	http.HandleFunc("/stream-key", fastws.NetUpgrade(wsHandler))

	ch := make(chan struct{}, 1)
	go func() {
		http.Serve(ln, nil)
		ch <- struct{}{}
	}()

	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().Nil(data)
		return json.Marshal(&binance.DatastreamReq{
			ListenKey: "stream-key",
		})
	}

	key, err := s.api.DataStream()
	s.Require().NoError(err)

	s.mock.Response = func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DatastreamReq{}, data)
		return nil, nil
	}

	ws, err := s.ws.AccountInfo(key)
	s.Require().NoError(err)

	updateType, actual, err := ws.Read()
	s.Require().NoError(err)
	s.Require().Equal(UpdateTypeBalanceUpdate, updateType)
	s.Require().EqualValues(expected, actual)

	err = s.api.DataStreamClose(key)
	s.Require().NoError(err)
	err = ws.Close()
	s.Require().NoError(err)
	err = ln.Close()
	s.Require().NoError(err)

	select {
	case <-ch:
	case <-time.After(time.Second * 5):
		s.Fail("timeout")
	}
}
