package protobuf

import parser "google.golang.org/protobuf/proto"
import  "core/proto"



func EncodeHeader(data []byte) (proto.MSAHeader,error) {
	var buf = proto.MSAHeader{}
	err := parser.Unmarshal(data,&buf)
	return buf,err
}

func DecodeHeader(header *proto.MSAHeader) ([]byte,error) {
	return parser.Marshal(header)
}

