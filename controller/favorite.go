package controller

import (
	"douyin/entity/request_entity"
	"douyin/errors_stuck"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	//token := c.Query("token")
	//
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
	//userId := c.PostForm("user_id")
	//token := c.PostForm("token")
	//videoId := c.PostForm("video_id")
	//actionType := c.PostForm("action_type")

	userId := c.Query("user_id")
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	userIdAtoi, err := strconv.Atoi(userId)
	videoIdAtoi, err := strconv.Atoi(videoId)
	actionTypeAtoi, err := strconv.Atoi(actionType)
	if err != nil {
		log.Println("Atoi err ", err)
	}

	action, err := service.FavoriteAction(request_entity.Video{
		UserId:     userIdAtoi,
		Token:      token,
		VideoId:    videoIdAtoi,
		ActionType: actionTypeAtoi,
	})
	if err != nil {
		log.Println(err)

		if err == errors_stuck.AlreadyExists {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "this is already favorite"})
		}
		return
	}

	if action {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
