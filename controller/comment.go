package controller

import (
	"douyin/entity"
	"douyin/entity/request_entity"
	"douyin/errors_stuck"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []entity.Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	//token := c.Query("token")

	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}

	//userId := c.PostForm("user_id")
	//token := c.PostForm("token")
	//videoId := c.PostForm("video_id")
	//actionType := c.PostForm("action_type")
	//commentText := c.PostForm("comment_text")
	//commentId := c.PostForm("comment_id")

	userId := c.Query("user_id")
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentId := c.Query("comment_id")

	userIdAtoi, err := strconv.Atoi(userId)
	videoIdAtoi, err := strconv.Atoi(videoId)
	actionTypeAtoi, err := strconv.Atoi(actionType)
	commentIdAtoi, err := strconv.Atoi(commentId)
	if err != nil {
		log.Println("atoi : ", err)
	}
	comment := request_entity.Comment{
		UserId:      userIdAtoi,
		Token:       token,
		VideoId:     videoIdAtoi,
		ActionType:  actionTypeAtoi,
		ContentText: commentText,
		ContentId:   commentIdAtoi,
	}

	err = service.CommentAction(comment)
	if err != nil {
		if err == errors_stuck.NoAction {
			c.JSON(http.StatusOK, Response{StatusCode: 3, StatusMsg: "no this action "})
			return
		}

		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "don't know what happened "})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0})

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {

	videoId := c.Query("video_id")

	videoIdAtoi, err := strconv.Atoi(videoId)
	if err != nil {
		return
	}

	list := service.GetCommentList(videoIdAtoi)

	index := 0
	for _, comment := range list {
		if comment.Id == 0 {
			index++
		} else {
			break
		}
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: list[index:],
	})
}
