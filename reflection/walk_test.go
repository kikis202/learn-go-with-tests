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
		{
			Name:          "array",
			Input:         [2]Profile{{"Riga", 23}, {"Tukums", 11}},
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

	t.Run("maps", func(t *testing.T) {
		aMap := map[string]string{
			"S": "Small",
			"M": "Medium",
			"L": "Large",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Small")
		assertContains(t, got, "Medium")
		assertContains(t, got, "Large")
	})

	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{"Riga", 23}
			aChannel <- Profile{"Tukums", 11}
			close(aChannel)
		}()

		var got []string
		want := []string{"Riga", "Tukums"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{"Riga", 23}, Profile{"Tukums", 11}
		}

		var got []string
		want := []string{"Riga", "Tukums"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	for _, val := range haystack {
		if val == needle {
			return
		}
	}

	t.Errorf("Expected %v to contain %q, but it didnt", haystack, needle)
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}
