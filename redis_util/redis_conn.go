package redis_util

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func GetConnection() redis.Conn {
	conn, err := redis.Dial("tcp",
		"r-2zei6m37sphtqeie68pd.redis.rds.aliyuncs.com:6379",
		redis.DialDatabase(1), //dialOption参数可以配置选择数据库、连接密码、心跳检测等等
		redis.DialPassword("Zgj19802480724"))
	if err != nil {
		fmt.Println("Connect to redis failed ,cause by >>>", err)
		return nil
	}
	return conn
}

func CloseConnection(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		log.Println("CloseConnection >>>>>", err)
		return
	}
}
