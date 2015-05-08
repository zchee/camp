package commands

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/zchee/camp/utils"
)

var Commands = []cli.Command{
	{
		Name:        "config",
		Usage:       "",
		Description: `config`,
		Action:      cmdConfig,
	},
	{
		Name:        "default",
		Usage:       "",
		Description: `default`,
		Action:      cmdDefault,
	},
	{
		Name:        "env",
		Usage:       "",
		Description: `env`,
		Action:      cmdEnv,
	},
	{
		Name:        "link",
		Usage:       "",
		Description: `link`,
		Action:      cmdLink,
	},
	{
		Name:        "start",
		Usage:       "Start camp server.",
		Description: "Argument is a container name.",
		Action:      cmdStart,
		Flags: []cli.Flag{
			cli.StringFlag{
				EnvVar: "CAMP_USERNAME",
				Name:   "username, u",
				Usage:  "SSH User Name",
				Value:  "docker",
			},
			cli.StringFlag{
				EnvVar: "CAMP_PASSWORD",
				Name:   "password, p",
				Usage:  "SSH Password",
				Value:  "tcuser",
			},
			cli.StringFlag{
				EnvVar: "CAMP_SERVER",
				Name:   "server, s",
				Usage:  "Server Address",
				Value:  utils.GetDockerHost(),
			},
			cli.StringFlag{
				EnvVar: "CAMP_LOCAL",
				Name:   "local, l",
				Usage:  "Local Address and Port",
				Value:  "",
			},
			cli.StringFlag{
				EnvVar: "CAMP_REMOTE",
				Name:   "remote, r",
				Usage:  "Remote Address and Port",
				Value:  "",
			},
		},
	},
	{
		Name:   "start",
		Usage:  "Start camp server",
		Action: cmdStart,
	},
	{
		Name:        "stop",
		Usage:       "",
		Description: `stop`,
		Action:      cmdStop,
	},
	{
		Name:        "unlink",
		Usage:       "",
		Description: `unlink`,
		Action:      cmdUnlink,
	},
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func cmdConfig(c *cli.Context) {
}

func cmdDefault(c *cli.Context) {
}

func cmdEnv(c *cli.Context) {
}

func cmdLink(c *cli.Context) {
}

func cmdStop(c *cli.Context) {
}

func cmdUnlink(c *cli.Context) {
}
