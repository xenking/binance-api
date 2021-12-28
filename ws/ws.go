package ws

import (
	"github.com/segmentio/encoding/json"
	"github.com/xenking/websocket"
)

type client struct {
	*websocket.Client
	streamErr error
}

func (c *client) Err() error {
	return c.streamErr
}

func (c *client) read(value interface{}) error {
	fr := websocket.AcquireFrame()
	fr.Reset()
	_, err := c.ReadFrame(fr)
	if err != nil {
		websocket.ReleaseFrame(fr)

		return err
	}
	err = json.Unmarshal(fr.Payload(), value)
	websocket.ReleaseFrame(fr)

	return err
}

func (c *client) stream(deferFunc func(), cb func(data []byte) error) {
	fr := websocket.AcquireFrame()
	defer websocket.ReleaseFrame(fr)
	defer deferFunc()

	var err error
	for {
		fr.Reset()
		_, err = c.ReadFrame(fr)
		if err != nil {
			c.streamErr = err

			return
		}
		err = cb(fr.Payload())
		if err != nil {
			c.streamErr = err

			return
		}
	}
}

// Depth is a wrapper for depth websocket
type Depth struct {
	client
}

// Read reads a depth update message from depth websocket
func (d *Depth) Read() (*DepthUpdate, error) {
	r := &DepthUpdate{}
	err := d.read(r)

	return r, err
}

// Stream stream a depth update message from depth websocket to channel
func (d *Depth) Stream() <-chan *DepthUpdate {
	updates := make(chan *DepthUpdate)
	cb := func(data []byte) error {
		u := &DepthUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go d.stream(func() {
		close(updates)
	}, cb)

	return updates
}

// DepthLevel is a wrapper for depth level websocket
type DepthLevel struct {
	client
}

// Read reads a depth update message from depth level websocket
func (d *DepthLevel) Read() (*DepthLevelUpdate, error) {
	r := &DepthLevelUpdate{}
	err := d.read(r)

	return r, err
}

// Stream stream a depth update message from depth level websocket to channel
func (d *DepthLevel) Stream() <-chan *DepthLevelUpdate {
	updates := make(chan *DepthLevelUpdate)
	cb := func(data []byte) error {
		u := &DepthLevelUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go d.stream(func() {
		close(updates)
	}, cb)

	return updates
}

// AllMarketTicker is a wrapper for all markets tickers websocket
type AllMarketTicker struct {
	client
}

// Read reads a market update message from all markets ticker websocket
func (t *AllMarketTicker) Read() (*AllMarketTickerUpdate, error) {
	r := &AllMarketTickerUpdate{}
	err := t.read(r)

	return r, err
}

// Stream stream a market update message from all markets ticker websocket to channel
func (t *AllMarketTicker) Stream() <-chan *AllMarketTickerUpdate {
	updates := make(chan *AllMarketTickerUpdate)
	cb := func(data []byte) error {
		u := &AllMarketTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.stream(func() {
		close(updates)
	}, cb)

	return updates
}

// IndivTicker is a wrapper for an individual ticker websocket
type IndivTicker struct {
	client
}

// Read reads a individual symbol update message from individual ticker websocket
func (t *IndivTicker) Read() (*IndivTickerUpdate, error) {
	r := &IndivTickerUpdate{}
	err := t.read(r)

	return r, err
}

// Stream stream a individual update message from individual ticker websocket to channel
func (t *IndivTicker) Stream() <-chan *IndivTickerUpdate {
	updates := make(chan *IndivTickerUpdate)
	cb := func(data []byte) error {
		u := &IndivTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.stream(func() {
		close(updates)
	}, cb)

	return updates
}

// AllMarketMiniTicker is a wrapper for all markets mini-tickers websocket
type AllMarketMiniTicker struct {
	client
}

// Read reads a market update message from all markets mini-ticker websocket
func (t *AllMarketMiniTicker) Read() (*AllMarketMiniTickerUpdate, error) {
	r := &AllMarketMiniTickerUpdate{}
	err := t.read(r)

	return r, err
}

// Stream stream a market update message from all markets mini-ticker websocket to channel
func (t *AllMarketMiniTicker) Stream() <-chan *AllMarketMiniTickerUpdate {
	updates := make(chan *AllMarketMiniTickerUpdate)
	cb := func(data []byte) error {
		u := &AllMarketMiniTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.stream(func() {
		close(updates)
	}, cb)

	return updates
}

// IndivMiniTicker is a wrapper for an individual mini-ticker websocket
type IndivMiniTicker struct {
	client
}

// Read reads a individual symbol update message from individual mini-ticker websocket
func (t *IndivMiniTicker) Read() (*IndivMiniTickerUpdate, error) {
	r := &IndivMiniTickerUpdate{}
	err := t.read(r)

	return r, err
}

// Stream stream a individual update message from individual mini-ticker websocket to channel
func (t *IndivMiniTicker) Stream() <-chan *IndivMiniTickerUpdate {
	updates := make(chan *IndivMiniTickerUpdate)
	cb := func(data []byte) error {
		u := &IndivMiniTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.stream(func() {
		close(updates)
	}, cb)

	return updates
}

// AllBookTicker is a wrapper for all book tickers websocket
type AllBookTicker struct {
	client
}

// Read reads a book update message from all book tickers websocket
func (t *AllBookTicker) Read() (*AllBookTickerUpdate, error) {
	r := &AllBookTickerUpdate{}
	err := t.read(r)

	return r, err
}

// Stream stream a book update message from all book tickers websocket to channel
func (t *AllBookTicker) Stream() <-chan *AllBookTickerUpdate {
	updates := make(chan *AllBookTickerUpdate)
	cb := func(data []byte) error {
		u := &AllBookTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.stream(func() {
		close(updates)
	}, cb)

	return updates
}

// IndivBookTicker is a wrapper for an individual book ticker websocket
type IndivBookTicker struct {
	client
}

// Read reads a individual book symbol update message from individual book ticker websocket
func (t *IndivBookTicker) Read() (*IndivBookTickerUpdate, error) {
	r := &IndivBookTickerUpdate{}
	err := t.read(r)

	return r, err
}

// Stream stream a individual book symbol update message from individual book ticker websocket to channel
func (t *IndivBookTicker) Stream() <-chan *IndivBookTickerUpdate {
	updates := make(chan *IndivBookTickerUpdate)
	cb := func(data []byte) error {
		u := &IndivBookTickerUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.stream(func() {
		close(updates)
	}, cb)

	return updates
}

// Klines is a wrapper for klines websocket
type Klines struct {
	client
}

// Read reads a klines update message from klines websocket
func (k *Klines) Read() (*KlinesUpdate, error) {
	r := &KlinesUpdate{}
	err := k.read(r)

	return r, err
}

// Stream stream a klines update message from klines websocket to channel
func (k *Klines) Stream() <-chan *KlinesUpdate {
	updates := make(chan *KlinesUpdate)
	cb := func(data []byte) error {
		u := &KlinesUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go k.stream(func() {
		close(updates)
	}, cb)

	return updates
}

// AggTrades is a wrapper for trades websocket
type AggTrades struct {
	client
}

// Read reads a trades update message from aggregated trades websocket
func (t *AggTrades) Read() (*AggTradeUpdate, error) {
	r := &AggTradeUpdate{}
	err := t.read(r)

	return r, err
}

// Stream stream a trades update message from aggregated trades websocket to channel
func (t *AggTrades) Stream() <-chan *AggTradeUpdate {
	updates := make(chan *AggTradeUpdate)
	cb := func(data []byte) error {
		u := &AggTradeUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.stream(func() {
		close(updates)
	}, cb)

	return updates
}

// Trades is a wrapper for trades websocket
type Trades struct {
	client
}

// Read reads a trades update message from trades websocket
func (t *Trades) Read() (*TradeUpdate, error) {
	r := &TradeUpdate{}
	err := t.read(r)

	return r, err
}

// Stream stream a trades update message from trades websocket to channel
func (t *Trades) Stream() <-chan *TradeUpdate {
	updates := make(chan *TradeUpdate)
	cb := func(data []byte) error {
		u := &TradeUpdate{}
		if err := json.Unmarshal(data, u); err != nil {
			return err
		}
		updates <- u

		return nil
	}
	go t.stream(func() {
		close(updates)
	}, cb)

	return updates
}

// AccountInfo is a wrapper for account info websocket
type AccountInfo struct {
	client
}

// Read reads a account info update message from account info websocket
// Remark: The websocket is used to update two different structs, which both are flat, hence every call to this function
// will return either one of the types initialized and the other one will be set to nil
func (i *AccountInfo) Read() (UpdateType, interface{}, error) {
	fr := websocket.AcquireFrame()
	defer websocket.ReleaseFrame(fr)
	fr.Reset()
	_, err := i.ReadFrame(fr)
	if err != nil {
		return UpdateTypeUnknown, nil, err
	}
	et := EventTypeUpdate{}
	err = json.Unmarshal(fr.Payload(), &et)
	if err != nil {
		return UpdateTypeUnknown, nil, err
	}

	var resp interface{}
	//nolint:exhaustive
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
		return et.EventType, fr.Payload(), nil
	}
	err = json.Unmarshal(fr.Payload(), resp)

	return et.EventType, resp, err
}
