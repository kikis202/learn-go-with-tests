package main_test

import (
	"testing"

	"github.com/kikis202/learn-go-with-tests/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	driver := go_specs_greeter.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
