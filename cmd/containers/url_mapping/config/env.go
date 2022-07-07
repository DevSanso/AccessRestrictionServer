package config

import (
	"os"
)


var (
	KafkaHost = os.Getenv("ACCESS_RESTRICTION_KAFKA_HOST")
	KafkaPort = os.Getenv("ACCESS_RESTRICTION_KAFKA_PORT")

	KafkaGatewayTopic = os.Getenv("ACCESS_RESTRICTION_GATEWAY_TOPIC")

	KafkaGroup = os.Getenv("ACCESS_RESTRICTION_URL_MAPPING_GROUP")
	KafkaTopic = os.Getenv("ACCESS_RESTRICTION_URL_MAPPING_TOPIC")
)