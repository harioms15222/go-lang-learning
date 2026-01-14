package main

import "log"

//

func Pointers() {
	var myStr string
	myStr = "Green"

	log.Println("myStr is set to ", myStr, &myStr)

	changeUsingPointer(&myStr)

	log.Println("after function call, myStr is set to ", myStr, &myStr)
}

func changeUsingPointer(s *string) {
	log.Println("pointer(ref) to s is", s)
	newValuse := "red"
	*s = newValuse
}
