package task5

import (
	"fmt"
	"errors"
)

/*
Есть 2 способа подсчёта счастливых билетов:
1. Простой — если на билете напечатано
шестизначное число, и сумма первых трёх цифр
равна сумме последних трёх, то этот билет
считается счастливым.
2. Сложный — если сумма чётных цифр билета равна
сумме нечётных цифр билета, то билет считается
счастливым.
Определить программно какой вариант подсчёта счастливых билетов даст их
большее количество на заданном интервале.
Входные параметры : границы min и max
Выход : элемент структурного типа с информацией о победившем методе и
количестве счастливых билетов для каждого способа подсчёта.
*/

type Params struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Result struct {
	Method int
	Count  int
}

// Returns error when params can't pass validation
func Validate(params Params) (err error) {
	if params.Min > params.Max {
		return errors.New(fmt.Sprintf("Min (%d) must be less than Max (%d)", params.Min, params.Max))
	}
	if params.Min > 999999 || params.Min < 0 || params.Max > 999999 {
		return errors.New(fmt.Sprintf("Min (%d) and Max (%d) must be in range from 0 to 999999", params.Min, params.Max))
	}
	return nil
}

func Run(params Params) (result Result, err error) {
	if err = Validate(params); err != nil {
		return
	}
	return BestCountingSuccessTickets(params.Min, params.Max), nil
}

func Demo(params []Params) {
	for _, param := range params {
		fmt.Printf("Received range. Min:%d, Max:%d\r\n", param.Min, param.Max)
		if result, err := Run(param); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Best method is %d, found %d lucky tickets\r\n", result.Method, result.Count)
		}
	}
}

func firstMethod(d []int) bool {
	return d[0]+d[1]+d[2] == d[3]+d[4]+d[5]
}

func secondMethod(d []int) bool {
	//return d[0]+d[2]+d[4] == d[1]+d[3]+d[5] // wrong understood task
	sumOdd, sumEven := 0, 0
	for _, v := range d {
		if (v % 2) == 0 {
			sumEven += v
		} else {
			sumOdd += v
		}
	}
	return sumOdd == sumEven
}

func BestCountingSuccessTickets(min, max int) (r Result) {

	firstMethodCounter := 0
	secondMethodCounter := 0
	for i := min; i <= max; i++ {
		digits := intToSliceInt(i)

		// Reverse slice
		for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
			digits[i], digits[j] = digits[j], digits[i]
		}

		digits = addBeginZeros(digits)

		if firstMethod(digits) {
			firstMethodCounter++
		}

		if secondMethod(digits) {
			secondMethodCounter++
		}
	}

	if firstMethodCounter > secondMethodCounter {
		return Result{
			Method: 1,
			Count:  firstMethodCounter,
		}
	} else {
		return Result{
			Method: 2,
			Count:  secondMethodCounter,
		}
	}

}

func intToSliceInt(num int) (nums []int) {
	for {
		nums = append(nums, num%10)
		num /= 10
		if num < 10 {
			nums = append(nums, num)
			break
		}
	}

	return
}

func addBeginZeros(nums []int) []int {
	if len(nums) < 6 {
		newSlice := make([]int, 6-len(nums), 6)
		newSlice = append(newSlice, nums...)

		return newSlice
	}
	return nums
}
