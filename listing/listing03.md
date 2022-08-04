Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil //type PathError = fs.PathError-> PathError записывает ошибку, а также операцию и путь к файлу, вызвавшие ее.(Структра с полями)
	return err
}

func main() {
	err := Foo()
	fmt.Println(err) 		//<nil>
	fmt.Println(err == nil)	// сравниваются два разных типа *fs.PathError и nil -> false
}
```

Ответ:
nil
false
```
...

```
