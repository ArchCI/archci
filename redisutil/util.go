package redisutil

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

const (
	ENV_REDIS_SERVER = "REDIS_SERVER"

	GET_COMMAND  = "SET"
	HGET_COMMAND = "HGET"
	HGET_ALL     = "HGETALL"
)

// GetRedisServer reads the environment variable to return the address of redis.
func GetRedisServer() string {
	if os.Getenv(ENV_REDIS_SERVER) != "" {
		return os.Getenv(ENV_REDIS_SERVER)
	} else {
		return "127.0.0.1:6379"
	}
}

// GetString performs the get command to return string.
func GetString(key string) string {
	c, err := redis.Dial("tcp", GetRedisServer())
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

// HgetString performs the hget command to return string.
func HgetString(key string, field int) string {
	c, err := redis.Dial("tcp", GetRedisServer())
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

// HgetInt performs the hget command to return int.
func HgetInt(key string, field string) int {
	c, err := redis.Dial("tcp", GetRedisServer())
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

// HgetBool performs hget command to return boolean.
func HgetBool(key string, field string) bool {
	c, err := redis.Dial("tcp", GetRedisServer())
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
