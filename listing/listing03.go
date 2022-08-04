package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil //type PathError = fs.PathError
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        //<nil>
	fmt.Println(err == nil) // сравниваются два разных типа *fs.PathError и nil -> false
}
