package dao_config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:853963abbe11@tcp(47.93.2.242:3306)/douyin?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
