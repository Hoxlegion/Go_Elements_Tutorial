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
			"struct with one string field",
			struct {
		        Name string
            }{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
                Name string
                City string
            }{"Chris", "Londen"},
			[]string{"Chris", "Londen"},
		},
		{
			"struct with non string field",
			Person {
                "Chris",
                Profile{33, "Londen"},
            },
			[]string{"Chris", "Londen"},
		},
		{
			"pointers to things",
			&Person {
			    "Chris",
                Profile{33, "London"},
			},
			[]string{"Chris", "London"},
        },
        {
            "slices",
            []Profile {
                {33, "London"},
                {34, "Reykjavik"},
            },
            []string{"London", "Reykjavik"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
