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
	OrderStatusNew      OrderStatus = "NEW"
	OrderStatusPartial  OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled   OrderStatus = "FILLED"
	OrderStatusCanceled OrderStatus = "CANCELED"
	OrderStatusPending  OrderStatus = "PENDING_CANCEL"
	OrderStatusRejected OrderStatus = "REJECTED"
	OrderStatusExpired  OrderStatus = "EXPIRED"
	OrderStatusReplaced OrderStatus = "REPLACED"
	OrderStatusTrade    OrderStatus = "TRADE"
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

type OrderReq struct {
	Symbol           string        `url:"symbol"`
	Side             OrderSide     `url:"side"`
	Type             OrderType     `url:"type"`
	Quantity         string        `url:"quantity,omitempty"`
	QuoteQuantity    string        `url:"quoteOrderQty,omitempty"`
	Price            string        `url:"price,omitempty"`
	TimeInForce      TimeInForce   `url:"timeInForce"`
	NewClientOrderId string        `url:"newClientOrderId,omitempty"`
	StopPrice        string        `url:"stopPrice,omitempty"`
	IcebergQty       string        `url:"icebergQty,omitempty"`
	OrderRespType    OrderRespType `url:"newOrderRespType,omitempty"`
}

type OrderRespAck struct {
	Symbol            string `json:"symbol"`
	OrderID           int    `json:"orderId"`
	OrigClientOrderID string `json:"origClientOrderId"`
	TransactTime      uint64 `json:"transactTime"`
}

type OrderRespResult struct {
	Symbol              string      `json:"symbol"`
	OrderID             int         `json:"orderId"`
	OrderListID         int         `json:"orderListId"`
	ClientOrderID       string      `json:"clientOrderId"`
	TransactTime        int64       `json:"transactTime"`
	Price               string      `json:"price"`
	OrigQty             string      `json:"origQty"`
	ExecutedQty         string      `json:"executedQty"`
	CummulativeQuoteQty string      `json:"cummulativeQuoteQty"`
	Status              OrderStatus `json:"status"`
	TimeInForce         string      `json:"timeInForce"`
	Type                OrderType   `json:"type"`
	Side                OrderSide   `json:"side"`
}

type OrderRespFull struct {
	Symbol              string              `json:"symbol"`
	OrderID             int                 `json:"orderId"`
	OrderListID         int                 `json:"orderListId"`
	ClientOrderID       string              `json:"clientOrderId"`
	TransactTime        int64               `json:"transactTime"`
	Price               string              `json:"price"`
	OrigQty             string              `json:"origQty"`
	ExecutedQty         string              `json:"executedQty"`
	CummulativeQuoteQty string              `json:"cummulativeQuoteQty"`
	Status              OrderStatus         `json:"status"`
	TimeInForce         string              `json:"timeInForce"`
	Type                OrderType           `json:"type"`
	Side                OrderSide           `json:"side"`
	Fills               []OrderRespFullFill `json:"fills"`
}

type OrderRespFullFill struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
}

type ServerTime struct {
	ServerTime uint64 `json:"serverTime"`
}

type KlineInterval string

const (
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

// DepthReq are used to specify symbol to retrieve order book for
type DepthReq struct {
	Symbol string `url:"symbol"` // Symbol is the symbol to fetch data for
	Limit  int    `url:"limit"`  // Limit is the number of order book items to retrieve. Max 100
}

// DepthElem represents a specific order in the order book
type DepthElem struct {
	Quantity decimal.Decimal `json:"quantity"`
	Price    decimal.Decimal `json:"price"`
}

// UnmarshalJSON unmarshal the given depth raw data and converts to depth struct
func (b *DepthElem) UnmarshalJSON(data []byte) error {
	if b == nil {
		return ErrNilUnmarshal
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
		return ErrInvalidJson
	}
	var err error
	b.Price, err = decimal.NewFromString(b2s(data[2:price]))
	b.Quantity, err = decimal.NewFromString(b2s(data[price+3 : qty]))
	return err
}

type Depth struct {
	LastUpdateID int         `json:"lastUpdateId"`
	Bids         []DepthElem `json:"bids"`
	Asks         []DepthElem `json:"asks"`
}

// TradeReq are used to specify symbol to get recent trades
type TradeReq struct {
	Symbol string `url:"symbol"` // Symbol is the symbol to fetch data for
	Limit  int    `url:"limit"`  // Limit is the maximal number of elements to receive. Default 500 Max 1000
}

type Trade struct {
	ID           int    `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

type KlinesReq struct {
	Symbol    string        `url:"symbol"`   // Symbol is the symbol to fetch data for
	Interval  KlineInterval `url:"interval"` // Interval is the interval for each kline/candlestick
	Limit     int           `url:"limit"`    // Limit is the maximal number of elements to receive. Max 500
	StartTime uint64        `url:"startTime,omitempty"`
	EndTime   uint64        `url:"endTime,omitempty"`
}

type Klines struct {
	OpenTime                 uint64
	OpenPrice                decimal.Decimal
	High                     decimal.Decimal
	Low                      decimal.Decimal
	ClosePrice               decimal.Decimal
	Volume                   decimal.Decimal
	CloseTime                uint64
	QuoteAssetVolume         decimal.Decimal
	Trades                   int
	TakerBuyBaseAssetVolume  decimal.Decimal
	TakerBuyQuoteAssetVolume decimal.Decimal
}

var (
	klinesQuote = []byte(`"`)
	klinesDelim = []byte(`"`)
)

// UnmarshalJSON unmarshal the given depth raw data and converts to depth struct
func (b *Klines) UnmarshalJSON(data []byte) error {
	if b == nil {
		return ErrNilUnmarshal
	}
	if len(data) == 0 {
		return nil
	}
	s := bytes.Replace(data, klinesQuote, nil, -1)
	s = bytes.Trim(s, "[]")
	tokens := bytes.Split(s, klinesDelim)
	if len(tokens) < 11 {
		return ErrInvalidJson
	}
	var err error
	u, err := fasthttp.ParseUint(tokens[0])
	if err != nil {
		return ErrInvalidJson
	}
	b.OpenTime = uint64(u)
	b.OpenPrice, err = decimal.NewFromString(b2s(tokens[1]))
	b.High, err = decimal.NewFromString(b2s(tokens[2]))
	b.Low, err = decimal.NewFromString(b2s(tokens[3]))
	b.ClosePrice, err = decimal.NewFromString(b2s(tokens[4]))
	b.Volume, err = decimal.NewFromString(b2s(tokens[5]))
	u, err = fasthttp.ParseUint(tokens[6])
	if err != nil {
		return ErrInvalidJson
	}
	b.CloseTime = uint64(u)
	b.QuoteAssetVolume, err = decimal.NewFromString(b2s(tokens[7]))
	b.Trades, err = fasthttp.ParseUint(tokens[9])
	b.TakerBuyBaseAssetVolume, err = decimal.NewFromString(b2s(tokens[9]))
	b.TakerBuyQuoteAssetVolume, err = decimal.NewFromString(b2s(tokens[10]))
	return err
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

type BookTicker struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

// TickerReq represents the request for a specified ticker
type TickerReq struct {
	Symbol string `url:"symbol"`
}

// TickerStats is the stats for a specific symbol
type TickerStats struct {
	PriceChange           string `json:"priceChange"`
	PriceChangePercentage string `json:"priceChangePercent"`
	WeightedAvgPrice      string `json:"weightedAvgPrice"`
	PrevClosePrice        string `json:"prevClosePrice"`
	LastPrice             string `json:"lastPrice"`
	BidPrice              string `json:"bidPrice"`
	AskPrice              string `json:"askPrice"`
	OpenPrice             string `json:"openPrice"`
	HighPrice             string `json:"highPrice"` // HighPrice is 24hr high price
	LowPrice              string `json:"lowPrice"`  // LowPrice is 24hr low price
	Volume                string `json:"volume"`
	QuoteVolume           string `json:"quoteVolume"`
	OpenTime              uint64 `json:"openTime"`
	CloseTime             uint64 `json:"closeTime"`
	FirstID               int    `json:"firstId"`
	LastID                int    `json:"lastId"`
	Count                 int    `json:"count"`
}

type TickerPriceReq struct {
	Symbol string `url:"symbol"`
}

type SymbolPrice struct {
	Symbol string
	Price  string
}

// QueryOrderReq represents the request for querying an order
// Remark: Either OrderID or OrigOrderiD must be set
type QueryOrderReq struct {
	Symbol            string `url:"symbol"`
	OrderID           int    `url:"orderId,omitempty"`
	OrigClientOrderId string `url:"origClientOrderId,omitempty"`
}

type QueryOrder struct {
	Symbol              string      `json:"symbol"`
	OrderID             int         `json:"orderId"`
	OrderListID         int         `json:"orderListId"`
	ClientOrderID       string      `json:"clientOrderId"`
	Price               string      `json:"price"`
	OrigQty             string      `json:"origQty"`
	ExecutedQty         string      `json:"executedQty"`
	CummulativeQuoteQty string      `json:"cummulativeQuoteQty"`
	Status              OrderStatus `json:"status"`
	TimeInForce         TimeInForce `json:"timeInForce"`
	Type                OrderType   `json:"type"`
	Side                OrderSide   `json:"side"`
	StopPrice           string      `json:"stopPrice"`
	IcebergQty          string      `json:"IcebergQty"`
	Time                uint64      `json:"time"`
	UpdateTime          uint64      `json:"updateTime"`
	OrigQuoteOrderQty   string      `json:"origQuoteOrderQty"`
}

// Remark: Either OrderID or OrigOrderID must be set
type CancelOrderReq struct {
	Symbol            string `url:"symbol"`
	OrderID           int    `url:"orderId,omitempty"`
	OrigClientOrderId string `url:"origClientOrderId,omitempty"`
	NewClientOrderId  string `url:"newClientOrderId,omitempty"`
}

type CancelOrder struct {
	Symbol              string      `json:"symbol"`
	OrigClientOrderID   string      `json:"origClientOrderId"`
	OrderID             int         `json:"orderId"`
	OrderListID         int         `json:"orderListId"`
	ClientOrderID       string      `json:"clientOrderId"`
	Price               string      `json:"price"`
	OrigQty             string      `json:"origQty"`
	ExecutedQty         string      `json:"executedQty"`
	CummulativeQuoteQty string      `json:"cummulativeQuoteQty"`
	Status              OrderStatus `json:"status"`
	TimeInForce         string      `json:"timeInForce"`
	Type                OrderType   `json:"type"`
	Side                OrderSide   `json:"side"`
}

type OpenOrdersReq struct {
	Symbol string `url:"symbol"`
}

type CancelOpenOrdersReq struct {
	Symbol string `url:"symbol"`
}

// AllOrdersReq represents the request used for querying orders of the given symbol
// Remark: If orderId is set, it will get orders >= that orderId. Otherwise most recent orders are returned
type AllOrdersReq struct {
	Symbol  string `url:"symbol"`  // Symbol is the symbol to fetch orders for
	OrderID int    `url:"orderId"` // OrderID, if set, will filter all recent orders newer from the given ID
	Limit   int    `url:"limit"`   // Limit is the maximal number of elements to receive. Max 500
}

type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type AccountInfo struct {
	MakerCommission  int        `json:"makerCommission"`
	TakerCommission  int        `json:"takerCommission"`
	BuyerCommission  int        `json:"buyerCommission"`
	SellerCommission int        `json:"sellerCommission"`
	CanTrade         bool       `json:"canTrade"`
	CanWithdraw      bool       `json:"canWithdraw"`
	CanDeposit       bool       `json:"canDeposit"`
	Balances         []*Balance `json:"balances"`
}

type AccountTradesReq struct {
	Symbol string `url:"symbol"`
	Limit  int    `url:"limit"`  // Limit is the maximal number of elements to receive. Max 500
	FromID int    `url:"fromId"` // FromID is trade ID to fetch from. Default gets most recent trades
}

type AccountTrades struct {
	ID              int    `json:"id"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            uint64 `json:"time"`
	Buyer           bool   `json:"isBuyer"`
	Maker           bool   `json:"isMaker"`
	BestMatch       bool   `json:"isBestMatch"`
}

type DatastreamReq struct {
	ListenKey string `json:"listenKey" url:"listenKey"`
}

type AggregatedTradeReq struct {
	Symbol    string `url:"symbol"` // Symbol is the symbol to fetch data for
	FromID    int    `url:"fromId"` // FromID to get aggregate trades from INCLUSIVE.
	Limit     int    `url:"limit"`  // Limit is the maximal number of elements to receive. Max 500
	StartTime uint64 `url:"startTime,omitempty"`
	EndTime   uint64 `url:"endTime,omitempty"`
}

type AggregatedTrade struct {
	TradeID      int    `json:"a"` // TradeID is the aggregate trade ID
	Price        string `json:"p"` // Price is the trade price
	Quantity     string `json:"q"` // Quantity is the trade quantity
	FirstTradeID int    `json:"f"`
	LastTradeID  int    `json:"l"`
	Time         uint64 `json:"T"`
	Maker        bool   `json:"m"` // Maker indicates if the buyer is the maker
	BestMatch    bool   `json:"M"` // BestMatch indicates if the trade was at the best price match
}

type ExchangeInfo struct {
	Symbols []SymbolInfo
}

type SymbolInfo struct {
	Symbol                   string             `json:"symbol"`
	Status                   SymbolStatus       `json:"status"`
	BaseAsset                string             `json:"baseAsset"`
	BaseAssetPrecision       int                `json:"baseAssetPrecision"`
	QuoteAsset               string             `json:"quoteAsset"`
	QuoteAssetPrecision      int                `json:"quoteAssetPrecision"`
	BaseCommissionPrecision  int                `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision int                `json:"quoteCommissionPrecision"`
	OrderTypes               []OrderType        `json:"orderTypes"`
	IcebergAllowed           bool               `json:"icebergAllowed"`
	OCOAllowed               bool               `json:"ocoAllowed"`
	QuoteQtyAllowed          bool               `json:"quoteOrderQtyMarketAllowed"`
	Filters                  []SymbolInfoFilter `json:"filters"`
}

type SymbolStatus string

const (
	SymbolStatusTrading SymbolStatus = "TRADING"
)

type FilterType string

const (
	FilterTypePrice       FilterType = "PRICE_FILTER"
	FilterTypeLotSize     FilterType = "LOT_SIZE"
	FilterTypeMinNotional FilterType = "MIN_NOTIONAL"
)

type SymbolInfoFilter struct {
	Type FilterType `json:"filterType"`

	// PRICE_FILTER parameters
	MinPrice string `json:"minPrice"`
	MaxPrice string `json:"maxPrice"`
	TickSize string `json:"tickSize"`

	// LOT_SIZE parameters
	MinQty   string `json:"minQty"`
	MaxQty   string `json:"maxQty"`
	StepSize string `json:"stepSize"`

	// MIN_NOTIONAL parameters
	MinNotional string `json:"minNotional"`
}
