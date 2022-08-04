package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Ustr(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if _, err := strconv.Atoi(string(s[0])); err == nil {
		return "", errors.New("некорректная строка")
	}
	sSlice := strings.Split(s, "")
	result := ""
	for i := 0; i < len(sSlice); i++ {
		if string(sSlice[i]) == `\` {
			result += sSlice[i+1]
			i++
		} else {
			if n, err := strconv.Atoi(string(sSlice[i])); err == nil {
				for j := 0; j < n-1; j++ {
					result += string(sSlice[i-1])
				}
			} else {
				result += sSlice[i]
			}
		}

	}
	return result, nil
}

func main() {
	var s string
	fmt.Println("Введите слово:")
	fmt.Scanf("%s\n", &s)
	r, _ := Ustr(s)

	fmt.Println(r)
}
