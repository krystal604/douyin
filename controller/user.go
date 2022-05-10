package controller

import (
	"douyin/entity"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	register, err := service.Register(entity.User{
		UserId:       0,
		UserName:     username,
		UserPassWord: password,
		UserToken:    0,
	})

	// 如果 失败
	if !register && err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			Token:    username + password,
		})
	}

}

func Login(c *gin.Context) {
	//username := c.Query("username")
	//password := c.Query("password")

	//token := username + password

	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 0},
	//		UserId:   user.Id,
	//		Token:    token,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}

func UserInfo(c *gin.Context) {
	//token := c.Query("token")
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}
