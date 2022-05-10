package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type User struct {
	userId       int
	userName     string
	userPassWord string
	userToken    int
}

//type UserDao interface {
//	SelectUserById(id int)
//}

func SelectUserById(id int) (ans User) {
	db, err := sql.Open("mysql", "root:853963abbe11@tcp(47.93.2.242:3306)/douyin?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	query, err := db.Query("select * from user where user_id = " + strconv.Itoa(id) + ";")
	if err != nil {
		log.Fatal(err)
	}

	for query.Next() {
		err := query.Scan(&ans.userId, &ans.userName, &ans.userPassWord, ans.userToken)
		if err != nil {
			return
		}
	}
	fmt.Println(ans.userName)
	return
}

//func main() {
//	db, err := sql.Open("mysql" , "root:853963abbe11@tcp(47.93.2.242:3306)/douyin?charset=utf8")
//	if err != nil {
//		log.Fatal(err)
//	}
//	query, err := db.Query("select * from user where user_id = '1'")
//	if err != nil {
//		log.Fatal(err)
//	}
//	columns, err := query.Columns()
//	fmt.Println(columns )
//}
