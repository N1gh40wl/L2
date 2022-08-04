package main

import (
	"develop/dev05/task5/pkg"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	app := &cli.App{
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:  "gogrep",
				Usage: "complete a task on the list",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "after", Aliases: []string{"A"}},
					&cli.StringFlag{Name: "before", Aliases: []string{"B"}},
					&cli.StringFlag{Name: "context", Aliases: []string{"C"}},
					&cli.BoolFlag{Name: "count", Aliases: []string{"c"}},
					&cli.BoolFlag{Name: "ignoreCase", Aliases: []string{"i"}},
					&cli.BoolFlag{Name: "invert", Aliases: []string{"v"}},
					&cli.BoolFlag{Name: "fixed", Aliases: []string{"F"}},
					&cli.BoolFlag{Name: "lineNum", Aliases: []string{"n"}},
				},
				Action: func(ctx *cli.Context) error {
					if len(ctx.Args().Slice()) < 1 {
						os.Exit(1)
					}
					
					obj := pkg.CreateGrep(ctx.String("after"), ctx.String("before"), ctx.String("context"),
						ctx.Bool("count"), ctx.Bool("ignoreCase"), ctx.Bool("invert"), ctx.Bool("fixed"),
						ctx.Bool("lineNum"), ctx.Args().Slice()[1:], ctx.Args().Slice()[0])

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
