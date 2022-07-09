package kafka_test

import (
	"core/config"
	"core/kafka"
	"core/proto"
	"core/utils/tests"
	"testing"
)

func TestSend(t *testing.T) {
	config.InitEnv()

	header := tests.NewHeader()
	var msg proto.MSAMessage
	msg.Header = header

	err := kafka.SendMessage(&msg)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

}
