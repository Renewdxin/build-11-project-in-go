package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting a server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "err:%v", err)
		return
	}
	fmt.Fprintf(writer, "POST request successfully\n")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(writer, "name = %v \naddress = %v", name, address)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 NOT FOUND", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer, "hello!")

}
