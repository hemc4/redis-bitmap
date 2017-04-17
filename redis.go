package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func newRedisPool(addr string, db int, password string) *redis.Pool {
	log.Println("adx:analytics:pool:", addr, db)
	return &redis.Pool{
		MaxIdle:   8,
		MaxActive: 120, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialDatabase(db), redis.DialPassword(password))
			if err != nil {
				log.Println("Could not establish connection with redis Server", addr)
				panic(err.Error())
			}
			return c, err
		},
	}
}
