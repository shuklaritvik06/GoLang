package main

import "fmt"

func Pointers() {
	fmt.Println("Let's Learn the Pointers")
	var a int = 23
	var b *int = &a
	fmt.Println(a, b)
	array := []int{1, 2, 3, 5}
	fmt.Println(&array[1], &array[2])
	fmt.Println(array)
	var str *myStruct
	str = &myStruct{name: "Ramesh"}
	fmt.Println(str)
	var str2 *myStruct
	str2 = new(myStruct)
	fmt.Println(str2)
	// (*str2).name = "Rakesh"
	str2.name = "Rakesh"
	fmt.Println(str2)
}

type myStruct struct {
	name string
}
