package models

import (
	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
)

func connectRedis() *redis.Conn {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		glog.Errorf("Connect to redis error", err)
	}
	return &c
}
