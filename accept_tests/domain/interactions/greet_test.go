package interactions_test

import (
	"testing"

	"github.com/kikis202/learn-go-with-tests/go-specs-greet/domain/interactions"
	"github.com/kikis202/learn-go-with-tests/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))
}
