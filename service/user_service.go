package service

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/errors_stuck"
	"log"
)

func Register(user entity.UserDao) (bool, error) {

	exist := dao.ExistUserName(user.UserName)
	if !exist {
		return false, errors_stuck.DoesNotExist
	}
	return dao.InsertUser(user), nil
}

func Login(userName string, userPassword string) (bool, error) {
	user, err := dao.SelectUserByName(userName)
	if err != nil {
		if err == errors_stuck.DoesNotExist {
			return false, err
		}
		log.Println(err)
		return false, err
	}
	if user.UserPassWord == userPassword {
		return true, nil
	} else {
		return false, errors_stuck.PassWordWrongs
	}
}

func GetUserByToken(token string) (entity.UserDao, error) {

	user, err := dao.SelectUserByToken(token)
	if err != nil {
		if err == errors_stuck.DoesNotExist {
			return entity.UserDao{}, err
		} else {
			log.Println(err)
		}
	}
	return user, nil
}

func GetUser(id int) entity.User {
	byId := dao.SelectUserById(id)

	user := entity.User{
		Id:            int64(byId.UserId),
		Name:          byId.UserName,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	return user
}
