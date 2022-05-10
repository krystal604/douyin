package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type User struct {
	userId int
	userName string
	userPassWord string
	userToken int
}

type UserDao interface {
	SelectUserById(id int)
}


func (u *User) SelectUserById(id int) (ans User){
	db, err := sql.Open("mysql" , "root:853963abbe11@tcp(47.93.2.242:3306)/douyin?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	query, err := db.Query("select * from user where id = " + strconv.Itoa(id) + ";")
	if err != nil {
		log.Fatal(err)
	}

	for query.Next() {
		err := query.Scan(&ans)
		if err != nil {
			return User{}
		}
	}
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
