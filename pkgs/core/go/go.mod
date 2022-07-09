module accessRestrictionCore

go 1.18


replace core => ./
require core v0.0.0
require (
	github.com/confluentinc/confluent-kafka-go v1.9.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
