package main

import (
	"develop/dev06/task6/pkg"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	app := &cli.App{
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:  "gocut",
				Usage: "complete a task on the list",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "fields", Aliases: []string{"f"}},
					&cli.StringFlag{Name: "delimiter", Aliases: []string{"d"}},
					&cli.BoolFlag{Name: "separated", Aliases: []string{"s"}},
				},
				Action: func(ctx *cli.Context) error {
					if len(ctx.Args().Slice()) < 1 {
						os.Exit(1)
					}
					obj := pkg.CreateCut(ctx.String("fields"), ctx.String("delimiter"), ctx.Bool("separated"), ctx.Args().Slice())
					err := obj.Run()
					if err != nil {
						return err
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
