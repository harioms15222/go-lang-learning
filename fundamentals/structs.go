package main

import (
	"log"
	"time"
)

func (m *User) PrintFirstName() {
	log.Println(m.FirstName)
}

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

	var user1 User
	user1.FirstName = "harry"

	log.Println(user.FirstName, user.LastName, user.kink, user.BirthDate)
	log.Println(user1.FirstName)

	user1.PrintFirstName()

}
