package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/ahaha0807/translation/src/script/translation"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "trans"
	app.Usage = "translation on command"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "english, e",
			Usage: "英語 to 日本語",
		},
	}

	app.Action = func(context *cli.Context) error {
		if len(context.Args()) == 0 {
			fmt.Println("USAGE: trans --[option] [words]")
			return nil
		}

		if !context.Bool("english") {
			result := translation.ToEnglish(context.Args().Get(0))
			fmt.Println(result)
			return nil
		} else {
			result := translation.ToJapanese(context.Args().Get(0))
			fmt.Println(result)
			return nil
		}
	}

	app.Run(os.Args)
}
