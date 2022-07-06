module url_mapping

go 1.18

replace core => ../../../pkgs/core/go

require core v0.0.0

require (
	github.com/doug-martin/goqu v5.0.0+incompatible // indirect
	github.com/mattn/go-sqlite3 v1.14.14 // indirect
)
