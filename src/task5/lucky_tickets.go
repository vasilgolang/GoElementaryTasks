package task5


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

type Result struct {
	Method int
	Count  int
}

func firstMethod(d []int) bool {
	return d[0]+d[1]+d[2] == d[3]+d[4]+d[5]
}

func secondMethod(d []int) bool {
	return d[0]+d[2]+d[4] == d[1]+d[3]+d[5]
}

func BestCountingSuccessTickets(min, max int) (r Result) {

	if min > 999999 || min < 0 || max > 999999 || min > max {
		return
	}

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

	// todo: to do something when both counter are equal
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
