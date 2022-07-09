module url_mapping

go 1.18

replace core => ../../../pkgs/core/go

require core v0.0.0

require (
	github.com/doug-martin/goqu v5.0.0+incompatible
	github.com/mattn/go-sqlite3 v1.14.14
)

require (
	github.com/c2fo/testify v0.0.0-20150827203832-fba96363964a // indirect
	github.com/confluentinc/confluent-kafka-go v1.9.1 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0 // indirect
	gopkg.in/doug-martin/goqu.v5 v5.0.0 // indirect
)
