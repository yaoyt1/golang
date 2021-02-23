package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//redisDial()
	redisBasic()
}

func redisDial() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("连接redis服务报错：%v", err)
		return
	}
	fmt.Println("连接redis服务器成功")
	defer conn.Close()
}

func redisBasic() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("连接redis服务报错：%v\n", err)
		return
	}
	defer conn.Close()

	reply, err := conn.Do("set", "abc", "100")
	if err != nil {
		fmt.Printf("set操作失败：%v\n", err)
		return
	}
	fmt.Sprintf("set reply:%#v\n", reply)

	r, err := redis.Int(conn.Do("get", "abc"))
	if err != nil {
		fmt.Printf("get操作失败：%v\n", err)
		return
	}
	fmt.Sprintf("get reply:%#v\n", r)
}
