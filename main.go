package main

import (
	"fmt"
	"./src/task1"
	"./src/task2"
	"./src/task3"
	"./src/task4"
	"./src/task5"
	"./src/task6"
	"./src/task7"
	"os"
	"io/ioutil"
	"encoding/json"
	"encoding/xml"
	"log"
	"path/filepath"
	"strings"
)

type Params1 struct {
	XMLName xml.Name  `xml:"task1params"`
	Width   int `json:"width" xml:"width"`
	Height  int `json:"height" xml:"height"`
	Symbol  string  `json:"symbol" xml:"symbol"`
}

type Params2 struct {
	Envelope1 task2.Envelope `json:"envelope1"`
	Envelope2 task2.Envelope `json:"envelope2"`
}

type Params3 struct {
	Triangles []task3.Triangle `json:"triangles"`
}

type Params4 struct {
	Number int `json:"number"`
}

type Params5 struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Params6 struct {
	Length int `json:"length"`
	Square int `json:"square"`
}

type Params7 struct {
	Context string `json:"context"` // context file name
}

type Params struct {
	XMLName xml.Name  `xml:"root"`
	Params1 []Params1 `json:"task1params" xml:"task1params"`
	Params2 []Params2 `json:"task2params" xml:"task2params"`
	Params3 []Params3 `json:"task3params" xml:"task3params"`
	Params4 []Params4 `json:"task4params" xml:"task4params"`
	Params5 []Params5 `json:"task5params" xml:"task5params"`
	Params6 []Params6 `json:"task6params" xml:"task6params"`
	Params7 []Params7 `json:"task7params" xml:"task7params"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Runs tasks\r\n", os.Args[0], "filename\r\n  filename - must contain JSON format")
		return
	}

	fileName := os.Args[1]
	extension := strings.ToLower(filepath.Ext(fileName))

	fmt.Println("FILE:", fileName)

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Fatal error: ", err)
	}

	var params Params
	switch extension {
	case ".json":
		err = json.Unmarshal(content, &params)
		if err != nil {
			log.Fatal("JSON error: ", err)
		}
	case ".xml":
		err = xml.Unmarshal(content, &params)
		if err != nil {
			log.Fatal("XML error: ", err)
		}
	default:
		fmt.Println("Unknown file extension")
		return
	}

	fmt.Printf("%#v\r\n", params)

	for _, p := range params.Params1 {
		fmt.Printf("\r\nTrying to execute task1 with params: %#v\r\n", p)
		fmt.Println(task1.ChessBoard(p.Width, p.Height, p.Symbol))
	}

	for _, p := range params.Params2 {
		fmt.Printf("\r\nTrying to execute task2 with params: %#v\r\n", p)
		fmt.Println(task2.CanEncloseEnvelopes(p.Envelope1, p.Envelope2))
	}

	for _, p := range params.Params3 {
		fmt.Printf("\r\nTrying to execute task3 with params: %#v\r\n", p)
		if task3.ValidateTriangles(p.Triangles) {
			fmt.Println("Triangles from max to min square:", task3.SortTriangles(p.Triangles))
		} else {
			fmt.Fprintln(os.Stderr, "One or more of triangles is not valid")
		}
	}

	for _, p := range params.Params4 {
		fmt.Printf("\r\nTrying to execute task4 with params: %#v\r\n", p)
		fmt.Println(task4.FindMaxPalindrome(p.Number))
	}

	for _, p := range params.Params5 {
		fmt.Printf("\r\nTrying to execute task5 with params: %#v\r\n", p)
		fmt.Println(task5.BestCountingSuccessTickets(p.Min, p.Max))
	}

	for _, p := range params.Params6 {
		fmt.Printf("\r\nTrying to execute task6 with params: %#v\r\n", p)
		if err = task6.WriteNumbers(p.Length, p.Square); err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("All is OK. Please see results in numbers.txt")
	}

	for _, p := range params.Params7 {
		fmt.Printf("\r\nTrying to execute task7 with params: %#v\r\n", p)
		restriction, err := task7.ParseContext(p.Context)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println(task7.Fib(restriction))

	}

}
