// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = time.Tick
	_ = context.Background
)

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paymentId\",\"type\":\"uint256\"}],\"name\":\"DeliverableMissing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paymentId\",\"type\":\"uint256\"}],\"name\":\"DisputedPayment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paymentId\",\"type\":\"uint256\"}],\"name\":\"NotApproved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentId\",\"type\":\"uint256\"}],\"name\":\"NotBeneficiary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentId\",\"type\":\"uint256\"}],\"name\":\"NotYourDebt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"owed\",\"type\":\"uint256\"}],\"name\":\"OverpaymentError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paymentId\",\"type\":\"uint256\"}],\"name\":\"Payed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_debtor\",\"type\":\"address\"}],\"name\":\"SlashedDebtor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dueTime\",\"type\":\"uint256\"}],\"name\":\"UnmaturedPayment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paymentId\",\"type\":\"uint256\"}],\"name\":\"Unpayed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"deliverableCID\",\"type\":\"string\"}],\"name\":\"DeliverableSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"paymentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"DisputedPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"originalPaymentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"milestonePaymentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"deliverableCID\",\"type\":\"string\"}],\"name\":\"MilestoneApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_debtor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"NewPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"SlashedPayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_splittedId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_debtor\",\"type\":\"address\"}],\"name\":\"SplitPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dueTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_agreementCID\",\"type\":\"string\"}],\"name\":\"createPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"}],\"name\":\"disputePayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"}],\"name\":\"payWithPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payed\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"dueTime\",\"type\":\"uint48\"},{\"internalType\":\"string\",\"name\":\"agreementCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"deliverableCID\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"}],\"name\":\"slashPayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"slashed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"slashedPayments\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"}],\"name\":\"splitAndApproveMilestone\",\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_deliverableCID\",\"type\":\"string\"}],\"name\":\"submitDeliverable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_paymentID\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"}],\"name\":\"undisputePayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bindings *BindingsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bindings *BindingsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bindings.Contract.BalanceOf(&_Bindings.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bindings *BindingsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bindings.Contract.BalanceOf(&_Bindings.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bindings *BindingsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bindings *BindingsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bindings.Contract.GetApproved(&_Bindings.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bindings *BindingsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bindings.Contract.GetApproved(&_Bindings.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bindings *BindingsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bindings *BindingsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bindings.Contract.IsApprovedForAll(&_Bindings.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bindings *BindingsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bindings.Contract.IsApprovedForAll(&_Bindings.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bindings *BindingsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bindings *BindingsSession) Name() (string, error) {
	return _Bindings.Contract.Name(&_Bindings.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bindings *BindingsCallerSession) Name() (string, error) {
	return _Bindings.Contract.Name(&_Bindings.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bindings *BindingsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bindings *BindingsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bindings.Contract.OwnerOf(&_Bindings.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bindings *BindingsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bindings.Contract.OwnerOf(&_Bindings.CallOpts, tokenId)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(address beneficiary, uint256 amount, uint256 payed, uint48 dueTime, string agreementCID, string deliverableCID, bool approved, bool disputed)
func (_Bindings *BindingsCaller) Payments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Beneficiary    common.Address
	Amount         *big.Int
	Payed          *big.Int
	DueTime        *big.Int
	AgreementCID   string
	DeliverableCID string
	Approved       bool
	Disputed       bool
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "payments", arg0)

	outstruct := new(struct {
		Beneficiary    common.Address
		Amount         *big.Int
		Payed          *big.Int
		DueTime        *big.Int
		AgreementCID   string
		DeliverableCID string
		Approved       bool
		Disputed       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Beneficiary = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Payed = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DueTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AgreementCID = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.DeliverableCID = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Approved = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Disputed = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(address beneficiary, uint256 amount, uint256 payed, uint48 dueTime, string agreementCID, string deliverableCID, bool approved, bool disputed)
func (_Bindings *BindingsSession) Payments(arg0 *big.Int) (struct {
	Beneficiary    common.Address
	Amount         *big.Int
	Payed          *big.Int
	DueTime        *big.Int
	AgreementCID   string
	DeliverableCID string
	Approved       bool
	Disputed       bool
}, error) {
	return _Bindings.Contract.Payments(&_Bindings.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(address beneficiary, uint256 amount, uint256 payed, uint48 dueTime, string agreementCID, string deliverableCID, bool approved, bool disputed)
func (_Bindings *BindingsCallerSession) Payments(arg0 *big.Int) (struct {
	Beneficiary    common.Address
	Amount         *big.Int
	Payed          *big.Int
	DueTime        *big.Int
	AgreementCID   string
	DeliverableCID string
	Approved       bool
	Disputed       bool
}, error) {
	return _Bindings.Contract.Payments(&_Bindings.CallOpts, arg0)
}

// Slashed is a free data retrieval call binding the contract method 0x8bcfa665.
//
// Solidity: function slashed(address ) view returns(uint256 slashedPayments)
func (_Bindings *BindingsCaller) Slashed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "slashed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Slashed is a free data retrieval call binding the contract method 0x8bcfa665.
//
// Solidity: function slashed(address ) view returns(uint256 slashedPayments)
func (_Bindings *BindingsSession) Slashed(arg0 common.Address) (*big.Int, error) {
	return _Bindings.Contract.Slashed(&_Bindings.CallOpts, arg0)
}

// Slashed is a free data retrieval call binding the contract method 0x8bcfa665.
//
// Solidity: function slashed(address ) view returns(uint256 slashedPayments)
func (_Bindings *BindingsCallerSession) Slashed(arg0 common.Address) (*big.Int, error) {
	return _Bindings.Contract.Slashed(&_Bindings.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bindings.Contract.SupportsInterface(&_Bindings.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bindings.Contract.SupportsInterface(&_Bindings.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bindings *BindingsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bindings *BindingsSession) Symbol() (string, error) {
	return _Bindings.Contract.Symbol(&_Bindings.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bindings *BindingsCallerSession) Symbol() (string, error) {
	return _Bindings.Contract.Symbol(&_Bindings.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _paymentID) view returns(string)
func (_Bindings *BindingsCaller) TokenURI(opts *bind.CallOpts, _paymentID *big.Int) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "tokenURI", _paymentID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _paymentID) view returns(string)
func (_Bindings *BindingsSession) TokenURI(_paymentID *big.Int) (string, error) {
	return _Bindings.Contract.TokenURI(&_Bindings.CallOpts, _paymentID)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _paymentID) view returns(string)
func (_Bindings *BindingsCallerSession) TokenURI(_paymentID *big.Int) (string, error) {
	return _Bindings.Contract.TokenURI(&_Bindings.CallOpts, _paymentID)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bindings *BindingsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bindings *BindingsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Approve(&_Bindings.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bindings *BindingsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Approve(&_Bindings.TransactOpts, to, tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _paymentId) returns()
func (_Bindings *BindingsTransactor) Claim(opts *bind.TransactOpts, _paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "claim", _paymentId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _paymentId) returns()
func (_Bindings *BindingsSession) Claim(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Claim(&_Bindings.TransactOpts, _paymentId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _paymentId) returns()
func (_Bindings *BindingsTransactorSession) Claim(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Claim(&_Bindings.TransactOpts, _paymentId)
}

// CreatePayment is a paid mutator transaction binding the contract method 0x67be6bed.
//
// Solidity: function createPayment(uint256 _dueTime, uint256 _amount, address _beneficiary, string _agreementCID) returns(uint256 id)
func (_Bindings *BindingsTransactor) CreatePayment(opts *bind.TransactOpts, _dueTime *big.Int, _amount *big.Int, _beneficiary common.Address, _agreementCID string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "createPayment", _dueTime, _amount, _beneficiary, _agreementCID)
}

// CreatePayment is a paid mutator transaction binding the contract method 0x67be6bed.
//
// Solidity: function createPayment(uint256 _dueTime, uint256 _amount, address _beneficiary, string _agreementCID) returns(uint256 id)
func (_Bindings *BindingsSession) CreatePayment(_dueTime *big.Int, _amount *big.Int, _beneficiary common.Address, _agreementCID string) (*types.Transaction, error) {
	return _Bindings.Contract.CreatePayment(&_Bindings.TransactOpts, _dueTime, _amount, _beneficiary, _agreementCID)
}

// CreatePayment is a paid mutator transaction binding the contract method 0x67be6bed.
//
// Solidity: function createPayment(uint256 _dueTime, uint256 _amount, address _beneficiary, string _agreementCID) returns(uint256 id)
func (_Bindings *BindingsTransactorSession) CreatePayment(_dueTime *big.Int, _amount *big.Int, _beneficiary common.Address, _agreementCID string) (*types.Transaction, error) {
	return _Bindings.Contract.CreatePayment(&_Bindings.TransactOpts, _dueTime, _amount, _beneficiary, _agreementCID)
}

// DisputePayment is a paid mutator transaction binding the contract method 0xd81f4874.
//
// Solidity: function disputePayment(uint256 _paymentId) returns()
func (_Bindings *BindingsTransactor) DisputePayment(opts *bind.TransactOpts, _paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "disputePayment", _paymentId)
}

// DisputePayment is a paid mutator transaction binding the contract method 0xd81f4874.
//
// Solidity: function disputePayment(uint256 _paymentId) returns()
func (_Bindings *BindingsSession) DisputePayment(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.DisputePayment(&_Bindings.TransactOpts, _paymentId)
}

// DisputePayment is a paid mutator transaction binding the contract method 0xd81f4874.
//
// Solidity: function disputePayment(uint256 _paymentId) returns()
func (_Bindings *BindingsTransactorSession) DisputePayment(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.DisputePayment(&_Bindings.TransactOpts, _paymentId)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(uint256 _paymentId) payable returns()
func (_Bindings *BindingsTransactor) Pay(opts *bind.TransactOpts, _paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "pay", _paymentId)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(uint256 _paymentId) payable returns()
func (_Bindings *BindingsSession) Pay(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Pay(&_Bindings.TransactOpts, _paymentId)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(uint256 _paymentId) payable returns()
func (_Bindings *BindingsTransactorSession) Pay(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Pay(&_Bindings.TransactOpts, _paymentId)
}

// PayWithPayment is a paid mutator transaction binding the contract method 0x823abe5b.
//
// Solidity: function payWithPayment(uint256 _withId, uint256 _paymentId) returns()
func (_Bindings *BindingsTransactor) PayWithPayment(opts *bind.TransactOpts, _withId *big.Int, _paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "payWithPayment", _withId, _paymentId)
}

// PayWithPayment is a paid mutator transaction binding the contract method 0x823abe5b.
//
// Solidity: function payWithPayment(uint256 _withId, uint256 _paymentId) returns()
func (_Bindings *BindingsSession) PayWithPayment(_withId *big.Int, _paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PayWithPayment(&_Bindings.TransactOpts, _withId, _paymentId)
}

// PayWithPayment is a paid mutator transaction binding the contract method 0x823abe5b.
//
// Solidity: function payWithPayment(uint256 _withId, uint256 _paymentId) returns()
func (_Bindings *BindingsTransactorSession) PayWithPayment(_withId *big.Int, _paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PayWithPayment(&_Bindings.TransactOpts, _withId, _paymentId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SafeTransferFrom(&_Bindings.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SafeTransferFrom(&_Bindings.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bindings *BindingsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bindings *BindingsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bindings.Contract.SafeTransferFrom0(&_Bindings.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bindings *BindingsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bindings.Contract.SafeTransferFrom0(&_Bindings.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bindings *BindingsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bindings *BindingsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bindings.Contract.SetApprovalForAll(&_Bindings.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bindings *BindingsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bindings.Contract.SetApprovalForAll(&_Bindings.TransactOpts, operator, approved)
}

// SlashPayer is a paid mutator transaction binding the contract method 0xc7f81873.
//
// Solidity: function slashPayer(uint256 _paymentId) returns()
func (_Bindings *BindingsTransactor) SlashPayer(opts *bind.TransactOpts, _paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "slashPayer", _paymentId)
}

// SlashPayer is a paid mutator transaction binding the contract method 0xc7f81873.
//
// Solidity: function slashPayer(uint256 _paymentId) returns()
func (_Bindings *BindingsSession) SlashPayer(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SlashPayer(&_Bindings.TransactOpts, _paymentId)
}

// SlashPayer is a paid mutator transaction binding the contract method 0xc7f81873.
//
// Solidity: function slashPayer(uint256 _paymentId) returns()
func (_Bindings *BindingsTransactorSession) SlashPayer(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SlashPayer(&_Bindings.TransactOpts, _paymentId)
}

// SplitAndApproveMilestone is a paid mutator transaction binding the contract method 0x28bda35f.
//
// Solidity: function splitAndApproveMilestone(uint256 _paymentId) payable returns()
func (_Bindings *BindingsTransactor) SplitAndApproveMilestone(opts *bind.TransactOpts, _paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "splitAndApproveMilestone", _paymentId)
}

// SplitAndApproveMilestone is a paid mutator transaction binding the contract method 0x28bda35f.
//
// Solidity: function splitAndApproveMilestone(uint256 _paymentId) payable returns()
func (_Bindings *BindingsSession) SplitAndApproveMilestone(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SplitAndApproveMilestone(&_Bindings.TransactOpts, _paymentId)
}

// SplitAndApproveMilestone is a paid mutator transaction binding the contract method 0x28bda35f.
//
// Solidity: function splitAndApproveMilestone(uint256 _paymentId) payable returns()
func (_Bindings *BindingsTransactorSession) SplitAndApproveMilestone(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SplitAndApproveMilestone(&_Bindings.TransactOpts, _paymentId)
}

// SubmitDeliverable is a paid mutator transaction binding the contract method 0x5e29172f.
//
// Solidity: function submitDeliverable(uint256 _paymentId, string _deliverableCID) returns()
func (_Bindings *BindingsTransactor) SubmitDeliverable(opts *bind.TransactOpts, _paymentId *big.Int, _deliverableCID string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "submitDeliverable", _paymentId, _deliverableCID)
}

// SubmitDeliverable is a paid mutator transaction binding the contract method 0x5e29172f.
//
// Solidity: function submitDeliverable(uint256 _paymentId, string _deliverableCID) returns()
func (_Bindings *BindingsSession) SubmitDeliverable(_paymentId *big.Int, _deliverableCID string) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitDeliverable(&_Bindings.TransactOpts, _paymentId, _deliverableCID)
}

// SubmitDeliverable is a paid mutator transaction binding the contract method 0x5e29172f.
//
// Solidity: function submitDeliverable(uint256 _paymentId, string _deliverableCID) returns()
func (_Bindings *BindingsTransactorSession) SubmitDeliverable(_paymentId *big.Int, _deliverableCID string) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitDeliverable(&_Bindings.TransactOpts, _paymentId, _deliverableCID)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferFrom(&_Bindings.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferFrom(&_Bindings.TransactOpts, from, to, tokenId)
}

// UndisputePayment is a paid mutator transaction binding the contract method 0x2f339d82.
//
// Solidity: function undisputePayment(uint256 _paymentId) returns()
func (_Bindings *BindingsTransactor) UndisputePayment(opts *bind.TransactOpts, _paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "undisputePayment", _paymentId)
}

// UndisputePayment is a paid mutator transaction binding the contract method 0x2f339d82.
//
// Solidity: function undisputePayment(uint256 _paymentId) returns()
func (_Bindings *BindingsSession) UndisputePayment(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UndisputePayment(&_Bindings.TransactOpts, _paymentId)
}

// UndisputePayment is a paid mutator transaction binding the contract method 0x2f339d82.
//
// Solidity: function undisputePayment(uint256 _paymentId) returns()
func (_Bindings *BindingsTransactorSession) UndisputePayment(_paymentId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UndisputePayment(&_Bindings.TransactOpts, _paymentId)
}

// BindingsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bindings contract.
type BindingsApprovalIterator struct {
	Event *BindingsApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsApproval represents a Approval event raised by the Bindings contract.
type BindingsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BindingsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsApprovalIterator{contract: _Bindings.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BindingsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsApproval)
				if err := _Bindings.contract.UnpackLog(event, "Approval", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) ParseApproval(log types.Log) (*BindingsApproval, error) {
	event := new(BindingsApproval)
	if err := _Bindings.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Bindings contract.
type BindingsApprovalForAllIterator struct {
	Event *BindingsApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsApprovalForAll represents a ApprovalForAll event raised by the Bindings contract.
type BindingsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bindings *BindingsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BindingsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BindingsApprovalForAllIterator{contract: _Bindings.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bindings *BindingsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BindingsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsApprovalForAll)
				if err := _Bindings.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bindings *BindingsFilterer) ParseApprovalForAll(log types.Log) (*BindingsApprovalForAll, error) {
	event := new(BindingsApprovalForAll)
	if err := _Bindings.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDeliverableSubmittedIterator is returned from FilterDeliverableSubmitted and is used to iterate over the raw logs and unpacked data for DeliverableSubmitted events raised by the Bindings contract.
type BindingsDeliverableSubmittedIterator struct {
	Event *BindingsDeliverableSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsDeliverableSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDeliverableSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsDeliverableSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsDeliverableSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDeliverableSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDeliverableSubmitted represents a DeliverableSubmitted event raised by the Bindings contract.
type BindingsDeliverableSubmitted struct {
	PaymentId      *big.Int
	DeliverableCID string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeliverableSubmitted is a free log retrieval operation binding the contract event 0xf04befa55a56a87de7073cea8ab22499ccb81b540f0999577cf892224e3f721f.
//
// Solidity: event DeliverableSubmitted(uint256 indexed _paymentId, string deliverableCID)
func (_Bindings *BindingsFilterer) FilterDeliverableSubmitted(opts *bind.FilterOpts, _paymentId []*big.Int) (*BindingsDeliverableSubmittedIterator, error) {

	var _paymentIdRule []interface{}
	for _, _paymentIdItem := range _paymentId {
		_paymentIdRule = append(_paymentIdRule, _paymentIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DeliverableSubmitted", _paymentIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsDeliverableSubmittedIterator{contract: _Bindings.contract, event: "DeliverableSubmitted", logs: logs, sub: sub}, nil
}

// WatchDeliverableSubmitted is a free log subscription operation binding the contract event 0xf04befa55a56a87de7073cea8ab22499ccb81b540f0999577cf892224e3f721f.
//
// Solidity: event DeliverableSubmitted(uint256 indexed _paymentId, string deliverableCID)
func (_Bindings *BindingsFilterer) WatchDeliverableSubmitted(opts *bind.WatchOpts, sink chan<- *BindingsDeliverableSubmitted, _paymentId []*big.Int) (event.Subscription, error) {

	var _paymentIdRule []interface{}
	for _, _paymentIdItem := range _paymentId {
		_paymentIdRule = append(_paymentIdRule, _paymentIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DeliverableSubmitted", _paymentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDeliverableSubmitted)
				if err := _Bindings.contract.UnpackLog(event, "DeliverableSubmitted", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeliverableSubmitted is a log parse operation binding the contract event 0xf04befa55a56a87de7073cea8ab22499ccb81b540f0999577cf892224e3f721f.
//
// Solidity: event DeliverableSubmitted(uint256 indexed _paymentId, string deliverableCID)
func (_Bindings *BindingsFilterer) ParseDeliverableSubmitted(log types.Log) (*BindingsDeliverableSubmitted, error) {
	event := new(BindingsDeliverableSubmitted)
	if err := _Bindings.contract.UnpackLog(event, "DeliverableSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDisputedPaymentIterator is returned from FilterDisputedPayment and is used to iterate over the raw logs and unpacked data for DisputedPayment events raised by the Bindings contract.
type BindingsDisputedPaymentIterator struct {
	Event *BindingsDisputedPayment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsDisputedPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDisputedPayment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsDisputedPayment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsDisputedPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDisputedPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDisputedPayment represents a DisputedPayment event raised by the Bindings contract.
type BindingsDisputedPayment struct {
	PaymentId   *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDisputedPayment is a free log retrieval operation binding the contract event 0xd07097a5ca93fdef2443ef261061ea72d4735efdd892a0dc932d8f64e8ff4b11.
//
// Solidity: event DisputedPayment(uint256 indexed paymentId, address beneficiary)
func (_Bindings *BindingsFilterer) FilterDisputedPayment(opts *bind.FilterOpts, paymentId []*big.Int) (*BindingsDisputedPaymentIterator, error) {

	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DisputedPayment", paymentIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsDisputedPaymentIterator{contract: _Bindings.contract, event: "DisputedPayment", logs: logs, sub: sub}, nil
}

// WatchDisputedPayment is a free log subscription operation binding the contract event 0xd07097a5ca93fdef2443ef261061ea72d4735efdd892a0dc932d8f64e8ff4b11.
//
// Solidity: event DisputedPayment(uint256 indexed paymentId, address beneficiary)
func (_Bindings *BindingsFilterer) WatchDisputedPayment(opts *bind.WatchOpts, sink chan<- *BindingsDisputedPayment, paymentId []*big.Int) (event.Subscription, error) {

	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DisputedPayment", paymentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDisputedPayment)
				if err := _Bindings.contract.UnpackLog(event, "DisputedPayment", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputedPayment is a log parse operation binding the contract event 0xd07097a5ca93fdef2443ef261061ea72d4735efdd892a0dc932d8f64e8ff4b11.
//
// Solidity: event DisputedPayment(uint256 indexed paymentId, address beneficiary)
func (_Bindings *BindingsFilterer) ParseDisputedPayment(log types.Log) (*BindingsDisputedPayment, error) {
	event := new(BindingsDisputedPayment)
	if err := _Bindings.contract.UnpackLog(event, "DisputedPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMilestoneApprovedIterator is returned from FilterMilestoneApproved and is used to iterate over the raw logs and unpacked data for MilestoneApproved events raised by the Bindings contract.
type BindingsMilestoneApprovedIterator struct {
	Event *BindingsMilestoneApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsMilestoneApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMilestoneApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsMilestoneApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsMilestoneApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMilestoneApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMilestoneApproved represents a MilestoneApproved event raised by the Bindings contract.
type BindingsMilestoneApproved struct {
	OriginalPaymentId  *big.Int
	MilestonePaymentId *big.Int
	Amount             *big.Int
	DeliverableCID     string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMilestoneApproved is a free log retrieval operation binding the contract event 0x146d1d8999eac6ccbcd5e7c0c5bb38a3b969845e029980e35390a2b3bbd604b1.
//
// Solidity: event MilestoneApproved(uint256 indexed originalPaymentId, uint256 indexed milestonePaymentId, uint256 amount, string deliverableCID)
func (_Bindings *BindingsFilterer) FilterMilestoneApproved(opts *bind.FilterOpts, originalPaymentId []*big.Int, milestonePaymentId []*big.Int) (*BindingsMilestoneApprovedIterator, error) {

	var originalPaymentIdRule []interface{}
	for _, originalPaymentIdItem := range originalPaymentId {
		originalPaymentIdRule = append(originalPaymentIdRule, originalPaymentIdItem)
	}
	var milestonePaymentIdRule []interface{}
	for _, milestonePaymentIdItem := range milestonePaymentId {
		milestonePaymentIdRule = append(milestonePaymentIdRule, milestonePaymentIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MilestoneApproved", originalPaymentIdRule, milestonePaymentIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMilestoneApprovedIterator{contract: _Bindings.contract, event: "MilestoneApproved", logs: logs, sub: sub}, nil
}

// WatchMilestoneApproved is a free log subscription operation binding the contract event 0x146d1d8999eac6ccbcd5e7c0c5bb38a3b969845e029980e35390a2b3bbd604b1.
//
// Solidity: event MilestoneApproved(uint256 indexed originalPaymentId, uint256 indexed milestonePaymentId, uint256 amount, string deliverableCID)
func (_Bindings *BindingsFilterer) WatchMilestoneApproved(opts *bind.WatchOpts, sink chan<- *BindingsMilestoneApproved, originalPaymentId []*big.Int, milestonePaymentId []*big.Int) (event.Subscription, error) {

	var originalPaymentIdRule []interface{}
	for _, originalPaymentIdItem := range originalPaymentId {
		originalPaymentIdRule = append(originalPaymentIdRule, originalPaymentIdItem)
	}
	var milestonePaymentIdRule []interface{}
	for _, milestonePaymentIdItem := range milestonePaymentId {
		milestonePaymentIdRule = append(milestonePaymentIdRule, milestonePaymentIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MilestoneApproved", originalPaymentIdRule, milestonePaymentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMilestoneApproved)
				if err := _Bindings.contract.UnpackLog(event, "MilestoneApproved", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMilestoneApproved is a log parse operation binding the contract event 0x146d1d8999eac6ccbcd5e7c0c5bb38a3b969845e029980e35390a2b3bbd604b1.
//
// Solidity: event MilestoneApproved(uint256 indexed originalPaymentId, uint256 indexed milestonePaymentId, uint256 amount, string deliverableCID)
func (_Bindings *BindingsFilterer) ParseMilestoneApproved(log types.Log) (*BindingsMilestoneApproved, error) {
	event := new(BindingsMilestoneApproved)
	if err := _Bindings.contract.UnpackLog(event, "MilestoneApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsNewPaymentIterator is returned from FilterNewPayment and is used to iterate over the raw logs and unpacked data for NewPayment events raised by the Bindings contract.
type BindingsNewPaymentIterator struct {
	Event *BindingsNewPayment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsNewPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNewPayment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsNewPayment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsNewPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNewPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNewPayment represents a NewPayment event raised by the Bindings contract.
type BindingsNewPayment struct {
	Beneficiary common.Address
	Debtor      common.Address
	Id          *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPayment is a free log retrieval operation binding the contract event 0x629777e97c687ba3e899fb62b1abe65a4fc512287be8092b1f429286afdfc51d.
//
// Solidity: event NewPayment(address _beneficiary, address _debtor, uint256 indexed _id)
func (_Bindings *BindingsFilterer) FilterNewPayment(opts *bind.FilterOpts, _id []*big.Int) (*BindingsNewPaymentIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NewPayment", _idRule)
	if err != nil {
		return nil, err
	}
	return &BindingsNewPaymentIterator{contract: _Bindings.contract, event: "NewPayment", logs: logs, sub: sub}, nil
}

// WatchNewPayment is a free log subscription operation binding the contract event 0x629777e97c687ba3e899fb62b1abe65a4fc512287be8092b1f429286afdfc51d.
//
// Solidity: event NewPayment(address _beneficiary, address _debtor, uint256 indexed _id)
func (_Bindings *BindingsFilterer) WatchNewPayment(opts *bind.WatchOpts, sink chan<- *BindingsNewPayment, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NewPayment", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNewPayment)
				if err := _Bindings.contract.UnpackLog(event, "NewPayment", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPayment is a log parse operation binding the contract event 0x629777e97c687ba3e899fb62b1abe65a4fc512287be8092b1f429286afdfc51d.
//
// Solidity: event NewPayment(address _beneficiary, address _debtor, uint256 indexed _id)
func (_Bindings *BindingsFilterer) ParseNewPayment(log types.Log) (*BindingsNewPayment, error) {
	event := new(BindingsNewPayment)
	if err := _Bindings.contract.UnpackLog(event, "NewPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsSlashedPayerIterator is returned from FilterSlashedPayer and is used to iterate over the raw logs and unpacked data for SlashedPayer events raised by the Bindings contract.
type BindingsSlashedPayerIterator struct {
	Event *BindingsSlashedPayer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsSlashedPayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSlashedPayer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsSlashedPayer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsSlashedPayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSlashedPayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSlashedPayer represents a SlashedPayer event raised by the Bindings contract.
type BindingsSlashedPayer struct {
	Payer common.Address
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSlashedPayer is a free log retrieval operation binding the contract event 0xd632b134083fb9656e20b3db160ef7be0c227c73832637aa3f03df7b56f36217.
//
// Solidity: event SlashedPayer(address _payer, uint256 indexed _id)
func (_Bindings *BindingsFilterer) FilterSlashedPayer(opts *bind.FilterOpts, _id []*big.Int) (*BindingsSlashedPayerIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SlashedPayer", _idRule)
	if err != nil {
		return nil, err
	}
	return &BindingsSlashedPayerIterator{contract: _Bindings.contract, event: "SlashedPayer", logs: logs, sub: sub}, nil
}

// WatchSlashedPayer is a free log subscription operation binding the contract event 0xd632b134083fb9656e20b3db160ef7be0c227c73832637aa3f03df7b56f36217.
//
// Solidity: event SlashedPayer(address _payer, uint256 indexed _id)
func (_Bindings *BindingsFilterer) WatchSlashedPayer(opts *bind.WatchOpts, sink chan<- *BindingsSlashedPayer, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SlashedPayer", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSlashedPayer)
				if err := _Bindings.contract.UnpackLog(event, "SlashedPayer", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashedPayer is a log parse operation binding the contract event 0xd632b134083fb9656e20b3db160ef7be0c227c73832637aa3f03df7b56f36217.
//
// Solidity: event SlashedPayer(address _payer, uint256 indexed _id)
func (_Bindings *BindingsFilterer) ParseSlashedPayer(log types.Log) (*BindingsSlashedPayer, error) {
	event := new(BindingsSlashedPayer)
	if err := _Bindings.contract.UnpackLog(event, "SlashedPayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsSplitPaymentIterator is returned from FilterSplitPayment and is used to iterate over the raw logs and unpacked data for SplitPayment events raised by the Bindings contract.
type BindingsSplitPaymentIterator struct {
	Event *BindingsSplitPayment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsSplitPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSplitPayment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsSplitPayment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsSplitPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSplitPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSplitPayment represents a SplitPayment event raised by the Bindings contract.
type BindingsSplitPayment struct {
	PaymentId  *big.Int
	SplittedId *big.Int
	Debtor     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSplitPayment is a free log retrieval operation binding the contract event 0x87655463e286e812fa68942951339fd32154c9327434e955699ab18898d588d5.
//
// Solidity: event SplitPayment(uint256 indexed _paymentId, uint256 indexed _splittedId, address _debtor)
func (_Bindings *BindingsFilterer) FilterSplitPayment(opts *bind.FilterOpts, _paymentId []*big.Int, _splittedId []*big.Int) (*BindingsSplitPaymentIterator, error) {

	var _paymentIdRule []interface{}
	for _, _paymentIdItem := range _paymentId {
		_paymentIdRule = append(_paymentIdRule, _paymentIdItem)
	}
	var _splittedIdRule []interface{}
	for _, _splittedIdItem := range _splittedId {
		_splittedIdRule = append(_splittedIdRule, _splittedIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SplitPayment", _paymentIdRule, _splittedIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsSplitPaymentIterator{contract: _Bindings.contract, event: "SplitPayment", logs: logs, sub: sub}, nil
}

// WatchSplitPayment is a free log subscription operation binding the contract event 0x87655463e286e812fa68942951339fd32154c9327434e955699ab18898d588d5.
//
// Solidity: event SplitPayment(uint256 indexed _paymentId, uint256 indexed _splittedId, address _debtor)
func (_Bindings *BindingsFilterer) WatchSplitPayment(opts *bind.WatchOpts, sink chan<- *BindingsSplitPayment, _paymentId []*big.Int, _splittedId []*big.Int) (event.Subscription, error) {

	var _paymentIdRule []interface{}
	for _, _paymentIdItem := range _paymentId {
		_paymentIdRule = append(_paymentIdRule, _paymentIdItem)
	}
	var _splittedIdRule []interface{}
	for _, _splittedIdItem := range _splittedId {
		_splittedIdRule = append(_splittedIdRule, _splittedIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SplitPayment", _paymentIdRule, _splittedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSplitPayment)
				if err := _Bindings.contract.UnpackLog(event, "SplitPayment", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSplitPayment is a log parse operation binding the contract event 0x87655463e286e812fa68942951339fd32154c9327434e955699ab18898d588d5.
//
// Solidity: event SplitPayment(uint256 indexed _paymentId, uint256 indexed _splittedId, address _debtor)
func (_Bindings *BindingsFilterer) ParseSplitPayment(log types.Log) (*BindingsSplitPayment, error) {
	event := new(BindingsSplitPayment)
	if err := _Bindings.contract.UnpackLog(event, "SplitPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bindings contract.
type BindingsTransferIterator struct {
	Event *BindingsTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTransfer represents a Transfer event raised by the Bindings contract.
type BindingsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BindingsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsTransferIterator{contract: _Bindings.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BindingsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTransfer)
				if err := _Bindings.contract.UnpackLog(event, "Transfer", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) ParseTransfer(log types.Log) (*BindingsTransfer, error) {
	event := new(BindingsTransfer)
	if err := _Bindings.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
