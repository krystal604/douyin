package dao

import (
	"database/sql"
	"douyin/dao/dao_config"
	"douyin/entity"
	"log"
)

func CountVideoCommentByVideoId(id int) int {
	return 0
}

func InsertComment(comment entity.CommentDao) {
	db := dao_config.GetDatabase()
	sqlStr := "insert into comment values (?,?,?,?)"
	prepare, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err)
		return
	}

	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			log.Println(err)
		}
	}(prepare)

	_, err = prepare.Exec(dao_config.AUTO_ID, comment.UserId, comment.VideoId, comment.ContentText)
	if err != nil {
		return
	}

}

func DeleteCommentById(id int) {
	db := dao_config.GetDatabase()
	sqlStr := "delete from comment where comment_id=?"
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
}

func SelectCommentListByVideoId(id int) (ans []entity.CommentDao) {

	ans = make([]entity.CommentDao, 0)

	db := dao_config.GetDatabase()
	sqlStr := "select * from comment where feed_id = ?"
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
		temp := entity.CommentDao{}
		query.Scan(&temp.Id, &temp.UserId, &temp.VideoId, &temp.ContentText)
		ans = append(ans, temp)
	}

	return
}
