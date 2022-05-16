package service

import (
	"douyin/dao"
	"douyin/entity/request_entity"
	"douyin/errors_stuck"
	"log"
)

func FavoriteAction(video request_entity.Video) (ans bool, err error) {
	selsct, err := dao.SelectVideoById(video.VideoId)
	if err != nil {
		return false, err
	}

	if video.ActionType == 1 {
		selsct.FavoriteCount += 1
		videoList, err := dao.SelectFavoriteVideoIdListByUserId(video.UserId)
		if err != nil {
			log.Println(err)
			return ans, err
		}

		flag := false // 用于判断喜欢中是否含有视频
		for _, id := range videoList {
			if id == video.VideoId {
				flag = true
				break
			}
		}

		//如果找到 重复关注
		if flag {
			err = errors_stuck.AlreadyExists
			return false, err
		}
		dao.InsertFavorite(video.UserId, video.VideoId)

	} else {
		selsct.FavoriteCount -= 1
		//应处理是否重复错误

		dao.CancelFavoriteByUserIdAAndFeedID(video.UserId, video.VideoId)
	}

	go dao.UpDateVideo(selsct)

	return true, nil
}

func IsFavoriteByToken(token string, videoId int) bool {
	if token == "" {
		return false
	}
	id := dao.SelectUserIdByToken(token)
	return dao.IsFavorite(id, videoId)
}
