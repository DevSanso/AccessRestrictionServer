package main

import (
	"log"
	"core/proto"
	"url_mapping/da"
)

func getMsgHeaderUrl(msg *proto.MSAMessage) string {
	return msg.GetHeader().GetUrl()
}

func main() {
	log.Println("Start Url Mapping Container")

	for msg := range recvStream {
		if msg.Err != nil {log.Fatalln(msg.Err)}

		url := getMsgHeaderUrl(&msg.Message)
		levels,err := da.SelectUrlLevel(db,url)

		if err != nil {
			
		}
	}
}