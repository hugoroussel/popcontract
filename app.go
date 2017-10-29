package main

import (
	"fmt"
	"log"
	"math/big"
	"gopkg.in/urfave/cli.v1"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"


	"github.com/dedis/cothority/cosi/check"
	_ "github.com/dedis/cothority/cosi/protocol"
	_ "github.com/dedis/cothority/cosi/service"

)



func main() {

	appCli := cli.NewApp()
	appCli.Name = "Proof-of-personhood party"
	appCli.Usage = "Handles party-creation, finalizing, pop-token creation, and verification"
	appCli.Version = "0.1"
	appCli.Commands = []cli.Command{}
	appCli.Commands = []cli.Command{
		commandAdmin,
		{
			Name:      "check",
			Aliases:   []string{"c"},
			Usage:     "Check if the servers in the group definition are up and running",
			ArgsUsage: "group.toml",
			Action: func(c *cli.Context) error {
				return check.Config(c.Args().First(), false)
			},
		},
	}


	//Generates Key
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	//Start simulated blockchain
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	sim := backends.NewSimulatedBackend(alloc)

	// deploy contract
	addr, _, contract, err := DeployPopcontract(auth, sim)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}

	fmt.Println("Address of the contract :", addr)

	//Test retrieval of the name of the party
	fmt.Println("Setting name of party")
	contract.SetName(auth, "DEDIS party")

	// simulate mining
	fmt.Println("Mining...")
	sim.Commit()

	retrieval, err := contract.NameOfParty(nil)
	fmt.Println("The name of your party is : ",retrieval)


}
