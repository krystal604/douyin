package service

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/entity/request_entity"
	"douyin/errors_stuck"
	"log"
)

func RelationActionAdd(request request_entity.RelationActionRequest) (bool, error) {
	_, err := dao.SelectRelationIdByTowUser(request.UserId, request.ToUserId)
	// 如果他不是不存在异常 那么证明 重复了
	if err != errors_stuck.DoesNotExist {
		return false, errors_stuck.AlreadyExists
	}

	relation := dao.AddRelation(request)
	return relation, nil
}

func RelationActionCancel(request request_entity.RelationActionRequest) (ans bool, err error) {
	id, err := dao.SelectRelationIdByTowUser(request.UserId, request.ToUserId)

	if err != nil {
		if err == errors_stuck.DoesNotExist {
			return
		}
		log.Println(err)
	}

	relation, err := dao.CancelRelation(id)
	if err != nil {
		log.Println(err)
	}

	return relation, nil
}

func GetFollowList(id int) (userList []entity.User, err error) {
	//先查询followUserList 以便确认 返回切片长度
	followUserList, err := dao.SelectFollowList(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//followUserList 初始化
	userList = make([]entity.User, len(followUserList))

	for i := 0; i < len(followUserList); i++ {
		user := dao.SelectUserById(followUserList[i])

		userList = append(userList, entity.User{
			Id:            int64(user.UserId),
			Name:          user.UserName,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		})
	}

	return
}
