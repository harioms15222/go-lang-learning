package main

import (
	"log"
	"sort"
)

func Slice() {
	var myStr string
	myStr = "fish"
	log.Println(myStr)

	// variables that store data in array format
	var myStudents []User
	hariom := User{
		FirstName: "Hariom",
	}

	kalu := User{
		FirstName: "Kalu",
	}

	// to add to array
	myStudents = append(myStudents, hariom)
	myStudents = append(myStudents, kalu)

	log.Println(myStudents)

	sort.Slice(myStudents, func(i, j int) bool {
		if myStudents[i].FirstName > myStudents[j].FirstName {
			return true
		} else {
			return false
		}
	})

	log.Println(myStudents)

	//////////////////////////////////////////////////
	println("/////////////////////////////////////////")
	println("order is maintained in a slice in which the elments are inserted")
	println("/////////////////////////////////////////")

	var marks []int
	marks = append(marks, 100)
	marks = append(marks, 11)
	marks = append(marks, 31341)

	log.Println("marks", marks)
	sort.Ints(marks)
	log.Println(marks)

	ages := []int{10, 20, 32, 28, 19}
	log.Println(ages)

	log.Println(ages[0:2])
	log.Println(ages[2:5])
}
