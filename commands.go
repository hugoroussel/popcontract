package main

import(
	"gopkg.in/urfave/cli.v1"

 )

var commandAdmin cli.Command

func init() {

// Preliminary commands. Was not able to make "Action" work, the contract objetct & the functions are not recognized
/*
	commandAdmin = cli.Command{
		Name: "administrator",
		Aliases: []string{"admin"},
		Usage: "Setting up a party",
		Subcommands: []cli.Command{
		{
			Name: "name",
			Usage: "set name of party",
			ArgsUsage : "name of party",
			Action: contract.SetName,
		},
		{
			Name: "setup",
			Usage: "setup the party configuration",
			ArgsUsage : "place #organizers organizersArray duration",
			Action: contract.SetConfiguration,
		},
	},
	}
*/
}
