package main

import (
	"log"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	CmdConfig,
	CmdDefault,
	CmdEnv,
	CmdLink,
	CmdStart,
	CmdStop,
	CmdUnlink,
}

var CmdConfig = cli.Command{
	Name:        "config",
	Usage:       "",
	Description: `config`,
	Action:      doConfig,
}

var CmdDefault = cli.Command{
	Name:        "default",
	Usage:       "",
	Description: `default`,
	Action:      doDefault,
}

var CmdEnv = cli.Command{
	Name:        "env",
	Usage:       "",
	Description: `env`,
	Action:      doEnv,
}

var CmdLink = cli.Command{
	Name:        "link",
	Usage:       "",
	Description: `link`,
	Action:      doLink,
}

var CmdStart = cli.Command{
	Name:        "start",
	Usage:       "",
	Description: `start`,
	Action:      doStart,
}

var CmdStop = cli.Command{
	Name:        "stop",
	Usage:       "",
	Description: `stop`,
	Action:      doStop,
}

var CmdUnlink = cli.Command{
	Name:        "unlink",
	Usage:       "",
	Description: `unlink`,
	Action:      doUnlink,
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
