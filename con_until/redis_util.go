package con_until

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/redis_util"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
)

func GetFollowAndFollower(user entity.UserDao) (follow int, follower int) {
	//获取redis 连接
	connection := redis_util.GetConnection()
	defer redis_util.CloseConnection(connection)

	//检查是否存在key值
	exists, err := redis.Bool(connection.Do("EXISTS", strconv.Itoa(user.UserId)+"_FollowCount"))

	if err != nil {
		log.Println(err)
	}

	var followCount int
	var followerCount int

	if exists {
		//获取关注及被关注人数
		v1, err := redis.Int(connection.Do("GET", strconv.Itoa(user.UserId)+"_FollowCount"))
		followerCount = v1

		if err != nil {
			log.Println(err)
		}

		v2, err := redis.Int(connection.Do("GET", strconv.Itoa(user.UserId)+"_FollowerCount"))
		followerCount = v2

		if err != nil {
			log.Println(err)
		}

	} else {
		//从数据库中获取并 存入redis

		follow, err := dao.CountFollow(user.UserId)
		followCount = follow
		if err != nil {
			log.Println(err)
		}

		fans, err := dao.CountFans(user.UserId)
		followerCount = fans
		if err != nil {
			log.Println(err)
		}

		//写入值{"test-Key":"test-Value"}
		_, err = connection.Do("SET", user.UserName+"_FollowCount", followCount)
		_, err = connection.Do("SET", user.UserName+"_FollowerCount", followerCount)
		if err != nil {
			fmt.Println("redis set value failed >>>", err)
		}

	}
	return followCount, followerCount
}
