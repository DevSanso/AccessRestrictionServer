package kafka_test

import (
	"core/proto"
	"testing"
	"url_mapping/config"
	"url_mapping/kafka"
	"url_mapping/utils/tests"
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
