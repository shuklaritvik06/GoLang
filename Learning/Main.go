package main

import (
	"fmt"
)

var name string = "Ritvik"

var (
	naam string = "Ramesh"
	age  int    = 22
)

func Main() {
	// var i int = 43
	// : => this is declaring a new variable here for the assignment
	// first letter small case variables => These variables are exposed at the package level and the package files can use them
	// first letter upper case variables => These variables are exposed at the global level
	// Acronym must be in capital letters to give much more visibility
	i := 43
	var k float64
	k = 230.34
	k = float64(i)
	fmt.Println(i, k)
	str := string(rune(52))
	fmt.Println(str)
	fmt.Printf("%v,%T\n", i, i)
	var flag bool = true
	fmt.Println(flag)
	value := 1 == 1
	fmt.Println(value)
	var unsign uint32 = 23
	fmt.Println(unsign)
	// Bit Manipulation
	left := 4                 // 0100
	right := 10               // 1010
	fmt.Println(left & right) // 0000
	fmt.Println(left ^ right) // 1110
	fmt.Println(left | right)
	fmt.Println(left &^ right)
	fmt.Println(left >> 3) // 2^2 / 2^3
	fmt.Println(left << 3) // 2^2 * 2^3
	fmt.Println(2.4e+14)
	var comp complex64 = 1 + 2i
	fmt.Println(comp)
	fmt.Println(real(comp), imag(comp))
	fmt.Println(complex(3.1, 4.5))
	s := "This is a string"
	fmt.Printf("\n%b\n", s[2])
	fmt.Println("Hello" + "Hi")
	fmt.Println([]byte(s))
	// CONSTANTS
	// These has to be known before the declaration and should not be determined on the compile time
	const myConst int = 23
	fmt.Println(myConst)
}
