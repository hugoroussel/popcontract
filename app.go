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
	"github.com/ethereum/go-ethereum/common"


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
	//fmt.Printf("Test : %x \n",PubKeyToAddress(key.PublicKey))

	//Organizers keys
	key_1, _ := crypto.GenerateKey()
	auth_1 := bind.NewKeyedTransactor(key_1)

	key_2, _ := crypto.GenerateKey()
	auth_2 := bind.NewKeyedTransactor(key_2)

	key_3, _ := crypto.GenerateKey()
	auth_3 := bind.NewKeyedTransactor(key_3)



	//Start simulated blockchain
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	alloc[auth_1.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	alloc[auth_2.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	alloc[auth_3.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	sim := backends.NewSimulatedBackend(alloc)


	fmt.Printf("Address of administrator %x \n", auth.From)

	p1 := crypto.PubkeyToAddress(key_1.PublicKey)
	p2 := crypto.PubkeyToAddress(key_2.PublicKey)
	p3 := crypto.PubkeyToAddress(key_3.PublicKey)



	// deploy contract
	addr, _, contract, err := DeployPopcontract(auth, sim)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}

	fmt.Printf("Address of the contract : %x", addr)

	contract.SetName(auth, "DEDIS party")

	// simulate miningNameOfParty(nil)
	fmt.Println("Mining...")
	sim.Commit()

	s := make([]common.Address,3)
  //fmt.Println("emp:", s)
	s[0] = p1
	s[1] = p2
	s[2] = p3

	//fmt.Println("emp:", s)

	i := big.NewInt(3)
	j := big.NewInt(60)

	contract.SetConfiguration(auth, "Lausanne" , i, s, j)

	sim.Commit()

	retrieval, err := contract.NameOfParty(nil)
	fmt.Printf("Name of party: %s \n", retrieval)

	place, err := contract.LocationOfParty(nil)
	fmt.Printf("Location of Party: %s \n", place)

	time, err := contract.EndOfParty(nil)
	fmt.Printf("Time %s \n", time)

	tab, err := contract.GetOrganizersAddresses(nil)
	fmt.Printf("Organizer 1 : %x \n", tab[0])
	fmt.Printf("Organizer 2 : %x \n", tab[1])
	fmt.Printf("Organizer 3 : %x \n", tab[2])

	//tab1, err := contract.SignedConfiguration(nil)
	//fmt.Printf("Signature 1 : %x \n", tab1[0])

	fmt.Printf("Msg signer 1 : %x \n", auth_1.From);

	contract.ConfigSignOrganizers(&bind.TransactOpts{
    From:     auth_1.From,
    Signer:   auth_1.Signer,
    GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(0),
	})

	fmt.Printf("Msg signer 2 : %x \n", auth_2.From);

	contract.ConfigSignOrganizers(&bind.TransactOpts{
    From:     auth_2.From,
    Signer:   auth_2.Signer,
    GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(0),
	})
	sim.Commit()

	fmt.Printf("Msg signer 3 : %x \n", auth_3.From);

	contract.ConfigSignOrganizers(&bind.TransactOpts{
		From:     auth_3.From,
		Signer:   auth_3.Signer,
		GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(0),
	})

	sim.Commit()

	tabSigned, err := contract.GetSignedConfiguration(nil)

	//fmt.Println(tabSigned)
	sim.Commit()

	fmt.Printf("Signed 1: %x \n", tabSigned[3])
	fmt.Printf("Signed 2: %x \n", tabSigned[4])
	fmt.Printf("Signed 3: %x \n", tabSigned[5])


	//Catches well the error if someone != from admin calls it
	contract.SignWholeConfiguration(&bind.TransactOpts{
		From:     auth_1.From,
    Signer:   auth_1.Signer,
    GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(0),
	})

	sim.Commit()

	signature, err := contract.Signed(nil)
	fmt.Printf("Whole config was signed : %t \n", signature)


	key_A, _ := crypto.GenerateKey()
	auth_A := bind.NewKeyedTransactor(key_A)
	key_B, _ := crypto.GenerateKey()
	auth_B := bind.NewKeyedTransactor(key_B)
	key_C, _ := crypto.GenerateKey()
	auth_C := bind.NewKeyedTransactor(key_C)



	attendees := make([]common.Address, 3)
	attendees[0] = auth_A.From
	attendees[1]=	auth_B.From
	attendees[2]= auth_C.From

	/*
	contract.DepositPublicKeys(&bind.TransactOpts{
		From:     auth_3.From,
		Signer:   auth_3.Signer,
		GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(0),
	}, _publicKeySet(auth_1, attendees))

	*/




}
