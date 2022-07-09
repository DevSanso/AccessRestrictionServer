package tests

import (
	"core/proto"
	"github.com/google/uuid"
)

var headerVal map[string]interface{} = map[string]interface{}{
	"header":        "application/json",
	"is_next":       false,
	"message_index": int32(1),
	"message_uuid":  uuid.New().String(),
	"sender_id":   "test",
	"state":         int32(proto.MSAHeader_Success),
	"type":          int32(proto.MSAHeader_General),
	"method":        "POST",
}

func NewHeader() *proto.MSAHeader {
	header := new(proto.MSAHeader)
	header.HttpHeader = headerVal["header"].(string)
	header.IsNext = headerVal["is_next"].(bool)
	header.MessageIndex = headerVal["message_index"].(int32)
	header.MessageUuid = headerVal["message_uuid"].(string)
	header.SenderId = headerVal["sender_id"].(string)
	header.State = proto.MSAHeader_State(headerVal["state"].(int32))
	header.Type = proto.MSAHeader_Type(headerVal["type"].(int32))
	header.Method = headerVal["method"].(string)
	return header
}