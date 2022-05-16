package controller

import (
	"douyin/entity"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []entity.Video `json:"video_list"`
	NextTime  int64          `json:"next_time"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	token := c.Query("token")

	feed, err := service.GetVideoFeed(token)
	if err != nil {
		return
	}

	index := 0
	for _, video := range feed {
		if video.Id == 0 {
			index++
		} else {
			break
		}
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: feed[index:],
		NextTime:  time.Now().Unix(),
	})

}
