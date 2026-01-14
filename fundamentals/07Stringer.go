package main

import "fmt"

// fmt package inplements an interface
// type Stringer interface {
// String() string
// }

// a good example of how interfaces are used
// this also shows th implications of implicit declarations of interface

type Person struct {
	name string
	age  int
}

// here Person implements Stringer method
// if this implements Stringer
// then any function the accepts varibles of this interface type
// can use this function and accept Person tyoe
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func StringerExample() {
	p1 := Person{"hariom", 28}
	p2 := Person{"Deepika", 28}

	fmt.Println(p1)
	fmt.Println(p2)

	anotherExample()
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func anotherExample() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
