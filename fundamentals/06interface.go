package main

import (
	"fmt"
	"log"
	"math"
)

// An interface type is defined as a set of method signatures.
type Animal interface {
	// a set of method signatures
	Says(string) string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
// comment these out to check what happens in assignments
func (d *Dog) Says(s string) string {
	return fmt.Sprintf("%s says %s", d.Name, s)
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

type Gorilla struct {
	Name      string
	NoOfTeeth int
}

// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
// comment these out to check what happens in assignments
func (g *Gorilla) Says(s string) string {
	return fmt.Sprintf("%s says %s", g.Name, s)
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}

func Interface() {
	doggy := Dog{"Kallu", "desi"}
	gorilla := Gorilla{"Don", 45}

	// here pointer type *Dog implements the interface so we can assign
	var animal1 Animal = &doggy
	introduce(animal1, "woof")

	var animal2 Animal = &gorilla
	introduce(animal2, "woof")

}

// now a varibale which has type of an interface
// can be used in place of any variable that implements every method that interface implements
func introduce(a Animal, sayStr string) {
	log.Printf("%s and has %d legs", a.Says(sayStr), a.NumberOfLegs())
}

// //////////////////////////////////////////////////////////////
// Calling a method on interface type value
// calls the method associated with the actual value only
/////////////////////////////////////////////////////////////////
// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
// (value, type)
////////////////////////////////////////////////////////////////

type I interface {
	// must implement M
	M()
}

type T struct {
	S string
}

// type T's pointer *T implements M
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// type F implements M
func (f F) M() {
	fmt.Println(f)
}

func Interface2() {
	// a value of inter face type
	// any type which implements the interface can be assigned
	// but when assigned to some other type then becomes like (concrete_value, type_of_that_var)
	var i I

	// sings *T implments M we can assign &T to i
	i = &T{"Hello"}
	describe(i)
	i.M()

	// since F implements M we can assign it to i
	i = F(math.Pi)
	describe(i)
	i.M()

	// interface value if assigned to a nil value then concrete value is nill under the hood
	// here the i won't be nill by default\
	// initalised f but value is not set
	var f *T
	i = f
	describe(i)

	// A nil interface value holds neither value nor concrete type.
	// Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	var iota I
	describe(iota)
	// the below line will give runtime error
	// iota.M()
}

func describe(i I) {
	// type %T wont be of interfac type but actual type of the assigned variables
	fmt.Printf("(%v, %T)\n", i, i)
}

// ////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////
// The interface type that specifies zero methods is known as the empty interface:
// An empty interface may hold values of any type. (Every type implements at least zero methods.)
// Empty interfaces are used by code that handles values of unknown type.
// For example, fmt.Print takes any number of arguments of type interface{}.
func EmptyInterface() {
	var i interface{}
	describe1(i)

	i = 42
	describe1(i)

	i = "hello"
	describe1(i)

}

func describe1(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////
// TYPE ASSERTION for interface to check what type of value it holds
func TypeAssertion() {
	// i can hold any type of value here because
	// it is type of EmptyInterface
	var i interface{} = "hello"

	// will give the value is it finds the right type
	s := i.(string)
	fmt.Println(s)

	// can also return report if the type matches
	s, ok := i.(string)
	fmt.Println(s, ok)

	// if type does not match
	// 1) return null value of that type
	// 2) and reportVariable is set to false
	f, ok := i.(float64)
	fmt.Println(f, ok)

	/////////////////////////////////////
	///// STRICT MODE
	////////////////////////////////////////
	// Here we are directly assigning f to the returned value and it will cause panic
	// g := i.(float64) // panic
	// fmt.Println(g)
}

func do(i interface{}) {
	// here we can send the type keyword to the type assertion syntax
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TypeSwitch() {
	do(21)
	do("hello")
	do(true)
}
