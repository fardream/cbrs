package quoteresponse

import (
	"github.com/fardream/cbrs/sandbox/enum"
	"github.com/fardream/cbrs/sandbox/field"
	"github.com/fardream/cbrs/sandbox/fix42"
	"github.com/fardream/cbrs/sandbox/tag"
	"github.com/quickfixgo/quickfix"
)

// QuoteResponse is the fix42 QuoteResponse type, MsgType = AJ.
type QuoteResponse struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

// FromMessage creates a QuoteResponse from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) QuoteResponse {
	return QuoteResponse{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m QuoteResponse) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a QuoteResponse initialized with the required fields for QuoteResponse.
func New(quoterespid field.QuoteRespIDField, quoteresptype field.QuoteRespTypeField, side field.SideField, symbol field.SymbolField) (m QuoteResponse) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AJ"))
	m.Set(quoterespid)
	m.Set(quoteresptype)
	m.Set(side)
	m.Set(symbol)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg QuoteResponse, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "AJ", r
}

// SetSide sets Side, Tag 54.
func (m QuoteResponse) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m QuoteResponse) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetQuoteRespID sets QuoteRespID, Tag 693.
func (m QuoteResponse) SetQuoteRespID(v string) {
	m.Set(field.NewQuoteRespID(v))
}

// SetQuoteRespType sets QuoteRespType, Tag 694.
func (m QuoteResponse) SetQuoteRespType(v enum.QuoteRespType) {
	m.Set(field.NewQuoteRespType(v))
}

// GetSide gets Side, Tag 54.
func (m QuoteResponse) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m QuoteResponse) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteRespID gets QuoteRespID, Tag 693.
func (m QuoteResponse) GetQuoteRespID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteRespIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteRespType gets QuoteRespType, Tag 694.
func (m QuoteResponse) GetQuoteRespType() (v enum.QuoteRespType, err quickfix.MessageRejectError) {
	var f field.QuoteRespTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSide returns true if Side is present, Tag 54.
func (m QuoteResponse) HasSide() bool {
	return m.Has(tag.Side)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m QuoteResponse) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasQuoteRespID returns true if QuoteRespID is present, Tag 693.
func (m QuoteResponse) HasQuoteRespID() bool {
	return m.Has(tag.QuoteRespID)
}

// HasQuoteRespType returns true if QuoteRespType is present, Tag 694.
func (m QuoteResponse) HasQuoteRespType() bool {
	return m.Has(tag.QuoteRespType)
}
