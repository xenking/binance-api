package binance_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/xenking/binance-api"
)

func TestClient(t *testing.T) {
	client := binance.NewClient("", "")

	suite.Run(t, &clientTestSuite{newBaseTestSuite(client)})
	suite.Run(t, &tickerTestSuite{newBaseTestSuite(client)})
}

func TestClientHTTP2(t *testing.T) {
	client, err := binance.NewClientHTTP2("", "")
	require.New(t).NoError(err)

	suite.Run(t, &clientTestSuite{newBaseTestSuite(client)})
	suite.Run(t, &tickerTestSuite{newBaseTestSuite(client)})
}

func TestMockClient(t *testing.T) {
	suite.Run(t, new(mockedAccountTestSuite))
	suite.Run(t, new(mockedOrderTestSuite))
	suite.Run(t, new(mockedOCOTestSuite))
}

type baseTestSuite struct {
	suite.Suite
	client *binance.Client
}

func newBaseTestSuite(client binance.RestClient) baseTestSuite {
	return baseTestSuite{
		client: binance.NewCustomClient(client),
	}
}

func (s *baseTestSuite) SetupSuite() {
	s.client = binance.NewClient("", "")
}

type clientTestSuite struct {
	baseTestSuite
}

func (s *clientTestSuite) TestTime() {
	_, e := s.client.Time()
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestPing() {
	s.Require().NoError(s.client.Ping())
}

func (s *clientTestSuite) TestExchangeInfo() {
	info, err := s.client.ExchangeInfo(nil)
	s.Require().NoError(err)
	s.Require().NotNil(info)
	s.Require().NotEmpty(info.Symbols)
}

func (s *clientTestSuite) TestExchangeInfoSymbol() {
	info, err := s.client.ExchangeInfo(&binance.ExchangeInfoReq{Symbol: "LTCBTC"})
	s.Require().NoError(err)
	s.Require().NotNil(info)
	s.Require().Len(info.Symbols, 1)
}

func (s *clientTestSuite) TestExchangeInfoSymbols() {
	info, err := s.client.ExchangeInfo(&binance.ExchangeInfoReq{Symbols: []string{"LTCBTC", "ETHBTC"}})
	s.Require().NoError(err)
	s.Require().NotNil(info)
	s.Require().Len(info.Symbols, 2)
}

func (s *clientTestSuite) TestExchangeInfoPermissions() {
	info, err := s.client.ExchangeInfo(&binance.ExchangeInfoReq{Permissions: []binance.PermissionType{binance.PermissionTypeSpot}})
	s.Require().NoError(err)
	s.Require().NotNil(info)
	s.Require().NotEmpty(info.Symbols)
}

func (s *clientTestSuite) TestDepth() {
	_, e := s.client.Depth(&binance.DepthReq{Symbol: "LTCBTC", Limit: 5})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestKlines() {
	resp, e := s.client.Klines(&binance.KlinesReq{Symbol: "LTCBTC", Interval: binance.KlineInterval1hour, Limit: 5})
	s.Require().NoError(e)
	s.Require().Len(resp, 5)
}

func (s *clientTestSuite) TestTrades() {
	_, e := s.client.Trades(&binance.TradeReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestAvgPrice() {
	_, e := s.client.AvgPrice(&binance.AvgPriceReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
}

func (s *clientTestSuite) TestPrices() {
	_, e := s.client.Prices(nil)
	s.Require().NoError(e)

	resp, e := s.client.Prices(&binance.TickerPricesReq{Symbols: []string{"LTCBTC", "ETHBTC"}})
	s.Require().NoError(e)
	s.Require().Len(resp, 2)
}

func (s *clientTestSuite) TestPrice() {
	_, e := s.client.Price(&binance.TickerPriceReq{Symbol: "LTCBTC"})
	s.Require().NoError(e)
}

type mockedClient struct {
	Response func(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error)
	window   int
}

func (m *mockedClient) UsedWeight() map[string]int64 {
	panic("not used")
}

func (m *mockedClient) OrderCount() map[string]int64 {
	panic("not used")
}

func (m *mockedClient) RetryAfter() int64 {
	panic("not used")
}

func (m *mockedClient) SetWindow(w int) {
	m.window = w
}

func (m *mockedClient) Do(method, endpoint string, data interface{}, sign, stream bool) ([]byte, error) {
	return m.Response(method, endpoint, data, sign, stream)
}

type mockedTestSuite struct {
	suite.Suite
	client *binance.Client
	mock   *mockedClient
}

func (s *mockedTestSuite) SetupSuite() {
	s.mock = &mockedClient{}
	s.client = binance.NewCustomClient(s.mock)
}

func (s *mockedTestSuite) SetupTest() {
	s.mock.Response = nil
}
