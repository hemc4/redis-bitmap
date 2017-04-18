package main

import (
	"log"
	"testing"
)

func generateMockAttendence(cDate string) {
	c := newRedisPool(redisAddr, 1, "").Get()
	defer c.Close()

	for i := 1; i <= size; i++ {
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
	yDate := getYesterdayDate()
	generateMockAttendence(yDate)
	count := getDailyCount(cDate)
	log.Println(count)
	ycount := getDailyCount(yDate)
	log.Println(ycount)

}
