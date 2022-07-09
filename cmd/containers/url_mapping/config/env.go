package config

import (
	"os"
)


var (
	KafkaAddr = os.Getenv("ACCESS_RESTRICTION_KAFKA_ADDR")


	KafkaGatewayTopic = os.Getenv("ACCESS_RESTRICTION_GATEWAY_TOPIC")

	KafkaGroup = os.Getenv("ACCESS_RESTRICTION_URL_MAPPING_GROUP")
	KafkaTopic = os.Getenv("ACCESS_RESTRICTION_URL_MAPPING_TOPIC")

	DataBaseSource = os.Getenv("ACCESS_RESTRICTION_URL_MAPPING_DB_SOURCE")
)




