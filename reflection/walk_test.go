package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name:          "Struct with one string field",
			Input:         struct{ Name string }{"Kristaps"},
			ExpectedCalls: []string{"Kristaps"},
		},
		{
			Name: "Struct with three string fields",
			Input: struct {
				Name    string
				City    string
				Country string
			}{"Kristaps", "Riga", "Latvia"},
			ExpectedCalls: []string{"Kristaps", "Riga", "Latvia"},
		},
		{
			Name: "Struct with non-string field",
			Input: struct {
				Name string
				Age  int
			}{"Kristaps", 23},
			ExpectedCalls: []string{"Kristaps"},
		},
		{
			Name: "Nested struct",
			Input: Person{
				Name: "Kristaps",
				Profile: Profile{
					City: "Riga",
					Age:  23,
				},
			},
			ExpectedCalls: []string{"Kristaps", "Riga"},
		},
		{
			Name: "Pointers",
			Input: &Person{
				Name: "Kristaps",
				Profile: Profile{
					City: "Riga",
					Age:  23,
				},
			},
			ExpectedCalls: []string{"Kristaps", "Riga"},
		},
		{
			Name:          "slice",
			Input:         []Profile{{"Riga", 23}, {"Tukums", 11}},
			ExpectedCalls: []string{"Riga", "Tukums"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, expected %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}
