package controller

import (
	"douyin/controller/controller_until"
	"douyin/entity"
	"douyin/errors_stuck"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
	User entity.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token := username + password
	//
	user := entity.UserDao{
		UserId:       0,
		UserName:     username,
		UserPassWord: password,
		UserToken:    token,
	}
	register, err := service.Register(user)

	// 如果 失败
	if !register || err != nil {
		if err != nil {
			if err == errors_stuck.DoesNotExist {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
				})
				return
			}
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "false"},
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		Token:    username + password,
	})

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	login, err := service.Login(username, password)
	if err != nil {
		if err == errors_stuck.PassWordWrongs {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "password is wrong"},
			})
		} else if err == errors_stuck.DoesNotExist {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
		} else {
			log.Println(err)
			return
		}
	}
	if login {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			Token:    token,
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	byToken, err := service.GetUserByToken(token)

	fmt.Println(byToken)
	if err != nil {
		if err == errors_stuck.DoesNotExist {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
			return
		} else {
			log.Println(err)
			return
		}
	}

	// 利用redis 查询countFollow
	followCount, followerCount := controller_until.GetFollowAndFollower(byToken)

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User: entity.User{
			Id:            int64(byToken.UserId),
			Name:          byToken.UserName,
			FollowCount:   int64(followCount),
			FollowerCount: int64(followerCount),
			IsFollow:      false,
		},
	})

}
