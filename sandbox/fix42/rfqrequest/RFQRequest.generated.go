package rfqrequest

import (
	"github.com/fardream/cbrs/sandbox/field"
	"github.com/fardream/cbrs/sandbox/fix42"
	"github.com/fardream/cbrs/sandbox/tag"
	"github.com/quickfixgo/quickfix"
)

// RFQRequest is the fix42 RFQRequest type, MsgType = AH.
type RFQRequest struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

// FromMessage creates a RFQRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) RFQRequest {
	return RFQRequest{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m RFQRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a RFQRequest initialized with the required fields for RFQRequest.
func New(rfqreqid field.RFQReqIDField) (m RFQRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AH"))
	m.Set(rfqreqid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg RFQRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "AH", r
}

// SetRFQReqID sets RFQReqID, Tag 644.
func (m RFQRequest) SetRFQReqID(v string) {
	m.Set(field.NewRFQReqID(v))
}

// GetRFQReqID gets RFQReqID, Tag 644.
func (m RFQRequest) GetRFQReqID() (v string, err quickfix.MessageRejectError) {
	var f field.RFQReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRFQReqID returns true if RFQReqID is present, Tag 644.
func (m RFQRequest) HasRFQReqID() bool {
	return m.Has(tag.RFQReqID)
}
