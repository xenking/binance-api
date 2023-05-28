package ws_test

import (
	"context"
	"math/rand"
	"net"
	"net/http"
	"time"

	websocket "github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/segmentio/encoding/json"

	"github.com/xenking/binance-api"
	"github.com/xenking/binance-api/ws"
)

type mockedClient struct {
	Callback func(method, endpoint string, data interface{}, sign, stream bool) ([]byte, error)
}

func (m *mockedClient) Do(method, endpoint string, data interface{}, sign, stream bool) ([]byte, error) {
	return m.Callback(method, endpoint, data, sign, stream)
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

func (m *mockedClient) SetWindow(_ int) {
	panic("not used")
}

type accountTestSuite struct {
	baseTestSuite
	api          *binance.Client
	mock         *mockedClient
	listener     net.Listener
	listenerDone chan struct{}
	expected     chan interface{}
}

func (s *accountTestSuite) SetupSuite() {
	s.mock = &mockedClient{}
	s.api = binance.NewCustomClient(s.mock)
	s.ws = ws.NewCustomClient("ws://localhost:9844/", nil)

	var err error
	s.listener, err = net.Listen("tcp", ":9844")
	s.Require().NoError(err)

	s.listenerDone = make(chan struct{}, 1)
	go func() {
		http.Serve(s.listener, nil) //nolint:errcheck // don't care about error here
		s.listenerDone <- struct{}{}
	}()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := websocket.UpgradeHTTP(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		go func() {
			defer conn.Close()

			writer := wsutil.NewWriter(conn, websocket.StateServerSide, websocket.OpText)
			enc := json.NewEncoder(writer)
			for ex := range s.expected {
				marshallErr := enc.Encode(ex)
				s.Require().NoError(marshallErr)
				flushErr := writer.Flush()
				s.Require().NoError(flushErr)
			}
		}()
	})

	http.HandleFunc("/stream-key", handler)
}

func (s *accountTestSuite) TearDownSuite() {
	err := s.listener.Close()
	s.Require().NoError(err)

	select {
	case <-s.listenerDone:
	case <-time.After(time.Second * 5):
		s.Fail("timeout")
	}
}

func (s *accountTestSuite) SetupTest() {
	s.mock.Callback = func(method, endpoint string, data interface{}, sign, stream bool) ([]byte, error) {
		s.Require().Nil(data)
		return json.Marshal(&binance.DataStream{
			ListenKey: "stream-key",
		})
	}
	s.expected = make(chan interface{})
}

func (s *accountTestSuite) TestAccountInfo_Read() {
	expected := []interface{}{
		&ws.BalanceUpdateEvent{
			EventType:    ws.AccountUpdateEventTypeBalanceUpdate,
			Time:         int64(rand.Uint32()),
			Asset:        "BTC",
			BalanceDelta: "1",
		},
		&ws.AccountUpdateEvent{
			Balances: []ws.AccountBalance{
				{
					Asset:  "ETH",
					Free:   "1",
					Locked: "0.5",
				},
			},
			EventType:  ws.AccountUpdateEventTypeOutboundAccountPosition,
			Time:       int64(rand.Uint32()),
			LastUpdate: int64(rand.Uint32()),
		},
		&ws.OrderUpdateEvent{
			EventType:        ws.AccountUpdateEventTypeOrderReport,
			Symbol:           "ETHBTC",
			Side:             "BUY",
			OrderType:        "LIMIT",
			TimeInForce:      "GTC",
			OrigQty:          "1",
			Price:            "3400",
			Status:           "FILLED",
			FilledQty:        "1",
			TotalFilledQty:   "1",
			FilledPrice:      "3400",
			Commission:       "0.00001",
			CommissionAsset:  "BTC",
			Time:             int64(rand.Uint32()),
			TradeTime:        int64(rand.Uint32()),
			OrderCreatedTime: int64(rand.Uint32()),
			TradeID:          rand.Int63(),
			OrderID:          int64(rand.Uint32()),
		},
		&ws.OCOOrderUpdateEvent{
			EventType: ws.AccountUpdateEventTypeOCOReport,
			Orders: []ws.OCOOrderUpdateEventOrder{
				{
					Symbol:  "ETH",
					OrderID: int64(rand.Uint32()),
				},
				{
					Symbol:  "BTC",
					OrderID: int64(rand.Uint32()),
				},
			},
			Symbol:          "ETHBTC",
			ContingencyType: "OCO",
			OCOStatus:       "EXEC_STARTED",
			OrderStatus:     "EXECUTING",
			OCORejectReason: "NONE",
			TransactTime:    int64(rand.Uint32()),
			OrderListID:     rand.Int63(),
			Time:            int64(rand.Uint32()),
		},
	}

	key, err := s.api.DataStream()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	info, err := s.ws.AccountInfo(ctx, key)
	s.Require().NoError(err)

	for _, ex := range expected {
		s.expected <- ex
		_, actual, err := info.Read()
		s.Require().NoError(err)
		s.Require().EqualValues(ex, actual)
	}
	close(s.expected)

	s.mock.Callback = func(method, endpoint string, data interface{}, sign, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DataStream{}, data)
		return nil, nil
	}

	err = s.api.DataStreamClose(key)
	s.Require().NoError(err)
	err = info.Close()
	s.Require().NoError(err)
}

func (s *accountTestSuite) TestAccountInfo_BalancesStream() {
	expected := []interface{}{
		&ws.BalanceUpdateEvent{
			EventType:    ws.AccountUpdateEventTypeBalanceUpdate,
			Time:         int64(rand.Uint32()),
			Asset:        "BTC",
			BalanceDelta: "1",
		},
		&ws.BalanceUpdateEvent{
			EventType:    ws.AccountUpdateEventTypeBalanceUpdate,
			Time:         int64(rand.Uint32()),
			Asset:        "ETH",
			BalanceDelta: "1",
		},
		&ws.BalanceUpdateEvent{
			EventType:    ws.AccountUpdateEventTypeBalanceUpdate,
			Time:         int64(rand.Uint32()),
			Asset:        "BTC",
			BalanceDelta: "2",
		},
	}

	key, err := s.api.DataStream()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	info, err := s.ws.AccountInfo(ctx, key)
	s.Require().NoError(err)

	stream := info.BalancesStream()
	for _, ex := range expected {
		s.expected <- ex
		actual := <-stream
		s.Require().NoError(err)
		s.Require().EqualValues(ex, actual)
	}
	close(s.expected)

	s.mock.Callback = func(method, endpoint string, data interface{}, sign, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DataStream{}, data)
		return nil, nil
	}

	err = s.api.DataStreamClose(key)
	s.Require().NoError(err)
	err = info.Close()
	s.Require().NoError(err)
}

func (s *accountTestSuite) TestAccountInfo_AccountStream() {
	expected := []interface{}{
		&ws.AccountUpdateEvent{
			Balances: []ws.AccountBalance{
				{
					Asset:  "ETH",
					Free:   "1",
					Locked: "0.5",
				},
			},
			EventType:  ws.AccountUpdateEventTypeOutboundAccountPosition,
			Time:       int64(rand.Uint32()),
			LastUpdate: int64(rand.Uint32()),
		},
		&ws.AccountUpdateEvent{
			Balances: []ws.AccountBalance{
				{
					Asset:  "BTC",
					Free:   "1",
					Locked: "0.5",
				},
			},
			EventType:  ws.AccountUpdateEventTypeOutboundAccountPosition,
			Time:       int64(rand.Uint32()),
			LastUpdate: int64(rand.Uint32()),
		},
	}

	key, err := s.api.DataStream()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	info, err := s.ws.AccountInfo(ctx, key)
	s.Require().NoError(err)

	stream := info.AccountStream()
	for _, ex := range expected {
		s.expected <- ex
		actual := <-stream
		s.Require().NoError(err)
		s.Require().EqualValues(ex, actual)
	}

	s.mock.Callback = func(method, endpoint string, data interface{}, sign, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DataStream{}, data)
		return nil, nil
	}

	err = s.api.DataStreamClose(key)
	s.Require().NoError(err)
	err = info.Close()
	s.Require().NoError(err)
}

func (s *accountTestSuite) TestAccountInfo_OrdersStream() {
	expected := []interface{}{
		&ws.OrderUpdateEvent{
			EventType:        ws.AccountUpdateEventTypeOrderReport,
			Symbol:           "ETHBTC",
			Side:             "BUY",
			OrderType:        "LIMIT",
			TimeInForce:      "GTC",
			OrigQty:          "1",
			Price:            "3400",
			Status:           "FILLED",
			FilledQty:        "1",
			TotalFilledQty:   "1",
			FilledPrice:      "3400",
			Commission:       "0.00001",
			CommissionAsset:  "BTC",
			Time:             int64(rand.Uint32()),
			TradeTime:        int64(rand.Uint32()),
			OrderCreatedTime: int64(rand.Uint32()),
			TradeID:          rand.Int63(),
			OrderID:          int64(rand.Uint32()),
		},
		&ws.OrderUpdateEvent{
			EventType:        ws.AccountUpdateEventTypeOrderReport,
			Symbol:           "ETHBTC",
			Side:             "BUY",
			OrderType:        "LIMIT",
			TimeInForce:      "GTC",
			OrigQty:          "1",
			Price:            "3500",
			Status:           "FILLED",
			FilledQty:        "1",
			TotalFilledQty:   "1",
			FilledPrice:      "3500",
			Commission:       "0.00001",
			CommissionAsset:  "BTC",
			Time:             int64(rand.Uint32()),
			TradeTime:        int64(rand.Uint32()),
			OrderCreatedTime: int64(rand.Uint32()),
			TradeID:          rand.Int63(),
			OrderID:          int64(rand.Uint32()),
		},
		&ws.OrderUpdateEvent{
			EventType:        ws.AccountUpdateEventTypeOrderReport,
			Symbol:           "ETHBTC",
			Side:             "BUY",
			OrderType:        "LIMIT",
			TimeInForce:      "GTC",
			OrigQty:          "1",
			Price:            "3600",
			Status:           "FILLED",
			FilledQty:        "1",
			TotalFilledQty:   "1",
			FilledPrice:      "3600",
			Commission:       "0.00001",
			CommissionAsset:  "BTC",
			Time:             int64(rand.Uint32()),
			TradeTime:        int64(rand.Uint32()),
			OrderCreatedTime: int64(rand.Uint32()),
			TradeID:          rand.Int63(),
			OrderID:          int64(rand.Uint32()),
		},
	}

	key, err := s.api.DataStream()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	info, err := s.ws.AccountInfo(ctx, key)
	s.Require().NoError(err)

	stream := info.OrdersStream()
	for _, ex := range expected {
		s.expected <- ex
		actual := <-stream
		s.Require().NoError(err)
		s.Require().EqualValues(ex, actual)
	}
	close(s.expected)

	s.mock.Callback = func(method, endpoint string, data interface{}, sign, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DataStream{}, data)
		return nil, nil
	}

	err = s.api.DataStreamClose(key)
	s.Require().NoError(err)
	err = info.Close()
	s.Require().NoError(err)
}

func (s *accountTestSuite) TestAccountInfo_OCOOrdersStream() {
	expected := []interface{}{
		&ws.OCOOrderUpdateEvent{
			EventType: ws.AccountUpdateEventTypeOCOReport,
			Orders: []ws.OCOOrderUpdateEventOrder{
				{
					Symbol:  "ETH",
					OrderID: int64(rand.Uint32()),
				},
				{
					Symbol:  "BTC",
					OrderID: int64(rand.Uint32()),
				},
			},
			Symbol:          "ETHBTC",
			ContingencyType: "OCO",
			OCOStatus:       "EXEC_STARTED",
			OrderStatus:     "EXECUTING",
			OCORejectReason: "NONE",
			TransactTime:    int64(rand.Uint32()),
			OrderListID:     rand.Int63(),
			Time:            int64(rand.Uint32()),
		},
		&ws.OCOOrderUpdateEvent{
			EventType: ws.AccountUpdateEventTypeOCOReport,
			Orders: []ws.OCOOrderUpdateEventOrder{
				{
					Symbol:  "ETH",
					OrderID: int64(rand.Uint32()),
				},
				{
					Symbol:  "BTC",
					OrderID: int64(rand.Uint32()),
				},
			},
			Symbol:          "ETHBTC",
			ContingencyType: "OCO",
			OCOStatus:       "EXEC_STARTED",
			OrderStatus:     "EXECUTING",
			OCORejectReason: "NONE",
			TransactTime:    int64(rand.Uint32()),
			OrderListID:     rand.Int63(),
			Time:            int64(rand.Uint32()),
		},
	}

	key, err := s.api.DataStream()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	info, err := s.ws.AccountInfo(ctx, key)
	s.Require().NoError(err)

	stream := info.OCOOrdersStream()
	for _, ex := range expected {
		s.expected <- ex
		actual := <-stream
		s.Require().NoError(err)
		s.Require().EqualValues(ex, actual)
	}
	close(s.expected)

	s.mock.Callback = func(method, endpoint string, data interface{}, sign, stream bool) ([]byte, error) {
		s.Require().IsType(binance.DataStream{}, data)
		return nil, nil
	}

	err = s.api.DataStreamClose(key)
	s.Require().NoError(err)
	err = info.Close()
	s.Require().NoError(err)
}
