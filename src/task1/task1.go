package task1

/*
1. Шахматная доска
Вывести шахматную доску с заданными размерами высоты и ширины, по принципу:
* * * * * *
* * * * * *
* * * * * *
* * * * * *
Входные параметры : длина, ширина, символ для отображения.
Выход : строка с представлением шахматной доски
*/

import (
	"errors"
	"unicode/utf8"
	"fmt"
)

type Params struct {
	Width  int `json:"width" xml:"width"`       // chess board width
	Height int `json:"height" xml:"height"`     // chess board height
	Symbol string  `json:"symbol" xml:"symbol"` // chess board symbol for white fields
}

func Demo(params Params) {
	fmt.Printf("Received params:\r\nWidth: %d\r\nHeight: %d\r\nSymbol: %s\r\n", params.Width, params.Height, params.Symbol)
	if result, err := Run(params); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:\r\n", result)
	}
}

// Returns error when params can't pass validation
func Validate(params Params) (err error) {
	// Check if width and height are positive numbers
	if params.Width < 0 || params.Height < 0 {
		return errors.New("Width and height must be more than 0")
	}
	// Check if symbol is rune
	if utf8.RuneCountInString(params.Symbol) != 1 {
		return errors.New("Symbol must be string with length equal 1")
	}
	// if all retrieved parameters are ok
	return nil
}

// Returns plain text chess board as a result and error when validation wasn't passed
func Run(params Params) (result string, err error) {
	if err := Validate(params); err != nil {
		return "", err
	}
	r, _ := utf8.DecodeRuneInString(params.Symbol) // r contains the first rune of the string
	return chessBoard(params.Width, params.Height, r), nil
}

// Returns text plain chess board
// Not exported function, because it mustn't run without validation
func chessBoard(width, height int, symbol rune) (board string) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// Detection "white" or "black" field of chess board
			if (i+j)%2 == 0 {
				board += string(symbol)
			} else {
				board += ` `
			}
		}
		board += "\r\n"
	}
	return
}
