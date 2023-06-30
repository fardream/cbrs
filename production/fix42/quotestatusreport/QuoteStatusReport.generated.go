package quotestatusreport

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/fardream/cbrs/sandbox/enum"
	"github.com/fardream/cbrs/sandbox/field"
	"github.com/fardream/cbrs/sandbox/fix42"
	"github.com/fardream/cbrs/sandbox/tag"
	"github.com/quickfixgo/quickfix"
)

// QuoteStatusReport is the fix42 QuoteStatusReport type, MsgType = AI.
type QuoteStatusReport struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

// FromMessage creates a QuoteStatusReport from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) QuoteStatusReport {
	return QuoteStatusReport{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m QuoteStatusReport) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a QuoteStatusReport initialized with the required fields for QuoteStatusReport.
func New(quotereqid field.QuoteReqIDField, symbol field.SymbolField, orderqty field.OrderQtyField, validuntiltime field.ValidUntilTimeField, expiretime field.ExpireTimeField, quotestatus field.QuoteStatusField) (m QuoteStatusReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AI"))
	m.Set(quotereqid)
	m.Set(symbol)
	m.Set(orderqty)
	m.Set(validuntiltime)
	m.Set(expiretime)
	m.Set(quotestatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg QuoteStatusReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "AI", r
}

// SetOrderQty sets OrderQty, Tag 38.
func (m QuoteStatusReport) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

// SetSide sets Side, Tag 54.
func (m QuoteStatusReport) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m QuoteStatusReport) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetValidUntilTime sets ValidUntilTime, Tag 62.
func (m QuoteStatusReport) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

// SetQuoteID sets QuoteID, Tag 117.
func (m QuoteStatusReport) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

// SetExpireTime sets ExpireTime, Tag 126.
func (m QuoteStatusReport) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

// SetQuoteReqID sets QuoteReqID, Tag 131.
func (m QuoteStatusReport) SetQuoteReqID(v string) {
	m.Set(field.NewQuoteReqID(v))
}

// SetBidPx sets BidPx, Tag 132.
func (m QuoteStatusReport) SetBidPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidPx(value, scale))
}

// SetOfferPx sets OfferPx, Tag 133.
func (m QuoteStatusReport) SetOfferPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferPx(value, scale))
}

// SetQuoteStatus sets QuoteStatus, Tag 297.
func (m QuoteStatusReport) SetQuoteStatus(v enum.QuoteStatus) {
	m.Set(field.NewQuoteStatus(v))
}

// SetQuoteRejectReason sets QuoteRejectReason, Tag 300.
func (m QuoteStatusReport) SetQuoteRejectReason(v enum.QuoteRejectReason) {
	m.Set(field.NewQuoteRejectReason(v))
}

// GetOrderQty gets OrderQty, Tag 38.
func (m QuoteStatusReport) GetOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSide gets Side, Tag 54.
func (m QuoteStatusReport) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m QuoteStatusReport) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetValidUntilTime gets ValidUntilTime, Tag 62.
func (m QuoteStatusReport) GetValidUntilTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ValidUntilTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteID gets QuoteID, Tag 117.
func (m QuoteStatusReport) GetQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExpireTime gets ExpireTime, Tag 126.
func (m QuoteStatusReport) GetExpireTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ExpireTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteReqID gets QuoteReqID, Tag 131.
func (m QuoteStatusReport) GetQuoteReqID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBidPx gets BidPx, Tag 132.
func (m QuoteStatusReport) GetBidPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOfferPx gets OfferPx, Tag 133.
func (m QuoteStatusReport) GetOfferPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteStatus gets QuoteStatus, Tag 297.
func (m QuoteStatusReport) GetQuoteStatus() (v enum.QuoteStatus, err quickfix.MessageRejectError) {
	var f field.QuoteStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteRejectReason gets QuoteRejectReason, Tag 300.
func (m QuoteStatusReport) GetQuoteRejectReason() (v enum.QuoteRejectReason, err quickfix.MessageRejectError) {
	var f field.QuoteRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasOrderQty returns true if OrderQty is present, Tag 38.
func (m QuoteStatusReport) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

// HasSide returns true if Side is present, Tag 54.
func (m QuoteStatusReport) HasSide() bool {
	return m.Has(tag.Side)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m QuoteStatusReport) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasValidUntilTime returns true if ValidUntilTime is present, Tag 62.
func (m QuoteStatusReport) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

// HasQuoteID returns true if QuoteID is present, Tag 117.
func (m QuoteStatusReport) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

// HasExpireTime returns true if ExpireTime is present, Tag 126.
func (m QuoteStatusReport) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

// HasQuoteReqID returns true if QuoteReqID is present, Tag 131.
func (m QuoteStatusReport) HasQuoteReqID() bool {
	return m.Has(tag.QuoteReqID)
}

// HasBidPx returns true if BidPx is present, Tag 132.
func (m QuoteStatusReport) HasBidPx() bool {
	return m.Has(tag.BidPx)
}

// HasOfferPx returns true if OfferPx is present, Tag 133.
func (m QuoteStatusReport) HasOfferPx() bool {
	return m.Has(tag.OfferPx)
}

// HasQuoteStatus returns true if QuoteStatus is present, Tag 297.
func (m QuoteStatusReport) HasQuoteStatus() bool {
	return m.Has(tag.QuoteStatus)
}

// HasQuoteRejectReason returns true if QuoteRejectReason is present, Tag 300.
func (m QuoteStatusReport) HasQuoteRejectReason() bool {
	return m.Has(tag.QuoteRejectReason)
}
