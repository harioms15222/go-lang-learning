package main

import (
	"log"
	"time"
)

type User struct {
	// will be only available to other files to other packages if start with capital letter
	FirstName string
	LastName  string
	PhoneNo   string
	Age       int
	BirthDate time.Time
	kink      string
}

func Struct() {
	var user User = User{
		FirstName: "hariom",
		LastName:  "sharma",
		kink:      "dauhda",
		BirthDate: time.Now(),
	}

	log.Println(user.FirstName, user.LastName, user.kink, user.BirthDate)
}
