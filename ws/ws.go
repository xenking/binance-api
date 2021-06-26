package ws

import (
	"github.com/segmentio/encoding/json"
	"github.com/valyala/bytebufferpool"
	"github.com/xenking/fastws"
)

type wsClient struct {
	conn *fastws.Conn
	err  error
}

func (w *wsClient) Close() error {
	return w.conn.Close()
}

// Depth is a wrapper for depth websocket
type Depth struct {
	wsClient
}

// Read reads a depth update message from depth websocket
func (d *Depth) Read() (*DepthUpdate, error) {
	buf := bytebufferpool.Get()
	_, data, err := d.conn.ReadMessage(buf.B)
	if err != nil {
		return nil, err
	}
	r := &DepthUpdate{}
	err = json.Unmarshal(data, r)
	bytebufferpool.Put(buf)
	return r, err
}

// Stream stream a depth update message from depth websocket to channel
func (d *Depth) Stream() <-chan DepthUpdate {
	updates := make(chan DepthUpdate)
	go func() {
		defer close(updates)
		var msg []byte
		var data []byte
		var err error
		for {
			_, data, err = d.conn.ReadMessage(msg[:0])
			if err != nil {
				d.err = err
				return
			}

			u := DepthUpdate{}
			err = json.Unmarshal(data, &u)
			if err != nil {
				d.err = err
				return
			}
			updates <- u
		}
	}()
	return updates
}

// DepthLevel is a wrapper for depth level websocket
type DepthLevel struct {
	wsClient
}

// Read reads a depth update message from depth level websocket
func (d *DepthLevel) Read() (*DepthLevelUpdate, error) {
	buf := bytebufferpool.Get()
	_, data, err := d.conn.ReadMessage(buf.B)
	if err != nil {
		return nil, err
	}
	r := &DepthLevelUpdate{}
	err = json.Unmarshal(data, r)
	bytebufferpool.Put(buf)
	return r, err
}

// Stream stream a depth update message from depth level websocket to channel
func (d *DepthLevel) Stream() <-chan DepthLevelUpdate {
	updates := make(chan DepthLevelUpdate)
	go func() {
		defer close(updates)
		var msg []byte
		var data []byte
		var err error
		for {
			_, data, err = d.conn.ReadMessage(msg[:0])
			if err != nil {
				d.err = err
				return
			}

			u := DepthLevelUpdate{}
			err = json.Unmarshal(data, &u)
			if err != nil {
				d.err = err
				return
			}
			updates <- u
		}
	}()
	return updates
}

// AllMarketTicker is a wrapper for all markets tickers websocket
type AllMarketTicker struct {
	wsClient
}

// Read reads a market update message from all markets ticker websocket
func (t *AllMarketTicker) Read() (*AllMarketTickerUpdate, error) {
	buf := bytebufferpool.Get()
	_, data, err := t.conn.ReadMessage(buf.B)
	if err != nil {
		return nil, err
	}
	r := &AllMarketTickerUpdate{}
	err = json.Unmarshal(data, r)
	bytebufferpool.Put(buf)
	return r, err
}

// Stream stream a market update message from all markets ticker websocket to channel
func (t *AllMarketTicker) Stream() <-chan AllMarketTickerUpdate {
	updates := make(chan AllMarketTickerUpdate)
	go func() {
		defer close(updates)
		var msg []byte
		var data []byte
		var err error
		for {
			_, data, err = t.conn.ReadMessage(msg[:0])
			if err != nil {
				t.err = err
				return
			}

			u := AllMarketTickerUpdate{}
			err = json.Unmarshal(data, &u)
			if err != nil {
				t.err = err
				return
			}
			updates <- u
		}
	}()
	return updates
}

// IndivTicker is a wrapper for an individual ticker websocket
type IndivTicker struct {
	wsClient
}

// Read reads a individual symbol update message from individual ticker websocket
func (t *IndivTicker) Read() (*IndivTickerUpdate, error) {
	buf := bytebufferpool.Get()
	_, data, err := t.conn.ReadMessage(buf.B)
	if err != nil {
		return nil, err
	}
	r := &IndivTickerUpdate{}
	err = json.Unmarshal(data, r)
	bytebufferpool.Put(buf)
	return r, err
}

// Stream stream a individual update message from individual ticker websocket to channel
func (t *IndivTicker) Stream() <-chan IndivTickerUpdate {
	updates := make(chan IndivTickerUpdate)
	go func() {
		defer close(updates)
		var msg []byte
		var data []byte
		var err error
		for {
			_, data, err = t.conn.ReadMessage(msg[:0])
			if err != nil {
				t.err = err
				return
			}

			u := IndivTickerUpdate{}
			err = json.Unmarshal(data, &u)
			if err != nil {
				t.err = err
				return
			}
			updates <- u
		}
	}()
	return updates
}

// IndivBookTicker is a wrapper for an individual book ticker websocket
type IndivBookTicker struct {
	wsClient
}

// Read reads a individual book symbol update message from individual book ticker websocket
func (t *IndivBookTicker) Read() (*IndivBookTickerUpdate, error) {
	buf := bytebufferpool.Get()
	_, data, err := t.conn.ReadMessage(buf.B)
	if err != nil {
		return nil, err
	}
	r := &IndivBookTickerUpdate{}
	err = json.Unmarshal(data, r)
	bytebufferpool.Put(buf)
	return r, err
}

// Stream stream a individual book symbol update message from individual book ticker websocket to channel
func (t *IndivBookTicker) Stream() <-chan IndivBookTickerUpdate {
	updates := make(chan IndivBookTickerUpdate)
	go func() {
		defer close(updates)
		var msg []byte
		var data []byte
		var err error
		for {
			_, data, err = t.conn.ReadMessage(msg[:0])
			if err != nil {
				t.err = err
				return
			}

			u := IndivBookTickerUpdate{}
			err = json.Unmarshal(data, &u)
			if err != nil {
				t.err = err
				return
			}
			updates <- u
		}
	}()
	return updates
}

// Klines is a wrapper for klines websocket
type Klines struct {
	wsClient
}

// Read reads a klines update message from klines websocket
func (k *Klines) Read() (*KlinesUpdate, error) {
	buf := bytebufferpool.Get()
	_, data, err := k.conn.ReadMessage(buf.B)
	if err != nil {
		return nil, err
	}
	r := &KlinesUpdate{}
	err = json.Unmarshal(data, r)
	bytebufferpool.Put(buf)
	return r, err
}

// Stream stream a klines update message from klines websocket to channel
func (k *Klines) Stream() <-chan KlinesUpdate {
	updates := make(chan KlinesUpdate)
	go func() {
		defer close(updates)
		var msg []byte
		var data []byte
		var err error
		for {
			_, data, err = k.conn.ReadMessage(msg[:0])
			if err != nil {
				k.err = err
				return
			}

			u := KlinesUpdate{}
			err = json.Unmarshal(data, &u)
			if err != nil {
				k.err = err
				return
			}
			updates <- u
		}
	}()
	return updates
}

// AggTrades is a wrapper for trades websocket
type AggTrades struct {
	wsClient
}

// Read reads a trades update message from aggregated trades websocket
func (t *AggTrades) Read() (*AggTradeUpdate, error) {
	buf := bytebufferpool.Get()
	_, data, err := t.conn.ReadMessage(buf.B)
	if err != nil {
		return nil, err
	}
	r := &AggTradeUpdate{}
	err = json.Unmarshal(data, r)
	bytebufferpool.Put(buf)
	return r, err
}

// Stream stream a trades update message from aggregated trades websocket to channel
func (t *AggTrades) Stream() <-chan AggTradeUpdate {
	updates := make(chan AggTradeUpdate)
	go func() {
		defer close(updates)
		var msg []byte
		var data []byte
		var err error
		for {
			_, data, err = t.conn.ReadMessage(msg[:0])
			if err != nil {
				t.err = err
				return
			}

			u := AggTradeUpdate{}
			err = json.Unmarshal(data, &u)
			if err != nil {
				t.err = err
				return
			}
			updates <- u
		}
	}()
	return updates
}

// Trades is a wrapper for trades websocket
type Trades struct {
	wsClient
}

// Read reads a trades update message from trades websocket
func (t *Trades) Read() (*TradeUpdate, error) {
	buf := bytebufferpool.Get()
	_, data, err := t.conn.ReadMessage(buf.B)
	if err != nil {
		return nil, err
	}
	r := &TradeUpdate{}
	err = json.Unmarshal(data, r)
	bytebufferpool.Put(buf)
	return r, err
}

// Stream stream a trades update message from trades websocket to channel
func (t *Trades) Stream() <-chan TradeUpdate {
	updates := make(chan TradeUpdate)
	go func() {
		defer close(updates)
		var msg []byte
		var data []byte
		var err error
		for {
			_, data, err = t.conn.ReadMessage(msg[:0])
			if err != nil {
				t.err = err
				return
			}

			u := TradeUpdate{}
			err = json.Unmarshal(data, &u)
			if err != nil {
				t.err = err
				return
			}
			updates <- u
		}
	}()
	return updates
}

// AccountInfo is a wrapper for account info websocket
type AccountInfo struct {
	wsClient
}

// Read reads a account info update message from account info websocket
// Remark: The websocket is used to update two different structs, which both are flat, hence every call to this function
// will return either one of the types initialized and the other one will be set to nil
func (i *AccountInfo) Read() (UpdateType, interface{}, error) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	_, data, err := i.conn.ReadMessage(buf.B)
	if err != nil {
		return UpdateTypeUnknown, nil, err
	}
	et := EventTypeUpdate{}
	if err = json.Unmarshal(data, &et); err != nil {
		return UpdateTypeUnknown, nil, err
	}
	var resp interface{}
	switch et.EventType {
	case UpdateTypeOutboundAccountInfo:
		resp = &AccountInfoUpdate{}
	case UpdateTypeOutboundAccountPosition:
		resp = &AccountUpdate{}
	case UpdateTypeBalanceUpdate:
		resp = &BalanceUpdate{}
	case UpdateTypeOrderReport:
		resp = &OrderUpdate{}
	case UpdateTypeOCOReport:
		return et.EventType, nil, nil
	default:
		return et.EventType, data, nil
	}
	return et.EventType, resp, json.Unmarshal(data, resp)
}
