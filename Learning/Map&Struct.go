package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	age        int
	name       string
	colleagues []string
}

// We can see the struct outside but not the entities

type Employee struct {
	ename    string
	edept    string
	esenior  string
	dmembers []string
	salaries []int
}

type Animal struct {
	name   string
	typeof string
}

type Snakes struct {
	Animal
	venomous bool
}

func MaoandStruct() {
	fmt.Printf("%v", "Hello")
	// In the case of map whenever we change any variable referenced to the main it will change both
	statePopulations := map[string]int{
		"hello": 23000,
		"hi":    25000,
	}
	fmt.Printf("%v", statePopulations)
	// Slices can't be a key for a map
	delete(statePopulations, "hi")
	fmt.Printf("%v", statePopulations["hi"])
	// When we use the named variable then we can omit any entity which we don't want to add
	doctor := Doctor{
		number: 8191919191,
		age:    41,
		name:   "Ramesh",
		colleagues: []string{
			"Rakesh",
			"Ramesh",
		},
	}
	fmt.Printf("%v", doctor)
	// When we do not use the name then we have to declare all in the order (Positional Syntax)
	empl := Employee{
		"Ramesh",
		"Google Maps",
		"Rakesh",
		[]string{
			"Rahul",
			"Mahesh",
		},
		[]int{
			500000,
			10000000,
		},
	}
	fmt.Printf("%v", empl)
	fmt.Printf("%v", empl.dmembers[1])
	// Anonymous Struct
	myStruct := struct{ name string }{name: "Rakesh"}
	fmt.Printf("%v", myStruct)
	// It will make the copies of the data
	anotherStruct := myStruct
	anotherStruct.name = "Ramesh"
	fmt.Printf("%v", anotherStruct)
	// If we use address of (&) operator then it will help us to manipulate the original data
	mySnake := Snakes{
		Animal: Animal{
			name:   "Snake",
			typeof: "Reptile",
		},
		venomous: true,
	}
	fmt.Printf("%v", mySnake)
	t := reflect.TypeOf(Animal{})
	f, _ := t.FieldByName("name")
	fmt.Printf("%v", f)
}
