package main

import "testing"

func TestA(t *testing.T) {
	cases := map[string]string{
		"sevendxbninefour2fourclmln":             "74",
		"5threesevenvnthreeqkcd2xkfhprfgdzseven": "57",
	}

	for i, o := range cases {
		if result := parseFirstAndLastNumberWordInString(i); result != o {
			t.Errorf("for input %s we should get %s, but got %s", i, o, result)
		}
	}
}
