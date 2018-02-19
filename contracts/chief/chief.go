// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chief

import (
	"math/big"
	"strings"

	"github.com/SmartMeshFoundation/SMChain/accounts/abi"
	"github.com/SmartMeshFoundation/SMChain/accounts/abi/bind"
	"github.com/SmartMeshFoundation/SMChain/common"
	"github.com/SmartMeshFoundation/SMChain/core/types"
	"fmt"
)

// TribeChiefABI is the input ABI used to generate the binding from.
const TribeChiefABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"int256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TribeChiefBin is the compiled bytecode used for deploying new contracts.
const TribeChiefBin = `0x606060405260408051908101604052600581527f302e302e31000000000000000000000000000000000000000000000000000000602082015260009080516200004d929160200190620002e5565b506014600255600f60035534156200006457600080fd5b60405162000f3238038062000f328339810160405280805160018054600160a060020a03191633600160a060020a0316179055919091019050600080808080808651955060008611156200012b57600094505b858510156200012557868581518110620000cd57fe5b90602001906020020151600160a060020a0381166000908152600460205260409020805460ff1916600117905593506200011784600364010000000062000a786200021582021704565b5050600190940193620000b7565b62000208565b5050734110bd1ff0b73fa12c259acf39c950277f266787600081905260046020527f052387a23a063359ef51016da56c6ef818568f4a38e27c2728fda35f7b8ae85e805460ff19166001179055905073adb3ea3ad356199206ca817b04fd668cc5196df273b94b3aa41609e3f59cbaff3c2c298c6cc4c50b81620001bf83600364010000000062000a786200021582021704565b5050620001e2826003620002156401000000000262000a78176401000000009004565b505062000205816003620002156401000000000262000a78176401000000009004565b50505b50505050505050620003b6565b600780546000918291606f9010620002345760009250829150620002dd565b80548190600181016200024883826200036a565b5060009182526020909120018054600160a060020a031916600160a060020a03871617905560018181018054909181016200028483826200036a565b50600091825260209091200184905560038101805460018101620002a983826200036a565b506000918252602080832043920191909155600160a060020a03871682526002830190526040902084905580549250600191505b509250929050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200032857805160ff191683800117855562000358565b8280016001018555821562000358579182015b82811115620003585782518255916020019190600101906200033b565b506200036692915062000396565b5090565b81548183558181151162000391576000838152602090206200039191810190830162000396565b505050565b620003b391905b808211156200036657600081556001016200039d565b90565b610b6c80620003c66000396000f3006060604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631c1b8772811461005b5780634e69d5601461007c57806354fd4d50146101b9575b600080fd5b341561006657600080fd5b61007a600160a060020a0360043516610243565b005b341561008757600080fd5b61008f6103f9565b604051808060200180602001806020018060200186815260200185810385528a818151815260200191508051906020019060200280838360005b838110156100e15780820151838201526020016100c9565b50505050905001858103845289818151815260200191508051906020019060200280838360005b83811015610120578082015183820152602001610108565b50505050905001858103835288818151815260200191508051906020019060200280838360005b8381101561015f578082015183820152602001610147565b50505050905001858103825287818151815260200191508051906020019060200280838360005b8381101561019e578082015183820152602001610186565b50505050905001995050505050505050505060405180910390f35b34156101c457600080fd5b6101cc610593565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102085780820151838201526020016101f0565b50505050905090810190601f1680156102355780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b33600160a060020a038116600090815260096020526040812054909182918291829182919082901361027457600080fd5b4360068190556002543397509011801561029957506002544381151561029657fe5b06155b156102ec57600b549450600019850193505b600084106102ec57600b80546102e09190869081106102c657fe5b600091825260209091200154600160a060020a031661063c565b600019909301926102ab565b600160a060020a038716158015906103085750610308876107a6565b1561031957610316876107ac565b50505b60075460065481151561032857fe5b06925060076000018381548110151561033d57fe5b600091825260209091200154600160a060020a038781169116146103d157600880548490811061036957fe5b60009182526020909120015460001901915060018213156103c357600880548391908590811061039557fe5b600091825260209091200155600654600a8054859081106103b257fe5b6000918252602090912001556103cc565b6103cc83610854565b6103f0565b6008805460039190859081106103e357fe5b6000918252602090912001555b50505050505050565b610401610a29565b610409610a29565b610411610a29565b610419610a29565b6000600b60000180548060200260200160405190810160405280929190818152602001828054801561047457602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610456575b5050505050945060076000018054806020026020016040519081016040528092919081815260200182805480156104d457602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116104b6575b50505050509350600760010180548060200260200160405190810160405280929190818152602001828054801561052a57602002820191906000526020600020905b815481526020019060010190808311610516575b50505050509250600760030180548060200260200160405190810160405280929190818152602001828054801561058057602002820191906000526020600020905b81548152602001906001019080831161056c575b5050505050915060065490509091929394565b61059b610a29565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106315780601f1061060657610100808354040283529160200191610631565b820191906000526020600020905b81548152906001019060200180831161061457829003601f168201915b505050505090505b90565b600b54600160a060020a0382166000908152600d602052604081205490808083111561079f5750506000198101805b6001840381101561071557600b80546001830190811061068757fe5b600091825260209091200154600b8054600160a060020a0390921691839081106106ad57fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600c8054600183019081106106e757fe5b600091825260209091200154600c80548390811061070157fe5b60009182526020909120015560010161066b565b600b8054600019860190811061072757fe5b60009182526020909120018054600160a060020a0319169055600c8054600019860190811061075257fe5b6000918252602082200155600b805460001901906107709082610a3b565b50600c805460001901906107849082610a3b565b50600160a060020a0385166000908152600d60205260408120555b5050505050565b50600190565b600b54600090819061014d90106107c85750600090508061084f565b6107d18361063c565b600b8054600181016107e38382610a3b565b5060009182526020909120018054600160a060020a031916600160a060020a038516179055600c80546001810161081a8382610a3b565b506000918252602080832043920191909155600b54600160a060020a0386168352600d90915260409091208190559150600190505b915091565b600754600081831061086557610a24565b600780546009916000918690811061087957fe5b6000918252602080832090910154600160a060020a0316835282019290925260400181205550815b600182038110156109825760078054600183019081106108bd57fe5b60009182526020909120015460078054600160a060020a0390921691839081106108e357fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600880546001830190811061091d57fe5b600091825260209091200154600880548390811061093757fe5b600091825260209091200155600a80546001830190811061095457fe5b600091825260209091200154600a80548390811061096e57fe5b6000918252602090912001556001016108a1565b60078054600019840190811061099457fe5b60009182526020909120018054600160a060020a03191690556008805460001984019081106109bf57fe5b6000918252602082200155600a805460001984019081106109dc57fe5b60009182526020822001556007805460001901906109fa9082610a3b565b50600880546000190190610a0e9082610a3b565b50600a80546000190190610a229082610a3b565b505b505050565b60206040519081016040526000815290565b815481835581811511610a2457600083815260209020610a2491810190830161063991905b80821115610a745760008155600101610a60565b5090565b600780546000918291606f9010610a955760009250829150610b38565b8054819060018101610aa78382610a3b565b5060009182526020909120018054600160a060020a031916600160a060020a0387161790556001818101805490918101610ae18382610a3b565b50600091825260209091200184905560038101805460018101610b048382610a3b565b506000918252602080832043920191909155600160a060020a03871682526002830190526040902084905580549250600191505b5092509290505600a165627a7a7230582049c90fe7698fb748d3197c554a4f650e111d07ff4896874ff12cd2202c2bc2bc0029`

// DeployTribeChief deploys a new Ethereum contract, binding an instance of TribeChief to it.
func DeployTribeChief(auth *bind.TransactOpts, backend bind.ContractBackend, genesisSigners []common.Address) (common.Address, *types.Transaction, *TribeChief, error) {
	fmt.Println("<><><><><><><><><><>",auth)
	parsed, err := abi.JSON(strings.NewReader(TribeChiefABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribeChiefBin), backend, genesisSigners)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TribeChief{TribeChiefCaller: TribeChiefCaller{contract: contract}, TribeChiefTransactor: TribeChiefTransactor{contract: contract}}, nil
}

// TribeChief is an auto generated Go binding around an Ethereum contract.
type TribeChief struct {
	TribeChiefCaller     // Read-only binding to the contract
	TribeChiefTransactor // Write-only binding to the contract
}

// TribeChiefCaller is an auto generated read-only Go binding around an Ethereum contract.
type TribeChiefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChiefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TribeChiefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribeChiefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TribeChiefSession struct {
	Contract     *TribeChief       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TribeChiefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TribeChiefCallerSession struct {
	Contract *TribeChiefCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TribeChiefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TribeChiefTransactorSession struct {
	Contract     *TribeChiefTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TribeChiefRaw is an auto generated low-level Go binding around an Ethereum contract.
type TribeChiefRaw struct {
	Contract *TribeChief // Generic contract binding to access the raw methods on
}

// TribeChiefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TribeChiefCallerRaw struct {
	Contract *TribeChiefCaller // Generic read-only contract binding to access the raw methods on
}

// TribeChiefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TribeChiefTransactorRaw struct {
	Contract *TribeChiefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTribeChief creates a new instance of TribeChief, bound to a specific deployed contract.
func NewTribeChief(address common.Address, backend bind.ContractBackend) (*TribeChief, error) {
	contract, err := bindTribeChief(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TribeChief{TribeChiefCaller: TribeChiefCaller{contract: contract}, TribeChiefTransactor: TribeChiefTransactor{contract: contract}}, nil
}

// NewTribeChiefCaller creates a new read-only instance of TribeChief, bound to a specific deployed contract.
func NewTribeChiefCaller(address common.Address, caller bind.ContractCaller) (*TribeChiefCaller, error) {
	contract, err := bindTribeChief(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TribeChiefCaller{contract: contract}, nil
}

// NewTribeChiefTransactor creates a new write-only instance of TribeChief, bound to a specific deployed contract.
func NewTribeChiefTransactor(address common.Address, transactor bind.ContractTransactor) (*TribeChiefTransactor, error) {
	contract, err := bindTribeChief(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TribeChiefTransactor{contract: contract}, nil
}

// bindTribeChief binds a generic wrapper to an already deployed contract.
func bindTribeChief(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TribeChiefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief *TribeChiefRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TribeChief.Contract.TribeChiefCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief *TribeChiefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief.Contract.TribeChiefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief *TribeChiefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief.Contract.TribeChiefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TribeChief *TribeChiefCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TribeChief.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TribeChief *TribeChiefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TribeChief.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TribeChief *TribeChiefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TribeChief.Contract.contract.Transact(opts, method, params...)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList int256[], numberList uint256[], number uint256)
func (_TribeChief *TribeChiefCaller) GetStatus(opts *bind.CallOpts) (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	ret := new(struct {
		VolunteerList []common.Address
		SignerList    []common.Address
		ScoreList     []*big.Int
		NumberList    []*big.Int
		Number        *big.Int
	})
	out := ret
	err := _TribeChief.contract.Call(opts, out, "getStatus")
	return *ret, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList int256[], numberList uint256[], number uint256)
func (_TribeChief *TribeChiefSession) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief.Contract.GetStatus(&_TribeChief.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(volunteerList address[], signerList address[], scoreList int256[], numberList uint256[], number uint256)
func (_TribeChief *TribeChiefCallerSession) GetStatus() (struct {
	VolunteerList []common.Address
	SignerList    []common.Address
	ScoreList     []*big.Int
	NumberList    []*big.Int
	Number        *big.Int
}, error) {
	return _TribeChief.Contract.GetStatus(&_TribeChief.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief *TribeChiefCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TribeChief.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief *TribeChiefSession) Version() (string, error) {
	return _TribeChief.Contract.Version(&_TribeChief.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TribeChief *TribeChiefCallerSession) Version() (string, error) {
	return _TribeChief.Contract.Version(&_TribeChief.CallOpts)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief *TribeChiefTransactor) Update(opts *bind.TransactOpts, volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief.contract.Transact(opts, "update", volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief *TribeChiefSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief.Contract.Update(&_TribeChief.TransactOpts, volunteer)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(volunteer address) returns()
func (_TribeChief *TribeChiefTransactorSession) Update(volunteer common.Address) (*types.Transaction, error) {
	return _TribeChief.Contract.Update(&_TribeChief.TransactOpts, volunteer)
}
