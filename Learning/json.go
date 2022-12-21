package main

// Creating and Consuming JSON
import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Price  int    `json:"price"`
	IsPaid bool   `json:"isPaid"`
}

func JSON() {
	fmt.Println("Let's Learn JSON")
	c1 := course{
		Name:   "GoLang",
		Author: "Rakesh",
		Price:  233,
		IsPaid: true,
	}
	// Convert to JSON
	finalJSON, _ := json.Marshal(c1)
	fmt.Println(string(finalJSON))
	// Convert from JSON
	var c2 course
	json.Unmarshal(finalJSON, &c2)
	fmt.Println(c2)
	// MarshalIndent
	finalJSON, _ = json.MarshalIndent(c1, "", " ")
	fmt.Println(string(finalJSON))
}

func DecodeJSON() {
	jsonData := []byte(
		`
		{
			"name": "GoLang",
			"author": "Rakesh",
			"price": 233,
			"isPaid": true
		}
		`,
	)
	var c1 course
	if json.Valid(jsonData) {
		fmt.Println("Valid JSON")
		json.Unmarshal(jsonData, &c1)
		fmt.Println(c1)
	}
}
