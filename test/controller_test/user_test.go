package controller_test

import (
	controller_until2 "douyin/con_until"
	"douyin/entity"
	"testing"
)

func TestSelectUserById(t *testing.T) {

	follow, follower := controller_until2.GetFollowAndFollower(entity.UserDao{
		UserId:       1,
		UserName:     "",
		UserPassWord: "",
		UserToken:    "",
	})

	t.Log("follow : ", follow)
	t.Log("follower : ", follower)
}
