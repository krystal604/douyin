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

type UserListResponse struct {
	Response
	UserList []entity.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {

	userId := c.PostForm("user_id")
	token := c.PostForm("token")
	toUserId := c.PostForm("to_user_id")
	actionType := c.PostForm("action_type")

	actionTypeAutoi, err := strconv.Atoi(actionType)
	if err != nil {
		log.Println(err)
	}

	userIdAutoi, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)
	}

	toUserIdAutoi, err := strconv.Atoi(toUserId)
	if err != nil {
		log.Println(err)
	}

	request := request_entity.RelationActionRequest{
		UserId:     userIdAutoi,
		Token:      token,
		ToUserId:   toUserIdAutoi,
		ActionType: actionType,
	}

	flag := false

	// type 1 关注 2 取消关注
	if actionTypeAutoi == 1 {
		//添加
		add, err := service.RelationActionAdd(request)

		//err处理
		//分支内的err传不到外面  和下面写了两次
		if err != nil {
			if err == errors_stuck.AlreadyExists {
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "this relation is already exists"})
			}
			log.Println(err)
			return
		}

		flag = add

	} else {
		//取消
		cancel, err := service.RelationActionCancel(request)

		if err != nil {
			if err == errors_stuck.AlreadyExists {
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "this relation is already exists"})
			}
			log.Println(err)
			return
		}
		flag = cancel
	}

	if err != nil {
		if err == errors_stuck.AlreadyExists {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "this relation is already exists"})
		}
		log.Println(err)
		return
	}

	// 如果返回的是false 操作失败
	if !flag {
		//code 非0 代表失败
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "false Don't know what happened"})
		return
	}

	//成功 code 0
	c.JSON(http.StatusOK, Response{StatusCode: 0})

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []entity.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {

	userId := c.Query("user_id")
	userIdAtoi, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)
		return
	}

	list, err := service.GetFollowList(userIdAtoi)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: list,
	})

}
