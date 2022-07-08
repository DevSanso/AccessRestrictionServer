package protobuf_test

import (
	"testing"
	"url_mapping/parser/protobuf"
	"core/proto"
)



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

