package main

import (
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/zchee/camp/commands"
	"github.com/zchee/camp/version"
)

var AppHelpTemplate = `Usage: {{.Name}} {{if .Flags}}[OPTIONS] {{end}}COMMAND [arg...]

{{.Usage}}

Version: {{.Version}}{{if or .Author .Email}}

Author:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}
{{if .Flags}}
Options:
  {{range .Flags}}{{.}}
  {{end}}{{end}}
Commands:
  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
Run '{{.Name}} COMMAND --help' for more information on a command.
`

var CommandHelpTemplate = `Usage: docker-machine {{.Name}}{{if .Flags}} [OPTIONS]{{end}} [arg...]

{{.Usage}}{{if .Description}}

Description:

   {{.Description}}{{end}}{{if .Flags}}
Options:
   {{range .Flags}}
   {{.}}{{end}}{{ end }}
`

func main() {
	for _, f := range os.Args {
		if f == "-D" || f == "--debug" || f == "-debug" {
			os.Setenv("DEBUG", "1")
		}
	}
	cli.AppHelpTemplate = AppHelpTemplate
	cli.CommandHelpTemplate = CommandHelpTemplate
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"
	app.Usage = "Port forwarding for Docker. Inspired Basecamp's Pow."
	app.Author = "zchee"
	app.Email = "zcheeee@gmail.com"
	app.Commands = commands.Commands
	app.CommandNotFound = cmdNotFound

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug mode",
		},
	}

	app.Run(os.Args)
}

func cmdNotFound(c *cli.Context, command string) {
	log.Errorf(
		"%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name,
		command,
		c.App.Name,
		c.App.Name,
	)
}
