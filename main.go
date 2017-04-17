package main

import (
	"fmt"
	"log"
)

const REDISADR = "127.0.0.1:6379"

func main() {
	fmt.Println("Redis based employee analytics")
	cDate := getCurrentDate()
	count := getDailyCount(cDate)
	fmt.Println("Total no employee came to office today is: ", count)

}

func getDailyCount(cDate string) int64 {
	c := newRedisPool(REDISADR, 1, "").Get()
	defer c.Close()

	res, err := c.Do("BITCOUNT", "attendence:"+cDate)

	if err != nil {
		log.Println("employee:redis:", err)
	}
	count := res.(int64)

	return count
}
