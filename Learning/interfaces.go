package main

import "fmt"

func Interfaces() {
	fmt.Println("Let's Learn Interfaces")
	var a Greet = myStruct{}
	a.greeting()
	var b IncrementCounter = Increment(0)
	fmt.Println(b.increment())
}

type Greet interface {
	greeting()
}

type Increment int

type IncrementCounter interface {
	increment() int
}

func (g myStruct) greeting() {
	fmt.Println("Hello World!")
}

func (i Increment) increment() int {
	sum := 0
	for ; i < 5; i++ {
		sum += int(i)
	}
	return sum
}

// type myStruct struct{}
