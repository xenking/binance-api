package binance

import (
	"github.com/segmentio/encoding/json"
)

type ValidationError struct {
	msg string
}

func (e ValidationError) Error() string {
	return e.msg
}

var (
	ErrNilRequest          = ValidationError{"nil request"}
	ErrEmptySymbol         = ValidationError{"symbol is not set"}
	ErrEmptyQuantity       = ValidationError{"quantity is not set"}
	ErrEmptyPrice          = ValidationError{"price is not set"}
	ErrEmptyStopPrice      = ValidationError{"stop price is not set"}
	ErrEmptySide           = ValidationError{"order side is not set"}
	ErrEmptyOrderID        = ValidationError{"order id is not set"}
	ErrMinStrategyType     = ValidationError{"strategy type can not be lower than 1000000"}
	ErrEmptyJSONResponse   = ValidationError{"empty json response"}
	ErrInvalidJSON         = ValidationError{"invalid json"}
	ErrInvalidTickerWindow = ValidationError{"invalid ticker window"}
	ErrInvalidOrderType    = ValidationError{"invalid order type"}
	// ErrIncorrectAccountEventType represents error when event type can't before determined
	ErrIncorrectAccountEventType = ValidationError{"incorrect account event type"}
)

type APIError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Error return error code and message
func (e *APIError) Error() string {
	bb, _ := json.Marshal(e)

	return b2s(bb)
}
