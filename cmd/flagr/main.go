package main

import (
	"github.com/cull-methi/flagr-cli/flagr/commands"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func main() {
	app.EnableBashCompletion = true
	app.Before = commands.BeforeMiddleware

	configureInfo()
	configureCommands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func configureInfo() {
	app.Name = "Flagr CLI"
	app.Usage = "A CLI for interacting with Flagr instances"
	app.Author = "Mitch Usher - usherm319@gmail.com"
	app.Version = "0.0.1"
}

func configureCommands() {
	subCmds := []cli.Command{}
	subCmds = append(subCmds, commands.QuerySubCommand())
	app.Commands = subCmds
}
