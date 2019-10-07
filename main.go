package main

import (
	"fmt"
	"redis-by-go/handle"
)

func main() {

	var mr handle.MyRedis
	mr.InitRedis()

	defer mr.Close()
	var key = "test-redigo111"
	is, err := mr.DelKey(key)
	if err != nil {
		fmt.Println("del err", err)
		return
	}
	fmt.Println(is)

	/*var ex handle.Expire
	ex.T = 60
	_,err := mr.SetStr(key,"tttt-1",ex)
	if err != nil {
		fmt.Println("Setstr error is :",err)
		return
	}

	res,err := mr.GetStr(key)
	if err != nil {
		fmt.Println("Getstr error is :",err)
		return
	}
	fmt.Println(res)
	return*/
}
