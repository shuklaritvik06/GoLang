package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func Goroutine() {
	fmt.Println("Let's Learn Goroutine")
	// Spin off a green thread and run this in that thread
	// Most prog languages use OS threads for the execution and this means they get individual call stack for their execution
	// This causes application to take time to setup
	// We have go scheduler which will map these goroutines onto the operating system thread for period of time, and the scheduler handle these tasks at low level
	go func() {
		fmt.Println("Hello World")
	}()
	fmt.Println("Hello World")
	wg.Add(1)
	msg := "Hello"
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()
	// Race between Goroutine
	for i := 0; i < 10; i++ {
		wg.Add(2)
		sayHello()
		increment()
	}
}

func sayHello() {
	fmt.Println("Hello", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
