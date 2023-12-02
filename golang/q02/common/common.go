package common

import (
	"strconv"
	"strings"
)

type Round struct {
	Red, Green, Blue int
}

type Game struct {
	Rounds []Round
}

func (this Game) MaxColors() Round {
	result := Round{}

	for _, round := range this.Rounds {
		if round.Red > result.Red {
			result.Red = round.Red
		}

		if round.Green > result.Green {
			result.Green = round.Green
		}

		if round.Blue > result.Blue {
			result.Blue = round.Blue
		}
	}

	return result
}

func NewGame(input string) Game {
	game := Game{}
	startIdx := strings.Index(input, ":") + 1
	gameParts := strings.Split(input[startIdx:], ";")

	for _, part := range gameParts {
		game.Rounds = append(game.Rounds, NewRound(part))
	}

	return game
}

func NewRound(input string) Round {
	roundParts := strings.Split(input, ",")
	round := Round{}

	for _, roundPart := range roundParts {
		cubeParts := strings.Split(strings.Trim(roundPart, " "), " ")
		amount, err := strconv.Atoi(cubeParts[0])
		if err != nil {
			panic(err)
		}

		switch cubeParts[1] {
		case "red":
			round.Red = amount
		case "green":
			round.Green = amount
		case "blue":
			round.Blue = amount
		}
	}

	return round
}

func ParseGames(lines []string) []Game {
	games := []Game{}

	for _, line := range lines {
		games = append(games, NewGame(line))
	}

	return games
}
