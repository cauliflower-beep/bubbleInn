package main

import (
"github.com/gomodule/redigo/redis"
"fmt"
)

func main()  {
	// 1. 链接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis Dial err:", err)
		return
	}
	fmt.Println("连接成功!")
	defer conn.Close()

	// 2. 操作数据库
	reply, err := conn.Do("set", "小花椰", "angel")

	// 3. 回复助手类函数. ---- 确定成具体的数据类型
	r, e := redis.String(reply, err)

	fmt.Println(r, e)
}
