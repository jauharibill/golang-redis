package main

import (
	"github.com/go-redis/redis/v8"
)

var red *redis.Client

func init() {
	red = redis.NewClient(&redis.Options{DB: 0, Password: "", Addr: ":6379"})
}

func Conn() *redis.Client {
	return red
}
