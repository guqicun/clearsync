// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package consensus_app

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ExitFormatAllocation is an auto generated low-level Go binding around an user-defined struct.
type ExitFormatAllocation struct {
	Destination    [32]byte
	Amount         *big.Int
	AllocationType uint8
	Metadata       []byte
}

// ExitFormatAssetMetadata is an auto generated low-level Go binding around an user-defined struct.
type ExitFormatAssetMetadata struct {
	AssetType uint8
	Metadata  []byte
}

// ExitFormatSingleAssetExit is an auto generated low-level Go binding around an user-defined struct.
type ExitFormatSingleAssetExit struct {
	Asset         common.Address
	AssetMetadata ExitFormatAssetMetadata
	Allocations   []ExitFormatAllocation
}

// INitroTypesFixedPart is an auto generated low-level Go binding around an user-defined struct.
type INitroTypesFixedPart struct {
	Participants      []common.Address
	ChannelNonce      uint64
	AppDefinition     common.Address
	ChallengeDuration *big.Int
}

// INitroTypesRecoveredVariablePart is an auto generated low-level Go binding around an user-defined struct.
type INitroTypesRecoveredVariablePart struct {
	VariablePart INitroTypesVariablePart
	SignedBy     *big.Int
}

// INitroTypesVariablePart is an auto generated low-level Go binding around an user-defined struct.
type INitroTypesVariablePart struct {
	Outcome []ExitFormatSingleAssetExit
	AppData []byte
	TurnNum *big.Int
	IsFinal bool
}

// ConsensusAppMetaData contains all meta data concerning the ConsensusApp contract.
var ConsensusAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"signedBy\",\"type\":\"uint256\"}],\"internalType\":\"structINitroTypes.RecoveredVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"signedBy\",\"type\":\"uint256\"}],\"internalType\":\"structINitroTypes.RecoveredVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"stateIsSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610922806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80639936d81214610030575b600080fd5b61004361003e36600461017c565b61005a565b604051610051929190610243565b60405180910390f35b6000606061008261006a876103c0565b6100748688610824565b61007d86610898565b61009f565b505060408051602081019091526000815260019094509492505050565b8151156100e05760405162461bcd60e51b815260206004820152600a60248201526907c70726f6f667c213d360b41b60448201526064015b60405180910390fd5b82515160208201516100f190610133565b60ff161461012e5760405162461bcd60e51b815260206004820152600a60248201526921756e616e696d6f757360b01b60448201526064016100d7565b505050565b6000805b821561015e576101486001846108ba565b9092169180610156816108cd565b915050610137565b92915050565b60006040828403121561017657600080fd5b50919050565b6000806000806060858703121561019257600080fd5b843567ffffffffffffffff808211156101aa57600080fd5b90860190608082890312156101be57600080fd5b909450602086013590808211156101d457600080fd5b818701915087601f8301126101e857600080fd5b8135818111156101f757600080fd5b8860208260051b850101111561020c57600080fd5b60208301955080945050604087013591508082111561022a57600080fd5b5061023787828801610164565b91505092959194509250565b82151581526000602060406020840152835180604085015260005b8181101561027a5785810183015185820160600152820161025e565b506000606082860101526060601f19601f830116850101925050509392505050565b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff811182821017156102d5576102d561029c565b60405290565b6040805190810167ffffffffffffffff811182821017156102d5576102d561029c565b6040516060810167ffffffffffffffff811182821017156102d5576102d561029c565b604051601f8201601f1916810167ffffffffffffffff8111828210171561034a5761034a61029c565b604052919050565b600067ffffffffffffffff82111561036c5761036c61029c565b5060051b60200190565b80356001600160a01b038116811461038d57600080fd5b919050565b803567ffffffffffffffff8116811461038d57600080fd5b803565ffffffffffff8116811461038d57600080fd5b6000608082360312156103d257600080fd5b6103da6102b2565b823567ffffffffffffffff8111156103f157600080fd5b830136601f82011261040257600080fd5b8035602061041761041283610352565b610321565b82815260059290921b8301810191818101903684111561043657600080fd5b938201935b8385101561045b5761044c85610376565b8252938201939082019061043b565b855250610469868201610392565b8185015250505061047c60408401610376565b604082015261048d606084016103aa565b606082015292915050565b600082601f8301126104a957600080fd5b813567ffffffffffffffff8111156104c3576104c361029c565b6104d6601f8201601f1916602001610321565b8181528460208386010111156104eb57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261051957600080fd5b8135602061052961041283610352565b82815260059290921b8401810191818101908684111561054857600080fd5b8286015b848110156105f157803567ffffffffffffffff8082111561056d5760008081fd5b908801906080828b03601f19018113156105875760008081fd5b61058f6102b2565b8784013581526040808501358983015260608086013560ff811681146105b55760008081fd5b838301529285013592848411156105ce57600091508182fd5b6105dc8e8b86890101610498565b9083015250865250505091830191830161054c565b509695505050505050565b8035801515811461038d57600080fd5b60006040828403121561061e57600080fd5b6106266102db565b905067ffffffffffffffff808335111561063f57600080fd5b823583016080818603121561065357600080fd5b61065b6102b2565b828235111561066957600080fd5b85601f83358401011261067b57600080fd5b61068b6104128335840135610352565b8235830180358083526020808401939260059290921b909101018810156106b157600080fd5b602084358501015b84358501803560051b016020018110156107c25785813511156106db57600080fd5b84358501813501601f196060828c03820112156106f757600080fd5b6106ff6102fe565b61070b60208401610376565b8152886040840135111561071e57600080fd5b60408301358301604083828f0301121561073757600080fd5b61073f6102db565b9250600460208201351061075257600080fd5b60208101358352896040820135111561076a57600080fd5b61077d8d60206040840135840101610498565b602084015250816020820152886060840135111561079a57600080fd5b6107ad8c60206060860135860101610508565b604082015285525050602092830192016106b9565b5082525060208201358310156107d757600080fd5b6107e78660208401358401610498565b60208201526107f8604083016103aa565b6040820152610809606083016105fc565b60608201528084525050506020820135602082015292915050565b600061083261041284610352565b80848252602080830192508560051b85013681111561085057600080fd5b855b8181101561088c57803567ffffffffffffffff8111156108725760008081fd5b61087e36828a0161060c565b865250938201938201610852565b50919695505050505050565b600061015e368361060c565b634e487b7160e01b600052601160045260246000fd5b8181038181111561015e5761015e6108a4565b600060ff821660ff81036108e3576108e36108a4565b6001019291505056fea26469706673582212205c109336fad9e53e762a48b6168d9e81e39613170e9f6d7b18aa3e48d343ed4764736f6c63430008160033",
}

// ConsensusAppABI is the input ABI used to generate the binding from.
// Deprecated: Use ConsensusAppMetaData.ABI instead.
var ConsensusAppABI = ConsensusAppMetaData.ABI

// ConsensusAppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConsensusAppMetaData.Bin instead.
var ConsensusAppBin = ConsensusAppMetaData.Bin

// DeployConsensusApp deploys a new Ethereum contract, binding an instance of ConsensusApp to it.
func DeployConsensusApp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ConsensusApp, error) {
	parsed, err := ConsensusAppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConsensusAppBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConsensusApp{ConsensusAppCaller: ConsensusAppCaller{contract: contract}, ConsensusAppTransactor: ConsensusAppTransactor{contract: contract}, ConsensusAppFilterer: ConsensusAppFilterer{contract: contract}}, nil
}

// ConsensusApp is an auto generated Go binding around an Ethereum contract.
type ConsensusApp struct {
	ConsensusAppCaller     // Read-only binding to the contract
	ConsensusAppTransactor // Write-only binding to the contract
	ConsensusAppFilterer   // Log filterer for contract events
}

// ConsensusAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsensusAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsensusAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsensusAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsensusAppSession struct {
	Contract     *ConsensusApp     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsensusAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsensusAppCallerSession struct {
	Contract *ConsensusAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ConsensusAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsensusAppTransactorSession struct {
	Contract     *ConsensusAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ConsensusAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsensusAppRaw struct {
	Contract *ConsensusApp // Generic contract binding to access the raw methods on
}

// ConsensusAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsensusAppCallerRaw struct {
	Contract *ConsensusAppCaller // Generic read-only contract binding to access the raw methods on
}

// ConsensusAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsensusAppTransactorRaw struct {
	Contract *ConsensusAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsensusApp creates a new instance of ConsensusApp, bound to a specific deployed contract.
func NewConsensusApp(address common.Address, backend bind.ContractBackend) (*ConsensusApp, error) {
	contract, err := bindConsensusApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConsensusApp{ConsensusAppCaller: ConsensusAppCaller{contract: contract}, ConsensusAppTransactor: ConsensusAppTransactor{contract: contract}, ConsensusAppFilterer: ConsensusAppFilterer{contract: contract}}, nil
}

// NewConsensusAppCaller creates a new read-only instance of ConsensusApp, bound to a specific deployed contract.
func NewConsensusAppCaller(address common.Address, caller bind.ContractCaller) (*ConsensusAppCaller, error) {
	contract, err := bindConsensusApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusAppCaller{contract: contract}, nil
}

// NewConsensusAppTransactor creates a new write-only instance of ConsensusApp, bound to a specific deployed contract.
func NewConsensusAppTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsensusAppTransactor, error) {
	contract, err := bindConsensusApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusAppTransactor{contract: contract}, nil
}

// NewConsensusAppFilterer creates a new log filterer instance of ConsensusApp, bound to a specific deployed contract.
func NewConsensusAppFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsensusAppFilterer, error) {
	contract, err := bindConsensusApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsensusAppFilterer{contract: contract}, nil
}

// bindConsensusApp binds a generic wrapper to an already deployed contract.
func bindConsensusApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConsensusAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsensusApp *ConsensusAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsensusApp.Contract.ConsensusAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsensusApp *ConsensusAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsensusApp.Contract.ConsensusAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsensusApp *ConsensusAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsensusApp.Contract.ConsensusAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsensusApp *ConsensusAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsensusApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsensusApp *ConsensusAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsensusApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsensusApp *ConsensusAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsensusApp.Contract.contract.Transact(opts, method, params...)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x9936d812.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256)[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256) candidate) pure returns(bool, string)
func (_ConsensusApp *ConsensusAppCaller) StateIsSupported(opts *bind.CallOpts, fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	var out []interface{}
	err := _ConsensusApp.contract.Call(opts, &out, "stateIsSupported", fixedPart, proof, candidate)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// StateIsSupported is a free data retrieval call binding the contract method 0x9936d812.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256)[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256) candidate) pure returns(bool, string)
func (_ConsensusApp *ConsensusAppSession) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	return _ConsensusApp.Contract.StateIsSupported(&_ConsensusApp.CallOpts, fixedPart, proof, candidate)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x9936d812.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256)[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),uint256) candidate) pure returns(bool, string)
func (_ConsensusApp *ConsensusAppCallerSession) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesRecoveredVariablePart, candidate INitroTypesRecoveredVariablePart) (bool, string, error) {
	return _ConsensusApp.Contract.StateIsSupported(&_ConsensusApp.CallOpts, fixedPart, proof, candidate)
}
