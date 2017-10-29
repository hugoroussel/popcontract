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
const MortalBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b60bb8061003b6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341c0e1b58114603c575b600080fd5b3415604657600080fd5b604c604e565b005b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161415608c5760005473ffffffffffffffffffffffffffffffffffffffff16ff5b5b5600a165627a7a72305820f95c091db2b604e6f5ffc8b5d2f658bf0f65bcdf38185af392061d634199efae0029`

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
const PopcontractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"place\",\"type\":\"string\"},{\"name\":\"organizers\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"address[]\"},{\"name\":\"durationInMinutes\",\"type\":\"uint256\"}],\"name\":\"setConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signWholeConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrganizersAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locationOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isOrganizer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_publicKeySet\",\"type\":\"address[]\"}],\"name\":\"depositPublicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"configSignOrganizers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"publicKeyConsensus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allSets\",\"outputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finalKeySet\",\"outputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// PopcontractBin is the compiled bytecode used for deploying new contracts.
const PopcontractBin = `0x6060604052600b8054600160a060020a0319169055341561001f57600080fd5b5b5b60008054600160a060020a03191633600160a060020a03161790555b600480546000919060ff19166001835b02179055505b5b610f84806100636000396000f300606060405236156100d75763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166270041481146100dc5780630c3f6acf1461018c578063210790c8146101c357806323266dfe146101ea57806330756ec31461025157806341676f15146102dc57806341c0e1b51461030f57806365472d7c1461032457806372ee91c2146103875780638e7d0f44146103ae578063c1b8471b146103d5578063c1f1a99214610407578063c47f002714610436578063d7e3a4d31461049b578063eff64cb3146104c0575b600080fd5b34156100e757600080fd5b61017860046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650509335935061054b92505050565b604051901515815260200160405180910390f35b341561019757600080fd5b61019f610604565b604051808260048111156101af57fe5b60ff16815260200191505060405180910390f35b34156101ce57600080fd5b61017861060d565b604051901515815260200160405180910390f35b34156101f557600080fd5b6101fd6106f9565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561023d5780820151818401525b602001610224565b505050509050019250505060405180910390f35b341561025c57600080fd5b610264610762565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102a15780820151818401525b602001610288565b50505050905090810190601f1680156102ce5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102e757600080fd5b610178600160a060020a0360043516610800565b604051901515815260200160405180910390f35b341561031a57600080fd5b6103226108b0565b005b341561032f57600080fd5b61017860046024813581810190830135806020818102016040519081016040528093929190818152602001838360200280828437509496506108d895505050505050565b604051901515815260200160405180910390f35b341561039257600080fd5b6101786109cc565b604051901515815260200160405180910390f35b34156103b957600080fd5b610178610ab5565b604051901515815260200160405180910390f35b34156103e057600080fd5b6103eb600435610b71565b604051600160a060020a03909116815260200160405180910390f35b341561041257600080fd5b6103eb610ba0565b604051600160a060020a03909116815260200160405180910390f35b341561044157600080fd5b61017860046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610baf95505050505050565b604051901515815260200160405180910390f35b34156104a657600080fd5b6104ae610bfd565b60405190815260200160405180910390f35b34156104cb57600080fd5b610264610c03565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102a15780820151818401525b602001610288565b50505050905090810190601f1680156102ce5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60048054600091829160ff169081111561056157fe5b81600481111561056d57fe5b14156100d75760005433600160a060020a03908116911614156100d75742603c84020160075560068680516105a6929160200190610ca1565b506008859055846105b8600982610d20565b5060098480516105cc929160200190610d4a565b50600480546001919060ff191682805b0217905550600191506105ef565b600080fd5b5b6105fa565b600080fd5b5b50949350505050565b60045460ff1681565b600480546000918291829160019160ff169081111561062857fe5b81600481111561063457fe5b14156100d7576000805490935033600160a060020a03908116911614156106e657600091505b6008548210156106c657600b54600a8054600160a060020a03909216918490811061068157fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a031614156106b557600080fd5b600192505b5b60019091019061065a565b82156106e657600480546002919060ff19166001835b0217905550600193505b5b5b6106f2565b600080fd5b5b50505090565b610701610db2565b600980548060200260200160405190810160405280929190818152602001828054801561075757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610739575b505050505090505b90565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107f85780601f106107cd576101008083540402835291602001916107f8565b820191906000526020600020905b8154815290600101906020018083116107db57829003601f168201915b505050505081565b60048054600091829160029160ff9091169081111561081b57fe5b81600481111561082757fe5b14156100d757600091505b6008548210156108995783600160a060020a031660098381548110151561085557fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a0316141561088d576001925061089e565b5b600190910190610832565b600092505b6108a8565b600080fd5b5b5050919050565b60005433600160a060020a03908116911614156108d557600054600160a060020a0316ff5b5b565b6004805460009160029160ff16908111156108ef57fe5b8160048111156108fb57fe5b14156100d75760075442116109b95761091333610800565b156109b457600180548082016109298382610dc4565b916000526020600020906002020160005b60408051908101604052600160a060020a03331681526020810187905291905081518154600160a060020a031916600160a060020a0391909116178155602082015181600101908051610991929160200190610d4a565b5050600480546003935090915060ff19166001835b0217905550600191506109b9565b600091505b5b5b6109c5565b600080fd5b5b50919050565b60048054600091829160019160ff909116908111156109e757fe5b8160048111156109f357fe5b14156100d7575b600854821015610aa0576009805483908110610a1257fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a031633600160a060020a03161415610a945733600a83815481101515610a5b57fe5b906000526020600020900160005b6101000a815481600160a060020a030219169083600160a060020a0316021790555060019250610aa5565b5b6001909101906109fa565b600092505b610aaf565b600080fd5b5b505090565b6004805460009160039160ff1690811115610acc57fe5b816004811115610ad857fe5b14156100d7576007544210610b605760015415610b5b5760018054fe5b906000526020600020906002020160005b50805460028054600160a060020a031916600160a060020a03909216919091178155600182018054610b3a91600391610e5e565b505060048054909150819060ff19166001825b021790555060019150610b60565b600091505b5b5b610b6c565b600080fd5b5b5090565b6001805482908110610b7f57fe5b906000526020600020906002020160005b5054600160a060020a0316905081565b600254600160a060020a031681565b60048054600091829160ff1690811115610bc557fe5b816004811115610bd157fe5b14156100d7576005838051610bea929160200190610ca1565b505b6109c5565b600080fd5b5b50919050565b60075481565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107f85780601f106107cd576101008083540402835291602001916107f8565b820191906000526020600020905b8154815290600101906020018083116107db57829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610ce257805160ff1916838001178555610d0f565b82800160010185558215610d0f579182015b82811115610d0f578251825591602001919060010190610cf4565b5b50610b6c929150610eaf565b5090565b815481835581811511610d4457600083815260209020610d44918101908301610eaf565b5b505050565b828054828255906000526020600020908101928215610da1579160200282015b82811115610da15782518254600160a060020a031916600160a060020a039190911617825560209290920191600190910190610d6a565b5b50610b6c929150610ed0565b5090565b60206040519081016040526000815290565b815481835581811511610d4457600202816002028360005260206000209182019101610d449190610efb565b5b505050565b828054828255906000526020600020908101928215610da1579160200282015b82811115610da15782518254600160a060020a031916600160a060020a039190911617825560209290920191600190910190610d6a565b5b50610b6c929150610ed0565b5090565b828054828255906000526020600020908101928215610da15760005260206000209182015b82811115610da1578254825591600101919060010190610e83565b5b50610b6c929150610ed0565b5090565b61075f91905b80821115610b6c5760008155600101610eb5565b5090565b90565b61075f91905b80821115610b6c578054600160a060020a0319168155600101610ed6565b5090565b90565b61075f91905b80821115610b6c578054600160a060020a03191681556000610f266001830182610f36565b50600201610f01565b5090565b90565b5080546000825590600052602060002090810190610f549190610eaf565b5b505600a165627a7a72305820e033143e974bba44bbe2b73e0b4462e6fad3f53b5b31c5ff6a70de689111da5f0029`

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

// FinalKeySet is a free data retrieval call binding the contract method 0xc1f1a992.
//
// Solidity: function finalKeySet() constant returns(sender address)
func (_Popcontract *PopcontractCaller) FinalKeySet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Popcontract.contract.Call(opts, out, "finalKeySet")
	return *ret0, err
}

// FinalKeySet is a free data retrieval call binding the contract method 0xc1f1a992.
//
// Solidity: function finalKeySet() constant returns(sender address)
func (_Popcontract *PopcontractSession) FinalKeySet() (common.Address, error) {
	return _Popcontract.Contract.FinalKeySet(&_Popcontract.CallOpts)
}

// FinalKeySet is a free data retrieval call binding the contract method 0xc1f1a992.
//
// Solidity: function finalKeySet() constant returns(sender address)
func (_Popcontract *PopcontractCallerSession) FinalKeySet() (common.Address, error) {
	return _Popcontract.Contract.FinalKeySet(&_Popcontract.CallOpts)
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

// ConfigSignOrganizers is a paid mutator transaction binding the contract method 0x72ee91c2.
//
// Solidity: function configSignOrganizers() returns(bool)
func (_Popcontract *PopcontractTransactor) ConfigSignOrganizers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "configSignOrganizers")
}

// ConfigSignOrganizers is a paid mutator transaction binding the contract method 0x72ee91c2.
//
// Solidity: function configSignOrganizers() returns(bool)
func (_Popcontract *PopcontractSession) ConfigSignOrganizers() (*types.Transaction, error) {
	return _Popcontract.Contract.ConfigSignOrganizers(&_Popcontract.TransactOpts)
}

// ConfigSignOrganizers is a paid mutator transaction binding the contract method 0x72ee91c2.
//
// Solidity: function configSignOrganizers() returns(bool)
func (_Popcontract *PopcontractTransactorSession) ConfigSignOrganizers() (*types.Transaction, error) {
	return _Popcontract.Contract.ConfigSignOrganizers(&_Popcontract.TransactOpts)
}

// DepositPublicKeys is a paid mutator transaction binding the contract method 0x65472d7c.
//
// Solidity: function depositPublicKeys(_publicKeySet address[]) returns(bool)
func (_Popcontract *PopcontractTransactor) DepositPublicKeys(opts *bind.TransactOpts, _publicKeySet []common.Address) (*types.Transaction, error) {
	return _Popcontract.contract.Transact(opts, "depositPublicKeys", _publicKeySet)
}

// DepositPublicKeys is a paid mutator transaction binding the contract method 0x65472d7c.
//
// Solidity: function depositPublicKeys(_publicKeySet address[]) returns(bool)
func (_Popcontract *PopcontractSession) DepositPublicKeys(_publicKeySet []common.Address) (*types.Transaction, error) {
	return _Popcontract.Contract.DepositPublicKeys(&_Popcontract.TransactOpts, _publicKeySet)
}

// DepositPublicKeys is a paid mutator transaction binding the contract method 0x65472d7c.
//
// Solidity: function depositPublicKeys(_publicKeySet address[]) returns(bool)
func (_Popcontract *PopcontractTransactorSession) DepositPublicKeys(_publicKeySet []common.Address) (*types.Transaction, error) {
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
