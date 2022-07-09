package kafka

import (
	"sync"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"core/parser/protobuf"
	"core/proto"
)

var producer *kafka.Producer
var once sync.Once
var gatewayTopic string

func initProducer(addr string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": addr,
	})
	if err != nil {
		panic(err)
	}
	producer = p
}
func InitProducer(addr,topic string) {
	once.Do(func() {
		initProducer(addr)
		gatewayTopic = topic
	})
}

func makeKafkaMessage(message []byte) *kafka.Message {
	return &kafka.Message{
		Value:          message,
		TopicPartition: kafka.TopicPartition{Topic: &gatewayTopic, Partition: kafka.PartitionAny},
	}
}
func SendMessage(message *proto.MSAMessage) error {
	defer producer.Flush(1000)
	data, err := protobuf.DecodeMessage(message)
	if err != nil {
		return err
	}

	return producer.Produce(makeKafkaMessage(data), nil)
}
