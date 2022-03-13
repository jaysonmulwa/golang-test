package test

import (
	"testing"

	customer "github.com/jaysonmulwa/jumia/internal/customer"
)

type test struct {
	input    string
	expected string
}

type resolver struct {
	input    string
	expected bool
}

func TestResolveCountry(t *testing.T) {

	tests := []test{
		{"(212) 698054317", "Morocco"},
		{"(212) 6007989253", "Morocco"}, //invalid
	}

	for _, test := range tests {
		if country, _ := customer.ResolveCountry(test.input); country == "Unknown" {
			t.Errorf("Expected country to be %s, got Unknown", test.expected)
		}
	}

}

func TestResolver(t *testing.T) {

	tests := []resolver{
		{"(212) 698054317", true},
		{"(212) 6007989253", false},
	}

	for _, test := range tests {
		if validity, _, _ := customer.Resolver(test.input); validity != test.expected {
			t.Errorf("Expected validity to be %T, got %T", test.expected, validity)
		}
	}

}
