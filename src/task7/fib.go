package task7

import (
	"io/ioutil"
	"strings"
	"github.com/pkg/errors"
	"strconv"
	"fmt"
)

/*
7. Ряд Фибоначчи для диапазона
Вывести все числа Фибоначчи, которые удовлетворяют переданному в функцию
ограничению: находятся в указанном диапазоне, либо имеют указанную длину.
Входные параметры : файл context со значениями min и max, либо с полем length
Выход : срез сгенерированных чисел
*/

/* К сожалению, не понял формат входящих параметров. Имеется ввиду физический файл с именем
"context" и в нем должно быть 2 числа min и max разделенных пробелами, либо одно число length?
Поэтому пока что делаю функцию которая принимает эти 3 параметра.
 */

type restriction struct {
	min, max, length int
}

func parseLine(s string) (result int, err error) {
	intVal, err := strconv.Atoi(s)
	if err != nil {
		return -1, errors.New("Restriction must be an integer")
	}
	return intVal, nil
}

func ParseContext(filename string) (r restriction, err error) {
	r = restriction{-1, -1, -1}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) > 2 {
		return r, errors.New("File " + filename + " must contain 1 or 2 strings")
	}
	for k, v := range lines {
		lines[k] = strings.TrimSpace(v)
	}
	if len(lines) == 2 {
		if r.min, err = parseLine(lines[0]); err != nil {
			return r, err
		}
		if r.max, err = parseLine(lines[1]); err != nil {
			return r, err
		}
	} else {
		if r.length, err = parseLine(lines[0]); err != nil {
			return r, err
		}
	}

	return
}

func Fib(r restriction) (res []int) {
	fmt.Printf("restriction:%#v\r\n", r)
	a, b := 0, 1
	for i := 0; ; i++ {
		a, b = b, a+b

		if r.length == -1 {
			if a < r.min {
				continue
			}
			if a > r.max {
				break
			}

		} else {
			l := len(strconv.Itoa(a))
			if l < r.length {
				continue
			}
			if l > r.length {
				break
			}

		}
		res = append(res, a)

	}
	return
}
