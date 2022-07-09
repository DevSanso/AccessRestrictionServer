package message

import (
	"core/proto"
	"encoding/json"
)

func writeOkHeader(header,src *proto.MSAHeader) {
	header.HttpHeader = src.HttpHeader
	header.IsNext = true
	header.MessageIndex = 0
	header.MessageUuid = src.MessageUuid
	header.Method = src.Method
	header.SenderId = "url_mapping"
	header.State = proto.MSAHeader_Success
	header.Type = proto.MSAHeader_General
	header.Url = src.Url
}

func makeMessageBody(levels []int) string {
	b,_ := json.Marshal(levels)
	return string(b)
}

func MakeOkMessage(origin *proto.MSAMessage,levels []int) proto.MSAMessage {
	var msg proto.MSAMessage
	msg.Header = new(proto.MSAHeader)
	writeOkHeader(msg.GetHeader(),origin.GetHeader())
	msg.HttpBody = origin.HttpBody
	msg.MessageBody = makeMessageBody(levels) 
	return msg
}