package main

import (
	"fmt"
	"net/url"
)

func URL() {
	myurl := "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj4"
	fmt.Println("Hello, World!")
	result, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(result.Path, result.Scheme, result.Port())
	qparams := result.Query()
	// Returns a map
	fmt.Println(qparams)
	for _, v := range qparams {
		fmt.Println(v)
	}
	fmt.Println(qparams["coursename"])
	patsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "coursename=reactjs&paymentid=ghbj4",
	}
	fmt.Println(patsOfUrl.String())
}
