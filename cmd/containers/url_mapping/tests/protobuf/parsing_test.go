package proto_parser_test

import (
	"testing"
	"url_mapping/parser/protobuf"
	"core/proto"
	"github.com/google/uuid"
)

var headerVal map[string]interface{} = map[string]interface{}{
	"header" : "application/json",
	"is_next" : false,
	"message_index": int32(1),
	"message_uuid" : uuid.New().String(),
	"register_id" : "test",
	"state" : int32(proto.MSAHeader_Success),
	"type" : int32(proto.MSAHeader_General),
	"method" : "POST",
} 
func newHeader() *proto.MSAHeader{
	header := new(proto.MSAHeader)
	header.HttpHeader = headerVal["header"].(string)
	header.IsNext = headerVal["is_next"].(bool)
	header.MessageIndex = headerVal["message_index"].(int32)
	header.MessageUuid = headerVal["message_uuid"].(string)
	header.RequesterId = headerVal["register_id"].(string)
	header.State = proto.MSAHeader_State(headerVal["state"].(int32))
	header.Type = proto.MSAHeader_Type(headerVal["type"].(int32))
	header.Method = headerVal["method"].(string)
	return header
}

func headerEq(h *proto.MSAHeader) bool {
	return h.HttpHeader == headerVal["header"].(string) &&
	h.IsNext == headerVal["is_next"].(bool) &&
	h.MessageIndex == headerVal["message_index"].(int32) &&
	h.MessageUuid == headerVal["message_uuid"].(string) &&
	h.RequesterId == headerVal["register_id"].(string) &&
	h.State == proto.MSAHeader_State(headerVal["state"].(int32)) &&
	h.Type == proto.MSAHeader_Type(headerVal["type"].(int32)) &&
	h.Method == headerVal["method"].(string)
}

func TestHeader(t *testing.T) {
	header := newHeader()
	bytes,err := protobuf.DecodeHeader(header)
	if err != nil {
		t.Error(err)
		return
	}
	encodeHeader, encodeErr := protobuf.EncodeHeader(bytes)
	if encodeErr != nil {
		t.Error(err)
		return
	}
	if !headerEq(&encodeHeader) {
		t.Error("not eq header data")
		t.Fail()
	}
}

func TestMessage(t *testing.T) {
	header := newHeader()
	message := &proto.MSAMessage{}
	message.Header = header
	message.HttpBody = "hello world"

	bytes,err := protobuf.DecodeMessage(message)
	if err != nil {
		t.Error(err)
		return
	}
	encodeMessage, encodeErr := protobuf.EncodeMessage(bytes)
	if encodeErr != nil {
		t.Error(err)
		return
	}
	if !headerEq(encodeMessage.GetHeader()){
		t.Error("not eq header data")
		t.Fail()
	}
	if message.GetHttpBody() != "hello world" {
		t.Error("not eq message body")
		t.Fail()
	}
}

