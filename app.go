package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/dedis/crypto.v0/abstract"
	"gopkg.in/dedis/onet.v1/app"
)

type Config struct {
	// config-file name
	Name string
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
	// Name and purpose of the party.
	Name string
	// DateTime of the party. It is in the following format, following UTC:
	//   YYYY-MM-DD HH:mm
	DateTime string
	// Location of the party
	Location string
}

type FinalStatement struct {
	// Desc is the description of the pop-party.
	Desc *PopDesc
	// Attendees holds a slice of all public keys of the attendees.
	Attendees []abstract.Point
	// Signature is created by all conodes responsible for that pop-party
	Signature []byte
	// Flag indicates, that party was merged
	Merged bool
}

type popDescToml struct {
	Name     string
	DateTime string
	Location string
}

type finalStatementToml struct {
	Desc      *popDescToml
	Attendees []string
	Signature string
	Merged    bool
}

type PopDescGroupToml struct {
	Name     string
	DateTime string
	Location string
	Servers  []*app.ServerToml `toml:"servers"`
}

var key = `{"address":"286485b3026d5d817f1f444060516b439b13dd2b","crypto":{"cipher":"aes-128-ctr","ciphertext":"d1a54d49808b658d9ea5a2c795c6a26741483699bf258d43a1d102dbfded867a","cipherparams":{"iv":"779603e70f888ee1496cbe19a7575cef"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"3fcc04ce5dbbfcdcdadeff6d5a69b05bb1931e435459105883ac64de5aefe271"},"mac":"d170e43d5e67e747bf766a112f48f649f574e6938ceb8c361dd73f7e2a586c16"},"id":"cad9bc2c-89c2-401f-bc5e-de4b4a11161d","version":3}`

var nonce int64 = 198

func main() {

	s := make([]common.Address, 1)
	s[0] = common.HexToAddress("0x286485b3026d5d817f1f444060516b439b13dd2b")
	fmt.Println(s)

	var keysetA [][32]byte
	//keysetA.push("AAAAC3NzaC1lZDI1NTE5AAAAIDD6LNlQsbLyr0Mp/6mJUCAILBbNGrJefiaVp+G6H97P")
	//keyset[0][32]("AAAAC3NzaC1lZDI1NTE5AAAAIDD6LNlQsbLyr0Mp/6mJUCAILBbNGrJefiaVp+G6H97P")
	//attendeeKey := "AAAAC3NzaC1lZDI1NTE5AAAAIDD6LNlQsbLyr0Mp/6mJUCAILBbNGrJefiaVp+G6H97P"
	//copy(keyset[0], attendeeKey)

	//	copy(keysetA[0][:], "AAAAC3NzaC1lZDI1NTE5AAAAIDD6LNlQsbLyr0Mp/6mJUCAILBbNGrJefiaVp+G6H97P")

	endPoint := "/home/hugo/.ethereum/rinkeby/geth.ipc"

	auth, err := bind.NewTransactor(strings.NewReader(key), "testpassword")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	//Deploys contract
	contract := orgConfig(auth, endPoint, Config{"First Party", "Nazareth", 1, s, 60})

	//Link to deployed contract
	//contract := orgLink(endPoint, testAddress)

	time.Sleep(30 * time.Second)

	st, _ := contract.CurrentState(nil)
	fmt.Println(returnState(uint(st)))
	fmt.Println(contract.CurrentState(nil))

	//organisator signs contract
	sign(auth, contract)

	time.Sleep(30 * time.Second)

	st, _ = contract.CurrentState(nil)
	fmt.Println(returnState(uint(st)))
	fmt.Println(contract.CurrentState(nil))

	//Administrator signs whole configuration
	signAdmin(auth, contract)

	time.Sleep(30 * time.Second)

	st, _ = contract.CurrentState(nil)
	fmt.Println(returnState(uint(st)))
	fmt.Println(contract.CurrentState(nil))

	//organizator deposit key set
	orgPublic(auth, contract, keysetA)

	time.Sleep(45 * time.Second)

	st, _ = contract.CurrentState(nil)
	fmt.Println(returnState(uint(st)))
	fmt.Println(contract.CurrentState(nil))

	fmt.Println(contract.AllSets(nil, big.NewInt(0)))

	//Administrator calls consensus function
	orgFinal(auth, contract)

	time.Sleep(30 * time.Second)

	st, _ = contract.CurrentState(nil)
	fmt.Println(returnState(uint(st)))
	fmt.Println(contract.CurrentState(nil))

}

func returnState(state uint) string {
	switch result := state; result {
	case 0:
		return "Initial State"
	case 1:
		return "configurationSet"
	case 2:
		return "configurationSigned"
	case 3:
		return "keyDeposited"
	case 4:
		return "locked"
	default:
		return "unknow"
	}
	return "error"
}

//Deploy & configure a new pop-party
func orgConfig(auth *bind.TransactOpts, network string, config Config) *Popcontract {

	conn, err := ethclient.Dial(network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}

	addr, _, contract, err := DeployPopcontract(&bind.TransactOpts{
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

	fmt.Printf("Contract deployed. Address : %x \n", addr)

	txe, err := contract.SetConfiguration(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(3000000),
		GasPrice: big.NewInt(90000000000),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(nonce),
	}, config.Name, config.Location, big.NewInt(config.NumberOfOrganizers), config.OrganizersAddresses, big.NewInt(config.Deadline))

	if err != nil {
		log.Fatalf("could not set configuration: %v", err)
	}
	nonce++

	fmt.Printf("Configuration set. Transaction : %v \n", txe)

	return contract

}

//connect to contract
func orgLink(network string, address common.Address) *Popcontract {

	conn, err := ethclient.Dial(network)
	if err != nil {
		log.Fatalf("could not connect to network: %v \n", err)
	}

	contract, err := NewPopcontract(address, conn)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v \n", err)
	}
	fmt.Println("Connected to contract.. \n")

	return contract
}

//Organizators sign configuration
func sign(auth *bind.TransactOpts, contract *Popcontract) {

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
}

//Administrator sign whole configuration
func signAdmin(auth *bind.TransactOpts, contract *Popcontract) {

	txe, err := contract.SignWholeConfiguration(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(3000000),
		GasPrice: big.NewInt(90000000000),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(nonce),
	})
	if err != nil {
		log.Fatalf("could not sign whole contract: %v", err)
	}
	nonce++
	fmt.Printf("Whole configuration signed. Transaction : %v \n", txe)
}

//add a new keyset
func orgPublic(auth *bind.TransactOpts, contract *Popcontract, keyset [][32]byte) error {

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

	return nil

}

//reach consensus
func orgFinal(auth *bind.TransactOpts, contract *Popcontract) error {

	txe, err := contract.PublicKeyConsensus(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(3000000),
		GasPrice: big.NewInt(90000000000),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(nonce),
	})
	if err != nil {
		log.Fatalf("could not reach consensus: %v", err)
	}
	nonce++
	fmt.Printf("Consensus reached. Transaction : %v \n", txe)

	return nil
}
