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
	VideoList []entity.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	feed, err := service.GetVideoFeed()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: feed[1:],
		NextTime:  time.Now().Unix(),
	})

}
