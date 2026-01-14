package main

import (
	"fmt"
	"strconv"
)

func Errors() {
	// Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
	// 	//
	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
}
