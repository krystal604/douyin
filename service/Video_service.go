package service

import (
	"douyin/dao"
	"douyin/entity"
	"log"
	"sync"
)

func GetVideoFeed(token string) (ans []entity.Video, err error) {
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
	var goGetVideo func(video entity.VideoDao)
	goGetVideo = func(video entity.VideoDao) {
		author := GetUser(video.Author)
		temp := entity.Video{
			Id:            int64(video.Id),
			Author:        author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: int64(dao.CountVideoFavoriteByVideoId(video.Id)),
			CommentCount:  int64(dao.CountVideoCommentByVideoId(video.Id)),
			IsFavorite:    IsFavoriteByToken(token, video.Id), // 由于未登录 默认认为为点赞
		}

		ch <- temp

		//mutex.Lock()
		//ans = append(ans , temp)
		//mutex.Unlock()
		group.Done()
	}

	for _, video := range videoList {
		//log.Println(runtime.NumGoroutine())
		group.Add(1)
		go goGetVideo(video)
	}

	var add func()
	add = func() {
		mutex.Lock()
		temp := <-ch
		ans = append(ans, temp)
		mutex.Unlock()
		group.Done()
	}
	group.Wait()
	for i := 0; i < len(videoList); i++ {
		group.Add(1)
		go add()
	}

	group.Wait()

	return
}

func GetPublishVideoList(id int) (ans []entity.Video, err error) {
	ans = make([]entity.Video, 0)

	feed, err := GetVideoFeed("")
	if err != nil {
		log.Println(err)
	}

	for _, video := range feed {
		if video.Author.Id == int64(id) {
			ans = append(ans, video)
		}
	}
	return
}
