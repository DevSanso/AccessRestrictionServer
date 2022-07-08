package protobuf_test

import (
	"testing"
	"url_mapping/parser/protobuf"
	"url_mapping/utils/tests"
)





func TestHeader(t *testing.T) {
	header := tests.NewHeader()
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
	if !tests.HeaderEq(&encodeHeader,header) {
		t.Error("not eq header data")
		t.Fail()
	}
}


