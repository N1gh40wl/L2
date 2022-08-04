package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func wgetgo(args []string) error {
	filepath := args[0]
	url:= args[1]
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	fmt.Println("Сайт сохранен в файл:", filepath)
	return nil
}

func main() {
	args := os.Args[1:]
	switch args[0] {
	case "wgetgo":
		err := wgetgo(args[1:])
		if err != nil {
			fmt.Errorf(err.Error())
		}
	default:
		fmt.Println("Команда не поддерживается")
	}

}
