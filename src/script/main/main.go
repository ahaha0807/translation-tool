package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/ahaha0807/translation/src/script/translation"
)

func main() {
	app := cli.NewApp()
	app.Name = "translation-ahaha0807"
	app.Usage = "translation on command"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		translation.ToEnglish("")
		return nil
	}

	app.Run(os.Args)
}
