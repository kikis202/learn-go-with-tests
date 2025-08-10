package main

import (
	"net/http"

	"github.com/kikis202/learn-go-with-tests/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	http.ListenAndServe(":8080", handler)

}
