package binance

import (
	"strconv"
	"strings"
)

const (
	ErrNilRequest   = "req is nil"
	ErrEmptySymbol  = "symbol are missing"
	ErrEmptyOrderID = "order id must be set"
)

type APIError struct {
	Code    int
	Message []byte
}

// Error return error code and message
func (e APIError) Error() string {
	var b strings.Builder
	b.WriteString("Code:")
	b.WriteString(strconv.Itoa(e.Code))
	b.WriteString(" \nBody:")
	b.Write(e.Message)
	return b.String()
}
