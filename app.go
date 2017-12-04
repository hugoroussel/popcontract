package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"

	"math/big"
	"os"

	"strings"

	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/dedis/onet.v1/app"
	"gopkg.in/dedis/onet.v1/log"
	"gopkg.in/dedis/onet.v1/network"
	cli "gopkg.in/urfave/cli.v1"
)

var nonce int64 = 229

func main() {
	appCli := cli.NewApp()
	appCli.Name = "Proof-of-personhood party"
	appCli.Usage = "Handles party-creation, finalizing, pop-token creation, and verification"
	appCli.Version = "0.1"
	appCli.Commands = []cli.Command{}
	appCli.Commands = []cli.Command{
		{
			Name:    "organizer",
			Aliases: []string{"org"},
			Usage:   "Organising a PoParty",
			Subcommands: []cli.Command{
				{
					Name:      "link",
					Aliases:   []string{"l"},
					Usage:     "deploy new pop contract",
					ArgsUsage: "private key and network path",
					Action:    orgLink,
				},
				{
					Name:      "config",
					Aliases:   []string{"c"},
					Usage:     "stores the configuration",
					ArgsUsage: "private key and pop_desc.toml",
					Action:    orgConfig,
				},
				{
					Name:      "public",
					Aliases:   []string{"p"},
					Usage:     "stores one or more public keys during the party",
					ArgsUsage: "key1,key2,key3 party_hash",
					Action:    orgPublic,
				},
				{
					Name:      "final",
					Aliases:   []string{"f"},
					Usage:     "finalizes the party",
					ArgsUsage: "party_hash",
					Action:    orgFinal,
				},
			},
		}}
	appCli.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "debug,d",
			Value: 0,
			Usage: "debug-level: 1 for terse, 5 for maximal",
		},
		cli.StringFlag{
			Name:  "config,c",
			Value: "~/.config/cothority/pop",
			Usage: "The configuration-directory of pop",
		},
	}
	appCli.Before = func(c *cli.Context) error {
		log.SetDebugVisible(c.Int("debug"))
		return nil
	}
	appCli.Run(os.Args)

}

type Config struct {
	//name of config-file
	name string
	//path to ipc
	network string
	//Address of contract
	Address string
}

type popDescToml struct {
	// config-file name
	name string
	//Location of party
	Location string
	//number of organizers
	NumberOfOrganizers int64
	//organizers Adresses in Ethereum format
	OrganizersAddresses []common.Address
	//Duration of party in minutes
	Deadline int64
}

type PopDesc struct {
	// config-file name
	name string
	//Location of party
	Location string
	//number of organizers
	NumberOfOrganizers int64
	//organizers Adresses in Ethereum format
	OrganizersAddresses []common.Address
	//Duration of party in minutes
	Deadline int64
	//Address of contract
	Address string
}

//connect to contract
func orgLink(c *cli.Context) error {
	log.Lvl3("Org: Link")
	if c.NArg() < 1 {
		log.Fatal("Please provide valid private key and geth.ipc path")
	}
	cfg := getConfig(c)
	key, _ := crypto.HexToECDSA(c.Args().Get(0))
	auth := bind.NewKeyedTransactor(key)
	network := c.Args().Get(1)
	conn, err := ethclient.Dial(network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	addr, _, _, err := DeployPopcontract(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(3000000),
		GasPrice: big.NewInt(90000000000),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(nonce),
	}, conn)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}
	nonce++
	cfg.Address = addr.String()
	cfg.network = network
	fmt.Printf("Successfully linked with %x : \n", addr)
	//log.Lvl3("Successfully linked with", addr)
	//write is not working here
	cfg.write()
	return nil

}

//Deploy & configure a new pop-party
func orgConfig(c *cli.Context) error {
	log.Lvl3("Org: Config")
	if c.NArg() < 2 {
		log.Fatal(`Please give valid admin private key & pop_desc.toml `)
	}
	cfg := getConfig(c)
	if cfg.Address == "" {
		log.Fatal("No address")
		return errors.New("No address found - please link first")
	}
	/*
		desc := &service.PopDesc{}
		pdFile := c.Args().Get(1)
		buf, err := ioutil.ReadFile(pdFile)
		log.ErrFatal(err, "While reading", pdFile)
		err = decodePopDesc(string(buf), desc)
		log.ErrFatal(err, "While decoding", pdFile)
		//log.ErrFatal(client.StoreConfig(cfg.Address, desc, cfg.OrgPrivate))
	*/
	var desc popDescToml
	if _, err := toml.Decode(c.Args().Get(1), &desc); err != nil {
		fmt.Println("Error decoding toml file.")
	}
	conn, err := ethclient.Dial(cfg.network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	contract, err := NewPopcontract(common.HexToAddress(cfg.Address), conn)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v \n", err)
	}
	fmt.Println("Connected to contract.. \n")
	key, _ := crypto.HexToECDSA(c.Args().Get(0))
	auth := bind.NewKeyedTransactor(key)

	txe, err := contract.SetConfiguration(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(3000000),
		GasPrice: big.NewInt(90000000000),
		Nonce:    big.NewInt(nonce),
	}, desc.name, desc.Location, big.NewInt(desc.NumberOfOrganizers), desc.OrganizersAddresses, big.NewInt(desc.Deadline))
	if err != nil {
		log.Fatalf("could not set configuration: %v", err)
	}
	nonce++
	fmt.Printf("Configuration set. Transaction : %v \n", txe)
	cfg.write()
	return nil
}

// adds a public key to the list
func orgPublic(c *cli.Context) error {
	if c.NArg() < 2 {
		log.Fatal("Please give a private key and the public keys to add")
	}
	log.Lvl3("Org: Adding public keys", c.Args().Get(1))
	str := c.Args().Get(1)
	if !strings.HasPrefix(str, "[") {
		str = "[" + str + "]"
	}
	// TODO: better cleanup rules
	str = strings.Replace(str, "\"", "", -1)
	str = strings.Replace(str, "[", "", -1)
	str = strings.Replace(str, "]", "", -1)
	str = strings.Replace(str, "\\", "", -1)
	log.Lvl3("Niceified public keys are:\n", str)
	keys := strings.Split(str, ",")
	cfg := getConfig(c)
	keyset := make([][32]byte, len(keys))
	for i := 0; i < len(keys); i++ {
		keyA := []byte(keys[i])
		copy(keyset[i][:], keyA)
	}
	conn, err := ethclient.Dial(cfg.network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	key, _ := crypto.HexToECDSA(c.Args().Get(0))
	auth := bind.NewKeyedTransactor(key)
	contract, err := NewPopcontract(common.HexToAddress(cfg.Address), conn)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v \n", err)
	}
	txe, err := contract.DepositPublicKeys(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(3000000),
		GasPrice: big.NewInt(90000000000),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(nonce),
	}, keyset)
	if err != nil {
		log.Fatalf("could not deposit keyset: %v", err)
	}
	nonce++
	fmt.Printf("Keyset Added. Transaction : %v \n", txe)
	cfg.write()
	return nil
}

//Organizators sign configuration
func sign(c *cli.Context) error {
	if c.NArg() < 1 {
		log.Fatal(`Please give valid private key `)
	}
	cfg := getConfig(c)
	conn, err := ethclient.Dial(cfg.network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	key, _ := crypto.HexToECDSA(c.Args().Get(0))
	auth := bind.NewKeyedTransactor(key)
	contract, err := NewPopcontract(common.HexToAddress(cfg.Address), conn)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v \n", err)
	}
	txe, err := contract.ConfigSignOrganizers(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(3000000),
		GasPrice: big.NewInt(90000000000),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(nonce),
	})
	if err != nil {
		log.Fatalf("could not sign contract: %v", err)
	}
	nonce++
	fmt.Printf("Configuration signed. Transaction : %v \n", txe)
	return nil
}

//Administrator sign whole configuration
func signAdmin(c *cli.Context) error {
	if c.NArg() < 1 {
		log.Fatal(`Please give valid private key `)
	}
	cfg := getConfig(c)
	conn, err := ethclient.Dial(cfg.network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	key, _ := crypto.HexToECDSA(c.Args().Get(0))
	auth := bind.NewKeyedTransactor(key)
	contract, err := NewPopcontract(common.HexToAddress(cfg.Address), conn)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v \n", err)
	}
	txe, err := contract.SignWholeConfiguration(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(3000000),
		GasPrice: big.NewInt(90000000000),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(nonce),
	})
	if err != nil {
		log.Fatalf("could not sign contract: %v", err)
	}
	nonce++
	fmt.Printf("Whole configuration signed. Transaction : %v \n", txe)
	return nil
}

//reach consensus
func orgFinal(c *cli.Context) error {
	if c.NArg() < 1 {
		log.Fatal(`Please give valid private key `)
	}
	cfg := getConfig(c)
	conn, err := ethclient.Dial(cfg.network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	key, _ := crypto.HexToECDSA(c.Args().Get(0))
	auth := bind.NewKeyedTransactor(key)
	contract, err := NewPopcontract(common.HexToAddress(cfg.Address), conn)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v \n", err)
	}
	txe, err := contract.PublicKeyConsensus(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(3000000),
		GasPrice: big.NewInt(90000000000),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(nonce),
	})
	if err != nil {
		log.Fatalf("Could not reach consensus: %v", err)
	}
	nonce++
	fmt.Printf("Asking for consensus. Transaction : %v \n", txe)

	return nil
}

// getConfigClient returns the configuration and a client-structure.
func getConfig(c *cli.Context) *Config {
	cfg, err := newConfig(path.Join(c.GlobalString("config"), "config.bin"))
	log.ErrFatal(err)
	return cfg
}

// newConfig tries to read the config and returns an organizer-
// config if it doesn't find anything.
func newConfig(fileConfig string) (*Config, error) {
	name := app.TildeToHome(fileConfig)
	if _, err := os.Stat(name); err != nil {
		return &Config{}, nil
	}
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("couldn't read %s: %s - please remove it",
			name, err)
	}
	_, msg, err := network.Unmarshal(buf)
	if err != nil {
		return nil, fmt.Errorf("error while reading file %s: %s",
			name, err)
	}
	cfg, ok := msg.(*Config)
	if !ok {
		log.Fatal("Wrong data-structure in file", name)
	}
	cfg.name = name
	return cfg, nil
}

// write saves the config to the given file.
func (cfg *Config) write() {
	buf, err := network.Marshal(cfg)
	log.ErrFatal(err)
	log.ErrFatal(ioutil.WriteFile(cfg.name, buf, 0660))
}
