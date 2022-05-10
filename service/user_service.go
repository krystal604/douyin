package service

import (
	"douyin/dao"
	"douyin/entity"
	"errors"
)

func Register(user entity.User) (bool, error) {

	exist := dao.ExistUserName(user.UserName)
	if exist {
		return false, errors.New("exist")
	}
	return dao.InsertUser(user), nil
}
