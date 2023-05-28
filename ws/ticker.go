package ws

import "github.com/segmentio/encoding/json"

// Depth is a wrapper for depth websocket
type Depth struct {
	Conn
}

// Read reads a depth update message from depth websocket
func (d *Depth) Read() (*DepthUpdate, error) {
	r := &DepthUpdate{}
	err := d.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a depth update message from depth websocket to channel
func (d *Depth) Stream() <-chan *DepthUpdate {
	updates := make(chan *DepthUpdate)
	go d.NewStream(func(dec *json.Decoder, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		u := &DepthUpdate{}
		err = dec.Decode(u)
		if err != nil {
			close(updates)
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

// DepthLevel is a wrapper for depth level websocket
type DepthLevel struct {
	Conn
}

// Read reads a depth update message from depth level websocket
func (d *DepthLevel) Read() (*DepthLevelUpdate, error) {
	r := &DepthLevelUpdate{}
	err := d.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a depth update message from depth level websocket to channel
func (d *DepthLevel) Stream() <-chan *DepthLevelUpdate {
	updates := make(chan *DepthLevelUpdate)
	go d.NewStream(func(dec *json.Decoder, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		u := &DepthLevelUpdate{}
		err = dec.Decode(u)
		if err != nil {
			close(updates)
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

// AllMarketTicker is a wrapper for all markets tickers websocket
type AllMarketTicker struct {
	Conn
}

// Read reads a market update message from all markets ticker websocket
func (t *AllMarketTicker) Read() (*AllMarketTickerUpdate, error) {
	r := &AllMarketTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a market update message from all markets ticker websocket to channel
func (t *AllMarketTicker) Stream() <-chan *AllMarketTickerUpdate {
	updates := make(chan *AllMarketTickerUpdate)
	go t.NewStream(func(dec *json.Decoder, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		u := &AllMarketTickerUpdate{}
		err = dec.Decode(u)
		if err != nil {
			close(updates)
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

// IndividualTicker is a wrapper for an Individualidual ticker websocket
type IndividualTicker struct {
	Conn
}

// Read reads a Individualidual symbol update message from Individualidual ticker websocket
func (t *IndividualTicker) Read() (*IndividualTickerUpdate, error) {
	r := &IndividualTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a Individualidual update message from Individualidual ticker websocket to channel
func (t *IndividualTicker) Stream() <-chan *IndividualTickerUpdate {
	updates := make(chan *IndividualTickerUpdate)
	go t.NewStream(func(dec *json.Decoder, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		u := &IndividualTickerUpdate{}
		err = dec.Decode(u)
		if err != nil {
			close(updates)
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

// AllMarketMiniTicker is a wrapper for all markets mini-tickers websocket
type AllMarketMiniTicker struct {
	Conn
}

// Read reads a market update message from all markets mini-ticker websocket
func (t *AllMarketMiniTicker) Read() (*AllMarketMiniTickerUpdate, error) {
	r := &AllMarketMiniTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a market update message from all markets mini-ticker websocket to channel
func (t *AllMarketMiniTicker) Stream() <-chan *AllMarketMiniTickerUpdate {
	updates := make(chan *AllMarketMiniTickerUpdate)
	go t.NewStream(func(dec *json.Decoder, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		u := &AllMarketMiniTickerUpdate{}
		err = dec.Decode(u)
		if err != nil {
			close(updates)
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

// IndividualMiniTicker is a wrapper for an Individualidual mini-ticker websocket
type IndividualMiniTicker struct {
	Conn
}

// Read reads a Individualidual symbol update message from Individualidual mini-ticker websocket
func (t *IndividualMiniTicker) Read() (*IndividualMiniTickerUpdate, error) {
	r := &IndividualMiniTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a Individualidual update message from Individualidual mini-ticker websocket to channel
func (t *IndividualMiniTicker) Stream() <-chan *IndividualMiniTickerUpdate {
	updates := make(chan *IndividualMiniTickerUpdate)
	go t.NewStream(func(dec *json.Decoder, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		u := &IndividualMiniTickerUpdate{}
		err = dec.Decode(u)
		if err != nil {
			close(updates)
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

// AllBookTicker is a wrapper for all book tickers websocket
type AllBookTicker struct {
	Conn
}

// Read reads a book update message from all book tickers websocket
func (t *AllBookTicker) Read() (*AllBookTickerUpdate, error) {
	r := &AllBookTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a book update message from all book tickers websocket to channel
func (t *AllBookTicker) Stream() <-chan *AllBookTickerUpdate {
	updates := make(chan *AllBookTickerUpdate)
	go t.NewStream(func(dec *json.Decoder, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		u := &AllBookTickerUpdate{}
		err = dec.Decode(u)
		if err != nil {
			close(updates)
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

// IndividualBookTicker is a wrapper for an Individualidual book ticker websocket
type IndividualBookTicker struct {
	Conn
}

// Read reads a Individualidual book symbol update message from Individualidual book ticker websocket
func (t *IndividualBookTicker) Read() (*IndividualBookTickerUpdate, error) {
	r := &IndividualBookTickerUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a Individualidual book symbol update message from Individualidual book ticker websocket to channel
func (t *IndividualBookTicker) Stream() <-chan *IndividualBookTickerUpdate {
	updates := make(chan *IndividualBookTickerUpdate)
	go t.NewStream(func(dec *json.Decoder, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		u := &IndividualBookTickerUpdate{}
		err = dec.Decode(u)
		if err != nil {
			close(updates)
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

// Klines is a wrapper for klines websocket
type Klines struct {
	Conn
}

// Read reads a klines update message from klines websocket
func (k *Klines) Read() (*KlinesUpdate, error) {
	r := &KlinesUpdate{}
	err := k.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a klines update message from klines websocket to channel
func (k *Klines) Stream() <-chan *KlinesUpdate {
	updates := make(chan *KlinesUpdate)
	go k.NewStream(func(dec *json.Decoder, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		u := &KlinesUpdate{}
		err = dec.Decode(u)
		if err != nil {
			close(updates)
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

// AggTrades is a wrapper for trades websocket
type AggTrades struct {
	Conn
}

// Read reads a trades update message from aggregated trades websocket
func (t *AggTrades) Read() (*AggTradeUpdate, error) {
	r := &AggTradeUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a trades update message from aggregated trades websocket to channel
func (t *AggTrades) Stream() <-chan *AggTradeUpdate {
	updates := make(chan *AggTradeUpdate)
	go t.NewStream(func(dec *json.Decoder, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		u := &AggTradeUpdate{}
		err = dec.Decode(u)
		if err != nil {
			close(updates)
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

// Trades is a wrapper for trades websocket
type Trades struct {
	Conn
}

// Read reads a trades update message from trades websocket
func (t *Trades) Read() (*TradeUpdate, error) {
	r := &TradeUpdate{}
	err := t.Conn.ReadValue(r)

	return r, err
}

// Stream NewStream a trades update message from trades websocket to channel
func (t *Trades) Stream() <-chan *TradeUpdate {
	updates := make(chan *TradeUpdate)
	go t.NewStream(func(dec *json.Decoder, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		u := &TradeUpdate{}
		err = dec.Decode(u)
		if err != nil {
			close(updates)
			return err
		}
		updates <- u

		return nil
	})

	return updates
}
