package service

import (
	"github.com/garyburd/redigo/redis"
)

var CONN

func NewRedis()  {
	conn, err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis connect error :",err)
		return
	}

}
