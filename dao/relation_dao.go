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

func SelectFollowList(id int) (idList []int, err error) {
	//idList 切片初始化
	idList = make([]int, 0)

	db := dao_config.GetDatabase()
	sqlStr := "select user_fans_id from fllow where user_id = ? "
	prepare, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println()
		return nil, err
	}
	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			log.Println(err)

		}
	}(prepare)

	query, err := prepare.Query(id)
	if err != nil {
		log.Println(err)
	}

	for query.Next() {
		var temp int
		err := query.Scan(&temp)
		if err != nil {
			return nil, err
		}
		idList = append(idList, temp)
	}

	return idList, err
}
func CountFans(id int) (count int, err error) {
	db := dao_config.GetDatabase()

	//粉丝有多少就查主用户次数 即被关注多少次

	sqlStr := "SELECT count(*) FROM fllow WHERE user_id = ?"
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

	query, err := prepare.Query(id)

	for query.Next() {
		err := query.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return
}

func CountFollow(id int) (count int, err error) {
	db := dao_config.GetDatabase()

	//关注多少人就是多少人的粉丝

	sqlStr := "SELECT count(*) FROM fllow WHERE user_fans_id = ?"
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

	query, err := prepare.Query(id)

	for query.Next() {
		err := query.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return
}
