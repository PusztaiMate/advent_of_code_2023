package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	INPUT_FILE_NAME = "input.txt"
)

type Draw struct {
	Red, Green, Blue int
}

func (d *Draw) String() string {
	return fmt.Sprintf("Draw(r:%d, g:%d, b:%d)", d.Red, d.Green, d.Blue)
}

type Game struct {
	Id    int
	Draws []Draw
}

func (g *Game) String() string {
	drawsAsStrings := []string{}
	for _, draw := range g.Draws {
		asString := draw.String()
		drawsAsStrings = append(drawsAsStrings, asString)
	}
	return fmt.Sprintf("Game (%d, [%s])", g.Id, strings.Join(drawsAsStrings, ", "))
}

func NewGame(id int, draws []Draw) *Game {
	return &Game{
		Id:    id,
		Draws: draws,
	}
}

func main() {
	f, err := os.Open(INPUT_FILE_NAME)
	if err != nil {
		panic("could not open input file")
	}

	data, err := io.ReadAll(f)
	if err != nil {
		panic("could not read input file")
	}

	content := string(data)
	games := make([]*Game, 0)

	for _, line := range strings.Split(content, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		game, err := parseLineIntoGame(line)
		if err != nil {
			panic(fmt.Sprintf("could not parse line '%s' into game %s", line, game))
		}

		games = append(games, game)
	}

	idSum := 0
	powerSum := 0
	// drawMax := Draw{12, 13, 14}
	for _, game := range games {
		idSum += game.Id
		minimalDraw := smallestNumberOfColorsNeeded(game)
		powerSum += powerOfCubes(minimalDraw)
	}

	fmt.Println("the sum of the ids of the possible games is", idSum)
	fmt.Println("the sum of power of the minimal draws is", powerSum)
}

func smallestNumberOfColorsNeeded(game *Game) Draw {
	minimalDraw := Draw{}

	for _, draw := range game.Draws {
		if draw.Red > minimalDraw.Red {
			minimalDraw.Red = draw.Red
		}
		if draw.Blue > minimalDraw.Blue {
			minimalDraw.Blue = draw.Blue
		}
		if draw.Green > minimalDraw.Green {
			minimalDraw.Green = draw.Green
		}
	}

	fmt.Printf("the minimal draw for %s is %s", game.String(), minimalDraw.String())

	return minimalDraw
}

func powerOfCubes(draw Draw) int {
	power := draw.Blue * draw.Red * draw.Green
	fmt.Println("the power of draw", draw, "is", power)
	return power
}

func isGamePossible(game *Game, drawMax Draw) bool {
	for _, draw := range game.Draws {
		if draw.Red > drawMax.Red || draw.Blue > drawMax.Blue || draw.Green > drawMax.Green {
			return false
		}
	}
	return true
}

func parseLineIntoGame(line string) (*Game, error) {
	gameAsString, drawsAsString := splitLineIntoGameAndDraws(line)
	game := parseLineToGame(gameAsString)
	draws := parseLineSegmentIntoDraws(drawsAsString)
	game.Draws = draws
	return game, nil
}

func splitLineIntoGameAndDraws(line string) (string, string) {
	splitted := strings.SplitN(line, ":", 2)
	game, draws := splitted[0], splitted[1]
	return strings.TrimSpace(game), strings.TrimSpace(draws)
}

func parseLineSegmentIntoDraws(lineSegment string) []Draw {
	draws := make([]Draw, 0)

	splitted := strings.Split(lineSegment, ";")
	for _, part := range splitted {
		draws = append(draws, parseLineSegmentIntoDraw(part))
	}

	return draws
}

func parseLineSegmentIntoDraw(lineSegment string) Draw {
	// 2 green, 7 blue, 1 red
	draw := Draw{}
	splitted := strings.Split(lineSegment, ",")
	for _, s := range splitted {
		s = strings.TrimSpace(s)
		ss := strings.Split(s, " ")
		numberAsString, color := ss[0], ss[1]
		number, err := strconv.Atoi(numberAsString)
		if err != nil {
			panic(fmt.Sprintf("could not parse %s", numberAsString))
		}
		switch color {
		case "red":
			draw.Red = number
			break
		case "blue":
			draw.Blue = number
			break
		case "green":
			draw.Green = number
			break
		}
	}

	return draw
}

func parseLineToGame(lineSegment string) *Game {
	splitted := strings.Split(lineSegment, " ")
	id, err := strconv.Atoi(splitted[1])
	if err != nil {
		panic(err)
	}
	return &Game{Id: id, Draws: []Draw{}}
}
