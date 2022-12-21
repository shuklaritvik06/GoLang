package main

import (
	"fmt"
	"time"
)

func Time() {
	fmt.Println("Hello, Time Package")
	currentTime := time.Now()
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))
	created := time.Date(2022, time.December, 16, 14, 39, 0, 0, time.UTC)
	fmt.Println(created.Format("01-02-2006 15:04:05 Monday"))
}

// new -> Allocate memory but not init (Zeroed Storage)
// make -> init (Non Zeroed Storage)
