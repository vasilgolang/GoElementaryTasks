package task3

import (
	"math"
	"sort"
	"fmt"
)

type Triangle struct {
	Vertices string
	A        float64
	B        float64
	C        float64
}

func (t Triangle) Validate() bool {
	return t.A > 0 && t.B > 0 && t.C > 0 && (t.A < t.B+t.C) && (t.B < t.A+t.C) && (t.C < t.A+t.B)
}

type SquareAndVerticales struct {
	Triangle *Triangle
	Square   float64
}

type Squares []SquareAndVerticales

func (s Squares) Len() int      { return len(s) }
func (s Squares) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Squares) Less(i, j int) bool {
	return s[i].Square < s[j].Square
}

func (t *Triangle) Square() float64 {
	p := (t.A + t.B + t.C) / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

func ValidateTriangles(triangles []Triangle) bool {
	for _, t := range triangles {
		if !t.Validate() {
			return false
		}
	}
	return true
}

func SortTriangles(triangles []Triangle) (names []string) {
	var tmp Squares
	for k := range triangles {
		tmp = append(tmp, SquareAndVerticales{
			Triangle: &triangles[k],
			Square:   triangles[k].Square(),
		})
		fmt.Printf("%#v Squeare: %v\r\n", triangles[k], triangles[k].Square())
	}

	sort.Sort(sort.Reverse(tmp)) // sort.Reverse() according to "Вывести треугольники в порядке _убывания_ их площади"

	for _, v := range tmp {
		names = append(names, v.Triangle.Vertices)
	}
	return
}
