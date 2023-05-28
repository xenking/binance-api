package ws_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/xenking/binance-api/ws"
)

func TestWSClient(t *testing.T) {
	suite.Run(t, new(tickerTestSuite))
	suite.Run(t, new(accountTestSuite))
}

type baseTestSuite struct {
	suite.Suite
	ws *ws.Client
}

func (s *baseTestSuite) SetupTest() {
	s.ws = ws.NewClient()
}
