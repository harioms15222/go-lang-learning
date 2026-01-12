package main

import "log"

var s1 = "seven"

func Vars() {
	// var s1 = "string"

	log.Println("s1 is", s1)
	saySomething1("night")
}

func saySomething1(s1 string) (string, string) {
	// NOTE: here the s1 passed into the function call will take precedence
	// scoping
	log.Println("s1 from saySomething1 is", s1)
	return s1, "something"
}
