package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func newRedisPool(addr string, db int, password string) *redis.Pool {
	//log.Println("redis:pool:", addr, db)
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

func getDailyCount(cDate string) int64 {
	c := newRedisPool(redisAddr, 1, "").Get()
	defer c.Close()

	res, err := c.Do("BITCOUNT", "attendence:"+cDate)

	if err != nil {
		log.Println("employee:redis:", err)
	}
	count := res.(int64)

	return count
}

func getDailyUser(cDate string, size int) string {
	c := newRedisPool(redisAddr, 1, "").Get()
	defer c.Close()

	res, err := redis.String(c.Do("GETRANGE", "attendence:"+cDate, 0, size))

	if err != nil {
		log.Println("employee:redis:", err)
	}

	return res
}

func getConsecutivePresent(cDate string, yDate string, size int) string {
	c := newRedisPool(redisAddr, 1, "").Get()
	defer c.Close()

	andKey := "attendence:" + cDate + ":" + yDate + ":and"
	_, err := c.Do("BITOP", "AND", andKey, "attendence:"+cDate, "attendence:"+yDate)

	if err != nil {
		log.Println("employee:redis:", err)
	}

	res, err := redis.String(c.Do("GETRANGE", andKey, 0, size))

	if err != nil {
		log.Println("employee:redis:", err)
	}

	return res
}

func getConsucutiveCount(cDate string, yDate string) int64 {
	c := newRedisPool(redisAddr, 1, "").Get()
	defer c.Close()

	andKey := "attendence:" + cDate + ":" + yDate + ":and"

	res, err := c.Do("BITCOUNT", andKey)

	if err != nil {
		log.Println("employee:redis:", err)
	}
	count := res.(int64)

	return count
}
