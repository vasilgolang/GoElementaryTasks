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
	"encoding/xml"
	"os"
	"strings"
	"path/filepath"
	"io/ioutil"
	"log"
	"encoding/json"
)




type Params struct {
	Params1 []task1.Params `json:"task1params" xml:"task1params"`
	Params2 []task2.Params `json:"task2params" xml:"task2params"`
	Params3 []task3.Params `json:"task3params" xml:"task3params"`
	Params4 []task4.Params `json:"task4params" xml:"task4params"`
	Params5 []task5.Params `json:"task5params" xml:"task5params"`
	Params6 []task6.Params `json:"task6params" xml:"task6params"`
	Params7 []task7.Params `json:"task7params" xml:"task7params"`
}

func main() {
	//var paramsTest []interface{}
	//p1 := task1.Params{Width: 30, Height: 20, Symbol: "Ж"}
	//p2 := task2.Params{task2.Envelope{5,5}, task2.Envelope{7,7}}
	//p21 := task2.Params{task2.Envelope{8,10}, task2.Envelope{7,7}}
	//
	//p5 := task5.Params{Min: 555, Max: 999999}
	//
	//paramsTest = append(paramsTest, p1)
	//paramsTest = append(paramsTest, p2)
	//paramsTest = append(paramsTest, p21)
	//paramsTest = append(paramsTest, p5)
	//
	//for k, v := range paramsTest {
	//	fmt.Printf("%d type(%T)\r\n", k, v)
	//
	//	//if fmt.Sprintf("%T") == "task1.Params" {
	//	switch v.(type) {
	//	case task1.Params:
	//		task1.Demo(v.(task1.Params))
	//	case task2.Params:
	//		task2.Demo(v.(task2.Params))
	//	default:
	//		fmt.Printf("Unknown params: %#v\r\n", v)
	//	}
	//
	//	//}
	//}
	//
	//return
	//result, err := task1.Run(p1)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println(result)
	//}
	//return
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
		//fmt.Println(task1.ChessBoard(p.Width, p.Height, p.Symbol))
	}

	for _, p := range params.Params2 {
		fmt.Printf("\r\nTrying to execute task2 with params: %#v\r\n", p)
		//fmt.Println(task2.CanEncloseEnvelopes(p.Envelope1, p.Envelope2))
	}

	for _, p := range params.Params3 {
		fmt.Printf("\r\nTrying to execute task3 with params: %#v\r\n", p)
		//if task3.ValidateTriangles(p.Triangles) {
		//	fmt.Println("Triangles from max to min square:", task3.SortTriangles(p.Triangles))
		//} else {
		//	fmt.Fprintln(os.Stderr, "One or more of triangles is not valid")
		//}
		task3.Demo(p)
	}

	for _, p := range params.Params4 {
		fmt.Printf("\r\nTrying to execute task4 with params: %#v\r\n", p)
		//fmt.Println(task4.FindMaxPalindrome(p.Number))
		task4.Demo(p)
	}

	for _, p := range params.Params5 {
		fmt.Printf("\r\nTrying to execute task5 with params: %#v\r\n", p)
		//fmt.Println(task5.BestCountingSuccessTickets(p.Min, p.Max))
		task5.Demo(p)
	}


	for _, p := range params.Params6 {
		fmt.Printf("\r\nTrying to execute task6 with params: %#v\r\n", p)
		task6.Demo(p)
		//if err = task6.WriteNumbers(p.Length, p.Square); err != nil {
		//	fmt.Println("Error:", err)
		//	continue
		//}
		//fmt.Println("All is OK. Please see results in numbers.txt")
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
