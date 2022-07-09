package main

import (
	"core/kafka"
	"core/proto"
	"log"
	"url_mapping/da"
	"url_mapping/message"
)

func getMsgHeaderUrl(msg *proto.MSAMessage) string {
	return msg.GetHeader().GetUrl()
}

func sendErrMessage(origin *proto.MSAMessage,err error) error {
	errMsg := message.MakeErrorMessage(origin,err)
	return kafka.SendMessage(&errMsg)
}

func sendOkMessage(origin *proto.MSAMessage,levels []int)error {
	msg := message.MakeOkMessage(origin,levels)
	return kafka.SendMessage(&msg)
}

func main() {
	log.Println("Start Url Mapping Container")

	for msg := range recvStream {
		if msg.Err != nil {log.Fatalln(msg.Err)}

		url := getMsgHeaderUrl(&msg.Message)
		levels,err := da.SelectUrlLevel(db,url)

		if err != nil {
			err = sendErrMessage(&msg.Message,err)
			if err != nil {log.Fatalln(err)}
			continue
		}

		if err = sendOkMessage(&msg.Message,levels);err != nil {
			log.Fatalln(err)
		}
	}
}