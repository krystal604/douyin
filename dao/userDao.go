package dao

import (
	"database/sql"
	"douyin/dao/dao_config"
	"douyin/entity"
	"douyin/errors_stuck"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

//type UserDao interface {
//	SelectUserById(id int)
//}

func SelectUserById(id int) (ans entity.User) {
	db := dao_config.GetDatabase()
	query, err := db.Query("select * from user where user_id = " + strconv.Itoa(id) + ";")
	if err != nil {
		log.Fatal(err)
	}

	for query.Next() {
		err := query.Scan(&ans.UserId, &ans.UserName, &ans.UserPassWord, ans.UserToken)
		if err != nil {
			return
		}
	}
	fmt.Println(ans.UserName)
	return
}

func InsertUser(user entity.User) (ans bool) {
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
	_, err = prepare.Exec(dao_config.AUTO_ID, user.UserName, user.UserPassWord, user.UserToken)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func SelectUserByName(name string) (ans entity.User, errs error) {
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
			log.Println("dao")
		}
	}(prepare)

	query, err := prepare.Query(name)
	if err != nil {
		log.Println(err)
	}

	flag := false // judge is run once

	for query.Next() {
		flag = true
		err := query.Scan(&ans.UserId, &ans.UserName, &ans.UserPassWord, &ans.UserToken)

		if err != nil {
			log.Println(err)
			log.Println("dao")
		}

	}

	// 没有任何数据返回DoesNotExist异常
	if !flag {
		errs = errors_stuck.DoesNotExist
		return
	}
	return
}

func SelectUserByToken(token string) (ans entity.User, errs error) {
	db := dao_config.GetDatabase()
	sqlStr := "select * from user where user_token = ? "
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

	query, err := prepare.Query(token)
	if err != nil {
		log.Println(err)
	}

	flag := false
	for query.Next() {
		flag = true
		err := query.Scan(&ans.UserId, &ans.UserName, &ans.UserPassWord, &ans.UserToken)

		if err != nil {
			log.Println(err)
		}

	}
	//如果没有进入循环 则没有查询到数据
	if !flag {
		errs = errors_stuck.DoesNotExist
	}
	return
}

/**
判断是否已经含有某个 名字
有返回true
*/
func ExistUserName(name string) bool {
	db := dao_config.GetDatabase()
	sqlStr := "SELECT count(*) FROM user WHERE user_name = ?"
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
		return true
	}
	var count int
	err = query.Scan(&count)
	if err != nil {
		return true
	}

	if count >= 1 {
		return true
	}

	return false
}
