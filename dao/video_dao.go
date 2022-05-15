package dao

import (
	"database/sql"
	"douyin/dao/dao_config"
	"douyin/entity"
	"douyin/errors_stuck"
	"log"
)

func SelectVideoById(id int) (ans entity.VideoDao, err error) {
	db := dao_config.GetDatabase()
	sqlStr := "select * from feed where feed_id = ? "
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

	flag := false
	for query.Next() {
		flag = true
		err := query.Scan(&ans.Id, &ans.PlayUrl, &ans.CoverUrl, &ans.FavoriteCount)

		if err != nil {
			log.Println(err)
			return ans, err
		}

	}
	//如果没有进入循环 则没有查询到数据
	if !flag {
		err = errors_stuck.DoesNotExist
		return
	}
	return
}

func UpDateVideo(video entity.VideoDao) {
	db := dao_config.GetDatabase()
	sqlStr := "update feed set feed_favorite_count = ? where feed_id = ? "
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

	_, err = prepare.Exec(video.FavoriteCount, video.Id)
	if err != nil {
		return
	}
}
