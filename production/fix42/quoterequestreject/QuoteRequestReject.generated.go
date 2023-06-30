package quoterequestreject

import (
	"github.com/shopspring/decimal"

	"github.com/fardream/cbrs/sandbox/enum"
	"github.com/fardream/cbrs/sandbox/field"
	"github.com/fardream/cbrs/sandbox/fix42"
	"github.com/fardream/cbrs/sandbox/tag"
	"github.com/quickfixgo/quickfix"
)

// QuoteRequestReject is the fix42 QuoteRequestReject type, MsgType = AG.
type QuoteRequestReject struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

// FromMessage creates a QuoteRequestReject from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) QuoteRequestReject {
	return QuoteRequestReject{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m QuoteRequestReject) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a QuoteRequestReject initialized with the required fields for QuoteRequestReject.
func New(quotereqid field.QuoteReqIDField, quoterequestrejectreason field.QuoteRequestRejectReasonField) (m QuoteRequestReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AG"))
	m.Set(quotereqid)
	m.Set(quoterequestrejectreason)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg QuoteRequestReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "AG", r
}

// SetQuoteReqID sets QuoteReqID, Tag 131.
func (m QuoteRequestReject) SetQuoteReqID(v string) {
	m.Set(field.NewQuoteReqID(v))
}

// SetNoRelatedSym sets NoRelatedSym, Tag 146.
func (m QuoteRequestReject) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

// SetQuoteRequestRejectReason sets QuoteRequestRejectReason, Tag 658.
func (m QuoteRequestReject) SetQuoteRequestRejectReason(v enum.QuoteRequestRejectReason) {
	m.Set(field.NewQuoteRequestRejectReason(v))
}

// GetQuoteReqID gets QuoteReqID, Tag 131.
func (m QuoteRequestReject) GetQuoteReqID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRelatedSym gets NoRelatedSym, Tag 146.
func (m QuoteRequestReject) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetQuoteRequestRejectReason gets QuoteRequestRejectReason, Tag 658.
func (m QuoteRequestReject) GetQuoteRequestRejectReason() (v enum.QuoteRequestRejectReason, err quickfix.MessageRejectError) {
	var f field.QuoteRequestRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasQuoteReqID returns true if QuoteReqID is present, Tag 131.
func (m QuoteRequestReject) HasQuoteReqID() bool {
	return m.Has(tag.QuoteReqID)
}

// HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146.
func (m QuoteRequestReject) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

// HasQuoteRequestRejectReason returns true if QuoteRequestRejectReason is present, Tag 658.
func (m QuoteRequestReject) HasQuoteRequestRejectReason() bool {
	return m.Has(tag.QuoteRequestRejectReason)
}

// NoRelatedSym is a repeating group element, Tag 146.
type NoRelatedSym struct {
	*quickfix.Group
}

// SetSymbol sets Symbol, Tag 55.
func (m NoRelatedSym) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetPrice sets Price, Tag 44.
func (m NoRelatedSym) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

// SetOrderQty sets OrderQty, Tag 38.
func (m NoRelatedSym) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

// GetSymbol gets Symbol, Tag 55.
func (m NoRelatedSym) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPrice gets Price, Tag 44.
func (m NoRelatedSym) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderQty gets OrderQty, Tag 38.
func (m NoRelatedSym) GetOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m NoRelatedSym) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasPrice returns true if Price is present, Tag 44.
func (m NoRelatedSym) HasPrice() bool {
	return m.Has(tag.Price)
}

// HasOrderQty returns true if OrderQty is present, Tag 38.
func (m NoRelatedSym) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

// NoRelatedSymRepeatingGroup is a repeating group, Tag 146.
type NoRelatedSymRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedSymRepeatingGroup returns an initialized, NoRelatedSymRepeatingGroup.
func NewNoRelatedSymRepeatingGroup() NoRelatedSymRepeatingGroup {
	return NoRelatedSymRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedSym,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.Price), quickfix.GroupElement(tag.OrderQty)}),
	}
}

// Add create and append a new NoRelatedSym to this group.
func (m NoRelatedSymRepeatingGroup) Add() NoRelatedSym {
	g := m.RepeatingGroup.Add()
	return NoRelatedSym{g}
}

// Get returns the ith NoRelatedSym in the NoRelatedSymRepeatinGroup.
func (m NoRelatedSymRepeatingGroup) Get(i int) NoRelatedSym {
	return NoRelatedSym{m.RepeatingGroup.Get(i)}
}
