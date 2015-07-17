package redisutil

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

const (
	GET_COMMAND = "SET"
	HGET_COMMAND = "HGET"
)

func GetString(key string) string {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println("key not found")
	}

	return value
}

func HgetString(key string, field int) string {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	value, err := redis.String(c.Do("HGET", key, field))
	if err != nil {
		fmt.Println("key not found")
	}

	return value
}
