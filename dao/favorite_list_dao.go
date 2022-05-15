package dao

import (
	"database/sql"
	"douyin/dao/dao_config"
	"log"
)

func SelectFavoriteVideoIdListByUserId(id int) (ans []int, err error) {
	ans = make([]int, 0)

	db := dao_config.GetDatabase()
	sqlStr := "select feed_id from favorite_list where user_id = ? "
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
	if err != nil {
		log.Println(err)
		return
	}

	for query.Next() {
		var temp int
		err := query.Scan(&temp)

		if err != nil {
			log.Println(err)
			return ans, err
		}
		ans = append(ans, temp)
	}

	return
}

func InsertFavorite(userId int, feedId int) {
	db := dao_config.GetDatabase()
	sqlStr := "insert into favorite_list(favorite_id , user_id , feed_id ) values (?,?,?)"
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

	_, err = prepare.Exec(dao_config.AUTO_ID, userId, feedId)
	if err != nil {
		log.Println(err)
		return
	}

}

func CancelFavoriteByUserIdAAndFeedID(userId int, feedId int) {
	db := dao_config.GetDatabase()
	sqlStr := "delete from favorite_list where user_id=? and feed_id = ?"
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

	_, err = prepare.Exec(userId, feedId)
	if err != nil {
		log.Println(err)
		return
	}
}
