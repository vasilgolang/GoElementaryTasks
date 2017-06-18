package task6

import (
	"math"
	"os"
	"strconv"
	"io/ioutil"
	"strings"
	"github.com/pkg/errors"
)

/*
Вывести в файл через запятую ряд длиной n, состоящий из натуральных чисел,
квадрат которых не меньше заданного m.
Входные параметры : длина и значение минимального квадрата
Выход : nil если сохранение удалось и err в противном случае
*/

func WriteNumbers(length int, square int) error {
	if square < 0 {
		return errors.New("Square can't be less than 0")
	}

	f, err := os.Create(`numbers.txt`)
	if err != nil {
		return err
	}
	defer f.Close()

	var numsS []string
	i := int(math.Sqrt(float64(square))) // маленькая оптимизация начального значения i :)
	for j := 0; j <= length; j, i = j+1, i+1 {
		if i*i < square {
			continue
		}

		numsS = append(numsS, strconv.Itoa(i))
	}

	return ioutil.WriteFile(`numbers.txt`, []byte(strings.Join(numsS, `,`)), 0777)
}
