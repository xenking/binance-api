package ws

import (
	"io"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/segmentio/encoding/json"
)

type Conn struct {
	conn net.Conn
}

func NewConn(conn net.Conn) Conn {
	return Conn{conn: conn}
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func (c *Conn) NetConn() net.Conn {
	return c.conn
}

func (c *Conn) ReadValue(value interface{}) error {
	h, r, err := wsutil.NextReader(c.conn, ws.StateClientSide)
	if err != nil {
		return err
	}
	if h.OpCode.IsControl() {
		return wsutil.ControlFrameHandler(c.conn, ws.StateClientSide)(h, r)
	}
	err = wsutil.WriteClientMessage(c.conn, ws.OpPong, nil)
	if err != nil {
		return err
	}
	return json.NewDecoder(r).Decode(value)
}

func (c *Conn) ReadRaw() ([]byte, error) {
	h, r, err := wsutil.NextReader(c.conn, ws.StateClientSide)
	if err != nil {
		return nil, err
	}
	if h.OpCode.IsControl() {
		return nil, wsutil.ControlFrameHandler(c.conn, ws.StateClientSide)(h, r)
	}
	err = wsutil.WriteClientMessage(c.conn, ws.OpPong, nil)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(r)
}

func (c *Conn) NewStream(callback func(dec *json.Decoder, err error) error) {
	var (
		state          = ws.StateClientSide
		reader         = wsutil.NewReader(c.conn, state)
		controlHandler = wsutil.ControlFrameHandler(c.conn, state)
		decoder        = json.NewDecoder(reader)
	)
	defer c.conn.Close()

	for {
		hdr, err := reader.NextFrame()
		if err != nil {
			_ = callback(nil, err)
			return
		}
		if hdr.OpCode.IsControl() {
			err = controlHandler(hdr, reader)
		}
		err = callback(decoder, err)
		if err != nil {
			return
		}
	}
}

func (c *Conn) NewStreamRaw(callback func(buf []byte, err error) error) {
	var (
		state          = ws.StateClientSide
		reader         = wsutil.NewReader(c.conn, state)
		controlHandler = wsutil.ControlFrameHandler(c.conn, state)
	)
	defer c.conn.Close()

	for {
		hdr, err := reader.NextFrame()
		if err != nil {
			_ = callback(nil, err)
			return
		}
		if hdr.OpCode.IsControl() {
			err = controlHandler(hdr, reader)
			err = callback(nil, err)
		}
		if err != nil {
			return
		}
		err = callback(io.ReadAll(reader))
		if err != nil {
			return
		}
	}
}
