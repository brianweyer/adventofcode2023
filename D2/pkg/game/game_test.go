package game

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIdFromLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIdFromLine(tt.args.line); got != tt.want {
				t.Errorf("getIdFromLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRoundsFromLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []Round
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRoundsFromLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRoundsFromLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseRound(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want Round
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRound(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPossible(t *testing.T) {
	type args struct {
		game  Game
		red   int
		green int
		blue  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPossible(tt.args.game, tt.args.red, tt.args.green, tt.args.blue); got != tt.want {
				t.Errorf("IsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinPossible(t *testing.T) {
	type args struct {
		game Game
	}
	tests := []struct {
		name string
		args args
		want MinPossibleResults
	}{
		name: "MinPossible case #1",
		args: args{
			game: Game{
				id: "100",
				rounds: []Round{{Blue: 1, Red: 3, Green: 12}, {Blue: 13, Red: 14, Green 2}}
			}
		}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPossible(tt.args.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
