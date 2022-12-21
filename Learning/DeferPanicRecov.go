package main

import "fmt"

func panicking() {
	fmt.Println("Entering Panicking Mode")
	// This will not stop the panicking sequence so deferred function should be there to capture the error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	panic("Something Happened")
}

func Defer() {
	fmt.Println("Let's Learn Defer Panic Recovery")
	fmt.Println("Start")
	// defer fmt.Println("Middle")
	fmt.Println("End")
	// Defer follows the LIFO so the last statement would be executed first
	// Panic
	// Defer always gets executed before the panicking
	// fmt.Println("Hello I am start")
	// defer fmt.Println("I am deferred")
	// panic("Something Happened")
	// Recover
	fmt.Println("Start")
	panicking()
	fmt.Println("End")
	// Defer => It is used for delayed execution
	// Panic => It is the condition occured when the program cannot continue at all
	// Recover => Used to recover from panics and capture the error
}
