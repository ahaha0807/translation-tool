package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
)

func main () {
	app := cli.NewApp()
	app.Name = "translation-ahaha0807"
	app.Usage = "translation on command"
	app.Version = "0.0.1"

	app.Action = func (context *cli.Context) error {
		fmt.Println("Hello world on cli")
		return nil
	}

	app.Run(os.Args)
}
