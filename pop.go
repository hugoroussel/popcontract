// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MortalABI is the input ABI used to generate the binding from.
const MortalABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// MortalBin is the compiled bytecode used for deploying new contracts.
const MortalBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b60bb8061003b6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341c0e1b58114603c575b600080fd5b3415604657600080fd5b604c604e565b005b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161415608c5760005473ffffffffffffffffffffffffffffffffffffffff16ff5b5b5600a165627a7a72305820fca37e3972c2e7b8a32d77bed5e3eb085c965d674371e7fb4cfc8f60ac5179d50029`

// DeployMortal deploys a new Ethereum contract, binding an instance of Mortal to it.
func DeployMortal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mortal, error) {
	parsed, err := abi.JSON(strings.NewReader(MortalABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MortalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mortal{MortalCaller: MortalCaller{contract: contract}, MortalTransactor: MortalTransactor{contract: contract}}, nil
}

// Mortal is an auto generated Go binding around an Ethereum contract.
type Mortal struct {
	MortalCaller     // Read-only binding to the contract
	MortalTransactor // Write-only binding to the contract
}

// MortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type MortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MortalSession struct {
	Contract     *Mortal           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MortalCallerSession struct {
	Contract *MortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MortalTransactorSession struct {
	Contract     *MortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type MortalRaw struct {
	Contract *Mortal // Generic contract binding to access the raw methods on
}

// MortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MortalCallerRaw struct {
	Contract *MortalCaller // Generic read-only contract binding to access the raw methods on
}

// MortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MortalTransactorRaw struct {
	Contract *MortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMortal creates a new instance of Mortal, bound to a specific deployed contract.
func NewMortal(address common.Address, backend bind.ContractBackend) (*Mortal, error) {
	contract, err := bindMortal(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mortal{MortalCaller: MortalCaller{contract: contract}, MortalTransactor: MortalTransactor{contract: contract}}, nil
}

// NewMortalCaller creates a new read-only instance of Mortal, bound to a specific deployed contract.
func NewMortalCaller(address common.Address, caller bind.ContractCaller) (*MortalCaller, error) {
	contract, err := bindMortal(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MortalCaller{contract: contract}, nil
}

// NewMortalTransactor creates a new write-only instance of Mortal, bound to a specific deployed contract.
func NewMortalTransactor(address common.Address, transactor bind.ContractTransactor) (*MortalTransactor, error) {
	contract, err := bindMortal(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MortalTransactor{contract: contract}, nil
}

// bindMortal binds a generic wrapper to an already deployed contract.
func bindMortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MortalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mortal *MortalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mortal.Contract.MortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mortal *MortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.Contract.MortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mortal *MortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mortal.Contract.MortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mortal *MortalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mortal *MortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mortal *MortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mortal.Contract.contract.Transact(opts, method, params...)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalSession) Kill() (*types.Transaction, error) {
	return _Mortal.Contract.Kill(&_Mortal.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalTransactorSession) Kill() (*types.Transaction, error) {
	return _Mortal.Contract.Kill(&_Mortal.TransactOpts)
}

// PopcontractABI is the input ABI used to generate the binding from.
const PopcontractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"place\",\"type\":\"string\"},{\"name\":\"organizers\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"address[]\"},{\"name\":\"durationInMinutes\",\"type\":\"uint256\"}],\"name\":\"setConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"finalKeySet\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signWholeConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrganizersAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locationOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isOrganizer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignedConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"configSignOrganizers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"publicKeyConsensus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signedConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"organizersAdresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allSets\",\"outputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_publicKeySet\",\"type\":\"bytes32[]\"}],\"name\":\"depositPublicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// PopcontractBin is the compiled bytecode used for deploying new contracts.
const PopcontractBin = `0x6060604052600a8054600160a060020a0319169055341561001f57600080fd5b5b5b60008054600160a060020a03191633600160a060020a03161790555b600380546000919060ff19166001835b02179055505b5b6112f5806100636000396000f300606060405236156101035763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166270041481146101085780630c3f6acf146101b85780631ae95ec4146101ef578063210790c81461021757806323266dfe1461023e578063232a6b9d146102a557806330756ec3146102cc57806341676f151461035757806341c0e1b51461038a5780635a52ecf61461039f57806372ee91c2146104065780638e7d0f44146104355780639f638dbf1461045c578063b9bcc99f1461048e578063c1b8471b146104c0578063c323f884146104f2578063c47f002714610555578063d7e3a4d3146105ba578063eff64cb3146105df575b600080fd5b341561011357600080fd5b6101a460046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650509335935061066a92505050565b604051901515815260200160405180910390f35b34156101c357600080fd5b6101cb610734565b604051808260048111156101db57fe5b60ff16815260200191505060405180910390f35b34156101fa57600080fd5b61020560043561073d565b60405190815260200160405180910390f35b341561022257600080fd5b6101a4610760565b604051901515815260200160405180910390f35b341561024957600080fd5b610251610881565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156102915780820151818401525b602001610278565b505050509050019250505060405180910390f35b34156102b057600080fd5b6101a46108ea565b604051901515815260200160405180910390f35b34156102d757600080fd5b6102df61090b565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561031c5780820151818401525b602001610303565b50505050905090810190601f1680156103495780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561036257600080fd5b6101a4600160a060020a03600435166109a9565b604051901515815260200160405180910390f35b341561039557600080fd5b61039d610a57565b005b34156103aa57600080fd5b610251610a7f565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156102915780820151818401525b602001610278565b505050509050019250505060405180910390f35b341561041157600080fd5b610419610ae8565b604051600160a060020a03909116815260200160405180910390f35b341561044057600080fd5b6101a4610bff565b604051901515815260200160405180910390f35b341561046757600080fd5b610419600435610cf3565b604051600160a060020a03909116815260200160405180910390f35b341561049957600080fd5b610419600435610d25565b604051600160a060020a03909116815260200160405180910390f35b34156104cb57600080fd5b610419600435610d57565b604051600160a060020a03909116815260200160405180910390f35b34156104fd57600080fd5b6101a46004602481358181019083013580602081810201604051908101604052809392919081815260200183836020028082843750949650610d8695505050505050565b604051901515815260200160405180910390f35b341561056057600080fd5b6101a460046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610ea495505050505050565b604051901515815260200160405180910390f35b34156105c557600080fd5b610205610f16565b60405190815260200160405180910390f35b34156105ea57600080fd5b6102df610f1c565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561031c5780820151818401525b602001610303565b50505050905090810190601f1680156103495780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600354600090819060ff16600481111561068057fe5b81600481111561068c57fe5b14156101035760005433600160a060020a03908116911614156101035742603c84020160065560058680516106c5929160200190610fba565b506007859055846106d7600882611039565b5060088480516106eb929160200190611063565b50600380546001919060ff191682805b02179055506000198501610710600982611039565b506001915061071f565b600080fd5b5b61072a565b600080fd5b5b50949350505050565b60035460ff1681565b600280548290811061074b57fe5b906000526020600020900160005b5054905081565b6003546000908190819060019060ff16600481111561077b57fe5b81600481111561078757fe5b1415610103576000805490935033600160a060020a039081169116141561086e57600091505b60095482101561084e57600a5460098054600160a060020a0390921691849081106107d457fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a0316141561080857600080fd5b600a805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055600192505b5b6001909101906107ad565b821561086e57600380546002919060ff19166001835b0217905550600193505b5b5b61087a565b600080fd5b5b50505090565b6108896110d8565b60088054806020026020016040519081016040528092919081815260200182805480156108df57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116108c1575b505050505090505b90565b600a5474010000000000000000000000000000000000000000900460ff1681565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109a15780601f10610976576101008083540402835291602001916109a1565b820191906000526020600020905b81548152906001019060200180831161098457829003601f168201915b505050505081565b600354600090819060029060ff1660048111156109c257fe5b8160048111156109ce57fe5b141561010357600091505b600754821015610a405783600160a060020a03166008838154811015156109fc57fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a03161415610a345760019250610a45565b5b6001909101906109d9565b600092505b610a4f565b600080fd5b5b5050919050565b60005433600160a060020a0390811691161415610a7c57600054600160a060020a0316ff5b5b565b610a876110d8565b60098054806020026020016040519081016040528092919081815260200182805480156108df57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116108c1575b505050505090505b90565b600354600090819060019060ff166004811115610b0157fe5b816004811115610b0d57fe5b141561010357600091505b600754821015610be0576008805483908110610b3057fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a031633600160a060020a03161415610bd4576009805460018101610b7b8382611039565b916000526020600020900160005b6008805486908110610b9757fe5b906000526020600020900160005b9054835461010093840a600160a060020a039390940a9091048216830292909102191617905550339250610bef565b5b600190910190610b18565b600a54600160a060020a031692505b610bf9565b600080fd5b5b505090565b600380546000919060ff166004811115610c1557fe5b816004811115610c2157fe5b1415610103576006544211610ce25760015415801590610c4f575060005433600160a060020a039081169116145b15610cdd57600a805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055600180546000198101908110610c9b57fe5b906000526020600020906002020160005b506001018054610cbe91600291611114565b50600380546004919060ff19166001835b021790555060019150610ce2565b600091505b5b5b610cee565b600080fd5b5b5090565b6009805482908110610d0157fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b6008805482908110610d0157fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b6001805482908110610d6557fe5b906000526020600020906002020160005b5054600160a060020a0316905081565b60035460009060029060ff166004811115610d9d57fe5b816004811115610da957fe5b1415610103576006544211610e9157610dc1336109a9565b15610e8c5760018054808201610dd78382611165565b916000526020600020906002020160005b60408051908101604052600160a060020a0333168152602081018790529190508151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391909116178155602082015181600101908051610e4c929160200190611197565b5060039250610e59915050565b60035460ff166004811115610e6a57fe5b14610e835760038054819060ff19166001825b02179055505b60019150610e91565b600091505b5b5b610e9d565b600080fd5b5b50919050565b600354600090819060ff166004811115610eba57fe5b816004811115610ec657fe5b14156101035760005433600160a060020a0390811691161415610e8c576004838051610ef6929160200190610fba565b5060019150610e91565b600091505b610e9d565b600080fd5b5b50919050565b60065481565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109a15780601f10610976576101008083540402835291602001916109a1565b820191906000526020600020905b81548152906001019060200180831161098457829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610ffb57805160ff1916838001178555611028565b82800160010185558215611028579182015b8281111561102857825182559160200191906001019061100d565b5b50610cee9291506111e5565b5090565b81548183558181151161105d5760008381526020902061105d9181019083016111e5565b5b505050565b8280548282559060005260206000209081019282156110c7579160200282015b828111156110c7578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617825560209290920191600190910190611083565b5b50610cee929150611206565b5090565b60206040519081016040526000815290565b81548183558181151161105d5760008381526020902061105d9181019083016111e5565b5b505050565b8280548282559060005260206000209081019282156110285760005260206000209182015b82811115611028578254825591600101919060010190611139565b5b50610cee9291506111e5565b5090565b81548183558181151161105d5760020281600202836000526020600020918201910161105d919061125f565b5b505050565b828054828255906000526020600020908101928215611028579160200282015b8281111561102857825182556020909201916001909101906111b7565b5b50610cee9291506111e5565b5090565b6108e791905b80821115610cee57600081556001016111eb565b5090565b90565b6108e791905b80821115610cee57805473ffffffffffffffffffffffffffffffffffffffff1916815560010161120c565b5090565b90565b6108e791905b80821115610cee57600081556001016111eb565b5090565b90565b6108e791905b80821115610cee57805473ffffffffffffffffffffffffffffffffffffffff19168155600061129760018301826112a7565b50600201611265565b5090565b90565b50805460008255906000526020600020908101906112c591906111e5565b5b505600a165627a7a723058207eaba8c9151398462fdb78fb220f0d99a0c09906e1a50aa5d153b5b500cd17330029`

// DeployPopcontract deploys a new Ethereum contract, binding an instance of Popcontract to it.
func DeployPopcontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Popcontract, error) {
	parsed, err := abi.JSON(strings.NewReader(PopcontractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PopcontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Popcontract{PopcontractCaller: PopcontractCaller{contract: contract}, PopcontractTransactor: PopcontractTransactor{contract: contract}}, nil
}

// Popcontract is an auto generated Go binding around an Ethereum contract.
type Popcontract struct {
	PopcontractCaller     // Read-only binding to the contract
	PopcontractTransactor // Write-only binding to the contract
}

// PopcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PopcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PopcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PopcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PopcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PopcontractSession struct {
	Contract     *Popcontract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PopcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PopcontractCallerSession struct {
	Contract *PopcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PopcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PopcontractTransactorSession struct {
	Contract     *PopcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PopcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PopcontractRaw struct {
	Contract *Popcontract // Generic contract binding to access the raw methods on
}

// PopcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PopcontractCallerRaw struct {
	Contract *PopcontractCaller // Generic read-only contract binding to access the raw methods on
}

// PopcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PopcontractTransactorRaw struct {
	Contract *PopcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPopcontract creates a new instance of Popcontract, bound to a specific deployed contract.
func NewPopcontract(address common.Address, backend bind.ContractBackend) (*Popcontract, error) {
	contract, err := bindPopcontract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Popcontract{PopcontractCaller: PopcontractCaller{contract: contract}, PopcontractTransactor: PopcontractTransactor{contract: contract}}, nil
}

// NewPopcontractCaller creates a new read-only instance of Popcontract, bound to a specific deployed contract.
func NewPopcontractCaller(address common.Address, caller bind.ContractCaller) (*PopcontractCaller, error) {
	contract, err := bindPopcontract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PopcontractCaller{contract: contract}, nil
}

// NewPopcontractTransactor creates a new write-only instance of Popcontract, bound to a specific deployed contract.
func NewPopcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*PopcontractTransactor, error) {
	contract, err := bindPopcontract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PopcontractTransactor{contract: contract}, nil
}

// bindPopcontract binds a generic wrapper to an already deployed contract.
func bindPopcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PopcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Popcontract *PopcontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Popcontract.Contract.PopcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Popcontract *PopcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.Contract.PopcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Popcontract *PopcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Popcontract.Contract.PopcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Popcontract *PopcontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Popcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Popcontract *PopcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Popcontract *PopcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Popcontract.Contract.contract.Transact(opts, method, params...)
}

// AllSets is a free data retrieval call binding the contract method 0xc1b8471b.
//
// Solidity: function allSets( uint256) constant returns(sender address)
func (_Popcontract *PopcontractCaller) AllSets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "allSets", arg0)
	return *ret0, err
}

// AllSets is a free data retrieval call binding the contract method 0xc1b8471b.
//
// Solidity: function allSets( uint256) constant returns(sender address)
func (_Popcontract *PopcontractSession) AllSets(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.AllSets(&_Popcontract.CallOpts, arg0)
}

// AllSets is a free data retrieval call binding the contract method 0xc1b8471b.
//
// Solidity: function allSets( uint256) constant returns(sender address)
func (_Popcontract *PopcontractCallerSession) AllSets(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.AllSets(&_Popcontract.CallOpts, arg0)
}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() constant returns(uint8)
func (_Popcontract *PopcontractCaller) CurrentState(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "currentState")
	return *ret0, err
}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() constant returns(uint8)
func (_Popcontract *PopcontractSession) CurrentState() (uint8, error) {
	return _Popcontract.Contract.CurrentState(&_Popcontract.CallOpts)
}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() constant returns(uint8)
func (_Popcontract *PopcontractCallerSession) CurrentState() (uint8, error) {
	return _Popcontract.Contract.CurrentState(&_Popcontract.CallOpts)
}

// EndOfParty is a free data retrieval call binding the contract method 0xd7e3a4d3.
//
// Solidity: function endOfParty() constant returns(uint256)
func (_Popcontract *PopcontractCaller) EndOfParty(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "endOfParty")
	return *ret0, err
}

// EndOfParty is a free data retrieval call binding the contract method 0xd7e3a4d3.
//
// Solidity: function endOfParty() constant returns(uint256)
func (_Popcontract *PopcontractSession) EndOfParty() (*big.Int, error) {
	return _Popcontract.Contract.EndOfParty(&_Popcontract.CallOpts)
}

// EndOfParty is a free data retrieval call binding the contract method 0xd7e3a4d3.
//
// Solidity: function endOfParty() constant returns(uint256)
func (_Popcontract *PopcontractCallerSession) EndOfParty() (*big.Int, error) {
	return _Popcontract.Contract.EndOfParty(&_Popcontract.CallOpts)
}

// FinalKeySet is a free data retrieval call binding the contract method 0x1ae95ec4.
//
// Solidity: function finalKeySet( uint256) constant returns(bytes32)
func (_Popcontract *PopcontractCaller) FinalKeySet(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "finalKeySet", arg0)
	return *ret0, err
}

// FinalKeySet is a free data retrieval call binding the contract method 0x1ae95ec4.
//
// Solidity: function finalKeySet( uint256) constant returns(bytes32)
func (_Popcontract *PopcontractSession) FinalKeySet(arg0 *big.Int) ([32]byte, error) {
	return _Popcontract.Contract.FinalKeySet(&_Popcontract.CallOpts, arg0)
}

// FinalKeySet is a free data retrieval call binding the contract method 0x1ae95ec4.
//
// Solidity: function finalKeySet( uint256) constant returns(bytes32)
func (_Popcontract *PopcontractCallerSession) FinalKeySet(arg0 *big.Int) ([32]byte, error) {
	return _Popcontract.Contract.FinalKeySet(&_Popcontract.CallOpts, arg0)
}

// GetOrganizersAddresses is a free data retrieval call binding the contract method 0x23266dfe.
//
// Solidity: function getOrganizersAddresses() constant returns(address[])
func (_Popcontract *PopcontractCaller) GetOrganizersAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "getOrganizersAddresses")
	return *ret0, err
}

// GetOrganizersAddresses is a free data retrieval call binding the contract method 0x23266dfe.
//
// Solidity: function getOrganizersAddresses() constant returns(address[])
func (_Popcontract *PopcontractSession) GetOrganizersAddresses() ([]common.Address, error) {
	return _Popcontract.Contract.GetOrganizersAddresses(&_Popcontract.CallOpts)
}

// GetOrganizersAddresses is a free data retrieval call binding the contract method 0x23266dfe.
//
// Solidity: function getOrganizersAddresses() constant returns(address[])
func (_Popcontract *PopcontractCallerSession) GetOrganizersAddresses() ([]common.Address, error) {
	return _Popcontract.Contract.GetOrganizersAddresses(&_Popcontract.CallOpts)
}

// GetSignedConfiguration is a free data retrieval call binding the contract method 0x5a52ecf6.
//
// Solidity: function getSignedConfiguration() constant returns(address[])
func (_Popcontract *PopcontractCaller) GetSignedConfiguration(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "getSignedConfiguration")
	return *ret0, err
}

// GetSignedConfiguration is a free data retrieval call binding the contract method 0x5a52ecf6.
//
// Solidity: function getSignedConfiguration() constant returns(address[])
func (_Popcontract *PopcontractSession) GetSignedConfiguration() ([]common.Address, error) {
	return _Popcontract.Contract.GetSignedConfiguration(&_Popcontract.CallOpts)
}

// GetSignedConfiguration is a free data retrieval call binding the contract method 0x5a52ecf6.
//
// Solidity: function getSignedConfiguration() constant returns(address[])
func (_Popcontract *PopcontractCallerSession) GetSignedConfiguration() ([]common.Address, error) {
	return _Popcontract.Contract.GetSignedConfiguration(&_Popcontract.CallOpts)
}

// LocationOfParty is a free data retrieval call binding the contract method 0x30756ec3.
//
// Solidity: function locationOfParty() constant returns(string)
func (_Popcontract *PopcontractCaller) LocationOfParty(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "locationOfParty")
	return *ret0, err
}

// LocationOfParty is a free data retrieval call binding the contract method 0x30756ec3.
//
// Solidity: function locationOfParty() constant returns(string)
func (_Popcontract *PopcontractSession) LocationOfParty() (string, error) {
	return _Popcontract.Contract.LocationOfParty(&_Popcontract.CallOpts)
}

// LocationOfParty is a free data retrieval call binding the contract method 0x30756ec3.
//
// Solidity: function locationOfParty() constant returns(string)
func (_Popcontract *PopcontractCallerSession) LocationOfParty() (string, error) {
	return _Popcontract.Contract.LocationOfParty(&_Popcontract.CallOpts)
}

// NameOfParty is a free data retrieval call binding the contract method 0xeff64cb3.
//
// Solidity: function nameOfParty() constant returns(string)
func (_Popcontract *PopcontractCaller) NameOfParty(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "nameOfParty")
	return *ret0, err
}

// NameOfParty is a free data retrieval call binding the contract method 0xeff64cb3.
//
// Solidity: function nameOfParty() constant returns(string)
func (_Popcontract *PopcontractSession) NameOfParty() (string, error) {
	return _Popcontract.Contract.NameOfParty(&_Popcontract.CallOpts)
}

// NameOfParty is a free data retrieval call binding the contract method 0xeff64cb3.
//
// Solidity: function nameOfParty() constant returns(string)
func (_Popcontract *PopcontractCallerSession) NameOfParty() (string, error) {
	return _Popcontract.Contract.NameOfParty(&_Popcontract.CallOpts)
}

// OrganizersAdresses is a free data retrieval call binding the contract method 0xb9bcc99f.
//
// Solidity: function organizersAdresses( uint256) constant returns(address)
func (_Popcontract *PopcontractCaller) OrganizersAdresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "organizersAdresses", arg0)
	return *ret0, err
}

// OrganizersAdresses is a free data retrieval call binding the contract method 0xb9bcc99f.
//
// Solidity: function organizersAdresses( uint256) constant returns(address)
func (_Popcontract *PopcontractSession) OrganizersAdresses(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.OrganizersAdresses(&_Popcontract.CallOpts, arg0)
}

// OrganizersAdresses is a free data retrieval call binding the contract method 0xb9bcc99f.
//
// Solidity: function organizersAdresses( uint256) constant returns(address)
func (_Popcontract *PopcontractCallerSession) OrganizersAdresses(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.OrganizersAdresses(&_Popcontract.CallOpts, arg0)
}

// Signed is a free data retrieval call binding the contract method 0x232a6b9d.
//
// Solidity: function signed() constant returns(bool)
func (_Popcontract *PopcontractCaller) Signed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "signed")
	return *ret0, err
}

// Signed is a free data retrieval call binding the contract method 0x232a6b9d.
//
// Solidity: function signed() constant returns(bool)
func (_Popcontract *PopcontractSession) Signed() (bool, error) {
	return _Popcontract.Contract.Signed(&_Popcontract.CallOpts)
}

// Signed is a free data retrieval call binding the contract method 0x232a6b9d.
//
// Solidity: function signed() constant returns(bool)
func (_Popcontract *PopcontractCallerSession) Signed() (bool, error) {
	return _Popcontract.Contract.Signed(&_Popcontract.CallOpts)
}

// SignedConfiguration is a free data retrieval call binding the contract method 0x9f638dbf.
//
// Solidity: function signedConfiguration( uint256) constant returns(address)
func (_Popcontract *PopcontractCaller) SignedConfiguration(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "signedConfiguration", arg0)
	return *ret0, err
}

// SignedConfiguration is a free data retrieval call binding the contract method 0x9f638dbf.
//
// Solidity: function signedConfiguration( uint256) constant returns(address)
func (_Popcontract *PopcontractSession) SignedConfiguration(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.SignedConfiguration(&_Popcontract.CallOpts, arg0)
}

// SignedConfiguration is a free data retrieval call binding the contract method 0x9f638dbf.
//
// Solidity: function signedConfiguration( uint256) constant returns(address)
func (_Popcontract *PopcontractCallerSession) SignedConfiguration(arg0 *big.Int) (common.Address, error) {
	return _Popcontract.Contract.SignedConfiguration(&_Popcontract.CallOpts, arg0)
}

// ConfigSignOrganizers is a paid mutator transaction binding the contract method 0x72ee91c2.
//
// Solidity: function configSignOrganizers() returns(address)
func (_Popcontract *PopcontractTransactor) ConfigSignOrganizers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "configSignOrganizers")
}

// ConfigSignOrganizers is a paid mutator transaction binding the contract method 0x72ee91c2.
//
// Solidity: function configSignOrganizers() returns(address)
func (_Popcontract *PopcontractSession) ConfigSignOrganizers() (*types.Transaction, error) {
	return _Popcontract.Contract.ConfigSignOrganizers(&_Popcontract.TransactOpts)
}

// ConfigSignOrganizers is a paid mutator transaction binding the contract method 0x72ee91c2.
//
// Solidity: function configSignOrganizers() returns(address)
func (_Popcontract *PopcontractTransactorSession) ConfigSignOrganizers() (*types.Transaction, error) {
	return _Popcontract.Contract.ConfigSignOrganizers(&_Popcontract.TransactOpts)
}

// DepositPublicKeys is a paid mutator transaction binding the contract method 0xc323f884.
//
// Solidity: function depositPublicKeys(_publicKeySet bytes32[]) returns(bool)
func (_Popcontract *PopcontractTransactor) DepositPublicKeys(opts *bind.TransactOpts, _publicKeySet [][32]byte) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "depositPublicKeys", _publicKeySet)
}

// DepositPublicKeys is a paid mutator transaction binding the contract method 0xc323f884.
//
// Solidity: function depositPublicKeys(_publicKeySet bytes32[]) returns(bool)
func (_Popcontract *PopcontractSession) DepositPublicKeys(_publicKeySet [][32]byte) (*types.Transaction, error) {
	return _Popcontract.Contract.DepositPublicKeys(&_Popcontract.TransactOpts, _publicKeySet)
}

// DepositPublicKeys is a paid mutator transaction binding the contract method 0xc323f884.
//
// Solidity: function depositPublicKeys(_publicKeySet bytes32[]) returns(bool)
func (_Popcontract *PopcontractTransactorSession) DepositPublicKeys(_publicKeySet [][32]byte) (*types.Transaction, error) {
	return _Popcontract.Contract.DepositPublicKeys(&_Popcontract.TransactOpts, _publicKeySet)
}

// IsOrganizer is a paid mutator transaction binding the contract method 0x41676f15.
//
// Solidity: function isOrganizer(sender address) returns(bool)
func (_Popcontract *PopcontractTransactor) IsOrganizer(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "isOrganizer", sender)
}

// IsOrganizer is a paid mutator transaction binding the contract method 0x41676f15.
//
// Solidity: function isOrganizer(sender address) returns(bool)
func (_Popcontract *PopcontractSession) IsOrganizer(sender common.Address) (*types.Transaction, error) {
	return _Popcontract.Contract.IsOrganizer(&_Popcontract.TransactOpts, sender)
}

// IsOrganizer is a paid mutator transaction binding the contract method 0x41676f15.
//
// Solidity: function isOrganizer(sender address) returns(bool)
func (_Popcontract *PopcontractTransactorSession) IsOrganizer(sender common.Address) (*types.Transaction, error) {
	return _Popcontract.Contract.IsOrganizer(&_Popcontract.TransactOpts, sender)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Popcontract *PopcontractTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Popcontract *PopcontractSession) Kill() (*types.Transaction, error) {
	return _Popcontract.Contract.Kill(&_Popcontract.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Popcontract *PopcontractTransactorSession) Kill() (*types.Transaction, error) {
	return _Popcontract.Contract.Kill(&_Popcontract.TransactOpts)
}

// PublicKeyConsensus is a paid mutator transaction binding the contract method 0x8e7d0f44.
//
// Solidity: function publicKeyConsensus() returns(bool)
func (_Popcontract *PopcontractTransactor) PublicKeyConsensus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "publicKeyConsensus")
}

// PublicKeyConsensus is a paid mutator transaction binding the contract method 0x8e7d0f44.
//
// Solidity: function publicKeyConsensus() returns(bool)
func (_Popcontract *PopcontractSession) PublicKeyConsensus() (*types.Transaction, error) {
	return _Popcontract.Contract.PublicKeyConsensus(&_Popcontract.TransactOpts)
}

// PublicKeyConsensus is a paid mutator transaction binding the contract method 0x8e7d0f44.
//
// Solidity: function publicKeyConsensus() returns(bool)
func (_Popcontract *PopcontractTransactorSession) PublicKeyConsensus() (*types.Transaction, error) {
	return _Popcontract.Contract.PublicKeyConsensus(&_Popcontract.TransactOpts)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0x00700414.
//
// Solidity: function setConfiguration(place string, organizers uint256, data address[], durationInMinutes uint256) returns(bool)
func (_Popcontract *PopcontractTransactor) SetConfiguration(opts *bind.TransactOpts, place string, organizers *big.Int, data []common.Address, durationInMinutes *big.Int) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "setConfiguration", place, organizers, data, durationInMinutes)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0x00700414.
//
// Solidity: function setConfiguration(place string, organizers uint256, data address[], durationInMinutes uint256) returns(bool)
func (_Popcontract *PopcontractSession) SetConfiguration(place string, organizers *big.Int, data []common.Address, durationInMinutes *big.Int) (*types.Transaction, error) {
	return _Popcontract.Contract.SetConfiguration(&_Popcontract.TransactOpts, place, organizers, data, durationInMinutes)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0x00700414.
//
// Solidity: function setConfiguration(place string, organizers uint256, data address[], durationInMinutes uint256) returns(bool)
func (_Popcontract *PopcontractTransactorSession) SetConfiguration(place string, organizers *big.Int, data []common.Address, durationInMinutes *big.Int) (*types.Transaction, error) {
	return _Popcontract.Contract.SetConfiguration(&_Popcontract.TransactOpts, place, organizers, data, durationInMinutes)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns(bool)
func (_Popcontract *PopcontractTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns(bool)
func (_Popcontract *PopcontractSession) SetName(_name string) (*types.Transaction, error) {
	return _Popcontract.Contract.SetName(&_Popcontract.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns(bool)
func (_Popcontract *PopcontractTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _Popcontract.Contract.SetName(&_Popcontract.TransactOpts, _name)
}

// SignWholeConfiguration is a paid mutator transaction binding the contract method 0x210790c8.
//
// Solidity: function signWholeConfiguration() returns(bool)
func (_Popcontract *PopcontractTransactor) SignWholeConfiguration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "signWholeConfiguration")
}

// SignWholeConfiguration is a paid mutator transaction binding the contract method 0x210790c8.
//
// Solidity: function signWholeConfiguration() returns(bool)
func (_Popcontract *PopcontractSession) SignWholeConfiguration() (*types.Transaction, error) {
	return _Popcontract.Contract.SignWholeConfiguration(&_Popcontract.TransactOpts)
}

// SignWholeConfiguration is a paid mutator transaction binding the contract method 0x210790c8.
//
// Solidity: function signWholeConfiguration() returns(bool)
func (_Popcontract *PopcontractTransactorSession) SignWholeConfiguration() (*types.Transaction, error) {
	return _Popcontract.Contract.SignWholeConfiguration(&_Popcontract.TransactOpts)
}
