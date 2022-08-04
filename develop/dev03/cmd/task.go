package main

import (
	"develop/dev03/task3/pkg"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// imports as package "cli"

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	app := &cli.App{
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:  "gosort",
				Usage: "complete a task on the list",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "key", Aliases: []string{"k"}},
					&cli.BoolFlag{Name: "number", Aliases: []string{"n"}},
					&cli.BoolFlag{Name: "reverse", Aliases: []string{"r"}},
					&cli.BoolFlag{Name: "unique", Aliases: []string{"u"}},
				},
				Action: func(ctx *cli.Context) error {
					if len(ctx.Args().Slice()) < 1 {
						os.Exit(1)
					}
					obj := pkg.CreateSort(ctx.String("key"), ctx.Bool("number"), ctx.Bool("reverse"), ctx.Bool("unique"), ctx.Args().Slice())

					err := obj.Run()
					if err != nil {
						fmt.Errorf("Error:", err)
					}
					return obj.Output()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
