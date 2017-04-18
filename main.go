package main

import (
	"fmt"
	"strconv"
)

const (
	redisAddr = "127.0.0.1:6379"
	size      = 100
)

func main() {

	fmt.Println("Redis based employee analytics")
	cDate := getCurrentDate()
	yDate := getYesterdayDate()

	count := getDailyCount(cDate)
	fmt.Println("Total no employee came to office today is: ", count)

	users := getDailyUser(cDate, size)
	bitArr := toBitArray(users, size)
	//fmt.Println(bitArr)
	psuers := PresentEmployees(bitArr)
	fmt.Println("Employee id's of today present employees", psuers)
	ausers := AbsentEmployees(bitArr)
	fmt.Println("Employee id's of today absent employees", ausers)

	ycount := getDailyCount(yDate)
	fmt.Println("Total no employee came to office yesterday is: ", ycount)
	yusers := getDailyUser(yDate, size)
	ybitArr := toBitArray(yusers, size)
	//fmt.Println(bitArr)
	ypsuers := PresentEmployees(ybitArr)
	fmt.Println("Employee id's of yesterday present employees", ypsuers)
	yausers := AbsentEmployees(ybitArr)
	fmt.Println("Employee id's of yesterday absent employees", yausers)

	conCount := getConsucutiveCount(cDate, yDate)
	fmt.Println("Total no employee came to office consecutive is: ", conCount)

	consecutePreUsers := getConsecutivePresent(cDate, yDate, size)
	bitArrConP := toBitArray(consecutePreUsers, size)
	conPresUsers := PresentEmployees(bitArrConP)
	fmt.Println("Consecutive present employees", conPresUsers)

}

func PresentEmployees(bitArr string) string {

	var s string
	for i, c := range bitArr {
		if c == '1' {
			s = s + strconv.Itoa(i+1) + ","
		}
	}

	return s

}

func AbsentEmployees(bitArr string) string {
	var s string
	for i, c := range bitArr {

		if c == '0' {
			s = s + strconv.Itoa(i+1) + ","
		}
	}
	return s

}
