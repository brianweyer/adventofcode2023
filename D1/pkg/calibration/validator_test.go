package calibration

import "testing"

func TestParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample validation 1", args: args{line: "1abc2"}, want: 12},
		{name: "sample validation 2", args: args{line: "pqr3stu8vwx"}, want: 38},
		{name: "sample validation 3", args: args{line: "a1b2c3d4e5f"}, want: 15},
		{name: "sample validation 4", args: args{line: "treb7uchet"}, want: 77},
		{name: "sample validation 5", args: args{line: "two1nine"}, want: 29},
		{name: "sample validation 6", args: args{line: "eightwothree"}, want: 83},
		{name: "sample validation 7", args: args{line: "abcone2threexyz"}, want: 13},
		{name: "sample validation 8", args: args{line: "xtwone3four"}, want: 24},
		{name: "sample validation 9", args: args{line: "4nineeightseven2"}, want: 42},
		{name: "sample validation 10", args: args{line: "zoneight234"}, want: 14},
		{name: "sample validation 11", args: args{line: "7pqrstsixteen"}, want: 76},
		{name: "sample validation 12", args: args{line: "2fiveshtds4oneightsjg"}, want: 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLine(tt.args.line); got != tt.want {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
