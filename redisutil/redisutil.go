package redisutil

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

const (
	GET_COMMAND = "SET"
	HGET_COMMAND = "HGET"
	HGET_ALL = "HGETALL"
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

func HgetInt(key string, field string) int {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	value, err := redis.Int(c.Do("HGET", key, field))
	if err != nil {
		fmt.Println("key not found")
	}

	return value
}

func HgetBool(key string, field string) bool {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	value, err := redis.Bool(c.Do("HGET", key, field))
	if err != nil {
		fmt.Println("key not found")
	}

	return value
}
