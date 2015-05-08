package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandConfig,
	commandDefault,
	commandEnv,
	commandLink,
	commandStart,
	commandStop,
	commandUnlink,
}

var commandConfig = cli.Command{
	Name:  "config",
	Usage: "",
	Description: `
`,
	Action: doConfig,
}

var commandDefault = cli.Command{
	Name:  "default",
	Usage: "",
	Description: `
`,
	Action: doDefault,
}

var commandEnv = cli.Command{
	Name:  "env",
	Usage: "",
	Description: `
`,
	Action: doEnv,
}

var commandLink = cli.Command{
	Name:  "link",
	Usage: "",
	Description: `
`,
	Action: doLink,
}

var commandStart = cli.Command{
	Name:  "start",
	Usage: "",
	Description: `
`,
	Action: doStart,
}

var commandStop = cli.Command{
	Name:  "stop",
	Usage: "",
	Description: `
`,
	Action: doStop,
}

var commandUnlink = cli.Command{
	Name:  "unlink",
	Usage: "",
	Description: `
`,
	Action: doUnlink,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doConfig(c *cli.Context) {
}

func doDefault(c *cli.Context) {
}

func doEnv(c *cli.Context) {
}

func doLink(c *cli.Context) {
}

func doStart(c *cli.Context) {
}

func doStop(c *cli.Context) {
}

func doUnlink(c *cli.Context) {
}
