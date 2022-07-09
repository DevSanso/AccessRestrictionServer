package sql_test

import (
	"testing"
	"url_mapping/sql"
)

func TestSelectUrlQuery(t *testing.T) {
	var eq = "SELECT \"url\", \"level\" FROM \"url_level\" WHERE (\"url\" = '/')"
	var query = sql.SelectUrlQuery("/")
	if eq != query {
		t.Error("SelectUrlQuery not match")
	}
}

func TestSelectUrlLevelsQuery(t *testing.T) {
	var eq = "SELECT \"level\" FROM \"url_level\" WHERE (\"url\" = '/index')"
	var query = sql.SelectUrlLevelsQuery("/index")
	if eq != query {
		t.Error("SelectUrlLevelsQuerynot match")
	}
}
