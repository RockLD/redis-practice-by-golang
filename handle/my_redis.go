package handle

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type MyRedis struct {
	C redis.Conn
}

func (mr *MyRedis) InitRedis() {
	var err error
	mr.C, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Redis conntect is error ", err)
		return
	}
	_, err = mr.C.Do("SELECT", 0)
	if err != nil {
		fmt.Println("Select db is error ", err)
		return
	}
}

func (mr *MyRedis) Close() {
	err := mr.C.Close()
	if err != nil {
		fmt.Println("Redis close is error ", err)
		return
	}
	return
}
