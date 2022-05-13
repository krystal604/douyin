package dao

import (
	"database/sql"
	"douyin/dao/dao_config"
	"douyin/entity/request_entity"
	"douyin/errors_stuck"
	"log"
)

func SelectRelationIdByTowUser(userId int, toUserId int) (ans int, errs error) {
	db := dao_config.GetDatabase()
	sqlStr := "select fllow_id from fllow where user_id = ? and  user_fans_id = ?"
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

	query, err := prepare.Query(toUserId, userId)

	if err != nil {
		log.Println(err)
	}

	flag := false // judge is run once

	for query.Next() {
		flag = true
		err := query.Scan(&ans)

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

func AddRelation(request request_entity.RelationActionRequest) bool {
	db := dao_config.GetDatabase()
	sqlStr := "insert into fllow(fllow_id , user_id , user_fans_id ) values (?,?,?)"
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

	//对方为主用户 自己为粉丝
	_, err = prepare.Exec(dao_config.AUTO_ID, request.ToUserId, request.UserId)
	if err != nil {
		log.Println(err)
	}

	return true
}

func CancelRelation(id int) (ans bool, err error) {
	db := dao_config.GetDatabase()
	sqlStr := "delete from fllow where fllow_id=?"
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

	_, err = prepare.Exec(id)
	if err != nil {
		log.Println(err)
	}
	return true, err
}
