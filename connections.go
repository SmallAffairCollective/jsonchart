package main

import (
	"github.com/mediocregopher/radix.v2/redis"
)

func connectRedis(redisHost string) *redis.Client {

	conn, er := redis.Dial("tcp", redisHost+":6379")
	check(er)

	//defer conn.Close()

	return conn

}
