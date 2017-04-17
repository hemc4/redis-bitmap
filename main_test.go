package main

import (
	"log"
	"testing"
)

func generateMockAttendence(cDate string) {
	c := newRedisPool(REDISADR, 1, "").Get()
	defer c.Close()

	for i := 1; i <= 10000000; i++ {
		value := getRandom()
		_, err := c.Do("SETBIT", "attendence:"+cDate, i, value)

		if err != nil {
			log.Println("employee:generateMockAttendence:", err)
		}
	}

}

func TestGetDailyCount(t *testing.T) {
	cDate := getCurrentDate()
	generateMockAttendence(cDate)
	count := getDailyCount(cDate)
	log.Println(count)

}
