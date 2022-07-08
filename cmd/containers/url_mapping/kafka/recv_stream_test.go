package kafka_test


import (
	"testing"
	"url_mapping/kafka"

)

func TestRecvStreamConnect(t *testing.T) {
	_ ,close := kafka.RecvStream()
	close()
	
}