// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trade

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// TradeABI is the input ABI used to generate the binding from.
const TradeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"vss\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"}],\"name\":\"isUserHoster\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"},{\"name\":\"_shard\",\"type\":\"string\"}],\"name\":\"uploadBuyerShardFromSeller\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"},{\"name\":\"_hoster\",\"type\":\"address\"}],\"name\":\"getSellerDecryptedShard\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"},{\"name\":\"_hosterID\",\"type\":\"address\"}],\"name\":\"getSellerShardByHosterID\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"},{\"name\":\"_shard\",\"type\":\"string\"}],\"name\":\"uploadSellerShardFromBuyer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"}],\"name\":\"getBuyer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"},{\"name\":\"_hoster\",\"type\":\"address\"}],\"name\":\"getBuyerDecryptedShard\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"},{\"name\":\"_hosterID\",\"type\":\"address\"}],\"name\":\"getBuyerShardByHosterID\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"},{\"name\":\"_decryptedShard\",\"type\":\"string\"}],\"name\":\"uploadSellerDecryptedShard\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"},{\"name\":\"_hosterNum\",\"type\":\"uint256\"}],\"name\":\"createNewTradeOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"}],\"name\":\"getShardHosters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxHosterNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"}],\"name\":\"getSeller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxHosterNum\",\"type\":\"uint256\"}],\"name\":\"updateMaxHosterNum\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"uint256\"},{\"name\":\"_decryptedShard\",\"type\":\"string\"}],\"name\":\"uploadBuyerDecryptedShard\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hoster\",\"type\":\"address\"}],\"name\":\"updateHosterContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hosterContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_feeManager\",\"type\":\"address\"},{\"name\":\"_vss\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"hosters\",\"type\":\"address[]\"}],\"name\":\"CreateOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"userType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"shards\",\"type\":\"string\"}],\"name\":\"UploadEncryptedShard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"privKey\",\"type\":\"string\"}],\"name\":\"RecoverPrivKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"hoster\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"userType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"decryptedShard\",\"type\":\"string\"}],\"name\":\"UploadDecryptedShards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Trade is an auto generated Go binding around an Ethereum contract.
type Trade struct {
	TradeCaller     // Read-only binding to the contract
	TradeTransactor // Write-only binding to the contract
	TradeFilterer   // Log filterer for contract events
}

// TradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradeSession struct {
	Contract     *Trade            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradeCallerSession struct {
	Contract *TradeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradeTransactorSession struct {
	Contract     *TradeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradeRaw struct {
	Contract *Trade // Generic contract binding to access the raw methods on
}

// TradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradeCallerRaw struct {
	Contract *TradeCaller // Generic read-only contract binding to access the raw methods on
}

// TradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradeTransactorRaw struct {
	Contract *TradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrade creates a new instance of Trade, bound to a specific deployed contract.
func NewTrade(address common.Address, backend bind.ContractBackend) (*Trade, error) {
	contract, err := bindTrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trade{TradeCaller: TradeCaller{contract: contract}, TradeTransactor: TradeTransactor{contract: contract}, TradeFilterer: TradeFilterer{contract: contract}}, nil
}

// NewTradeCaller creates a new read-only instance of Trade, bound to a specific deployed contract.
func NewTradeCaller(address common.Address, caller bind.ContractCaller) (*TradeCaller, error) {
	contract, err := bindTrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradeCaller{contract: contract}, nil
}

// NewTradeTransactor creates a new write-only instance of Trade, bound to a specific deployed contract.
func NewTradeTransactor(address common.Address, transactor bind.ContractTransactor) (*TradeTransactor, error) {
	contract, err := bindTrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradeTransactor{contract: contract}, nil
}

// NewTradeFilterer creates a new log filterer instance of Trade, bound to a specific deployed contract.
func NewTradeFilterer(address common.Address, filterer bind.ContractFilterer) (*TradeFilterer, error) {
	contract, err := bindTrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradeFilterer{contract: contract}, nil
}

// bindTrade binds a generic wrapper to an already deployed contract.
func bindTrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.TradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transact(opts, method, params...)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() constant returns(address)
func (_Trade *TradeCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "feeManager")
	return *ret0, err
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() constant returns(address)
func (_Trade *TradeSession) FeeManager() (common.Address, error) {
	return _Trade.Contract.FeeManager(&_Trade.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() constant returns(address)
func (_Trade *TradeCallerSession) FeeManager() (common.Address, error) {
	return _Trade.Contract.FeeManager(&_Trade.CallOpts)
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(_orderID uint256) constant returns(address)
func (_Trade *TradeCaller) GetBuyer(opts *bind.CallOpts, _orderID *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "getBuyer", _orderID)
	return *ret0, err
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(_orderID uint256) constant returns(address)
func (_Trade *TradeSession) GetBuyer(_orderID *big.Int) (common.Address, error) {
	return _Trade.Contract.GetBuyer(&_Trade.CallOpts, _orderID)
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(_orderID uint256) constant returns(address)
func (_Trade *TradeCallerSession) GetBuyer(_orderID *big.Int) (common.Address, error) {
	return _Trade.Contract.GetBuyer(&_Trade.CallOpts, _orderID)
}

// GetBuyerDecryptedShard is a free data retrieval call binding the contract method 0x6da647e4.
//
// Solidity: function getBuyerDecryptedShard(_orderID uint256, _hoster address) constant returns(string)
func (_Trade *TradeCaller) GetBuyerDecryptedShard(opts *bind.CallOpts, _orderID *big.Int, _hoster common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "getBuyerDecryptedShard", _orderID, _hoster)
	return *ret0, err
}

// GetBuyerDecryptedShard is a free data retrieval call binding the contract method 0x6da647e4.
//
// Solidity: function getBuyerDecryptedShard(_orderID uint256, _hoster address) constant returns(string)
func (_Trade *TradeSession) GetBuyerDecryptedShard(_orderID *big.Int, _hoster common.Address) (string, error) {
	return _Trade.Contract.GetBuyerDecryptedShard(&_Trade.CallOpts, _orderID, _hoster)
}

// GetBuyerDecryptedShard is a free data retrieval call binding the contract method 0x6da647e4.
//
// Solidity: function getBuyerDecryptedShard(_orderID uint256, _hoster address) constant returns(string)
func (_Trade *TradeCallerSession) GetBuyerDecryptedShard(_orderID *big.Int, _hoster common.Address) (string, error) {
	return _Trade.Contract.GetBuyerDecryptedShard(&_Trade.CallOpts, _orderID, _hoster)
}

// GetBuyerShardByHosterID is a free data retrieval call binding the contract method 0x79f25afa.
//
// Solidity: function getBuyerShardByHosterID(_orderID uint256, _hosterID address) constant returns(string)
func (_Trade *TradeCaller) GetBuyerShardByHosterID(opts *bind.CallOpts, _orderID *big.Int, _hosterID common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "getBuyerShardByHosterID", _orderID, _hosterID)
	return *ret0, err
}

// GetBuyerShardByHosterID is a free data retrieval call binding the contract method 0x79f25afa.
//
// Solidity: function getBuyerShardByHosterID(_orderID uint256, _hosterID address) constant returns(string)
func (_Trade *TradeSession) GetBuyerShardByHosterID(_orderID *big.Int, _hosterID common.Address) (string, error) {
	return _Trade.Contract.GetBuyerShardByHosterID(&_Trade.CallOpts, _orderID, _hosterID)
}

// GetBuyerShardByHosterID is a free data retrieval call binding the contract method 0x79f25afa.
//
// Solidity: function getBuyerShardByHosterID(_orderID uint256, _hosterID address) constant returns(string)
func (_Trade *TradeCallerSession) GetBuyerShardByHosterID(_orderID *big.Int, _hosterID common.Address) (string, error) {
	return _Trade.Contract.GetBuyerShardByHosterID(&_Trade.CallOpts, _orderID, _hosterID)
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(_orderID uint256) constant returns(address)
func (_Trade *TradeCaller) GetSeller(opts *bind.CallOpts, _orderID *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "getSeller", _orderID)
	return *ret0, err
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(_orderID uint256) constant returns(address)
func (_Trade *TradeSession) GetSeller(_orderID *big.Int) (common.Address, error) {
	return _Trade.Contract.GetSeller(&_Trade.CallOpts, _orderID)
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(_orderID uint256) constant returns(address)
func (_Trade *TradeCallerSession) GetSeller(_orderID *big.Int) (common.Address, error) {
	return _Trade.Contract.GetSeller(&_Trade.CallOpts, _orderID)
}

// GetSellerDecryptedShard is a free data retrieval call binding the contract method 0x34a733ec.
//
// Solidity: function getSellerDecryptedShard(_orderID uint256, _hoster address) constant returns(string)
func (_Trade *TradeCaller) GetSellerDecryptedShard(opts *bind.CallOpts, _orderID *big.Int, _hoster common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "getSellerDecryptedShard", _orderID, _hoster)
	return *ret0, err
}

// GetSellerDecryptedShard is a free data retrieval call binding the contract method 0x34a733ec.
//
// Solidity: function getSellerDecryptedShard(_orderID uint256, _hoster address) constant returns(string)
func (_Trade *TradeSession) GetSellerDecryptedShard(_orderID *big.Int, _hoster common.Address) (string, error) {
	return _Trade.Contract.GetSellerDecryptedShard(&_Trade.CallOpts, _orderID, _hoster)
}

// GetSellerDecryptedShard is a free data retrieval call binding the contract method 0x34a733ec.
//
// Solidity: function getSellerDecryptedShard(_orderID uint256, _hoster address) constant returns(string)
func (_Trade *TradeCallerSession) GetSellerDecryptedShard(_orderID *big.Int, _hoster common.Address) (string, error) {
	return _Trade.Contract.GetSellerDecryptedShard(&_Trade.CallOpts, _orderID, _hoster)
}

// GetSellerShardByHosterID is a free data retrieval call binding the contract method 0x3c8047a1.
//
// Solidity: function getSellerShardByHosterID(_orderID uint256, _hosterID address) constant returns(string)
func (_Trade *TradeCaller) GetSellerShardByHosterID(opts *bind.CallOpts, _orderID *big.Int, _hosterID common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "getSellerShardByHosterID", _orderID, _hosterID)
	return *ret0, err
}

// GetSellerShardByHosterID is a free data retrieval call binding the contract method 0x3c8047a1.
//
// Solidity: function getSellerShardByHosterID(_orderID uint256, _hosterID address) constant returns(string)
func (_Trade *TradeSession) GetSellerShardByHosterID(_orderID *big.Int, _hosterID common.Address) (string, error) {
	return _Trade.Contract.GetSellerShardByHosterID(&_Trade.CallOpts, _orderID, _hosterID)
}

// GetSellerShardByHosterID is a free data retrieval call binding the contract method 0x3c8047a1.
//
// Solidity: function getSellerShardByHosterID(_orderID uint256, _hosterID address) constant returns(string)
func (_Trade *TradeCallerSession) GetSellerShardByHosterID(_orderID *big.Int, _hosterID common.Address) (string, error) {
	return _Trade.Contract.GetSellerShardByHosterID(&_Trade.CallOpts, _orderID, _hosterID)
}

// GetShardHosters is a free data retrieval call binding the contract method 0xc6af3400.
//
// Solidity: function getShardHosters(_orderID uint256) constant returns(address[])
func (_Trade *TradeCaller) GetShardHosters(opts *bind.CallOpts, _orderID *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "getShardHosters", _orderID)
	return *ret0, err
}

// GetShardHosters is a free data retrieval call binding the contract method 0xc6af3400.
//
// Solidity: function getShardHosters(_orderID uint256) constant returns(address[])
func (_Trade *TradeSession) GetShardHosters(_orderID *big.Int) ([]common.Address, error) {
	return _Trade.Contract.GetShardHosters(&_Trade.CallOpts, _orderID)
}

// GetShardHosters is a free data retrieval call binding the contract method 0xc6af3400.
//
// Solidity: function getShardHosters(_orderID uint256) constant returns(address[])
func (_Trade *TradeCallerSession) GetShardHosters(_orderID *big.Int) ([]common.Address, error) {
	return _Trade.Contract.GetShardHosters(&_Trade.CallOpts, _orderID)
}

// HosterContract is a free data retrieval call binding the contract method 0xf6f66483.
//
// Solidity: function hosterContract() constant returns(address)
func (_Trade *TradeCaller) HosterContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "hosterContract")
	return *ret0, err
}

// HosterContract is a free data retrieval call binding the contract method 0xf6f66483.
//
// Solidity: function hosterContract() constant returns(address)
func (_Trade *TradeSession) HosterContract() (common.Address, error) {
	return _Trade.Contract.HosterContract(&_Trade.CallOpts)
}

// HosterContract is a free data retrieval call binding the contract method 0xf6f66483.
//
// Solidity: function hosterContract() constant returns(address)
func (_Trade *TradeCallerSession) HosterContract() (common.Address, error) {
	return _Trade.Contract.HosterContract(&_Trade.CallOpts)
}

// IsUserHoster is a free data retrieval call binding the contract method 0x18b28435.
//
// Solidity: function isUserHoster(_orderID uint256) constant returns(bool)
func (_Trade *TradeCaller) IsUserHoster(opts *bind.CallOpts, _orderID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "isUserHoster", _orderID)
	return *ret0, err
}

// IsUserHoster is a free data retrieval call binding the contract method 0x18b28435.
//
// Solidity: function isUserHoster(_orderID uint256) constant returns(bool)
func (_Trade *TradeSession) IsUserHoster(_orderID *big.Int) (bool, error) {
	return _Trade.Contract.IsUserHoster(&_Trade.CallOpts, _orderID)
}

// IsUserHoster is a free data retrieval call binding the contract method 0x18b28435.
//
// Solidity: function isUserHoster(_orderID uint256) constant returns(bool)
func (_Trade *TradeCallerSession) IsUserHoster(_orderID *big.Int) (bool, error) {
	return _Trade.Contract.IsUserHoster(&_Trade.CallOpts, _orderID)
}

// MaxHosterNum is a free data retrieval call binding the contract method 0xd2cfe671.
//
// Solidity: function maxHosterNum() constant returns(uint256)
func (_Trade *TradeCaller) MaxHosterNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "maxHosterNum")
	return *ret0, err
}

// MaxHosterNum is a free data retrieval call binding the contract method 0xd2cfe671.
//
// Solidity: function maxHosterNum() constant returns(uint256)
func (_Trade *TradeSession) MaxHosterNum() (*big.Int, error) {
	return _Trade.Contract.MaxHosterNum(&_Trade.CallOpts)
}

// MaxHosterNum is a free data retrieval call binding the contract method 0xd2cfe671.
//
// Solidity: function maxHosterNum() constant returns(uint256)
func (_Trade *TradeCallerSession) MaxHosterNum() (*big.Int, error) {
	return _Trade.Contract.MaxHosterNum(&_Trade.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Trade *TradeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Trade *TradeSession) Owner() (common.Address, error) {
	return _Trade.Contract.Owner(&_Trade.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Trade *TradeCallerSession) Owner() (common.Address, error) {
	return _Trade.Contract.Owner(&_Trade.CallOpts)
}

// Vss is a free data retrieval call binding the contract method 0x0a157e02.
//
// Solidity: function vss() constant returns(address)
func (_Trade *TradeCaller) Vss(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Trade.contract.Call(opts, out, "vss")
	return *ret0, err
}

// Vss is a free data retrieval call binding the contract method 0x0a157e02.
//
// Solidity: function vss() constant returns(address)
func (_Trade *TradeSession) Vss() (common.Address, error) {
	return _Trade.Contract.Vss(&_Trade.CallOpts)
}

// Vss is a free data retrieval call binding the contract method 0x0a157e02.
//
// Solidity: function vss() constant returns(address)
func (_Trade *TradeCallerSession) Vss() (common.Address, error) {
	return _Trade.Contract.Vss(&_Trade.CallOpts)
}

// CreateNewTradeOrder is a paid mutator transaction binding the contract method 0xbc5cdb29.
//
// Solidity: function createNewTradeOrder(_orderID uint256, _seller address, _hosterNum uint256) returns(bool)
func (_Trade *TradeTransactor) CreateNewTradeOrder(opts *bind.TransactOpts, _orderID *big.Int, _seller common.Address, _hosterNum *big.Int) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "createNewTradeOrder", _orderID, _seller, _hosterNum)
}

// CreateNewTradeOrder is a paid mutator transaction binding the contract method 0xbc5cdb29.
//
// Solidity: function createNewTradeOrder(_orderID uint256, _seller address, _hosterNum uint256) returns(bool)
func (_Trade *TradeSession) CreateNewTradeOrder(_orderID *big.Int, _seller common.Address, _hosterNum *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.CreateNewTradeOrder(&_Trade.TransactOpts, _orderID, _seller, _hosterNum)
}

// CreateNewTradeOrder is a paid mutator transaction binding the contract method 0xbc5cdb29.
//
// Solidity: function createNewTradeOrder(_orderID uint256, _seller address, _hosterNum uint256) returns(bool)
func (_Trade *TradeTransactorSession) CreateNewTradeOrder(_orderID *big.Int, _seller common.Address, _hosterNum *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.CreateNewTradeOrder(&_Trade.TransactOpts, _orderID, _seller, _hosterNum)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trade *TradeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trade *TradeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trade.Contract.RenounceOwnership(&_Trade.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trade *TradeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trade.Contract.RenounceOwnership(&_Trade.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Trade *TradeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Trade *TradeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trade.Contract.TransferOwnership(&_Trade.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Trade *TradeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trade.Contract.TransferOwnership(&_Trade.TransactOpts, newOwner)
}

// UpdateHosterContract is a paid mutator transaction binding the contract method 0xf52706a0.
//
// Solidity: function updateHosterContract(_hoster address) returns(bool)
func (_Trade *TradeTransactor) UpdateHosterContract(opts *bind.TransactOpts, _hoster common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "updateHosterContract", _hoster)
}

// UpdateHosterContract is a paid mutator transaction binding the contract method 0xf52706a0.
//
// Solidity: function updateHosterContract(_hoster address) returns(bool)
func (_Trade *TradeSession) UpdateHosterContract(_hoster common.Address) (*types.Transaction, error) {
	return _Trade.Contract.UpdateHosterContract(&_Trade.TransactOpts, _hoster)
}

// UpdateHosterContract is a paid mutator transaction binding the contract method 0xf52706a0.
//
// Solidity: function updateHosterContract(_hoster address) returns(bool)
func (_Trade *TradeTransactorSession) UpdateHosterContract(_hoster common.Address) (*types.Transaction, error) {
	return _Trade.Contract.UpdateHosterContract(&_Trade.TransactOpts, _hoster)
}

// UpdateMaxHosterNum is a paid mutator transaction binding the contract method 0xdaa674b1.
//
// Solidity: function updateMaxHosterNum(_maxHosterNum uint256) returns(bool)
func (_Trade *TradeTransactor) UpdateMaxHosterNum(opts *bind.TransactOpts, _maxHosterNum *big.Int) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "updateMaxHosterNum", _maxHosterNum)
}

// UpdateMaxHosterNum is a paid mutator transaction binding the contract method 0xdaa674b1.
//
// Solidity: function updateMaxHosterNum(_maxHosterNum uint256) returns(bool)
func (_Trade *TradeSession) UpdateMaxHosterNum(_maxHosterNum *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.UpdateMaxHosterNum(&_Trade.TransactOpts, _maxHosterNum)
}

// UpdateMaxHosterNum is a paid mutator transaction binding the contract method 0xdaa674b1.
//
// Solidity: function updateMaxHosterNum(_maxHosterNum uint256) returns(bool)
func (_Trade *TradeTransactorSession) UpdateMaxHosterNum(_maxHosterNum *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.UpdateMaxHosterNum(&_Trade.TransactOpts, _maxHosterNum)
}

// UploadBuyerDecryptedShard is a paid mutator transaction binding the contract method 0xe370e0ac.
//
// Solidity: function uploadBuyerDecryptedShard(_orderID uint256, _decryptedShard string) returns(bool)
func (_Trade *TradeTransactor) UploadBuyerDecryptedShard(opts *bind.TransactOpts, _orderID *big.Int, _decryptedShard string) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "uploadBuyerDecryptedShard", _orderID, _decryptedShard)
}

// UploadBuyerDecryptedShard is a paid mutator transaction binding the contract method 0xe370e0ac.
//
// Solidity: function uploadBuyerDecryptedShard(_orderID uint256, _decryptedShard string) returns(bool)
func (_Trade *TradeSession) UploadBuyerDecryptedShard(_orderID *big.Int, _decryptedShard string) (*types.Transaction, error) {
	return _Trade.Contract.UploadBuyerDecryptedShard(&_Trade.TransactOpts, _orderID, _decryptedShard)
}

// UploadBuyerDecryptedShard is a paid mutator transaction binding the contract method 0xe370e0ac.
//
// Solidity: function uploadBuyerDecryptedShard(_orderID uint256, _decryptedShard string) returns(bool)
func (_Trade *TradeTransactorSession) UploadBuyerDecryptedShard(_orderID *big.Int, _decryptedShard string) (*types.Transaction, error) {
	return _Trade.Contract.UploadBuyerDecryptedShard(&_Trade.TransactOpts, _orderID, _decryptedShard)
}

// UploadBuyerShardFromSeller is a paid mutator transaction binding the contract method 0x1c17b9b2.
//
// Solidity: function uploadBuyerShardFromSeller(_orderID uint256, _shard string) returns(bool)
func (_Trade *TradeTransactor) UploadBuyerShardFromSeller(opts *bind.TransactOpts, _orderID *big.Int, _shard string) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "uploadBuyerShardFromSeller", _orderID, _shard)
}

// UploadBuyerShardFromSeller is a paid mutator transaction binding the contract method 0x1c17b9b2.
//
// Solidity: function uploadBuyerShardFromSeller(_orderID uint256, _shard string) returns(bool)
func (_Trade *TradeSession) UploadBuyerShardFromSeller(_orderID *big.Int, _shard string) (*types.Transaction, error) {
	return _Trade.Contract.UploadBuyerShardFromSeller(&_Trade.TransactOpts, _orderID, _shard)
}

// UploadBuyerShardFromSeller is a paid mutator transaction binding the contract method 0x1c17b9b2.
//
// Solidity: function uploadBuyerShardFromSeller(_orderID uint256, _shard string) returns(bool)
func (_Trade *TradeTransactorSession) UploadBuyerShardFromSeller(_orderID *big.Int, _shard string) (*types.Transaction, error) {
	return _Trade.Contract.UploadBuyerShardFromSeller(&_Trade.TransactOpts, _orderID, _shard)
}

// UploadSellerDecryptedShard is a paid mutator transaction binding the contract method 0xadd2f133.
//
// Solidity: function uploadSellerDecryptedShard(_orderID uint256, _decryptedShard string) returns(bool)
func (_Trade *TradeTransactor) UploadSellerDecryptedShard(opts *bind.TransactOpts, _orderID *big.Int, _decryptedShard string) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "uploadSellerDecryptedShard", _orderID, _decryptedShard)
}

// UploadSellerDecryptedShard is a paid mutator transaction binding the contract method 0xadd2f133.
//
// Solidity: function uploadSellerDecryptedShard(_orderID uint256, _decryptedShard string) returns(bool)
func (_Trade *TradeSession) UploadSellerDecryptedShard(_orderID *big.Int, _decryptedShard string) (*types.Transaction, error) {
	return _Trade.Contract.UploadSellerDecryptedShard(&_Trade.TransactOpts, _orderID, _decryptedShard)
}

// UploadSellerDecryptedShard is a paid mutator transaction binding the contract method 0xadd2f133.
//
// Solidity: function uploadSellerDecryptedShard(_orderID uint256, _decryptedShard string) returns(bool)
func (_Trade *TradeTransactorSession) UploadSellerDecryptedShard(_orderID *big.Int, _decryptedShard string) (*types.Transaction, error) {
	return _Trade.Contract.UploadSellerDecryptedShard(&_Trade.TransactOpts, _orderID, _decryptedShard)
}

// UploadSellerShardFromBuyer is a paid mutator transaction binding the contract method 0x3e4b1302.
//
// Solidity: function uploadSellerShardFromBuyer(_orderID uint256, _shard string) returns(bool)
func (_Trade *TradeTransactor) UploadSellerShardFromBuyer(opts *bind.TransactOpts, _orderID *big.Int, _shard string) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "uploadSellerShardFromBuyer", _orderID, _shard)
}

// UploadSellerShardFromBuyer is a paid mutator transaction binding the contract method 0x3e4b1302.
//
// Solidity: function uploadSellerShardFromBuyer(_orderID uint256, _shard string) returns(bool)
func (_Trade *TradeSession) UploadSellerShardFromBuyer(_orderID *big.Int, _shard string) (*types.Transaction, error) {
	return _Trade.Contract.UploadSellerShardFromBuyer(&_Trade.TransactOpts, _orderID, _shard)
}

// UploadSellerShardFromBuyer is a paid mutator transaction binding the contract method 0x3e4b1302.
//
// Solidity: function uploadSellerShardFromBuyer(_orderID uint256, _shard string) returns(bool)
func (_Trade *TradeTransactorSession) UploadSellerShardFromBuyer(_orderID *big.Int, _shard string) (*types.Transaction, error) {
	return _Trade.Contract.UploadSellerShardFromBuyer(&_Trade.TransactOpts, _orderID, _shard)
}

// TradeCreateOrderIterator is returned from FilterCreateOrder and is used to iterate over the raw logs and unpacked data for CreateOrder events raised by the Trade contract.
type TradeCreateOrderIterator struct {
	Event *TradeCreateOrder // Event containing the contract specifics and raw log

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
func (it *TradeCreateOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeCreateOrder)
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
		it.Event = new(TradeCreateOrder)
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
func (it *TradeCreateOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeCreateOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeCreateOrder represents a CreateOrder event raised by the Trade contract.
type TradeCreateOrder struct {
	OrderID *big.Int
	Buyer   common.Address
	Seller  common.Address
	Hosters []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateOrder is a free log retrieval operation binding the contract event 0xcfc45bb8f7b7e40d083374fa1c4b4848e9df33830a2a6b9301493647bb9d1f24.
//
// Solidity: e CreateOrder(orderID uint256, buyer indexed address, seller indexed address, hosters address[])
func (_Trade *TradeFilterer) FilterCreateOrder(opts *bind.FilterOpts, buyer []common.Address, seller []common.Address) (*TradeCreateOrderIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Trade.contract.FilterLogs(opts, "CreateOrder", buyerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &TradeCreateOrderIterator{contract: _Trade.contract, event: "CreateOrder", logs: logs, sub: sub}, nil
}

// WatchCreateOrder is a free log subscription operation binding the contract event 0xcfc45bb8f7b7e40d083374fa1c4b4848e9df33830a2a6b9301493647bb9d1f24.
//
// Solidity: e CreateOrder(orderID uint256, buyer indexed address, seller indexed address, hosters address[])
func (_Trade *TradeFilterer) WatchCreateOrder(opts *bind.WatchOpts, sink chan<- *TradeCreateOrder, buyer []common.Address, seller []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Trade.contract.WatchLogs(opts, "CreateOrder", buyerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeCreateOrder)
				if err := _Trade.contract.UnpackLog(event, "CreateOrder", log); err != nil {
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

// TradeOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Trade contract.
type TradeOwnershipRenouncedIterator struct {
	Event *TradeOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *TradeOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeOwnershipRenounced)
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
		it.Event = new(TradeOwnershipRenounced)
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
func (it *TradeOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeOwnershipRenounced represents a OwnershipRenounced event raised by the Trade contract.
type TradeOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Trade *TradeFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*TradeOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Trade.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TradeOwnershipRenouncedIterator{contract: _Trade.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Trade *TradeFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *TradeOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Trade.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeOwnershipRenounced)
				if err := _Trade.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// TradeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Trade contract.
type TradeOwnershipTransferredIterator struct {
	Event *TradeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TradeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeOwnershipTransferred)
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
		it.Event = new(TradeOwnershipTransferred)
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
func (it *TradeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeOwnershipTransferred represents a OwnershipTransferred event raised by the Trade contract.
type TradeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Trade *TradeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TradeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trade.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TradeOwnershipTransferredIterator{contract: _Trade.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Trade *TradeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TradeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trade.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeOwnershipTransferred)
				if err := _Trade.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TradeRecoverPrivKeyIterator is returned from FilterRecoverPrivKey and is used to iterate over the raw logs and unpacked data for RecoverPrivKey events raised by the Trade contract.
type TradeRecoverPrivKeyIterator struct {
	Event *TradeRecoverPrivKey // Event containing the contract specifics and raw log

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
func (it *TradeRecoverPrivKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeRecoverPrivKey)
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
		it.Event = new(TradeRecoverPrivKey)
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
func (it *TradeRecoverPrivKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeRecoverPrivKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeRecoverPrivKey represents a RecoverPrivKey event raised by the Trade contract.
type TradeRecoverPrivKey struct {
	OrderID  *big.Int
	UserType uint8
	PrivKey  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRecoverPrivKey is a free log retrieval operation binding the contract event 0xd3efd4cb85578995aa8efa07b7f7a9f59b4e5c71f230b7839759d72793e23b41.
//
// Solidity: e RecoverPrivKey(orderID uint256, userType uint8, privKey string)
func (_Trade *TradeFilterer) FilterRecoverPrivKey(opts *bind.FilterOpts) (*TradeRecoverPrivKeyIterator, error) {

	logs, sub, err := _Trade.contract.FilterLogs(opts, "RecoverPrivKey")
	if err != nil {
		return nil, err
	}
	return &TradeRecoverPrivKeyIterator{contract: _Trade.contract, event: "RecoverPrivKey", logs: logs, sub: sub}, nil
}

// WatchRecoverPrivKey is a free log subscription operation binding the contract event 0xd3efd4cb85578995aa8efa07b7f7a9f59b4e5c71f230b7839759d72793e23b41.
//
// Solidity: e RecoverPrivKey(orderID uint256, userType uint8, privKey string)
func (_Trade *TradeFilterer) WatchRecoverPrivKey(opts *bind.WatchOpts, sink chan<- *TradeRecoverPrivKey) (event.Subscription, error) {

	logs, sub, err := _Trade.contract.WatchLogs(opts, "RecoverPrivKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeRecoverPrivKey)
				if err := _Trade.contract.UnpackLog(event, "RecoverPrivKey", log); err != nil {
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

// TradeUploadDecryptedShardsIterator is returned from FilterUploadDecryptedShards and is used to iterate over the raw logs and unpacked data for UploadDecryptedShards events raised by the Trade contract.
type TradeUploadDecryptedShardsIterator struct {
	Event *TradeUploadDecryptedShards // Event containing the contract specifics and raw log

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
func (it *TradeUploadDecryptedShardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeUploadDecryptedShards)
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
		it.Event = new(TradeUploadDecryptedShards)
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
func (it *TradeUploadDecryptedShardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeUploadDecryptedShardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeUploadDecryptedShards represents a UploadDecryptedShards event raised by the Trade contract.
type TradeUploadDecryptedShards struct {
	OrderID        *big.Int
	Hoster         common.Address
	UserType       uint8
	DecryptedShard string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUploadDecryptedShards is a free log retrieval operation binding the contract event 0x24538d8d2b02982080c2a529355335e2ed33e94f5f6f2e9a6be5e54a64af2d8d.
//
// Solidity: e UploadDecryptedShards(orderID uint256, hoster indexed address, userType uint8, decryptedShard string)
func (_Trade *TradeFilterer) FilterUploadDecryptedShards(opts *bind.FilterOpts, hoster []common.Address) (*TradeUploadDecryptedShardsIterator, error) {

	var hosterRule []interface{}
	for _, hosterItem := range hoster {
		hosterRule = append(hosterRule, hosterItem)
	}

	logs, sub, err := _Trade.contract.FilterLogs(opts, "UploadDecryptedShards", hosterRule)
	if err != nil {
		return nil, err
	}
	return &TradeUploadDecryptedShardsIterator{contract: _Trade.contract, event: "UploadDecryptedShards", logs: logs, sub: sub}, nil
}

// WatchUploadDecryptedShards is a free log subscription operation binding the contract event 0x24538d8d2b02982080c2a529355335e2ed33e94f5f6f2e9a6be5e54a64af2d8d.
//
// Solidity: e UploadDecryptedShards(orderID uint256, hoster indexed address, userType uint8, decryptedShard string)
func (_Trade *TradeFilterer) WatchUploadDecryptedShards(opts *bind.WatchOpts, sink chan<- *TradeUploadDecryptedShards, hoster []common.Address) (event.Subscription, error) {

	var hosterRule []interface{}
	for _, hosterItem := range hoster {
		hosterRule = append(hosterRule, hosterItem)
	}

	logs, sub, err := _Trade.contract.WatchLogs(opts, "UploadDecryptedShards", hosterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeUploadDecryptedShards)
				if err := _Trade.contract.UnpackLog(event, "UploadDecryptedShards", log); err != nil {
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

// TradeUploadEncryptedShardIterator is returned from FilterUploadEncryptedShard and is used to iterate over the raw logs and unpacked data for UploadEncryptedShard events raised by the Trade contract.
type TradeUploadEncryptedShardIterator struct {
	Event *TradeUploadEncryptedShard // Event containing the contract specifics and raw log

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
func (it *TradeUploadEncryptedShardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeUploadEncryptedShard)
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
		it.Event = new(TradeUploadEncryptedShard)
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
func (it *TradeUploadEncryptedShardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeUploadEncryptedShardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeUploadEncryptedShard represents a UploadEncryptedShard event raised by the Trade contract.
type TradeUploadEncryptedShard struct {
	OrderID  *big.Int
	Who      common.Address
	UserType uint8
	Shards   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUploadEncryptedShard is a free log retrieval operation binding the contract event 0x4491fee63e1b2614fdd79a3cb9ab4eaf6cf253b366ee5df2f03adabbcc54c85c.
//
// Solidity: e UploadEncryptedShard(orderID uint256, who indexed address, userType uint8, shards string)
func (_Trade *TradeFilterer) FilterUploadEncryptedShard(opts *bind.FilterOpts, who []common.Address) (*TradeUploadEncryptedShardIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Trade.contract.FilterLogs(opts, "UploadEncryptedShard", whoRule)
	if err != nil {
		return nil, err
	}
	return &TradeUploadEncryptedShardIterator{contract: _Trade.contract, event: "UploadEncryptedShard", logs: logs, sub: sub}, nil
}

// WatchUploadEncryptedShard is a free log subscription operation binding the contract event 0x4491fee63e1b2614fdd79a3cb9ab4eaf6cf253b366ee5df2f03adabbcc54c85c.
//
// Solidity: e UploadEncryptedShard(orderID uint256, who indexed address, userType uint8, shards string)
func (_Trade *TradeFilterer) WatchUploadEncryptedShard(opts *bind.WatchOpts, sink chan<- *TradeUploadEncryptedShard, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Trade.contract.WatchLogs(opts, "UploadEncryptedShard", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeUploadEncryptedShard)
				if err := _Trade.contract.UnpackLog(event, "UploadEncryptedShard", log); err != nil {
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
