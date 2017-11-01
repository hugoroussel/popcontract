package main

import(
	"gopkg.in/urfave/cli.v1"
 )

var commandAdmin, commandOrg, commandAttendee cli.Command

func init() {
/*

// Preliminary commands. Was not able to make "Action" work, the contract objetct & the functions are not recognized
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
		{
			Name: "sign",
			Usage: "sign whole configuration",
			Action: contract.SignWholeConfiguration,
		},
		{
			Name: "consensus",
			Usage: "choose final set to keep",
			Action: contract.publicKeyConsensus,
		},
	},
	}

	commandOrg = cli.Command{
	Name: "organizers",
	Aliases: []string{org},
	Usage: "Signing config, deposit of keypairs, reaching consensus",
	Subcommands: []cli.Command{
	{
		Name : "deposit",
		Usage : "deposit of participant public keys",
		ArgsUsage : "set of keys",
		Action : contract.depositPublicKeys,
	},
	{
		Name : "sign",
		Usage : "agree to the setup",
		Action : contract.configSignOrganizers,
	},
	},
	}

	commandAttendee = cli.Command{
		Name:    "attendee",
		Aliases: []string{"att"},
		Usage:   "attendee of a pop-party",
		Subcommands: []cli.Command{
			{
				Name:    "create",
				Aliases: []string{"cr"},
				Usage:   "create a private/public key pair",
				Action:  attCreate,
			},
			{
				Name:      "join",
				Aliases:   []string{"j"},
				Usage:     "join a poparty",
				ArgsUsage: "private_key party_hash",
				Action:    attJoin,
				Flags: []cli.Flag{
					cli.BoolTFlag{
						Name:  "yes,y",
						Usage: "disable asking",
					},
				},
			},
			{
				Name:      "sign",
				Aliases:   []string{"s"},
				Usage:     "sign a message and its context",
				ArgsUsage: "message context party_hash",
				Action:    attSign,
			},
			{
				Name:      "verify",
				Aliases:   []string{"v"},
				Usage:     "verifies a tag and a signature",
				ArgsUsage: "message context signature tag party_hash",
				Action:    attVerify,
			},
		},
	}
*/

}
