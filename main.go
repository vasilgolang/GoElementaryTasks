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
)

func main() {
	fmt.Println("\r\n1) Chess board")
	fmt.Println(task1.ChessBoard(0, 0, '*'))
	fmt.Println(task1.ChessBoard(-1, 0, '*'))
	fmt.Println(task1.ChessBoard(2, 2, '*'))
	fmt.Println(task1.ChessBoard(3, 2, '*'))
	fmt.Println(task1.ChessBoard(10, 10, 'X'))

	fmt.Println("\r\n2) Envelopes")
	e1 := task2.Envelope{3, 5}
	e2 := task2.Envelope{3.1, 5.1}
	fmt.Println(task2.CanEncloseEnvelopes(e1, e2))

	fmt.Println("\r\n3) Triangles")
	var triangles = []task3.Triangle{
		{"ABC", 10, 20, 22.36},
		{"DEF", 100, 200, 223.6},
		{"KLM", 1, 2, 2.236},
		{"OPQ", 3, 3, 3},
	}
	fmt.Println(task3.SortTriangles(triangles))

	fmt.Println("\r\n4) Palindrome")
	fmt.Println(task4.FindPalindrome(9123456789))
	fmt.Println(task4.FindPalindrome(12234444437))
	fmt.Println(task4.FindPalindrome(12388321))
	fmt.Println(task4.FindPalindrome(2505))
	fmt.Println(task4.FindPalindrome(12834432))
	fmt.Println(task4.FindPalindrome(55))
	fmt.Println(task4.FindPalindrome(505))

	fmt.Println("\r\n5) Success Tickets")
	fmt.Println(task5.BestCountingSuccessTickets(1, 999999))

	fmt.Println("\r\n6) Number sequence")
	fmt.Println(task6.WriteNumbers(10, 1000))
	fmt.Println(task6.WriteNumbers(10, -1000))

	fmt.Println("\r\n7) Fibonacci")
	restriction, err := task7.ParseContext(`context1.txt`)
	if err != nil {
		fmt.Fprintln(os.Stderr,    err)
	}
	fmt.Println(task7.Fib(restriction))
	fmt.Println()

	restriction, err = task7.ParseContext(`context2.txt`)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(task7.Fib(restriction))

}
