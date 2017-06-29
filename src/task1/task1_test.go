package task1

import (
	"testing"
	"unicode/utf8"
)

var tests = []struct {
	input     Params
	waitError bool
	board     string
}{
	{
		input:     Params{Width: 3, Height: 3, Symbol: "*"},
		waitError: true,
		board:     "* *\r\n * \r\n* *\r\n",
	},
	{
		input:     Params{Width: -1, Height: 10, Symbol: "*"},
		waitError: false,
	},
	{
		input:     Params{Width: 5, Height: -10, Symbol: "*"},
		waitError: false,
	},
	{
		input:     Params{Width: 10, Height: 10, Symbol: "ab"},
		waitError: false,
	},
}

func TestValidate(t *testing.T) {
	for _, test := range tests {
		if err := Validate(test.input); (err == nil) != test.waitError {
			t.Errorf("%v != %v", err, test.waitError)
		}
	}
}

func TestRun(t *testing.T) {
	for _, test := range tests {
		if board, err := Run(test.input); (err == nil) != test.waitError || board != test.board {
			t.Errorf("Wait: %v = %v\r\nWait board:\r\n%sGet board:\r\n%s", err, test.waitError, test.board, board)
		}
	}
}

func TestChessBoard(t *testing.T) {
	for _, test := range tests {
		if err := Validate(test.input); err == nil {
			r, _ := utf8.DecodeRuneInString(test.input.Symbol) // r contains the first rune of the string
			if board := chessBoard(test.input.Width, test.input.Height, r); board != test.board {
				t.Errorf("Wait board:\r\n%sGet board:\r\n%s", test.board, board)
			}
		}
	}
}
