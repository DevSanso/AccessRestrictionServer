package parser

import parser "google.golang.org/protobuf/proto"
import  "core/proto"


func EncodeMessage(data []byte) (proto.MSAMessage,error) {
	var buf = proto.MSAMessage{}
	err := parser.Unmarshal(data,&buf)
	return buf,err
}

func DecodeMessage(mssage *proto.MSAMessage) ([]byte,error) {
	return parser.Marshal(mssage)
}

