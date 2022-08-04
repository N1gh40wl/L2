package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func makeKey(val string) string {
	letters := strings.Split(val, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func checkSlice(s string, sl []string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}

func sortMap(NSmap map[string][]string) map[string][]string {
	result := make(map[string][]string)
	for _, v := range NSmap {
		result[v[0]] = v
		sort.Strings(v)
	}
	return result
}

func FindQuantities(s []string) map[string][]string {
	result := make(map[string][]string)
	var falseKeys []string
	for i := 0; i < len(s); i++ {
		val := strings.ToLower(s[i])
		val = makeKey(val)
		if checkSlice(val, falseKeys) {
			result[val] = append(result[val], s[i])
		} else {
			falseKeys = append(falseKeys, val)
			result[val] = []string{s[i]}
		}
	}
	result = sortMap(result)
	return result
}

func main() {

}
