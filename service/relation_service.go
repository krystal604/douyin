package service

import (
	"douyin/dao"
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
