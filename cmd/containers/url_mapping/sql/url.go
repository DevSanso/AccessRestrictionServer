package sql

import (
	"github.com/doug-martin/goqu"
)


func SelectUrlQuery(url string) string {
	s,_,_ := goqu.From("url_level").Select("url","level").Where(goqu.Ex{"url" : url}).ToSql()
	return s
}

func SelectUrlLevelsQuery(url string) string {
	s,_,_ := goqu.From("url_level").Select("level").Where(goqu.Ex{"url" : url}).ToSql()
	return s
}