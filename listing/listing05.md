Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {// сравниваются два разных типа *main.customError и <nil>
		println("error")
		ret<nil>
	println("ok")
}
```

Ответ:
error
```
...

```
