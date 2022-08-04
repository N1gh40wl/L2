package pattern

import (
	"fmt"
	"sort"
)

type sortAlg interface {
	sort(a *array)
}

type baloon struct {
}

func (b *baloon) sort(a *array) {
	fmt.Println("Сортировка пузырьком")
	for i := 0; i+1 < a.len; i++ {
		for j := 0; j+1 < a.len-i; j++ {
			if a.vals[j+1] < a.vals[j] {
				k := a.vals[j]
				a.vals[j] = a.vals[j+1]
				a.vals[j+1] = k
			}
		}
	}
}

type cocktail struct {
}

func (c *cocktail) sort(a *array) {
	fmt.Println("Шейкерная сортировка")
	left := 0
	right := a.len - 1
	for left <= right {
		for i := right; i > left; i-- {
			if a.vals[i-1] > a.vals[i] {
				k := a.vals[i-1]
				a.vals[i-1] = a.vals[i]
				a.vals[i] = k
			}
		}
		left++
		for i := left; i < right; i++ {
			if a.vals[i] > a.vals[i+1] {
				k := a.vals[i]
				a.vals[i] = a.vals[i+1]
				a.vals[i+1] = k
			}
		}
		right--

	}
}

type quick struct {
}

func (q *quick) sort(a *array) {
	fmt.Println("Быстрая сортировка")
	sort.Ints(a.vals)
}

type array struct {
	vals    []int
	sortAlg sortAlg
	len     int
}

func initArray(s sortAlg) *array {
	vals := make([]int, 0)
	return &array{
		vals:    vals,
		sortAlg: s,
		len:     0,
	}
}

func (a *array) setAlg(s sortAlg) {
	a.sortAlg = s
}

func (a *array) add(vals ...int) {
	for _, v := range vals {
		a.vals = append(a.vals, v)
		a.len++
	}
}

func (a *array) clear() {
	vals := make([]int, 0)
	a.vals = vals
	a.len = 0
}

func (a *array) show() {
	a.sort()
	fmt.Println(a.vals)
}

func (a *array) sort() {
	a.sortAlg.sort(a)
}

func RunStrategy() {
	test := []int{2, 5, 41, 1, 32, 21, 13, 4}
	fmt.Println("Исходный массив")
	fmt.Println(test)
	ballon := &baloon{}
	arr := initArray(ballon)
	arr.add(test...)
	arr.show()
	arr.clear()
	arr.add(test...)
	cocktail := &cocktail{}
	arr.setAlg(cocktail)
	arr.show()
	arr.clear()
	arr.add(test...)
	quick := &quick{}
	arr.setAlg(quick)
	arr.show()
}
