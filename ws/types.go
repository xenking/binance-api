package ws

import (
	"bytes"

	"github.com/go-faster/errors"

	"github.com/xenking/binance-api"
)

// UpdateType represents type of account update event
type UpdateType string

const (
	// UpdateTypeUnknown default for unknown type
	UpdateTypeUnknown     UpdateType = "unknown"
	UpdateTypeDepth       UpdateType = "depthUpdate"
	UpdateTypeIndivTicker UpdateType = "24hrTicker"
	UpdateTypeKline       UpdateType = "kline"
	UpdateTypeAggTrades   UpdateType = "aggTrade"
	UpdateTypeTrades      UpdateType = "trade"

	UpdateTypeOutboundAccountInfo     UpdateType = "outboundAccountInfo"
	UpdateTypeOutboundAccountPosition UpdateType = "outboundAccountPosition"
	UpdateTypeOrderReport             UpdateType = "executionReport"
	UpdateTypeBalanceUpdate           UpdateType = "balanceUpdate"
	UpdateTypeOCOReport               UpdateType = "listStatus"
)

// FrequencyType is a interval for Depth update
type FrequencyType string

const (
	// Frequency1000ms is default frequency
	Frequency1000ms = "@1000ms"

	// Frequency100ms for fastest updates
	Frequency100ms = "@100ms"
)

// IndivTickerUpdate represents incoming ticker websocket feed
type IndivTickerUpdate struct {
	EventType     UpdateType `json:"e"` // EventType represents the update type
	Time          uint64     `json:"E"` // Time represents the event time
	Symbol        string     `json:"s"` // Symbol represents the symbol related to the update
	Price         string     `json:"p"` // Price is the order price
	PricePercent  string     `json:"P"` // Price percent change
	WeightedPrice string     `json:"w"` // Weighted average price
	FirstTrade    string     `json:"x"` // First trade(F)-1 price (first trade before the 24hr rolling window)
	LastPrice     string     `json:"c"` // Last price
	LastQty       string     `json:"Q"` // Last quantity
	BestBidPrice  string     `json:"b"` // Best bid price
	BestBidQty    string     `json:"B"` // Best bid quantity
	BestAskPrice  string     `json:"a"` // Best ask price
	BestAskQty    string     `json:"A"` // Best ask quantity
	OpenPrice     string     `json:"o"` // Open price
	HighPrice     string     `json:"h"` // High price
	LowPrice      string     `json:"l"` // Low price
	VolumeBase    string     `json:"v"` // Total traded base asset volume
	VolumeQuote   string     `json:"q"` // Total traded quote asset volume
	StatisticOT   uint64     `json:"O"` // Statistics open time
	StatisticsCT  uint64     `json:"C"` // Statistics close time
	FirstTradeID  uint64     `json:"F"` // First trade ID
	LastTradeID   uint64     `json:"L"` // Last trade ID
	TotalTrades   int        `json:"n"` // Total number of trades
}

// AllMarketTickerUpdate represents incoming ticker websocket feed for all tickers
type AllMarketTickerUpdate []IndivTickerUpdate

// IndivBookTickerUpdate represents incoming book ticker websocket feed
type IndivBookTickerUpdate struct {
	UpdateID int    `json:"u"` // UpdateID to sync up with updateID in /ws/v3/depth
	Symbol   string `json:"s"` // Symbol represents the symbol related to the update
	BidPrice string `json:"b"` // BidPrice
	BidQty   string `json:"B"` // BidQty
	AskPrice string `json:"a"` // AskPrice
	AskQty   string `json:"A"` // AskQty
}

// AllBookTickerUpdate represents incoming ticker websocket feed for all book tickers
type AllBookTickerUpdate IndivBookTickerUpdate

// IndivMiniTickerUpdate represents incoming mini-ticker websocket feed
type IndivMiniTickerUpdate struct {
	EventType   UpdateType `json:"e"` // EventType represents the update type
	Time        uint64     `json:"E"` // Time represents the event time
	Symbol      string     `json:"s"` // Symbol represents the symbol related to the update
	LastPrice   string     `json:"c"` // Last price
	OpenPrice   string     `json:"o"` // Open price
	HighPrice   string     `json:"h"` // High price
	LowPrice    string     `json:"l"` // Low price
	VolumeBase  string     `json:"v"` // Total traded base asset volume
	VolumeQuote string     `json:"q"` // Total traded quote asset volume
}

// AllMarketMiniTickerUpdate represents incoming mini-ticker websocket feed for all tickers
type AllMarketMiniTickerUpdate []IndivMiniTickerUpdate

// DepthUpdate represents the incoming messages for depth websocket updates
type DepthUpdate struct {
	EventType     UpdateType          `json:"e"` // EventType represents the update type
	Time          uint64              `json:"E"` // Time represents the event time
	Symbol        string              `json:"s"` // Symbol represents the symbol related to the update
	FirstUpdateID uint64              `json:"U"` // FirstTradeID in event
	FinalUpdateID uint64              `json:"u"` // FirstTradeID in event to sync in /ws/v3/depth
	Bids          []binance.DepthElem `json:"b"` // Bids is a list of bids for symbol
	Asks          []binance.DepthElem `json:"a"` // Asks is a list of asks for symbol
}

// DepthLevelUpdate represents the incoming messages for depth level websocket updates
type DepthLevelUpdate struct {
	LastUpdateID uint64              `json:"lastUpdateId"` // EventType represents the update type
	Bids         []binance.DepthElem `json:"bids"`         // Bids is a list of bids for symbol
	Asks         []binance.DepthElem `json:"asks"`         // Asks is a list of asks for symbol
}

// KlinesUpdate represents the incoming messages for klines websocket updates
type KlinesUpdate struct {
	EventType UpdateType `json:"e"` // EventType represents the update type
	Time      uint64     `json:"E"` // Time represents the event time
	Symbol    string     `json:"s"` // Symbol represents the symbol related to the update
	Kline     struct {
		StartTime    uint64                `json:"t"` // StartTime is the start time of this bar
		EndTime      uint64                `json:"T"` // EndTime is the end time of this bar
		Symbol       string                `json:"s"` // Symbol represents the symbol related to this kline
		Interval     binance.KlineInterval `json:"i"` // Interval is the kline interval
		FirstTradeID uint64                `json:"f"` // FirstTradeID is the first trade ID
		LastTradeID  uint64                `json:"L"` // LastTradeID is the first trade ID

		OpenPrice            string `json:"o"` // OpenPrice represents the open price for this bar
		ClosePrice           string `json:"c"` // ClosePrice represents the close price for this bar
		High                 string `json:"h"` // High represents the highest price for this bar
		Low                  string `json:"l"` // Low represents the lowest price for this bar
		Volume               string `json:"v"` // Volume is the trades volume for this bar
		Trades               int    `json:"n"` // Trades is the number of conducted trades
		Final                bool   `json:"x"` // Final indicates whether this bar is final or yet may receive updates
		VolumeQuote          string `json:"q"` // VolumeQuote indicates the quote volume for the symbol
		VolumeActiveBuy      string `json:"V"` // VolumeActiveBuy represents the volume of active buy
		VolumeQuoteActiveBuy string `json:"Q"` // VolumeQuoteActiveBuy represents the quote volume of active buy
	} `json:"k"` // Kline is the kline update
}

// AggTradeUpdate represents the incoming messages for aggregated trades websocket updates
type AggTradeUpdate struct {
	EventType             UpdateType `json:"e"` // EventType represents the update type
	Time                  uint64     `json:"E"` // Time represents the event time
	Symbol                string     `json:"s"` // Symbol represents the symbol related to the update
	TradeID               uint64     `json:"a"` // TradeID is the aggregated trade ID
	Price                 string     `json:"p"` // Price is the trade price
	Quantity              string     `json:"q"` // Quantity is the trade quantity
	FirstBreakDownTradeID uint64     `json:"f"` // FirstBreakDownTradeID is the first breakdown trade ID
	LastBreakDownTradeID  uint64     `json:"l"` // LastBreakDownTradeID is the last breakdown trade ID
	TradeTime             uint64     `json:"T"` // Time is the trade time
	Maker                 bool       `json:"m"` // Maker indicates whether buyer is a maker
}

// ErrIncorrectEventType represents error when event type can't before determined
var ErrIncorrectEventType = errors.New("cant unmarshal event type")

// EventTypeUpdate represents only incoming event type
type EventTypeUpdate struct {
	EventType UpdateType `json:"e"` // EventType represents the update type
}

// UnmarshalJSON need to getting partial json data
func (e *EventTypeUpdate) UnmarshalJSON(buf []byte) error {
	start := bytes.IndexByte(buf, ':')
	end := bytes.IndexByte(buf, ',')
	if end < start+2 || end-start < 3 {
		e.EventType = UpdateTypeUnknown

		return ErrIncorrectEventType
	}
	e.EventType = UpdateType(b2s(buf[start+2 : end-1]))

	return nil
}

// TradeUpdate represents the incoming messages for trades websocket updates
type TradeUpdate struct {
	EventType UpdateType `json:"e"` // EventType represents the update type
	Time      uint64     `json:"E"` // Time represents the event time
	Symbol    string     `json:"s"` // Symbol represents the symbol related to the update
	TradeID   uint64     `json:"t"` // TradeID is the aggregated trade ID
	Price     string     `json:"p"` // Price is the trade price
	Quantity  string     `json:"q"` // Quantity is the trade quantity
	BuyerID   int        `json:"b"` // BuyerID is the buyer trade ID
	SellerID  int        `json:"a"` // SellerID is the seller trade ID
	TradeTime uint64     `json:"T"` // Time is the trade time
	Maker     bool       `json:"m"` // Maker indicates whether buyer is a maker
}

// AccountInfoUpdate represents the incoming messages for account info websocket updates
type AccountInfoUpdate struct {
	EventType        UpdateType `json:"e"` // EventType represents the update type
	Time             uint64     `json:"E"` // Time represents the event time
	MakerCommission  int        `json:"m"` // MakerCommission is the maker commission for the account
	TakerCommission  int        `json:"t"` // TakerCommission is the taker commission for the account
	BuyerCommission  int        `json:"b"` // BuyerCommission is the buyer commission for the account
	SellerCommission int        `json:"s"` // SellerCommission is the seller commission for the account
	CanTrade         bool       `json:"T"`
	CanWithdraw      bool       `json:"W"`
	CanDeposit       bool       `json:"D"`
	Balances         []*struct {
		Asset  string `json:"a"`
		Free   string `json:"f"`
		Locked string `json:"l"`
	} `json:"B"`
}

// AccountUpdate represents the incoming messages for account websocket updates
type AccountUpdate struct {
	EventType  UpdateType `json:"e"` // EventType represents the update type
	Time       uint64     `json:"E"` // Time represents the event time
	LastUpdate uint64     `json:"u"` // LastUpdate represents last account update

	Balances []*struct {
		Asset  string `json:"a"`
		Free   string `json:"f"`
		Locked string `json:"l"`
	} `json:"B"`
}

// BalanceUpdate represents the incoming message for account balances websocket updates
type BalanceUpdate struct {
	EventType    UpdateType `json:"e"` // EventType represents the update type
	Time         uint64     `json:"E"` // Time represents the event time
	Asset        string     `json:"a"` // Asset
	BalanceDelta string     `json:"d"` // Balance Delta
	ClearTime    uint64     `json:"T"` // Clear Time
}

// OrderUpdate represents the incoming messages for account orders websocket updates
type OrderUpdate struct {
	EventType           UpdateType           `json:"e"` // EventType represents the update type
	Time                uint64               `json:"E"` // Time represents the event time
	Symbol              string               `json:"s"` // Symbol represents the symbol related to the update
	NewClientOrderID    string               `json:"c"` // NewClientOrderID is the new client order ID
	Side                binance.OrderSide    `json:"S"` // Side is the order side
	OrderType           binance.OrderType    `json:"o"` // OrderType represents the order type
	TimeInForce         binance.TimeInForce  `json:"f"` // TimeInForce represents the order TIF type
	OrigQty             string               `json:"q"` // OrigQty represents the order original quantity
	Price               string               `json:"p"` // Price is the order price
	StopPrice           string               `json:"P"`
	IcebergQty          string               `json:"F"`
	OrderListID         int64                `json:"g"`
	OrigClientOrderID   string               `json:"C"`
	ExecutionType       binance.OrderStatus  `json:"x"` // ExecutionType represents the execution type for the order
	Status              binance.OrderStatus  `json:"X"` // Status represents the order status for the order
	Error               binance.OrderFailure `json:"r"` // Error represents an order rejection reason
	OrderID             uint64               `json:"i"` // OrderID represents the order ID
	FilledQty           string               `json:"l"` // FilledQty represents the quantity of the last filled trade
	TotalFilledQty      string               `json:"z"` // TotalFilledQty is the accumulated quantity of filled trades on this order
	FilledPrice         string               `json:"L"` // FilledPrice is the price of last filled trade
	Commission          string               `json:"n"` // Commission is the commission for the trade
	CommissionAsset     string               `json:"N"` // CommissionAsset is the asset on which commission is taken
	TradeTime           uint64               `json:"T"` // TradeTime is the trade time
	TradeID             uint64               `json:"t"` // TradeID represents the trade ID
	Maker               bool                 `json:"m"` // Maker represents whether buyer is maker or not
	OrderCreatedTime    uint64               `json:"O"` // OrderTime represents the order time
	QuoteTotalFilledQty string               `json:"Z"`
	QuoteFilledQty      string               `json:"Y"`
	QuoteQty            string               `json:"Q"`
}
