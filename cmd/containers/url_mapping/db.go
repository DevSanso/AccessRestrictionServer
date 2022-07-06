package main


import	"database/sql"
import	_ "github.com/mattn/go-sqlite3"



func OpenDB(source string) (*sql.DB,error) {
	return sql.Open("sqlite3",source)
}