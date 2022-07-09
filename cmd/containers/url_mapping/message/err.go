package message


import (
	"core/proto"
)

func errHeader(wHeader,src *proto.MSAHeader) {
	wHeader.IsNext = false
	wHeader.State = proto.MSAHeader_Error
	wHeader.Type = proto.MSAHeader_General
	
}

func MakeErrorMessage(origin *proto.MSAMessage,err error) {
	var dst  proto.MSAMessage
	errHeader(dst.GetHeader(),origin.GetHeader())

	
}