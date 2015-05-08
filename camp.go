package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "camp"
	app.Version = Version
	app.Usage = ""
	app.Author = "zhee"
	app.Email = "zcheeee@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
