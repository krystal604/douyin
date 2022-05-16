package service

import (
	"douyin/dao"
	"douyin/entity"
	"sync"
)

func GetVideoFeed() (ans []entity.Video, err error) {
	videoList := dao.GetAllVideo()

	ans = make([]entity.Video, len(videoList))

	//开辟管道
	//缓冲区长度为list长度
	var ch chan entity.Video
	ch = make(chan entity.Video, len(videoList))
	//关闭管道
	defer close(ch)

	group := sync.WaitGroup{}
	mutex := sync.Mutex{}
	// 并发查询函数
	var goGetVideo func(video *entity.VideoDao)
	goGetVideo = func(video *entity.VideoDao) {

		author := GetUser(video.Author)
		temp := entity.Video{
			Id:            int64(video.Id),
			Author:        author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: int64(dao.CountVideoFavoriteByVideoId(video.Id)),
			CommentCount:  int64(dao.CountVideoCommentByVideoId(video.Id)),
			IsFavorite:    false, // 由于未登录 默认认为为点赞
		}

		ch <- temp
		group.Done()

	}

	for _, video := range videoList {
		//log.Println(runtime.NumGoroutine())
		group.Add(1)
		go goGetVideo(&video)
	}

	var add func()
	add = func() {
		mutex.Lock()
		temp := <-ch
		ans = append(ans, temp)
		ch <- temp
		mutex.Unlock()
	}

	for i := 0; i < len(videoList); i++ {
		go add()
	}

	group.Wait()
	return
}