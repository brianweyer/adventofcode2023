package calibration

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var NUMBERS = map[string] int {
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}



func ParseLine(line string) int {
	first := findFirstNumber(line)
	last := findLastNumber(line)
	parsed, err := strconv.ParseInt(fmt.Sprintf("%s%s", first, last), 10, 0)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	return int(parsed)
}

func findFirstNumber(line string) string {
  findFirstExpression := regexp.MustCompile(`(\d)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)`)
	match := replaceNumberWordsStrings(findFirstExpression.FindString(line))
	return match
}

func findLastNumber(line string) string {
	var last string
	var furthestIndex int
	for word, number := range NUMBERS {

		wordLastIndex := strings.LastIndex(line, word)
		if wordLastIndex >= furthestIndex {
			last = word
			furthestIndex = wordLastIndex
		}
		numberLastIndex := strings.LastIndex(line, fmt.Sprint(number))
		if numberLastIndex >= furthestIndex {
			last = fmt.Sprint(number)
			furthestIndex = numberLastIndex
		}
	}
 
	return replaceNumberWordsStrings(last)
}

func replaceNumberWordsStrings(match string) string {
	switch match {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}
	return match
}
