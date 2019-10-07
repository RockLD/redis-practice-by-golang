package handle

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Expire struct {
	T int
}

func (mr *MyRedis) SetStr(key, value string, expire Expire) (interface{}, error) {
	action := "SET"
	fmt.Println("action = ", action)
	var reply interface{}
	var err error
	if expire.T != 0 {
		reply, err = mr.C.Do(action, key, value, "EX", expire.T)
	} else {
		reply, err = mr.C.Do(action, key, value)
	}

	if err != nil {
		return "", err
	}
	return reply, nil
}

func (mr *MyRedis) GetStr(key string) (string, error) {
	res, err := redis.String(mr.C.Do("GET", key))
	if err != nil {
		return "", err
	}
	return res, nil
}

func (mr *MyRedis) ExistKey(key string) bool {
	res, err := redis.Int64(mr.C.Do("EXISTS", key))
	if err != nil {
		fmt.Println("exists error is ", err)
		return false
	}
	fmt.Println("exists result is ", res)
	if res != 1 {
		return false
	}

	return true
}

func (mr *MyRedis) DelKey(key string) (bool, error) {
	res, err := mr.C.Do("DEL", key)
	if err != nil {
		return false, err
	}
	fmt.Println("del result is ", res)
	return true, nil
}
