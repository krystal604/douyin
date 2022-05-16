package service

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/entity/request_entity"
	"douyin/errors_stuck"
)

func CommentAction(comment request_entity.Comment) error {

	if comment.ActionType == 1 {
		dao.InsertComment(entity.CommentDao{
			Id:          comment.ContentId,
			UserId:      comment.UserId,
			VideoId:     comment.VideoId,
			ContentText: comment.ContentText,
		})
	} else if comment.ActionType == 2 {
		dao.DeleteCommentById(comment.ContentId)
	} else {
		return errors_stuck.NoAction
	}

	return nil
}

func GetCommentList(id int) (ans []entity.Comment) {
	commentList := dao.SelectCommentListByVideoId(id)
	ans = make([]entity.Comment, len(commentList))

	for _, commentDao := range commentList {
		temp := entity.Comment{
			Id:         int64(commentDao.Id),
			User:       GetUser(commentDao.UserId),
			Content:    commentDao.ContentText,
			CreateDate: "05-16",
		}
		ans = append(ans, temp)
	}

	return
}
