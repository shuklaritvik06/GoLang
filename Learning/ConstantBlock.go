package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	isAdmin = 1 << iota
	isHeadquarters
)

func ConstantBlock() {
	fmt.Println(a, b, c)
	fmt.Println(isAdmin, isHeadquarters)
}
