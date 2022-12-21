package main

import "fmt"

func Arrays() {
	fmt.Println("Let's Learn Arrays")
	marks := [4]int{1, 2, 3, 4}
	fmt.Printf("%v", marks)
	var student [3]string
	student[0] = "Ritvik"
	fmt.Printf("%v", student)
	fmt.Printf("%v", len(student))
	identityMatrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {6, 7, 8}}
	fmt.Printf("%v", identityMatrix)
	var startArray [3]int
	fmt.Printf("%v", startArray) // 0 0 0
	fmt.Printf("%v", [...]int{1, 2, 3, 4})
	temp := &marks
	fmt.Printf("%v", temp)
	fmt.Printf("%v", temp[0:2])
	fmt.Printf("%v", temp[2:])
	a := make([]int, 4, 100)
	fmt.Printf("%v", cap(a))
	a = append(a, 1)
	fmt.Printf("%v", a)
	// a = append(a, []int{4, 5, 6, 7, 8, 9}...)
}
