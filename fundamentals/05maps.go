package main

import "log"

func Maps() {
	// means that its not usable
	var myMap map[string]string

	// panic: assignment to entry in nil map
	// myMap["hi"] = "dajh"

	// myMap.hi undefined (type map[string]string has no field or method hi)
	// myMap.hi = "np"

	log.Println("My Map", myMap)

	myMapInitialized := make(map[string]string)
	myMapInitialized["hi"] = "sdajh"

	//  myMapInitialized.dj undefined (type map[string]string has no field or method dj)
	// myMapInitialized.dj = "duhsad"

	log.Println("My Map Initialized", myMapInitialized)

	//myMapInitialized.hi undefined (type map[string]string has no field or method hi)
	// log.Println("My Map Initialized", myMapInitialized.hi)

	log.Println("My Map Initialized", myMapInitialized["hi"])

	myMapInitialized["hi"] = "idk"
	log.Println("My Map Initialized", myMapInitialized["hi"])

	userMap := make(map[string]*User)
	id := "100"

	hariom := User{
		FirstName: "Hariom",
		LastName:  "Sharma",
	}

	userMap[id] = &hariom

	log.Println(userMap)
	// to call a
	userMap[id].PrintFirstName()
}
