package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", &handler{dur}) //TIMEOUTS?
}
