package main

import (
	"os"
	"time"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hafen"
	app.Usage = "Install & run command line tools as a docker containers"

	app.Commands = []cli.Command{
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search for a cli tool",
			Action:  SearchAction,
		},

		{
			Name:   "info",
			Usage:  "Show cli tool's info",
			Action: InfoAction,
		},

		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "Install a cli tool",
			Action:  InstallAction,
		},

		{
			Name:   "uninstall",
			Usage:  "Uninstall a cli tool",
			Action: UninstallAction,
		},

		{
			Name:   "switch",
			Usage:  "Switch to a different version of a cli tool",
			Action: SwitchAction,
		},
	}

	app.Run(os.Args)
}

// SearchAction is for searching in registry for a cli tool by specified query
//
// Example output:
// ```
// Searching for docker-compose.. docker-compose not found.
// ```
func SearchAction(c *cli.Context) {
	print("Searching for ", c.Args()[0], "..")
	time.Sleep(300 * time.Millisecond)
	println("", c.Args()[0], "not found.")
}

// InfoAction is for fetching information from registry for a specified cli tool
//
// Example output:
// ```
// docker-compose:
// Download url (git):    https://github.com/docker/compose
// Download url (tar):    https://github.com/docker/compose/archive/{{version}}.tar.gz
// Download url (zip):    https://github.com/docker/compose/archive/{{version}}.zip
// Docker build context:  https://github.com/heimathafen/recipes; /docker-compose
// Run script:            https://github.com/heimathafen/recipes; /docker-compose/run.sh
// ```
func InfoAction(c *cli.Context) {
	println(c.Args()[0] + ":")
	println("Download url (git):   ", "https://github.com/docker/compose")
	println("Download url (tar):   ", "https://github.com/docker/compose/archive/{{version}}.tar.gz")
	println("Download url (zip):   ", "https://github.com/docker/compose/archive/{{version}}.zip")
	println("Docker build context: ", "https://github.com/heimathafen/recipes;", "/docker-compose")
	println("Run script:           ", "https://github.com/heimathafen/recipes;", "/docker-compose/run.sh")
}

// InstallAction is for installing a specified cli tool as a docker container
//
// Example output:
// ```
// Installing docker-compose.. Done.
// ```
func InstallAction(c *cli.Context) {
	print("Installing ", c.Args()[0], "..")
	time.Sleep(500 * time.Millisecond)
	println(" Done.")
}

// UninstallAction is for uninstalling a specified cli tool
//
// Example output:
// ```
// Uninstalling docker-compose.. Done.
// ```
func UninstallAction(c *cli.Context) {
	print("Uninstalling ", c.Args()[0], "..")
	time.Sleep(200 * time.Millisecond)
	println(" Done.")
}

// SwitchAction is for switching a specified cli tool to specified version if
// it is installed
//
// Example output:
// ```
// Switching to docker-compose==1.3.1.. Done.
// ```
func SwitchAction(c *cli.Context) {
	print("Switching to ", c.Args()[0], "==", c.Args()[1], "..")
	time.Sleep(100 * time.Millisecond)
	println(" Done.")
}
