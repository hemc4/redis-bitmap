package main

import (
	"log"
	"math/rand"
	"time"
)

func getRandom() bool {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	randNo := r1.Intn(100)
	boolNo := randNo % 2

	log.Printf("%d", boolNo)
	if boolNo == 1 {
		return true
	}
	return false

}
