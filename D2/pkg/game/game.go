package game

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Round struct {
	Red int
	Green int
	Blue int
}

type Game struct {
	Id int
	rounds []Round
}

func ParseLine(line string) Game {
	splitLine := strings.Split(line, ":")
	game := Game{
		Id: getIdFromLine(splitLine[0]),
		rounds: getRoundsFromLine(splitLine[1]),
	}
	return game
}

func getIdFromLine(line string) int {
	expression := regexp.MustCompile(`Game (\d+)`)
	match := expression.FindAllStringSubmatch(line, -1)
	id, err := strconv.ParseInt(match[0][1], 10, 0)
	if err != nil {
		fmt.Printf("error parsing id from line: %s", err)
	}
	return int(id)
}

func getRoundsFromLine(line string) []Round {
	rounds := []Round{}
	splitRounds := strings.Split(line, ";")
	fmt.Println("\nSplit Rounds: ", splitRounds)
	for _, round := range splitRounds {
		if round != "" {
			rounds = append(rounds, parseRound(round))
		}
	}
	return rounds
}

func parseRound(input string) Round {
	fmt.Println(input)
	round := Round{}
	matches := strings.Split(input, ",")
	countRegexp := regexp.MustCompile(`(\d+)|(\w+)`)
	for _, match := range matches {
		countMatches := countRegexp.FindAllStringSubmatch(match, -1)
		fmt.Println("submatches count", countMatches[0][0])
		fmt.Println("submatches color", countMatches[1][0])
		parsedCount, err := strconv.ParseInt(countMatches[0][0], 10, 0)
		if err != nil {
			fmt.Println("error: ", err)
		}
		switch countMatches[1][0] {
		case "red":
			round.Red = int(parsedCount)
		case "blue":
			round.Blue = int(parsedCount)
		case "green":
			round.Green = int(parsedCount)
		}
		
	}

	return round
}

func IsPossible(game Game, red, green, blue int) bool {
	isPossible := true
	for _, round := range game.rounds {
		if blue < round.Blue || red < round.Red || green < round.Green {
			isPossible = false
		}
	}
	return isPossible
}

type MinPossibleResults struct {
	Red int
	Blue int
	Green int
}

func MinPossible(game Game) MinPossibleResults {
	minPossible := MinPossibleResults{
		Red: 0,
		Blue: 0,
		Green: 0,
	}
	for _, round := range game.rounds {
		if minPossible.Blue < round.Blue {
			minPossible.Blue = round.Blue
		}
		if minPossible.Red < round.Red {
			minPossible.Red = round.Red
		}
		if minPossible.Green < round.Green {
			minPossible.Green = round.Green
		}
	}
	return minPossible
}