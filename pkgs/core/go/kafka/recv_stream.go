package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"core/proto"

	"core/parser/protobuf"
)

type RecvMessage struct {
	Message proto.MSAMessage
	Err     error
}

func mkKafkaConfig(addr,group string) *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": addr,
		"group.id":          group,
	}
}

func initRunningFlag() *bool {
	loopFlag := new(bool)
	*loopFlag = true
	return loopFlag
}

func loop(runningflag *bool, c *kafka.Consumer, channel chan RecvMessage) {
	for *runningflag {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			channel <- RecvMessage{Err: err}
			continue
		}
		encodeMsg, encodeErr := protobuf.EncodeMessage(msg.Value)
		channel <- RecvMessage{
			encodeMsg,
			encodeErr,
		}
	}
	c.Close()
}

func NewRecvStream(addr,group,topic string) (channel <-chan RecvMessage, close func()) {
	c, cerr := kafka.NewConsumer(mkKafkaConfig(addr,group))
	if cerr != nil {
		panic(cerr)
	}

	c.SubscribeTopics([]string{topic}, nil)

	loopFlag := initRunningFlag()
	recvChan := make(chan RecvMessage)

	go loop(loopFlag, c, recvChan)

	channel = recvChan
	close = func() { *loopFlag = false }

	return
}
