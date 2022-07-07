package kafka

import (
	
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"core/proto"

	"url_mapping/parser/protobuf"
	"url_mapping/config"
)

type RecvMessage struct {
	message proto.MSAMessage
	err error
}

func mkKafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": config.KafkaHost,
		"group.id": config.KafkaGroup,
	}
}

func initRunningFlag() *bool {
	loopFlag := new(bool)
	*loopFlag = true
	return loopFlag
}

func loop(runningflag *bool,c *kafka.Consumer,channel chan RecvMessage) {
	for *runningflag {
		msg ,err := c.ReadMessage(-1)
		if err != nil {
			channel <- RecvMessage{err : err}
			continue
		}
		encodeMsg,encodeErr := protobuf.EncodeMessage(msg.Value)
		channel <- RecvMessage{
			encodeMsg,
			encodeErr,
		}
	}
}

func RecvStream() (channel <-chan RecvMessage,close func()) {
	c,cerr := kafka.NewConsumer(mkKafkaConfig())
	if cerr != nil {
		panic(cerr)
	}

	c.SubscribeTopics([]string{config.KafkaTopic},nil)

	loopFlag := initRunningFlag()
	recvChan := make(chan RecvMessage)

	go loop(loopFlag,c,recvChan)

	channel = recvChan
	close = func() {*loopFlag = false}

	return
}