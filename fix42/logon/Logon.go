//Package logon msg type = A.
package logon

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a Logon wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//EncryptMethod is a required field for Logon.
func (m Message) EncryptMethod() (*field.EncryptMethodField, quickfix.MessageRejectError) {
	f := &field.EncryptMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptMethod reads a EncryptMethod from Logon.
func (m Message) GetEncryptMethod(f *field.EncryptMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//HeartBtInt is a required field for Logon.
func (m Message) HeartBtInt() (*field.HeartBtIntField, quickfix.MessageRejectError) {
	f := &field.HeartBtIntField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHeartBtInt reads a HeartBtInt from Logon.
func (m Message) GetHeartBtInt(f *field.HeartBtIntField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for Logon.
func (m Message) RawDataLength() (*field.RawDataLengthField, quickfix.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from Logon.
func (m Message) GetRawDataLength(f *field.RawDataLengthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for Logon.
func (m Message) RawData() (*field.RawDataField, quickfix.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from Logon.
func (m Message) GetRawData(f *field.RawDataField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ResetSeqNumFlag is a non-required field for Logon.
func (m Message) ResetSeqNumFlag() (*field.ResetSeqNumFlagField, quickfix.MessageRejectError) {
	f := &field.ResetSeqNumFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetResetSeqNumFlag reads a ResetSeqNumFlag from Logon.
func (m Message) GetResetSeqNumFlag(f *field.ResetSeqNumFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaxMessageSize is a non-required field for Logon.
func (m Message) MaxMessageSize() (*field.MaxMessageSizeField, quickfix.MessageRejectError) {
	f := &field.MaxMessageSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxMessageSize reads a MaxMessageSize from Logon.
func (m Message) GetMaxMessageSize(f *field.MaxMessageSizeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoMsgTypes is a non-required field for Logon.
func (m Message) NoMsgTypes() (*field.NoMsgTypesField, quickfix.MessageRejectError) {
	f := &field.NoMsgTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMsgTypes reads a NoMsgTypes from Logon.
func (m Message) GetNoMsgTypes(f *field.NoMsgTypesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds Logon messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for Logon.
func Builder(
	encryptmethod *field.EncryptMethodField,
	heartbtint *field.HeartBtIntField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("A"))
	builder.Body().Set(encryptmethod)
	builder.Body().Set(heartbtint)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX42, "A", r
}
