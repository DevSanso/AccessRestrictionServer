module accessRestrictionCore

go 1.18

replace core => ./

require (
	core v0.0.0
	github.com/confluentinc/confluent-kafka-go v1.9.1
	github.com/google/uuid v1.3.0
	google.golang.org/protobuf v1.28.0
)
