package kafka_test

import (
	"core/kafka"
	"testing"
)

func TestRecvStreamConnect(t *testing.T) {
	_, close := kafka.NewRecvStream("localhost:9092","test","test")
	close()

}
