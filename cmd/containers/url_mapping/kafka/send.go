package kafka

import (
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"core/proto"
	"url_mapping/config"
	"url_mapping/parser/protobuf"
)
var producer *kafka.Producer
var once sync.Once

func initProducer() {
	p,err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaAddr, 
	});
	if err != nil {
		panic(err)
	}
	producer = p
}

func makeKafkaMessage(message []byte) *kafka.Message {
	return &kafka.Message{
		Value : message,
		TopicPartition: kafka.TopicPartition{Topic : &config.KafkaGatewayTopic,Partition: kafka.PartitionAny},
	}
}
func SendMessage(message *proto.MSAMessage) error {
	once.Do(initProducer)
	defer producer.Flush(1000)
	data,err := protobuf.DecodeMessage(message)
	if err != nil {return err}

	return producer.Produce(makeKafkaMessage(data),nil)
}