package main

import (
	"testing"
)

func equalGames(first, second *Game) bool {
	if first.Id != second.Id {
		return false
	}
	hasSameDraws := sameDraws(first.Draws, second.Draws)
	return hasSameDraws
}

func equalDraws(first, second Draw) bool {
	return first.Red == second.Red && first.Blue == second.Blue && first.Green == second.Green
}

func sameDraws(first, second []Draw) bool {
	if len(first) != len(second) {
		return false
	}

	for _, d1 := range first {
		foundInSecond := false
		for _, d2 := range second {
			if equalDraws(d1, d2) {
				foundInSecond = true
				break
			}
		}
		if !foundInSecond {
			return false
		}
	}

	return true
}

func TestInputParsing(t *testing.T) {
	cases := map[string]*Game{
		"Game 1: 6 green, 3 blue; 3 red, 1 green; 4 green, 3 red, 5 blue": {
			Id: 1,
			Draws: []Draw{
				{
					Red:   0,
					Green: 6,
					Blue:  3,
				},
				{
					Red:   3,
					Green: 1,
					Blue:  0,
				},
				{
					Red:   3,
					Green: 4,
					Blue:  5,
				},
			},
		},
		"Game 2: 2 red, 7 green; 13 green, 2 blue, 4 red; 4 green, 5 red, 1 blue; 1 blue, 9 red, 1 green": {
			Id: 2,
			Draws: []Draw{
				{
					Red:   2,
					Green: 7,
					Blue:  0,
				},
				{
					Red:   4,
					Green: 13,
					Blue:  2,
				},
				{
					Red:   5,
					Green: 4,
					Blue:  1,
				},
				{
					Red:   9,
					Green: 1,
					Blue:  1,
				},
			},
		}}

	for input, want := range cases {
		if got, _ := parseLineIntoGame(input); !equalGames(got, want) {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}

func TestDrawsAreEqual(t *testing.T) {
	cases := map[Draw]Draw{
		{Red: 1, Green: 2, Blue: 3}: {Red: 1, Green: 2, Blue: 3},
		{Blue: 2, Green: 2}:         {Red: 0, Green: 2, Blue: 2},
	}

	for first, second := range cases {
		if !equalDraws(first, second) {
			t.Errorf("%s and %s should be equal", first.String(), second.String())
		}
	}
}

func TestDrawsAreNotEqual(t *testing.T) {
	cases := map[Draw]Draw{
		{Red: 1, Green: 2, Blue: 2}: {Red: 1, Green: 2, Blue: 3},
		{Red: 1, Green: 1, Blue: 3}: {Red: 1, Green: 2, Blue: 3},
		{Red: 0, Green: 2, Blue: 3}: {Red: 1, Green: 2, Blue: 3},
	}

	for first, second := range cases {
		if equalDraws(first, second) {
			t.Errorf("%s and %s should not be equal", first.String(), second.String())
		}
	}
}
