package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/brianweyer/adventofcode2023/D1/pkg/calibration"
)

func main() {
  // Open calibration file
  puzzleInput, err := os.Open("./puzzle_input.txt")
  if err != nil {
    panic(err)
  }
  defer puzzleInput.Close()

  scanner := bufio.NewScanner(puzzleInput)

  total := 0

  for scanner.Scan() {
    lineText := scanner.Text()
    lineNumber := calibration.ParseLine(lineText)
    total += int(lineNumber)
  }

  fmt.Println("TOTAL: ", total)
}

