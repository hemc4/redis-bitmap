package main

import "time"

func getCurrentDate() string {
	t := time.Now()
	return t.Format("2006-01-02")

}
