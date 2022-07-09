package kafka_test

import (
	"core/kafka"
	"testing"
)

func TestRecvStreamConnect(t *testing.T) {
	_, close := kafka.RecvStream()
	close()

}
