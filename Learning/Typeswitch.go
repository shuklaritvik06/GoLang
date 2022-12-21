package main

import "fmt"

func Typeswitch() {
	fmt.Println("Let's Learn Typeswitch")
	var i interface{} = [1]string{"Ritvik"} // As we are using interface here we can assign any type to i
	// In order to compare the type of array we have to keep same size and same data type of entries
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case [1]string:
		fmt.Println("array")
	default:
		fmt.Println("default")
	}
}
