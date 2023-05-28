package binance

// Endpoints with NONE security
const (
	EndpointAggTrades    = "/api/v3/aggTrades"
	EndpointAvgPrice     = "/api/v3/avgPrice"
	EndpointDepth        = "/api/v3/depth"
	EndpointExchangeInfo = "/api/v3/exchangeInfo"
	EndpointKlines       = "/api/v3/klines"
	EndpointPing         = "/api/v3/ping"
	EndpointTicker       = "/api/v3/ticker"
	EndpointTicker24h    = "/api/v3/ticker/24hr"
	EndpointTickerPrice  = "/api/v3/ticker/price"
	EndpointTickerBook   = "/api/v3/ticker/bookTicker"
	EndpointTime         = "/api/v3/time"
	EndpointTrades       = "/api/v3/trades"
	EndpointUIKlines     = "/api/v3/uiKlines"
)

// Endpoints with SIGNED security
const (
	EndpointAccountTrades      = "/api/v3/myTrades"
	EndpointRateLimit          = "/api/v3/rateLimit/order"
	EndpointMyPreventedMatches = "/api/v3/myPreventedMatches"
	EndpointHistoricalTrades   = "/api/v3/historicalTrades"
	EndpointOrder              = "/api/v3/order"
	EndpointOrderTest          = "/api/v3/order/test"
	EndpointCancelReplaceOrder = "/api/v3/order/cancelReplace"
	EndpointOrdersAll          = "/api/v3/allOrders"
	EndpointOpenOrders         = "/api/v3/openOrders"
	EndpointOCOOrder           = "/api/v3/order/oco"
	EndpointOCOOrders          = "/api/v3/orderList"
	EndpointOCOOrdersAll       = "/api/v3/allOrderList"
	EndpointOpenOCOOrders      = "/api/v3/openOrderList"
)

const (
	EndpointAccount    = "/api/v3/account"
	EndpointDataStream = "/api/v3/userDataStream"
)
