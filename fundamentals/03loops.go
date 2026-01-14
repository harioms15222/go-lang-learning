package main

import (
	"log"
	"time"
)

func Loops() {
	// we only have a single type of loop --> for loop

	// 1)
	for i := 0; i <= 5; i++ {
		log.Println("i:", i)
	}

	// 2) ranging over data
	animals := []string{"dog", "fish", "cat"}
	for i, animal := range animals {
		log.Println(i, animal)
	}

	// 3) using blank identifier to ignore the returned values
	animals2 := []string{"dog", "fish", "cat"}
	for _, animal := range animals2 {
		log.Println(animal)
	}

	pet_name_map := make(map[string]string)
	pet_name_map["dog"] = "tommy"
	pet_name_map["cat"] = "chunmun"
	pet_name_map["mouse"] = "jerry"

	for animal, name := range pet_name_map {
		log.Println(animal, "has name: ", name)
	}

	// string is actually a slice of runes (int32 alias)
	str := "I am great!"
	str = "ade"
	for index, ltr := range str {
		log.Println(index, "i has charcter", ltr, string(ltr))
	}

	// range over custom objects
	hariom := User{
		FirstName: "Hariom",
		LastName:  "Sharma",
		Age:       29,
		BirthDate: time.Now(),
	}

	kalu := User{
		FirstName: "Kalu",
		LastName:  "Sharma",
		Age:       21,
		BirthDate: time.Now(),
	}

	var users []User
	users = append(users, hariom)
	users = append(users, kalu)

	for i, data := range users {
		log.Printf("for %d th index we have %s user", i, data.FirstName)
	}

}
