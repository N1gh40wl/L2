package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах


*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	path, _ := filepath.Abs(".")
	fmt.Print(path, " > ")
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		switch command[0] {

		case "cd":
			err := os.Chdir(command[1])
			if err != nil {
				fmt.Println("Incorrect path")
			}
		case "echo":
			for i := 1; i < len(command); i++ {
				fmt.Fprint(os.Stdout, command[i], " ")
			}
			fmt.Println()
		case "quit":
			return
		default:
			cmd := exec.Command(command[0], command[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Run()
		}

		path, _ = filepath.Abs(".")
		fmt.Print(path, " > ")
	}
}
