package ws

import "github.com/segmentio/encoding/json"

// AccountInfo is a wrapper for account info websocket
type AccountInfo struct {
	Conn
}

// Read reads a account info update message from account info websocket
// Remark: The websocket is used to update two different structs, which both are flat, hence every call to this function
// will return either one of the types initialized and the other one will be set to nil
func (i *AccountInfo) Read() (AccountUpdateEventType, interface{}, error) {
	payload, err := i.Conn.ReadRaw()
	if err != nil {
		return AccountUpdateEventTypeUnknown, nil, err
	}
	et := UpdateEventType{}
	err = json.Unmarshal(payload, &et)
	if err != nil {
		return AccountUpdateEventTypeUnknown, nil, err
	}

	var resp interface{}
	switch et.EventType {
	case AccountUpdateEventTypeOutboundAccountPosition:
		resp = &AccountUpdateEvent{}
	case AccountUpdateEventTypeBalanceUpdate:
		resp = &BalanceUpdateEvent{}
	case AccountUpdateEventTypeOrderReport:
		resp = &OrderUpdateEvent{}
	case AccountUpdateEventTypeOCOReport:
		resp = &OCOOrderUpdateEvent{}
	default:
		buf := make([]byte, len(payload))
		copy(buf, payload)
		return et.EventType, buf, nil
	}
	err = json.Unmarshal(payload, resp)

	return et.EventType, resp, err
}

func (i *AccountInfo) OrdersStream() <-chan *OrderUpdateEvent {
	updates := make(chan *OrderUpdateEvent)

	go i.NewStreamRaw(func(payload []byte, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		var event UpdateEventType
		err = event.UnmarshalJSON(payload)
		if err != nil {
			return err
		}

		if event.EventType != AccountUpdateEventTypeOrderReport {
			return nil
		}

		u := &OrderUpdateEvent{}
		err = json.Unmarshal(payload, u)
		if err != nil {
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

func (i *AccountInfo) OCOOrdersStream() <-chan *OCOOrderUpdateEvent {
	updates := make(chan *OCOOrderUpdateEvent)

	go i.NewStreamRaw(func(payload []byte, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		var event UpdateEventType
		err = event.UnmarshalJSON(payload)
		if err != nil {
			return err
		}

		if event.EventType != AccountUpdateEventTypeOCOReport {
			return nil
		}

		u := &OCOOrderUpdateEvent{}
		err = json.Unmarshal(payload, u)
		if err != nil {
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

func (i *AccountInfo) BalancesStream() <-chan *BalanceUpdateEvent {
	updates := make(chan *BalanceUpdateEvent)

	go i.NewStreamRaw(func(payload []byte, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		var event UpdateEventType
		err = event.UnmarshalJSON(payload)
		if err != nil {
			return err
		}

		if event.EventType != AccountUpdateEventTypeBalanceUpdate {
			return nil
		}

		u := &BalanceUpdateEvent{}
		err = json.Unmarshal(payload, u)
		if err != nil {
			return err
		}
		updates <- u

		return nil
	})

	return updates
}

func (i *AccountInfo) AccountStream() <-chan *AccountUpdateEvent {
	updates := make(chan *AccountUpdateEvent)

	go i.NewStreamRaw(func(payload []byte, err error) error {
		if err != nil {
			close(updates)
			return err
		}

		var event UpdateEventType
		err = event.UnmarshalJSON(payload)
		if err != nil {
			return err
		}

		if event.EventType != AccountUpdateEventTypeOutboundAccountPosition {
			return nil
		}

		u := &AccountUpdateEvent{}
		err = json.Unmarshal(payload, u)
		if err != nil {
			return err
		}
		updates <- u

		return nil
	})

	return updates
}
