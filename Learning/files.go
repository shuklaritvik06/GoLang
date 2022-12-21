package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Files() {
	fmt.Println("Let's manipulate files")
	content := "Hello I am RS"
	data := []byte(content)
	data2 := []byte(content)
	file, err := os.Create("./file")
	if err != nil {
		fmt.Println(err)
	}
	file.Write([]byte(content))
	file.Read(data)
	fmt.Println(string(data))
	data2, err2 := ioutil.ReadFile("./file")
	fmt.Println(data2)
	if err2 != nil {
		fmt.Println(err2)
	}
}
