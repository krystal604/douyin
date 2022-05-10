package dao

import (
	"database/sql"
	"douyin/dao/dao_config"
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
	db := dao_config.GetDatabase()
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

func InsertUser(user User) (ans bool) {
	db := dao_config.GetDatabase()
	sqlStr := "insert into user(user_id , user_name , user_password , user_token ) values (?,?,?,?)"
	prepare, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err)
		return false
	}

	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			ans = false
			log.Println(err)
		}
	}(prepare)

	//执行SQL，填加站位值
	_, err = prepare.Exec(dao_config.AUTO_ID, user.userName, user.userPassWord, user.userToken)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func SelectUserByName(name string) (ans User, errs error) {
	db := dao_config.GetDatabase()
	sqlStr := "select * from user where user_name = ? "
	prepare, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println()
	}
	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			log.Println(err)

		}
	}(prepare)

	query, err := prepare.Query(name)
	if err != nil {
		log.Println(err)
	}

	if query == nil || !query.Next() {
		errs = dao_config.NilSelectError
		return
	}

	for query.Next() {
		err := query.Scan(&ans.userId, &ans.userName, &ans.userPassWord, &ans.userToken)
		if err != nil {
			log.Println(err)
		}
	}
	return
}
