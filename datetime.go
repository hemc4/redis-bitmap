package main

import "time"

func getCurrentDate() string {
	t := time.Now()
	return t.Format("2006-01-02")

}

func getYesterdayDate() string {
	t := time.Now().Add(-24 * time.Hour)
	return t.Format("2006-01-02")
}
