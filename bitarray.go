package main

import "fmt"

func toBitArray(byteArray string, size int) string {

	//byteArray := "\x04Uliw\xf5\xefx\xe2\x83\x8b\xa8\xb0"
	//for i := 0; i < len(byteArray); i++ {
	//fmt.Printf("%x ", byteArray[i])
	//}
	//fmt.Println()

	var bitstr string
	extra := (len(byteArray) * 8) % size
	//fmt.Println("extra:", extra)
	for i := 0; i < len(byteArray); i++ {
		s := fmt.Sprintf("%08b", byteArray[i])
		//fmt.Println(s)
		if i == 0 {
			s = s[1:len(s)]
		} else if i == len(byteArray)-1 {
			s = s[0 : len(s)-extra+1]
		}
		//fmt.Println(s)
		bitstr += s
		//fmt.Println(bitstr)
		//fmt.Println()

	}
	return bitstr
}
