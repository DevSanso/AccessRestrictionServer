package rows


import (
	"database/sql"
	"url_mapping/structure"
)


func Decode(row *sql.Rows) []structure.UrlLevelMapping {
	result :=  make([]structure.UrlLevelMapping,0)
	for row.Next() {
		obj := structure.UrlLevelMapping{}
		err := row.Scan(&obj.Url,&obj.Level)
		if err != nil {panic(err)}
	}
	if err := row.Err();err != nil {panic(err)}
	return result
}