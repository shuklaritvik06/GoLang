package main

import "fmt"

func Loops() {
	fmt.Println("Let's Learn Loops")
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	z := 0
	for ; z < 5; z++ {
		fmt.Println(z)
	}
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	for {
		fmt.Println("Hello")
		break
	}
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
Loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				continue
			}
			if i*j == 3 {
				// Helpful for breaking out of both the loop
				break Loop
			}
			fmt.Println(i, j)
		}
	}
	// For-Range Loop
	// For-Range loop is used to iterate over the elements of an array, slice, map, string, channel or a collection.
	// The syntax of for-range loop is:
	// for index, value := range collection {
	//     // body of for loop
	// }
	// The index variable is used to get the index of the element in the collection. The value variable is used to get the value of the element from the collection.
	for k, v := range []int{1, 2, 4, 5} {
		fmt.Println(k, v)
	}
	for k, v := range map[string]int{"Ritvik": 1, "Rahul": 2} {
		fmt.Println(k, v)
	}
	for k, v := range "Ritvik" {
		// value is ASCII
		fmt.Println(k, v)
	}
}
