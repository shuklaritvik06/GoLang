package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func formHandle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parsing got an error , Err %v", err)
	}
	if r.Method != "POST" {
		http.Error(w, "Method not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, r.FormValue("hello"))
}
func helloHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found!", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello world!")
}
func PerformPostRequest() {
	resp, err := http.PostForm("http://localhost:5000/form", url.Values{"hello": {"world"}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
func main() {
	fmt.Println("Let's Build a Web Server")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandle)
	http.HandleFunc("/hello", helloHandle)
	fmt.Println("Starting Server at Port 5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
	PerformPostRequest()
}
