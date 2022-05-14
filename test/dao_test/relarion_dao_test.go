package dao

import (
	"douyin/dao"
	"testing"
)

func TestCountFollow(t *testing.T) {

	follow, err := dao.CountFollow(1)
	if err != nil {
		t.Log(err)
	}

	t.Log("关注人数为", follow)
}

func TestCountFans(t *testing.T) {

	follow, err := dao.CountFans(2)
	if err != nil {
		t.Log(err)
	}

	t.Log("粉丝数为", follow)
}
