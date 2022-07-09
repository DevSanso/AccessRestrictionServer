package message

import (
	"fmt"
	"core/proto"
)

func linkHeader(name string) *proto.MSAHeader {
	var header  = new(proto.MSAHeader)
	header.HttpHeader = ""
	header.IsNext = false
	header.MessageIndex = 0
	header.MessageUuid = ""
	header.Method = ""
	header.SenderId = name
	header.State = proto.MSAHeader_Success
	header.Type = proto.MSAHeader_LinkRequest
	header.Url = ""
	return header
}

func LinkMessage(name string,level int) proto.MSAMessage{
	var message proto.MSAMessage
	message.Header = linkHeader(name)
	message.MessageBody = fmt.Sprintf("%d",level)
	return message
}