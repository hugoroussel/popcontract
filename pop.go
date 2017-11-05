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
const MortalBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b60bb8061003b6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341c0e1b58114603c575b600080fd5b3415604657600080fd5b604c604e565b005b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161415608c5760005473ffffffffffffffffffffffffffffffffffffffff16ff5b5b5600a165627a7a7230582043f3d353b752a925a1efeda23c79dd35586a9eddbf0984a8359c389b16766d010029`

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
const PopcontractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"place\",\"type\":\"string\"},{\"name\":\"organizers\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"address[]\"},{\"name\":\"durationInMinutes\",\"type\":\"uint256\"}],\"name\":\"setConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signWholeConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrganizersAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locationOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isOrganizer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignedConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_publicKeySet\",\"type\":\"address[]\"}],\"name\":\"depositPublicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"configSignOrganizers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"publicKeyConsensus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signedConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"organizersAdresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allSets\",\"outputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finalKeySet\",\"outputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameOfParty\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// PopcontractBin is the compiled bytecode used for deploying new contracts.
const PopcontractBin = `0x6060604052600b8054600160a060020a0319908116909155600c80549091166012179055341561002e57600080fd5b5b5b60008054600160a060020a03191633600160a060020a03161790555b600480546000919060ff19166001835b02179055505b5b611283806100726000396000f300606060405236156101035763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166270041481146101085780630c3f6acf146101b8578063210790c8146101ef57806323266dfe14610216578063232a6b9d1461027d57806330756ec3146102a457806341676f151461032f57806341c0e1b5146103625780635a52ecf61461037757806365472d7c146103de57806372ee91c2146104415780638e7d0f44146104705780639f638dbf14610497578063b9bcc99f146104c9578063c1b8471b146104fb578063c1f1a9921461052d578063c47f00271461055c578063d7e3a4d3146105c1578063eff64cb3146105e6575b600080fd5b341561011357600080fd5b6101a460046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650509335935061067192505050565b604051901515815260200160405180910390f35b34156101c357600080fd5b6101cb610755565b604051808260048111156101db57fe5b60ff16815260200191505060405180910390f35b34156101fa57600080fd5b6101a461075e565b604051901515815260200160405180910390f35b341561022157600080fd5b610229610882565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156102695780820151818401525b602001610250565b505050509050019250505060405180910390f35b341561028857600080fd5b6101a46108eb565b604051901515815260200160405180910390f35b34156102af57600080fd5b6102b761090c565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102f45780820151818401525b6020016102db565b50505050905090810190601f1680156103215780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561033a57600080fd5b6101a4600160a060020a03600435166109aa565b604051901515815260200160405180910390f35b341561036d57600080fd5b610375610a5a565b005b341561038257600080fd5b610229610a82565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156102695780820151818401525b602001610250565b505050509050019250505060405180910390f35b34156103e957600080fd5b6101a46004602481358181019083013580602081810201604051908101604052809392919081815260200183836020028082843750949650610aeb95505050505050565b604051901515815260200160405180910390f35b341561044c57600080fd5b610454610bfe565b604051600160a060020a03909116815260200160405180910390f35b341561047b57600080fd5b6101a4610d0a565b604051901515815260200160405180910390f35b34156104a257600080fd5b610454600435610de2565b604051600160a060020a03909116815260200160405180910390f35b34156104d457600080fd5b610454600435610e14565b604051600160a060020a03909116815260200160405180910390f35b341561050657600080fd5b610454600435610e46565b604051600160a060020a03909116815260200160405180910390f35b341561053857600080fd5b610454610e75565b604051600160a060020a03909116815260200160405180910390f35b341561056757600080fd5b6101a460046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610e8495505050505050565b604051901515815260200160405180910390f35b34156105cc57600080fd5b6105d4610ed2565b60405190815260200160405180910390f35b34156105f157600080fd5b6102b7610ed8565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102f45780820151818401525b6020016102db565b50505050905090810190601f1680156103215780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60048054600091829160ff169081111561068757fe5b81600481111561069357fe5b14156101035760005433600160a060020a03908116911614156101035742603c84020160075560068680516106cc929160200190610f76565b506008859055846106de600982610ff5565b5060098480516106f292916020019061101f565b50600480546001919060ff191682805b021790555084610713600a82610ff5565b50600c805474ff00000000000000000000000000000000000000001916905560019150610740565b600080fd5b5b61074b565b600080fd5b5b50949350505050565b60045460ff1681565b600480546000918291829160019160ff169081111561077957fe5b81600481111561078557fe5b1415610103576000805490935033600160a060020a039081169116141561086f57600091505b60085482101561084f57600b54600a8054600160a060020a0390921691600385019081106107d557fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a0316141561080957600080fd5b600c805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055600192505b5b6001909101906107ab565b821561086f57600480546002919060ff19166001835b0217905550600193505b5b5b61087b565b600080fd5b5b50505090565b61088a611087565b60098054806020026020016040519081016040528092919081815260200182805480156108e057602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116108c2575b505050505090505b90565b600c5474010000000000000000000000000000000000000000900460ff1681565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109a25780601f10610977576101008083540402835291602001916109a2565b820191906000526020600020905b81548152906001019060200180831161098557829003601f168201915b505050505081565b60048054600091829160029160ff909116908111156109c557fe5b8160048111156109d157fe5b141561010357600091505b600854821015610a435783600160a060020a03166009838154811015156109ff57fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a03161415610a375760019250610a48565b5b6001909101906109dc565b600092505b610a52565b600080fd5b5b5050919050565b60005433600160a060020a0390811691161415610a7f57600054600160a060020a0316ff5b5b565b610a8a611087565b600a8054806020026020016040519081016040528092919081815260200182805480156108e057602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116108c2575b505050505090505b90565b6004805460009160029160ff1690811115610b0257fe5b816004811115610b0e57fe5b1415610103576007544211610beb57610b26336109aa565b15610be65760018054808201610b3c8382611099565b916000526020600020906002020160005b60408051908101604052600160a060020a03331681526020810187905291905081518154600160a060020a031916600160a060020a0391909116178155602082015181600101908051610ba492916020019061101f565b5060039250610bb1915050565b6004805460ff1690811115610bc257fe5b14610bdd57600480546003919060ff19166001835b02179055505b60019150610beb565b600091505b5b5b610bf7565b600080fd5b5b50919050565b600480546000918291829160019160ff1690811115610c1957fe5b816004811115610c2557fe5b14156101035760009250600091505b600854821015610c95576009805483908110610c4c57fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a031633600160a060020a03161415610c89578192505b5b600190910190610c34565b600a805460018101610ca78382610ff5565b916000526020600020900160005b6009805487908110610cc357fe5b906000526020600020900160005b9054835461010093840a600160a060020a039390940a90910482168302929091021916179055505b61087b565b600080fd5b5b50505090565b6004805460009160039160ff1690811115610d2157fe5b816004811115610d2d57fe5b1415610103576007544210610dd15760015415801590610d5b575060005433600160a060020a039081169116145b15610dcc5760018054fe5b906000526020600020906002020160005b50805460028054600160a060020a031916600160a060020a03909216919091178155600182018054610dab9160039161115d565b505060048054909150819060ff19166001825b021790555060019150610dd1565b600091505b5b5b610ddd565b600080fd5b5b5090565b600a805482908110610df057fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b6009805482908110610df057fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b6001805482908110610e5457fe5b906000526020600020906002020160005b5054600160a060020a0316905081565b600254600160a060020a031681565b60048054600091829160ff1690811115610e9a57fe5b816004811115610ea657fe5b1415610103576005838051610ebf929160200190610f76565b505b610bf7565b600080fd5b5b50919050565b60075481565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109a25780601f10610977576101008083540402835291602001916109a2565b820191906000526020600020905b81548152906001019060200180831161098557829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610fb757805160ff1916838001178555610fe4565b82800160010185558215610fe4579182015b82811115610fe4578251825591602001919060010190610fc9565b5b50610ddd9291506111ae565b5090565b815481835581811511611019576000838152602090206110199181019083016111ae565b5b505050565b828054828255906000526020600020908101928215611076579160200282015b828111156110765782518254600160a060020a031916600160a060020a03919091161782556020929092019160019091019061103f565b5b50610ddd9291506111cf565b5090565b60206040519081016040526000815290565b8154818355818115116110195760020281600202836000526020600020918201910161101991906111fa565b5b505050565b828054828255906000526020600020908101928215611076579160200282015b828111156110765782518254600160a060020a031916600160a060020a03919091161782556020929092019160019091019061103f565b5b50610ddd9291506111cf565b5090565b815481835581811511611019576000838152602090206110199181019083016111ae565b5b505050565b8280548282559060005260206000209081019282156110765760005260206000209182015b82811115611076578254825591600101919060010190611182565b5b50610ddd9291506111cf565b5090565b6108e891905b80821115610ddd57600081556001016111b4565b5090565b90565b6108e891905b80821115610ddd578054600160a060020a03191681556001016111d5565b5090565b90565b6108e891905b80821115610ddd578054600160a060020a031916815560006112256001830182611235565b50600201611200565b5090565b90565b508054600082559060005260206000209081019061125391906111ae565b5b505600a165627a7a7230582086eabb14f25fc2c3aa7ac319daee9a8e0fcb29368a4f09d6fa0620e0c4c561e80029`

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
