package binance

import (
	"github.com/pkg/errors"
	"github.com/segmentio/encoding/json"
)

var (
	ErrNilRequest   = errors.New("request is nil")
	ErrEmptySymbol  = errors.New("symbol are missing")
	ErrEmptyOrderID = errors.New("order id must be set")
	ErrEmptyLimit   = errors.New("empty price or quantity")
	ErrEmptyMarket  = errors.New("quantity or quote quantity expected")
	ErrNilUnmarshal = errors.New("UnmarshalJSON on nil pointer")
	ErrInvalidJSON  = errors.New("invalid json")
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
