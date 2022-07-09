package da

import (
	"database/sql"
	sqlQuery "url_mapping/sql"
	parser "url_mapping/parser/rows"
	"url_mapping/structure"
)

func SelectUrlLevel(db *sql.DB,url string) ([]structure.UrlLevelMapping,error) {
	rows,err :=  db.Query(sqlQuery.SelectUrlQuery(url))
	if err != nil {return nil,err}
	return parser.Decode(rows),nil
}