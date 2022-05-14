package controller_test

import (
	"douyin/controller/controller_until"
	"douyin/entity"
	"testing"
)

func TestSelectUserById(t *testing.T) {

	follow, follower := controller_until.GetFollowAndFollower(entity.UserDao{
		UserId:       1,
		UserName:     "",
		UserPassWord: "",
		UserToken:    "",
	})

	t.Log("follow : ", follow)
	t.Log("follower : ", follower)
}
