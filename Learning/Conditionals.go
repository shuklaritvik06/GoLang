package main

import (
	"fmt"
	"math"
)

func returnTrue() bool {
	return true
}

func Conditionals() {
	fmt.Println("Let's Learn Conditionals")
	if true {
		fmt.Printf("Hello I am in true conditional")
	}
	if returnTrue() {
		fmt.Println("Hello I am in true conditional")
	}
	age := 23
	if age >= 20 {
		fmt.Println("20 se jyada umar")
	} else if age >= 30 {
		fmt.Printf("30 se jyada umar")
	} else {
		fmt.Printf("pata nhi")
	}
	if age == 23 && returnTrue() {
		fmt.Printf("Done")
	}
	quan := math.Pow(2, 3)
	fmt.Println(quan)
	switch 2 {
	case 1:
		fmt.Printf("Hello I am 1")
		break
	case 2:
		fmt.Printf("Hello I am 2")
		break
	default:
		fmt.Printf("I am default")
	}
	switch 3 {
	case 1:
		fmt.Printf("Hello I am 1")
		break
	case 2:
	case 3:
		fmt.Printf("Hello I am 2 or 3")
		break
	default:
		fmt.Printf("I am default")
	}
	switch 3 {
	case 1, 2, 3:
		fmt.Printf("Hello I am 1 or 2 or 3")
		break
	default:
		fmt.Printf("I am default")
	}
	switch i := 2 + 1; i {
	case 1, 2, 3:
		fmt.Printf("Hello I am 1 or 2 or 3")
		break
	default:
		fmt.Printf("I am default")
	}
	i := 3
	switch {
	case i <= 3:
		{
			fmt.Printf("Hello I am 1 or 2 or 3")
		}
	default:
		{
			fmt.Printf("I am default")
		}
	}
	switch 2 {
	case 1:
		fmt.Printf("Hello I am 1")
		fallthrough
	case 2:
		fmt.Printf("Hello I am 2")
	default:
		fmt.Printf("I am default")
	}
}
