package main

import (
	"core/kafka"
	"url_mapping/config"
	"database/sql"
)

var (
	recvStream <-chan kafka.RecvMessage
	recvCloseFunc func()
	db *sql.DB
)

func init() {
	d,err := OpenDB(config.DataBaseSource)
	if err != nil {panic(err)}
	db = d
	
	kafka.InitProducer(config.KafkaAddr,config.KafkaGatewayTopic)
	rs,close := kafka.NewRecvStream(config.KafkaAddr,config.KafkaGroup,config.KafkaTopic)
	recvStream = rs
	recvCloseFunc = close
}

