package main

import (
	"bufio"

	"fmt"
	"os"

	"github.com/brianweyer/adventofcode2023/D2/pkg/game"
)

func main() {
  scanner := readPuzzleInput()

  total := 0
  impossible := 0
  possible := 0
  minPossiblePowers := 0

  for scanner.Scan() {
    lineText := scanner.Text()
    parsedGame := game.ParseLine(lineText)
    fmt.Println(parsedGame)
    if game.IsPossible(parsedGame, 12, 13, 14) {
      total += parsedGame.Id
      possible += 1
    } else {
      impossible += 1
    }

    minPossible := game.MinPossible(parsedGame)
    minPossiblePowers += minPossible.Blue * minPossible.Red * minPossible.Green
  }

  fmt.Println("Possible: ", possible)
  fmt.Println("Impossible: ", impossible)
  fmt.Println("Min Possible Powers: ", minPossiblePowers)
  fmt.Println("TOTAL: ", total)
}

func readPuzzleInput() *bufio.Scanner {
  puzzleInput, err := os.Open("./puzzle_input.txt")
  if err != nil {
    panic(err)
  }
  defer puzzleInput.Close()

  return bufio.NewScanner(puzzleInput)
}