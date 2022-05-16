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
		err := query.Scan(&ans.Id, &ans.PlayUrl, &ans.CoverUrl, &ans.FavoriteCount, &ans.Author)

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

func InsertVideo(video entity.VideoDao) {
	db := dao_config.GetDatabase()
	sqlStr := "insert into feed values (?,?,?,?,?)"
	prepare, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err)
	}

	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			log.Println(err)
		}
	}(prepare)

	_, err = prepare.Exec(dao_config.AUTO_ID, video.PlayUrl, video.CoverUrl, 0, video.Author)
	if err != nil {
		return
	}

}

func GetAllVideo() []entity.VideoDao {
	ans := make([]entity.VideoDao, 0)

	db := dao_config.GetDatabase()
	sqlStr := "select * from feed "
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

	query, err := prepare.Query()
	if err != nil {
		return nil
	}

	for query.Next() {
		var temp entity.VideoDao
		err := query.Scan(&temp.Id, &temp.PlayUrl, &temp.CoverUrl, &temp.FavoriteCount, &temp.Author)
		if err != nil {
			log.Println(err)
			return nil
		}
		ans = append(ans, temp)
	}

	return ans
}
