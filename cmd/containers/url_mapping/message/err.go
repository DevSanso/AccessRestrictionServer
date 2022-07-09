package message


import (
	"core/proto"
)

func writeErrHeader(wHeader,src *proto.MSAHeader) {
	wHeader.IsNext = false
	wHeader.State = proto.MSAHeader_Error
	wHeader.Type = proto.MSAHeader_General
	wHeader.SenderId = "url_mapping"
	wHeader.MessageIndex = 0
	wHeader.MessageUuid = src.GetMessageUuid()
	wHeader.Method = src.Method
}

func MakeErrorMessage(origin *proto.MSAMessage,err error) proto.MSAMessage {
	var dst  proto.MSAMessage
	dst.Header = new(proto.MSAHeader)
	writeErrHeader(dst.GetHeader(),origin.GetHeader())
	dst.MessageBody = err.Error()
	return dst
	
}