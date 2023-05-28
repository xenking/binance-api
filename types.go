package binance

import (
	"bytes"

	"github.com/valyala/fasthttp"
	"github.com/xenking/decimal"
)

// OrderType represents the order type
type OrderType string

const (
	OrderTypeMarket          OrderType = "MARKET"
	OrderTypeLimit           OrderType = "LIMIT"
	OrderTypeStopLoss        OrderType = "STOP_LOSS"
	OrderTypeStopLossLimit   OrderType = "STOP_LOSS_LIMIT"
	OrderTypeTakeProfit      OrderType = "TAKE_PROFIT"
	OrderTypeTakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"
	OrderTypeLimitMaker      OrderType = "LIMIT_MAKER"
)

type OrderStatus string

const (
	OrderStatusNew            OrderStatus = "NEW"
	OrderStatusPartial        OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled         OrderStatus = "FILLED"
	OrderStatusCanceled       OrderStatus = "CANCELED"
	OrderStatusPending        OrderStatus = "PENDING_CANCEL"
	OrderStatusRejected       OrderStatus = "REJECTED"
	OrderStatusExpired        OrderStatus = "EXPIRED"
	OrderStatusExpiredInMatch OrderStatus = "EXPIRED_IN_MATCH"
)

type OrderFailure string

const (
	OrderFailureNone              OrderFailure = "NONE"
	OrderFailureUnknownInstrument OrderFailure = "UNKNOWN_INSTRUMENT"
	OrderFailureMarketClosed      OrderFailure = "MARKET_CLOSED"
	OrderFailurePriceExceed       OrderFailure = "PRICE_QTY_EXCEED_HARD_LIMITS"
	OrderFailureUnknownOrder      OrderFailure = "UNKNOWN_ORDER"
	OrderFailureDuplicate         OrderFailure = "DUPLICATE_ORDER"
	OrderFailureUnknownAccount    OrderFailure = "UNKNOWN_ACCOUNT"
	OrderFailureInsufficientFunds OrderFailure = "INSUFFICIENT_BALANCE"
	OrderFailureAccountInaactive  OrderFailure = "ACCOUNT_INACTIVE"
	OrderFailureAccountSettle     OrderFailure = "ACCOUNT_CANNOT_SETTLE"
)

type OrderSide string

const (
	OrderSideBuy  OrderSide = "BUY"
	OrderSideSell OrderSide = "SELL"
)

type TimeInForce string

const (
	TimeInForceGTC TimeInForce = "GTC" // Good Till Cancel
	TimeInForceIOC TimeInForce = "IOC" // Immediate or Cancel
	TimeInForceFOK TimeInForce = "FOK" // Fill or Kill
)

type OrderRespType string

const (
	OrderRespTypeAsk    = "ASK"
	OrderRespTypeResult = "RESULT"
	OrderRespTypeFull   = "FULL"
)

const MinStrategyType = 1000000

type OrderReq struct {
	Symbol                  string                  `url:"symbol"`
	Side                    OrderSide               `url:"side"`
	Type                    OrderType               `url:"type"`
	TimeInForce             TimeInForce             `url:"timeInForce,omitempty"`
	Quantity                string                  `url:"quantity,omitempty"`
	QuoteQuantity           string                  `url:"quoteOrderQty,omitempty"`
	Price                   string                  `url:"price,omitempty"`
	NewClientOrderID        string                  `url:"newClientOrderId,omitempty"`
	StrategyID              int                     `url:"strategyId,omitempty"`
	StrategyType            int                     `url:"strategyType,omitempty"` // Should be more than 1000000
	StopPrice               string                  `url:"stopPrice,omitempty"`
	TrailingDelta           int64                   `url:"trailingDelta,omitempty"` // Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
	IcebergQty              string                  `url:"icebergQty,omitempty"`
	OrderRespType           OrderRespType           `url:"newOrderRespType,omitempty"`
	SelfTradePreventionMode SelfTradePreventionMode `url:"selfTradePreventionMode,omitempty"`
}

type OrderRespAck struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	OrderListID   int64  `json:"orderListId"`
	ClientOrderID string `json:"clientOrderId"`
	TransactTime  int64  `json:"transactTime"`
}

type OrderRespResult struct {
	Symbol                  string                  `json:"symbol"`
	OrderID                 int64                   `json:"orderId"`
	OrderListID             int                     `json:"orderListId"`
	ClientOrderID           string                  `json:"clientOrderId"`
	TransactTime            int64                   `json:"transactTime"`
	Price                   string                  `json:"price"`
	OrigQty                 string                  `json:"origQty"`
	ExecutedQty             string                  `json:"executedQty"`
	CummulativeQuoteQty     string                  `json:"cummulativeQuoteQty"`
	Status                  OrderStatus             `json:"status"`
	TimeInForce             string                  `json:"timeInForce"`
	Type                    OrderType               `json:"type"`
	Side                    OrderSide               `json:"side"`
	WorkingTime             int64                   `json:"workingTime"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode,omitempty"`
	IcebergQty              string                  `json:"IcebergQty"`
	StopPrice               string                  `json:"stopPrice"`
	PreventedMatchID        int64                   `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string                  `json:"preventedQuantity,omitempty"`
	StrategyID              int                     `json:"strategyId,omitempty"`
	StrategyType            int                     `json:"strategyType,omitempty"`
	TrailingDelta           int                     `json:"trailingDelta,omitempty"`
	TrailingTime            int64                   `json:"trailingTime,omitempty"`
}

type OrderRespFull struct {
	Symbol                  string                  `json:"symbol"`
	OrderID                 int64                   `json:"orderId"`
	OrderListID             int64                   `json:"orderListId"`
	ClientOrderID           string                  `json:"clientOrderId"`
	TransactTime            int64                   `json:"transactTime"`
	Price                   string                  `json:"price"`
	OrigQty                 string                  `json:"origQty"`
	ExecutedQty             string                  `json:"executedQty"`
	CummulativeQuoteQty     string                  `json:"cummulativeQuoteQty"`
	Status                  OrderStatus             `json:"status"`
	TimeInForce             string                  `json:"timeInForce"`
	Type                    OrderType               `json:"type"`
	Side                    OrderSide               `json:"side"`
	WorkingTime             int64                   `json:"workingTime"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode,omitempty"`
	IcebergQty              string                  `json:"IcebergQty"`
	StopPrice               string                  `json:"stopPrice"`
	PreventedMatchID        int64                   `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string                  `json:"preventedQuantity,omitempty"`
	StrategyID              int                     `json:"strategyId,omitempty"`
	StrategyType            int                     `json:"strategyType,omitempty"`
	TrailingDelta           int                     `json:"trailingDelta,omitempty"`
	TrailingTime            int64                   `json:"trailingTime,omitempty"`
	Fills                   []OrderRespFullFill     `json:"fills"`
}

type OrderRespFullFill struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	TradeID         int64  `json:"tradeId"`
}

type ServerTime struct {
	ServerTime int64 `json:"serverTime"`
}

type KlineInterval string

const (
	KlineInterval1sec   KlineInterval = "1s"
	KlineInterval1min   KlineInterval = "1m"
	KlineInterval3min   KlineInterval = "3m"
	KlineInterval5min   KlineInterval = "5m"
	KlineInterval15min  KlineInterval = "15m"
	KlineInterval30min  KlineInterval = "30m"
	KlineInterval1hour  KlineInterval = "1h"
	KlineInterval2hour  KlineInterval = "2h"
	KlineInterval4hour  KlineInterval = "4h"
	KlineInterval6hour  KlineInterval = "6h"
	KlineInterval8hour  KlineInterval = "8h"
	KlineInterval12hour KlineInterval = "12h"
	KlineInterval1day   KlineInterval = "1d"
	KlineInterval3day   KlineInterval = "3d"
	KlineInterval1week  KlineInterval = "1w"
	KlineInterval1month KlineInterval = "1M"
)

const (
	DefaultDepthLimit = 100
	MaxDepthLimit     = 5000
)

// DepthReq are used to specify symbol to retrieve order book for
type DepthReq struct {
	Symbol string `url:"symbol"`          // Symbol is the symbol to fetch data for
	Limit  int    `url:"limit,omitempty"` // Limit is the number of order book items to retrieve. Default 100; Max 5000
}

// DepthElem represents a specific order in the order book
type DepthElem struct {
	Price    decimal.Decimal `json:"price"`
	Quantity decimal.Decimal `json:"quantity"`
}

// UnmarshalJSON unmarshal the given depth raw data and converts to depth struct
func (b *DepthElem) UnmarshalJSON(data []byte) error {
	if b == nil {
		return ErrEmptyJSONResponse
	}

	if len(data) <= 4 {
		return nil
	}
	qty, price := 3, 0
	next := false
	for qty < len(data)-1 {
		if data[qty] == '"' {
			if next {
				break
			}
			next = true
			price = qty
			qty += 3

			continue
		}
		qty++
	}
	if price < 3 || qty < 4 || !next {
		return ErrInvalidJSON
	}
	var err error
	b.Price, err = decimal.NewFromString(b2s(data[2:price]))
	if err != nil {
		return err
	}
	b.Quantity, err = decimal.NewFromString(b2s(data[price+3 : qty]))

	return err
}

type Depth struct {
	LastUpdateID int64       `json:"lastUpdateId"`
	Bids         []DepthElem `json:"bids"`
	Asks         []DepthElem `json:"asks"`
}

const (
	DefaultTradesLimit = 500
	MaxTradesLimit     = 1000
)

// TradeReq are used to specify symbol to get recent trades
type TradeReq struct {
	Symbol string `url:"symbol"`          // Symbol is the symbol to fetch data for
	Limit  int    `url:"limit,omitempty"` // Limit is the maximal number of elements to receive. Default 500; Max 1000
}

// HistoricalTradeReq are used to specify symbol to get older trades
type HistoricalTradeReq struct {
	Symbol string `url:"symbol"`           // Symbol is the symbol to fetch data for
	Limit  int    `url:"limit,omitempty"`  // Limit is the maximal number of elements to receive. Default 500; Max 1000
	FromID int64  `url:"fromId,omitempty"` // FromID is trade ID to fetch from. Default gets most recent trades
}

type Trade struct {
	ID           int64  `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

const (
	DefaultKlinesLimit = 500
	MaxKlinesLimit     = 1000
)

type KlinesReq struct {
	Symbol    string        `url:"symbol"`          // Symbol is the symbol to fetch data for
	Interval  KlineInterval `url:"interval"`        // Interval is the interval for each kline/candlestick
	Limit     int           `url:"limit,omitempty"` // Limit is the maximal number of elements to receive. Default 500; Max 1000
	StartTime int64         `url:"startTime,omitempty"`
	EndTime   int64         `url:"endTime,omitempty"`
}

type Kline struct {
	OpenPrice                decimal.Decimal
	HighPrice                decimal.Decimal
	LowPrice                 decimal.Decimal
	ClosePrice               decimal.Decimal
	Volume                   decimal.Decimal
	QuoteAssetVolume         decimal.Decimal
	TakerBuyBaseAssetVolume  decimal.Decimal
	TakerBuyQuoteAssetVolume decimal.Decimal
	OpenTime                 int64
	CloseTime                int64
	Trades                   int
}

var (
	klinesQuote = []byte(`"`)
	klinesDelim = []byte(`,`)
)

// UnmarshalJSON unmarshal the given depth raw data and converts to depth struct
func (b *Kline) UnmarshalJSON(data []byte) error {
	if b == nil {
		return ErrEmptyJSONResponse
	}
	if len(data) == 0 {
		return nil
	}
	s := bytes.ReplaceAll(data, klinesQuote, nil)
	tokens := bytes.Split(s, klinesDelim)
	if len(tokens) < 11 {
		return ErrInvalidJSON
	}
	u, err := fasthttp.ParseUint(tokens[0][1:])
	if err != nil {
		return ErrInvalidJSON
	}
	b.OpenTime = int64(u)
	b.OpenPrice, err = decimal.NewFromString(b2s(tokens[1]))
	if err != nil {
		return ErrInvalidJSON
	}
	b.HighPrice, err = decimal.NewFromString(b2s(tokens[2]))
	if err != nil {
		return ErrInvalidJSON
	}
	b.LowPrice, err = decimal.NewFromString(b2s(tokens[3]))
	if err != nil {
		return ErrInvalidJSON
	}
	b.ClosePrice, err = decimal.NewFromString(b2s(tokens[4]))
	if err != nil {
		return ErrInvalidJSON
	}
	b.Volume, err = decimal.NewFromString(b2s(tokens[5]))
	if err != nil {
		return ErrInvalidJSON
	}
	u, err = fasthttp.ParseUint(tokens[6])
	if err != nil {
		return ErrInvalidJSON
	}
	b.CloseTime = int64(u)
	b.QuoteAssetVolume, err = decimal.NewFromString(b2s(tokens[7]))
	if err != nil {
		return ErrInvalidJSON
	}
	b.Trades, err = fasthttp.ParseUint(tokens[8])
	if err != nil {
		return ErrInvalidJSON
	}
	b.TakerBuyBaseAssetVolume, err = decimal.NewFromString(b2s(tokens[9]))
	if err != nil {
		return ErrInvalidJSON
	}
	b.TakerBuyQuoteAssetVolume, err = decimal.NewFromString(b2s(tokens[10]))
	if err != nil {
		return ErrInvalidJSON
	}

	return nil
}

type AvgPriceReq struct {
	Symbol string `url:"symbol"`
}

type AvgPrice struct {
	Mins  int    `json:"mins"`
	Price string `json:"price"`
}

type BookTickerReq struct {
	Symbol string `url:"symbol"`
}

type BookTickersReq struct {
	Symbols []string `url:"symbols,omitempty"`
}

type BookTicker struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

type TickerReq struct {
	Symbol     string           `url:"symbol"`
	WindowSize TickerWindowSize `url:"windowSize,omitempty"`
	RespType   TickerRespType   `url:"type,omitempty"`
}

type TickersReq struct {
	Symbols    []string         `url:"symbols,omitempty"`
	WindowSize TickerWindowSize `url:"windowSize,omitempty"`
	RespType   TickerRespType   `url:"type,omitempty"`
}

// TickerWindowSize used to compute statistics will be no more than 59999ms from the requested windowSize.
// Defaults to 1d if no parameter provided
// Supported windowSize values:
// 1m,2m....59m for minutes
// 1h, 2h....23h - for hours
// 1d...7d - for days
// Units cannot be combined (e.g. 1d2h is not allowed)
type TickerWindowSize string

func (s TickerWindowSize) IsValid() bool {
	if len(s) < 2 || len(s) > 3 {
		return false
	}

	switch s[(len(s) - 1)] {
	case 'm':
		if s[0] < '1' || s[0] > '9' {
			return false
		}
		if len(s) < 3 {
			break
		}
		if s[0] > '5' || s[0] < '1' {
			return false
		}
	case 'h':
		if s[0] < '1' || s[0] > '9' {
			return false
		}
		if len(s) < 3 {
			break
		}
		if s[0] == '2' && s[1] > '3' || (s[0] != '1' && s[0] != '2') {
			return false
		}
	case 'd':
		if s[0] < '1' || s[0] > '7' {
			return false
		}
	default:
		return false
	}

	return true
}

// Ticker24hReq represents the request for a specified ticker
type Ticker24hReq struct {
	Symbol   string         `url:"symbol"`
	RespType TickerRespType `url:"type,omitempty"`
}

// Tickers24hReq represents the request for a specified tickers
type Tickers24hReq struct {
	Symbols  []string       `url:"symbols,omitempty"`
	RespType TickerRespType `url:"type,omitempty"`
}

type TickerRespType string

const (
	TickerRespTypeFull TickerRespType = "FULL"
	TickerRespTypeMini TickerRespType = "MINI"
)

// TickerStatFull is the stats for a specific symbol
type TickerStatFull struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"` // HighPrice is 24hr high price
	LowPrice           string `json:"lowPrice"`  // LowPrice is 24hr low price
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int64  `json:"firstId"`
	LastID             int64  `json:"lastId"`
	Count              int    `json:"count"`
}

type TickerStatMini struct {
	Symbol      string `json:"symbol"`
	OpenPrice   string `json:"openPrice"`
	HighPrice   string `json:"highPrice"` // HighPrice is 24hr high price
	LowPrice    string `json:"lowPrice"`  // LowPrice is 24hr low price
	LastPrice   string `json:"lastPrice"`
	Volume      string `json:"volume"`
	QuoteVolume string `json:"quoteVolume"`
	OpenTime    int64  `json:"openTime"`
	CloseTime   int64  `json:"closeTime"`
	FirstID     int64  `json:"firstId"`
	LastID      int64  `json:"lastId"`
	Count       int    `json:"count"`
}

type TickerStat struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"` // HighPrice is 24hr high price
	LowPrice           string `json:"lowPrice"`  // LowPrice is 24hr low price
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int64  `json:"firstId"`
	LastID             int64  `json:"lastId"`
	Count              int    `json:"count"`
}

type TickerPriceReq struct {
	Symbol string `url:"symbol"`
}

type TickerPricesReq struct {
	Symbols []string `url:"symbols,omitempty"`
}

type SymbolPrice struct {
	Symbol string
	Price  string
}

// QueryOrderReq represents the request for querying an order
// Remark: Either OrderID or OrigOrderiD must be set
type QueryOrderReq struct {
	Symbol            string `url:"symbol"`
	OrderID           int64  `url:"orderId,omitempty"`
	OrigClientOrderID string `url:"origClientOrderId,omitempty"`
}

type QueryOrder struct {
	Symbol                  string                  `json:"symbol"`
	OrderID                 int64                   `json:"orderId"`
	OrderListID             int64                   `json:"orderListId"`
	ClientOrderID           string                  `json:"clientOrderId"`
	Price                   string                  `json:"price"`
	OrigQty                 string                  `json:"origQty"`
	ExecutedQty             string                  `json:"executedQty"`
	CummulativeQuoteQty     string                  `json:"cummulativeQuoteQty"`
	Status                  OrderStatus             `json:"status"`
	TimeInForce             TimeInForce             `json:"timeInForce"`
	Type                    OrderType               `json:"type"`
	Side                    OrderSide               `json:"side"`
	Time                    int64                   `json:"time"`
	UpdateTime              int64                   `json:"updateTime"`
	IsWorking               bool                    `json:"isWorking"`
	WorkingTime             int64                   `json:"workingTime"`
	OrigQuoteOrderQty       string                  `json:"origQuoteOrderQty"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode,omitempty"`
	IcebergQty              string                  `json:"IcebergQty"`
	StopPrice               string                  `json:"stopPrice"`
	PreventedMatchID        int64                   `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string                  `json:"preventedQuantity,omitempty"`
	StrategyID              int                     `json:"strategyId,omitempty"`
	StrategyType            int                     `json:"strategyType,omitempty"`
	TrailingDelta           int                     `json:"trailingDelta,omitempty"`
	TrailingTime            int64                   `json:"trailingTime,omitempty"`
}

// Remark: Either OrderID or OrigOrderID must be set
type CancelOrderReq struct {
	Symbol             string            `url:"symbol"`
	OrderID            int64             `url:"orderId,omitempty"`
	OrigClientOrderID  string            `url:"origClientOrderId,omitempty"`
	NewClientOrderID   string            `url:"newClientOrderId,omitempty"`
	CancelRestrictions CancelRestriction `url:"cancelRestrictions,omitempty"`
}

type CancelRestriction string

const (
	CancelRestrictionOnlyNew             CancelRestriction = "ONLY_NEW"
	CancelRestrictionOnlyPartiallyFilled CancelRestriction = "ONLY_PARTIALLY_FILLED"
)

type CancelOrder struct {
	Symbol                  string                  `json:"symbol"`
	OrigClientOrderID       string                  `json:"origClientOrderId"`
	OrderID                 int64                   `json:"orderId"`
	OrderListID             int64                   `json:"orderListId"`
	ClientOrderID           string                  `json:"clientOrderId"`
	Price                   string                  `json:"price"`
	OrigQty                 string                  `json:"origQty"`
	ExecutedQty             string                  `json:"executedQty"`
	CummulativeQuoteQty     string                  `json:"cummulativeQuoteQty"`
	Status                  OrderStatus             `json:"status"`
	TimeInForce             TimeInForce             `json:"timeInForce"`
	Type                    OrderType               `json:"type"`
	Side                    OrderSide               `json:"side"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode,omitempty"`
	IcebergQty              string                  `json:"IcebergQty"`
	StopPrice               string                  `json:"stopPrice"`
	PreventedMatchID        int64                   `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string                  `json:"preventedQuantity,omitempty"`
	StrategyID              int                     `json:"strategyId,omitempty"`
	StrategyType            int                     `json:"strategyType,omitempty"`
	TrailingDelta           int                     `json:"trailingDelta,omitempty"`
	TrailingTime            int64                   `json:"trailingTime,omitempty"`
}

type CancelOCOOrderType struct {
	ContingencyType ContingencyType `json:"contingencyType"`
}

type CancelOCOOrder struct{}

type CancelReplaceResult string

const (
	CancelReplaceResultSuccess      CancelReplaceResult = "SUCCESS"
	CancelReplaceResultFailure      CancelReplaceResult = "FAILURE"
	CancelReplaceResultNotAttempted CancelReplaceResult = "NOT_ATTEMPTED"
)

type CancelReplaceMode string

const (
	CancelReplaceModeStopOnFailure CancelReplaceMode = "STOP_ON_FAILURE"
	CancelReplaceModeAllowFailure  CancelReplaceMode = "ALLOW_FAILURE"
)

// Note: Either CancelOrderID or CancelOrigClientOrderID must be set
type CancelReplaceOrderReq struct {
	OrderReq
	CancelRestrictions      CancelRestriction `url:"cancelRestrictions,omitempty"`
	CancelReplaceMode       CancelReplaceMode `url:"cancelReplaceMode"`
	CancelOrderID           int64             `url:"cancelOrderId,omitempty"`
	CancelOrigClientOrderID string            `url:"cancelOrigClientOrderId,omitempty"`
	CancelNewClientOrderID  string            `url:"cancelNewClientOrderId,omitempty"`
}

type CancelReplaceOrder struct {
	CancelResponse   CancelOrder         `json:"cancelResponse"`
	NewOrderResponse *OrderRespFull      `json:"newOrderResponse,omitempty"`
	CancelResult     CancelReplaceResult `json:"cancelResult"`
	NewOrderResult   CancelReplaceResult `json:"newOrderResult"`
}

type OpenOrdersReq struct {
	Symbol string `url:"symbol,omitempty"`
}

type CancelOpenOrdersReq struct {
	Symbol string `url:"symbol"`
}

const (
	DefaultOrderLimit = 500
	MaxOrderLimit     = 1000
)

// AllOrdersReq represents the request used for querying orders of the given symbol
// Remark: If orderId is set, it will get orders >= that orderId. Otherwise most recent orders are returned
type AllOrdersReq struct {
	Symbol    string `url:"symbol"`            // Symbol is the symbol to fetch orders for
	OrderID   int64  `url:"orderId,omitempty"` // OrderID, if set, will filter all recent orders newer from the given ID
	Limit     int    `url:"limit,omitempty"`   // Limit is the maximal number of elements to receive. Default 500; Max 1000
	StartTime int64  `url:"startTime,omitempty"`
	EndTime   int64  `url:"endTime,omitempty"`
}

type OCOReq struct {
	Symbol                  string                  `url:"symbol"`
	Side                    OrderSide               `url:"side"`
	Type                    OrderType               `url:"type"`
	Quantity                string                  `url:"quantity"`
	Price                   string                  `url:"price"`
	StopPrice               string                  `url:"stopPrice"`
	ListClientOrderID       string                  `url:"listClientOrderId,omitempty"`
	LimitClientOrderID      string                  `url:"limitClientOrderId,omitempty"`
	LimitStrategyID         int                     `url:"limitStrategyId,omitempty"` // Should be more than 1000000
	LimitStrategyType       int                     `url:"limitStrategyType,omitempty"`
	LimitIcebergQty         string                  `url:"limitIcebergQty,omitempty"`
	TrailingDelta           int64                   `url:"trailingDelta,omitempty"`
	StopClientOrderID       string                  `url:"stopClientOrderId,omitempty"`
	StopStrategyID          int                     `url:"stopStrategyId,omitempty"` // Should be more than 1000000
	StopStrategyType        int                     `url:"stopStrategyType,omitempty"`
	StopLimitPrice          string                  `url:"stopLimitPrice,omitempty"`
	StopIcebergQty          string                  `url:"stopIcebergQty,omitempty"`
	StopLimitTimeInForce    TimeInForce             `url:"stopLimitTimeInForce,omitempty"`
	OrderRespType           OrderRespType           `url:"newOrderRespType,omitempty"`
	SelfTradePreventionMode SelfTradePreventionMode `url:"selfTradePreventionMode,omitempty"`
}

type OCOOrder struct {
	Symbol            string               `json:"symbol"`
	OrderListID       int64                `json:"orderListId"`
	ContingencyType   ContingencyType      `json:"contingencyType"`
	ListStatusType    OCOStatus            `json:"listStatusType"`
	ListOrderStatus   OCOOrderStatus       `json:"listOrderStatus"`
	ListClientOrderID string               `json:"listClientOrderId"`
	TransactionTime   int64                `json:"transactionTime"`
	Orders            []OCOOrderStatusResp `json:"orders"`
	OrderReports      []OrderRespFull      `json:"orderReports"`
}

type OCOOrderStatusResp struct {
	Symbol        string `json:"symbol"`
	OrderID       int    `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}

// Remark: Either OrderListID or ListClientOrderID must be set
type CancelOCOReq struct {
	Symbol            string `url:"symbol"`
	OrderListID       int64  `url:"orderListId,omitempty"`
	ListClientOrderID string `url:"listClientOrderId,omitempty"`
	NewClientOrderID  string `url:"newClientOrderId,omitempty"`
}

// QueryOCOReq represents the request for querying an OCO order
// Remark: Either OrderID or OrigOrderiD must be set
type QueryOCOReq struct {
	OrderListID       int64  `url:"orderListId,omitempty"`
	ListClientOrderID string `url:"listClientOrderId,omitempty"`
}

// AllOCOReq represents the request used for querying OCO orders
type AllOCOReq struct {
	FromID    int64 `url:"fromId,omitempty"` // If supplied, neither startTime or endTime can be provided
	Limit     int   `url:"limit,omitempty"`  // Limit is the maximal number of elements to receive. Default 500; Max 1000
	StartTime int64 `url:"startTime,omitempty"`
	EndTime   int64 `url:"endTime,omitempty"`
}

type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type PermissionType string

const (
	PermissionTypeSpot      PermissionType = "SPOT"
	PermissionTypeMargin    PermissionType = "MARGIN"
	PermissionTypeLeveraged PermissionType = "LEVERAGED"
	PermissionTypeGRP2      PermissionType = "TRD_GRP_002"
	PermissionTypeGRP3      PermissionType = "TRD_GRP_003"
	PermissionTypeGRP4      PermissionType = "TRD_GRP_004"
	PermissionTypeGRP5      PermissionType = "TRD_GRP_005"
	PermissionTypeGRP6      PermissionType = "TRD_GRP_006"
	PermissionTypeGRP7      PermissionType = "TRD_GRP_007"
)

type AccountInfo struct {
	MakerCommission  int              `json:"makerCommission"`
	TakerCommission  int              `json:"takerCommission"`
	BuyerCommission  int              `json:"buyerCommission"`
	SellerCommission int              `json:"sellerCommission"`
	CanTrade         bool             `json:"canTrade"`
	CanWithdraw      bool             `json:"canWithdraw"`
	CanDeposit       bool             `json:"canDeposit"`
	AccountType      PermissionType   `json:"accountType"`
	Balances         []*Balance       `json:"balances"`
	Permissions      []PermissionType `json:"permissions"`
}

const MaxAccountTradesLimit = 500

type AccountTradesReq struct {
	Symbol    string `url:"symbol"`
	OrderID   string `url:"orderId,omitempty"` // OrderID can only be used in combination with symbol
	Limit     int    `url:"limit,omitempty"`   // Limit is the maximal number of elements to receive. Default 500; Max 1000
	FromID    int64  `url:"fromId,omitempty"`  // FromID is trade ID to fetch from. Default gets most recent trades
	StartTime int64  `url:"startTime,omitempty"`
	EndTime   int64  `url:"endTime,omitempty"`
}

type AccountTrade struct {
	ID              int64  `json:"id"`
	OrderID         int64  `json:"orderId"`
	OrderListID     int64  `json:"orderListId"`
	Symbol          string `json:"symbol"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	Buyer           bool   `json:"isBuyer"`
	Maker           bool   `json:"isMaker"`
	BestMatch       bool   `json:"isBestMatch"`
}

type DataStream struct {
	ListenKey string `json:"listenKey" url:"listenKey"`
}

type MyPreventedMatchesReq struct {
	Symbol           string `url:"symbol"`
	PreventedMatchID int64  `url:"preventedMatchId,omitempty"`
	FromID           int64  `url:"fromId,omitempty"` // FromID is trade ID to fetch from. Default gets most recent trades
	StartTime        int64  `url:"startTime,omitempty"`
	EndTime          int64  `url:"endTime,omitempty"`
}

type AggregatedTradeReq struct {
	Symbol    string `url:"symbol"`              // Symbol is the symbol to fetch data for
	FromID    int64  `url:"fromId"`              // FromID to get aggregate trades from INCLUSIVE.
	StartTime int64  `url:"startTime,omitempty"` // StartTime timestamp in ms to get aggregate trades from INCLUSIVE.
	EndTime   int64  `url:"endTime,omitempty"`   // EndTime timestamp in ms to get aggregate trades until INCLUSIVE.
	Limit     int    `url:"limit,omitempty"`     // Limit is the maximal number of elements to receive. Default 500; Max 1000
}

type AggregatedTrade struct {
	TradeID      int64  `json:"a"` // TradeID is the aggregate trade ID
	Price        string `json:"p"` // Price is the trade price
	Quantity     string `json:"q"` // Quantity is the trade quantity
	FirstTradeID int64  `json:"f"` // FirstTradeID is the first trade ID
	LastTradeID  int64  `json:"l"` // LastTradeID is the last trade ID
	Timestamp    int64  `json:"T"` // Timestamp is the trade timestamp
	Maker        bool   `json:"m"` // Maker indicates if the buyer is the maker
	BestMatch    bool   `json:"M"` // BestMatch indicates if the trade was at the best price match
}

type ExchangeInfoReq struct {
	Symbol      string           `url:"symbol,omitempty"`
	Symbols     []string         `url:"symbols,omitempty"`
	Permissions []PermissionType `url:"permissions,omitempty"`
}

type ExchangeInfo struct {
	Timezone        string            `json:"timezone"`
	ServerTime      int64             `json:"serverTime"`
	RateLimits      []*RateLimit      `json:"rateLimits"`
	ExchangeFilters []*ExchangeFilter `json:"exchangeFilters"`
	Symbols         []*SymbolInfo     `json:"symbols"`
}

type RateLimitType string

const (
	RateLimitTypeRequestWeight RateLimitType = "REQUEST_WEIGHT"
	RateLimitTypeOrders        RateLimitType = "ORDERS"
	RateLimitTypeRawRequests   RateLimitType = "RAW_REQUESTS"
)

type RateLimitInterval string

const (
	RateLimitIntervalSecond RateLimitInterval = "SECOND"
	RateLimitIntervalHour   RateLimitInterval = "HOUR"
	RateLimitIntervalMinute RateLimitInterval = "MINUTE"
	RateLimitIntervalDay    RateLimitInterval = "DAY"
)

var RateLimitIntervalLetter = map[byte]RateLimitInterval{
	's': RateLimitIntervalSecond,
	'S': RateLimitIntervalSecond,
	'h': RateLimitIntervalHour,
	'H': RateLimitIntervalHour,
	'm': RateLimitIntervalMinute,
	'M': RateLimitIntervalMinute,
	'd': RateLimitIntervalDay,
	'D': RateLimitIntervalDay,
}

type RateLimit struct {
	Type        RateLimitType     `json:"rateLimitType"`
	Interval    RateLimitInterval `json:"interval"`
	IntervalNum int               `json:"intervalNum"`
	Limit       int               `json:"limit"`
	Count       int               `json:"count,omitempty"`
}

type ExchangeFilterType string

const (
	ExchangeFilterTypeMaxNumOrders  ExchangeFilterType = "EXCHANGE_MAX_NUM_ORDERS"
	ExchangeFilterTypeMaxAlgoOrders ExchangeFilterType = "EXCHANGE_MAX_ALGO_ORDERS"
)

type ExchangeFilter struct {
	Type ExchangeFilterType `json:"filterType"`

	// EXCHANGE_MAX_NUM_ORDERS parameters
	MaxNumOrders int `json:"maxNumOrders"`

	// EXCHANGE_MAX_ALGO_ORDERS parameters
	MaxNumAlgoOrders int `json:"maxNumAlgoOrders"`
}

type SymbolInfo struct {
	Symbol                          string                    `json:"symbol"`
	Status                          SymbolStatus              `json:"status"`
	BaseAsset                       string                    `json:"baseAsset"`
	BaseAssetPrecision              int                       `json:"baseAssetPrecision"`
	QuoteAsset                      string                    `json:"quoteAsset"`
	QuotePrecision                  int                       `json:"quotePrecision"`
	QuoteAssetPrecision             int                       `json:"quoteAssetPrecision"`
	BaseCommissionPrecision         int                       `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision        int                       `json:"quoteCommissionPrecision"`
	OrderTypes                      []OrderType               `json:"orderTypes"`
	IcebergAllowed                  bool                      `json:"icebergAllowed"`
	OCOAllowed                      bool                      `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed      bool                      `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop               bool                      `json:"allowTrailingStop"`
	CancelReplaceAllowed            bool                      `json:"cancelReplaceAllowed"`
	IsSpotTradingAllowed            bool                      `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed          bool                      `json:"isMarginTradingAllowed"`
	Filters                         []SymbolInfoFilter        `json:"filters"`
	Permissions                     []PermissionType          `json:"permissions"`
	DefaultSelfTradePreventionMode  SelfTradePreventionMode   `json:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []SelfTradePreventionMode `json:"allowedSelfTradePreventionModes"`
}

type SymbolStatus string

const (
	SymbolStatusPreTrading   SymbolStatus = "PRE_TRADING"
	SymbolStatusTrading      SymbolStatus = "TRADING"
	SymbolStatusPostTrading  SymbolStatus = "POST_TRADING"
	SymbolStatusEndOfDay     SymbolStatus = "END_OF_DAY"
	SymbolStatusHalt         SymbolStatus = "HALT"
	SymbolStatusAuctionMatch SymbolStatus = "AUCTION_MATCH"
	SymbolStatusBreak        SymbolStatus = "BREAK"
)

type FilterType string

const (
	FilterTypePrice               FilterType = "PRICE_FILTER"
	FilterTypePercentPrice        FilterType = "PERCENT_PRICE"
	FilterTypeLotSize             FilterType = "LOT_SIZE"
	FilterTypeMinNotional         FilterType = "MIN_NOTIONAL"
	FilterTypeIcebergParts        FilterType = "ICEBERG_PARTS"
	FilterTypeMarketLotSize       FilterType = "MARKET_LOT_SIZE"
	FilterTypeMaxNumOrders        FilterType = "MAX_NUM_ORDERS"
	FilterTypeMaxNumAlgoOrders    FilterType = "MAX_NUM_ALGO_ORDERS"
	FilterTypeMaxNumIcebergOrders FilterType = "MAX_NUM_ICEBERG_ORDERS"
	FilterTypeMaxPosition         FilterType = "MAX_POSITION"
)

type SymbolInfoFilter struct {
	Type FilterType `json:"filterType"`

	// PRICE_FILTER parameters
	MinPrice string `json:"minPrice"`
	MaxPrice string `json:"maxPrice"`
	TickSize string `json:"tickSize"`

	// PERCENT_PRICE parameters
	MultiplierUp   string `json:"multiplierUp"`
	MultiplierDown string `json:"multiplierDown"`
	AvgPriceMins   int    `json:"avgPriceMins"`

	// LOT_SIZE or MARKET_LOT_SIZE parameters
	MinQty   string `json:"minQty"`
	MaxQty   string `json:"maxQty"`
	StepSize string `json:"stepSize"`

	// MIN_NOTIONAL parameter
	MinNotional   string `json:"minNotional"`
	ApplyToMarket bool   `json:"applyToMarket"`

	// ICEBERG_PARTS parameter
	IcebergLimit int `json:"limit"`

	// TRAILING_DELTA parameter
	MinTrailingAboveDelta int `json:"minTrailingAboveDelta"`
	MaxTrailingAboveDelta int `json:"maxTrailingAboveDelta"`
	MinTrailingBelowDelta int `json:"minTrailingBelowDelta"`
	MaxTrailingBelowDelta int `json:"maxTrailingBelowDelta"`

	// MAX_NUM_ORDERS parameter
	MaxNumOrders int `json:"maxNumOrders"`

	// MAX_NUM_ALGO_ORDERS parameter
	MaxNumAlgoOrders int `json:"maxNumAlgoOrders"`

	// MAX_NUM_ICEBERG_ORDERS parameter
	MaxNumIcebergOrders int `json:"maxNumIcebergOrders"`

	// MAX_POSITION parameter
	MaxPosition string `json:"maxPosition"`
}

type OCOStatus string

const (
	OCOStatusResponse    OCOStatus = "RESPONSE"
	OCOStatusExecStarted OCOStatus = "EXEC_STARTED"
	OCOStatusAllDone     OCOStatus = "ALL_DONE"
)

type OCOOrderStatus string

const (
	OCOOrderStatusExecuting OCOOrderStatus = "EXECUTING"
	OCOOrderStatusAllDone   OCOOrderStatus = "ALL_DONE"
	OCOOrderStatusReject    OCOOrderStatus = "REJECT"
)

type ContingencyType string

const (
	ContingencyTypeOCO ContingencyType = "OCO"
)

type SelfTradePreventionMode string

const (
	SelfTradePreventionModeNone        SelfTradePreventionMode = "NONE"
	SelfTradePreventionModeExpireTaker SelfTradePreventionMode = "EXPIRE_TAKER"
	SelfTradePreventionModeExpireMaker SelfTradePreventionMode = "EXPIRE_MAKER"
	SelfTradePreventionModeExpireBoth  SelfTradePreventionMode = "EXPIRE_BOTH"
)
