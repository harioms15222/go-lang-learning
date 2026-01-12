// always starts with a package name
package main

import "fmt"

// package level variables
// dont raise an error
var pkgLvlVar int

// program always starts with main
// only things that start with capital letter are exported
func Functions() {
	// fmt.Println("hello")

	var whatToSay string
	var i int

	whatToSay = "hello world!"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to ", i)

	// to use formatted strings
	// have to declare is not using inference using :=
	var formattedString string = fmt.Sprintf("i is set to %d", i)
	fmt.Println(formattedString)

	// here no need to declare use inference here
	message, intReturned := saySomething()
	fmt.Println(message, intReturned)
}

// functions can return more than one thing
func saySomething() (string, int) {
	return "something", 100
}
