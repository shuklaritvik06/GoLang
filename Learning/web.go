package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://google.com"
	fmt.Println("Let's play with web")
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(response)
	ioutil.ReadAll(response.Body)
	response.Body.Close() // Our responsibility to close the connection
}
