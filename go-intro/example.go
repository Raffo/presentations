package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Gophers\n")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
