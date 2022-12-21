package main

import "fmt"

func sayMessage(msg *string, name string) {
	*msg = "Hi"
	fmt.Println(name, *msg)
}

// return can be specified using (variable_name datatype)
// and only return keyword will do the work here, not necessary to specify the name

func getSlice(values ...int) int {
	fmt.Println(values)
	sum := 0
	for _, v := range values {
		sum = sum + v
	}
	return sum
}

func panickingDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()
	if b == 0 {
		panic("Cannot divide by zero")
	}
	fmt.Println(a / b)
}

// Methods in go
// Methods are functions with receiver
// Receiver can be of any type
// This will have the copy of the person struct as p in this context and can be accessed by this function
func (p person) printFullName() {
	fmt.Println(p.firstName, p.lastName)
}

type person struct {
	firstName string
	lastName  string
}

func function() {
	fmt.Println("Let's learn functions in go")
	greeting := "Hello"
	sayMessage(&greeting, "Ramesh")
	sum := getSlice(1, 2, 3, 4, 5)
	fmt.Println(sum)
	panickingDivide(10, 0)
	p := person{firstName: "Ramesh", lastName: "Kumar"}
	p.printFullName()
	// If we try to change the value of any entity in the function it will not change globally as we are not passing reference
	// But if we pass the reference of the variable it will change globally
	// For that we have to write reciever as a pointer (p *person)
}
