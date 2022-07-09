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
		result = append(result,obj)
	}
	if err := row.Err();err != nil {panic(err)}
	return result
}

func DecodeLevel(row *sql.Rows) []int{
	result :=  make([]int,0)
	for row.Next() {
		var level int
		err := row.Scan(&level)
		if err != nil {panic(err)}
		result = append(result,level)
	}
	if err := row.Err();err != nil {panic(err)}
	return result
}