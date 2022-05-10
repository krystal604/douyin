package service

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/errors_stuck"
	"errors"
	"log"
)

func Register(user entity.User) (bool, error) {

	exist := dao.ExistUserName(user.UserName)
	if exist {
		return false, errors.New("exist")
	}
	return dao.InsertUser(user), nil
}

func Login(userName string, userPassword string) (bool, error) {
	user, err := dao.SelectUserByName(userName)
	if err != nil {
		log.Println(err)
		if err == errors_stuck.DoesNotExist {
			return false, err
		}
		return false, err
	}
	if user.UserPassWord == userPassword {
		return true, nil
	} else {
		return false, errors_stuck.PassWordWrongs
	}
}
