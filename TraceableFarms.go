// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TraceableFarms

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
)

// TraceableFarmsAccreditation is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsAccreditation struct {
	Id                string
	Name              string
	AccreditationType TraceableFarmsAccreditationType
}

// TraceableFarmsAccreditationCompany is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsAccreditationCompany struct {
	Accreditation TraceableFarmsAccreditation
	Checker       string
}

// TraceableFarmsAccreditationType is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsAccreditationType struct {
	Id   string
	Name string
}

// TraceableFarmsBatch is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsBatch struct {
	Id                 string
	Number             string
	Date               string
	ProductName        string
	ProductVariety     string
	ProductDescription string
	ProductPhotoUrl    string
	Company            TraceableFarmsCompany
}

// TraceableFarmsBatchOrigin is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsBatchOrigin struct {
	Description         string
	Location            string
	LocationCoordinates string
}

// TraceableFarmsBatchProcess is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsBatchProcess struct {
	Name        string
	Description string
}

// TraceableFarmsCompany is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsCompany struct {
	Nif                      string
	BussinessName            string
	Location                 string
	LocationCoordinates      string
	InformationalResourceUrl string
	Consortium               TraceableFarmsConsortium
}

// TraceableFarmsConsortium is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsConsortium struct {
	Id   string
	Name string
}

// TraceableFarmsFootprintType is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsFootprintType struct {
	Id                    string
	Name                  string
	UnitMeasurementName   string
	UnitMeasurementSymbol string
}

// TraceableFarmsMetaData contains all meta data concerning the TraceableFarms contract.
var TraceableFarmsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getAccreditation\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.AccreditationType\",\"name\":\"accreditationType\",\"type\":\"tuple\"}],\"internalType\":\"structTraceableFarms.Accreditation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nif\",\"type\":\"string\"}],\"name\":\"getAccreditationCompany\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.AccreditationType\",\"name\":\"accreditationType\",\"type\":\"tuple\"}],\"internalType\":\"structTraceableFarms.Accreditation\",\"name\":\"accreditation\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"checker\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.AccreditationCompany[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getAccreditationType\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.AccreditationType\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productVariety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productDescription\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productPhotoUrl\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"nif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bussinessName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"locationCoordinates\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"informationalResourceUrl\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.Consortium\",\"name\":\"consortium\",\"type\":\"tuple\"}],\"internalType\":\"structTraceableFarms.Company\",\"name\":\"company\",\"type\":\"tuple\"}],\"internalType\":\"structTraceableFarms.Batch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_batchId\",\"type\":\"string\"}],\"name\":\"getBatchOrigin\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"locationCoordinates\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.BatchOrigin[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_batchId\",\"type\":\"string\"}],\"name\":\"getBatchProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.BatchProcess[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nif\",\"type\":\"string\"}],\"name\":\"getCompany\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bussinessName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"locationCoordinates\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"informationalResourceUrl\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.Consortium\",\"name\":\"consortium\",\"type\":\"tuple\"}],\"internalType\":\"structTraceableFarms.Company\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getConsortium\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.Consortium\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getFootprintType\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unitMeasurementName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unitMeasurementSymbol\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.FootprintType\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_accreditationTypeId\",\"type\":\"string\"}],\"name\":\"setAccreditation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_accreditationId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_checker\",\"type\":\"string\"}],\"name\":\"setAccreditationCompany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setAccreditationType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_number\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_productName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_productVariety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_productDescription\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_productPhotoUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_companyId\",\"type\":\"string\"}],\"name\":\"setBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_batchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_locationCoordinates\",\"type\":\"string\"}],\"name\":\"setBatchOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_batchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"setBatchProcess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_bussinessName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_locationCoordinates\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_informationalResourceUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_consortiumId\",\"type\":\"string\"}],\"name\":\"setCompany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setConsortium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_unitMeasurementName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_unitMeasurementSymbol\",\"type\":\"string\"}],\"name\":\"setFootprintType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TraceableFarmsABI is the input ABI used to generate the binding from.
// Deprecated: Use TraceableFarmsMetaData.ABI instead.
var TraceableFarmsABI = TraceableFarmsMetaData.ABI

// TraceableFarms is an auto generated Go binding around an Ethereum contract.
type TraceableFarms struct {
	TraceableFarmsCaller     // Read-only binding to the contract
	TraceableFarmsTransactor // Write-only binding to the contract
	TraceableFarmsFilterer   // Log filterer for contract events
}

// TraceableFarmsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraceableFarmsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceableFarmsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraceableFarmsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceableFarmsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraceableFarmsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceableFarmsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraceableFarmsSession struct {
	Contract     *TraceableFarms   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceableFarmsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraceableFarmsCallerSession struct {
	Contract *TraceableFarmsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TraceableFarmsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraceableFarmsTransactorSession struct {
	Contract     *TraceableFarmsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TraceableFarmsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraceableFarmsRaw struct {
	Contract *TraceableFarms // Generic contract binding to access the raw methods on
}

// TraceableFarmsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraceableFarmsCallerRaw struct {
	Contract *TraceableFarmsCaller // Generic read-only contract binding to access the raw methods on
}

// TraceableFarmsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraceableFarmsTransactorRaw struct {
	Contract *TraceableFarmsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraceableFarms creates a new instance of TraceableFarms, bound to a specific deployed contract.
func NewTraceableFarms(address common.Address, backend bind.ContractBackend) (*TraceableFarms, error) {
	contract, err := bindTraceableFarms(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraceableFarms{TraceableFarmsCaller: TraceableFarmsCaller{contract: contract}, TraceableFarmsTransactor: TraceableFarmsTransactor{contract: contract}, TraceableFarmsFilterer: TraceableFarmsFilterer{contract: contract}}, nil
}

// NewTraceableFarmsCaller creates a new read-only instance of TraceableFarms, bound to a specific deployed contract.
func NewTraceableFarmsCaller(address common.Address, caller bind.ContractCaller) (*TraceableFarmsCaller, error) {
	contract, err := bindTraceableFarms(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraceableFarmsCaller{contract: contract}, nil
}

// NewTraceableFarmsTransactor creates a new write-only instance of TraceableFarms, bound to a specific deployed contract.
func NewTraceableFarmsTransactor(address common.Address, transactor bind.ContractTransactor) (*TraceableFarmsTransactor, error) {
	contract, err := bindTraceableFarms(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraceableFarmsTransactor{contract: contract}, nil
}

// NewTraceableFarmsFilterer creates a new log filterer instance of TraceableFarms, bound to a specific deployed contract.
func NewTraceableFarmsFilterer(address common.Address, filterer bind.ContractFilterer) (*TraceableFarmsFilterer, error) {
	contract, err := bindTraceableFarms(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraceableFarmsFilterer{contract: contract}, nil
}

// bindTraceableFarms binds a generic wrapper to an already deployed contract.
func bindTraceableFarms(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraceableFarmsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraceableFarms *TraceableFarmsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraceableFarms.Contract.TraceableFarmsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraceableFarms *TraceableFarmsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraceableFarms.Contract.TraceableFarmsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraceableFarms *TraceableFarmsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraceableFarms.Contract.TraceableFarmsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraceableFarms *TraceableFarmsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraceableFarms.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraceableFarms *TraceableFarmsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraceableFarms.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraceableFarms *TraceableFarmsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraceableFarms.Contract.contract.Transact(opts, method, params...)
}

// GetAccreditation is a free data retrieval call binding the contract method 0x01384d8c.
//
// Solidity: function getAccreditation(string _id) view returns((string,string,(string,string)))
func (_TraceableFarms *TraceableFarmsCaller) GetAccreditation(opts *bind.CallOpts, _id string) (TraceableFarmsAccreditation, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getAccreditation", _id)

	if err != nil {
		return *new(TraceableFarmsAccreditation), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsAccreditation)).(*TraceableFarmsAccreditation)

	return out0, err

}

// GetAccreditation is a free data retrieval call binding the contract method 0x01384d8c.
//
// Solidity: function getAccreditation(string _id) view returns((string,string,(string,string)))
func (_TraceableFarms *TraceableFarmsSession) GetAccreditation(_id string) (TraceableFarmsAccreditation, error) {
	return _TraceableFarms.Contract.GetAccreditation(&_TraceableFarms.CallOpts, _id)
}

// GetAccreditation is a free data retrieval call binding the contract method 0x01384d8c.
//
// Solidity: function getAccreditation(string _id) view returns((string,string,(string,string)))
func (_TraceableFarms *TraceableFarmsCallerSession) GetAccreditation(_id string) (TraceableFarmsAccreditation, error) {
	return _TraceableFarms.Contract.GetAccreditation(&_TraceableFarms.CallOpts, _id)
}

// GetAccreditationCompany is a free data retrieval call binding the contract method 0x460123bc.
//
// Solidity: function getAccreditationCompany(string _nif) view returns(((string,string,(string,string)),string)[])
func (_TraceableFarms *TraceableFarmsCaller) GetAccreditationCompany(opts *bind.CallOpts, _nif string) ([]TraceableFarmsAccreditationCompany, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getAccreditationCompany", _nif)

	if err != nil {
		return *new([]TraceableFarmsAccreditationCompany), err
	}

	out0 := *abi.ConvertType(out[0], new([]TraceableFarmsAccreditationCompany)).(*[]TraceableFarmsAccreditationCompany)

	return out0, err

}

// GetAccreditationCompany is a free data retrieval call binding the contract method 0x460123bc.
//
// Solidity: function getAccreditationCompany(string _nif) view returns(((string,string,(string,string)),string)[])
func (_TraceableFarms *TraceableFarmsSession) GetAccreditationCompany(_nif string) ([]TraceableFarmsAccreditationCompany, error) {
	return _TraceableFarms.Contract.GetAccreditationCompany(&_TraceableFarms.CallOpts, _nif)
}

// GetAccreditationCompany is a free data retrieval call binding the contract method 0x460123bc.
//
// Solidity: function getAccreditationCompany(string _nif) view returns(((string,string,(string,string)),string)[])
func (_TraceableFarms *TraceableFarmsCallerSession) GetAccreditationCompany(_nif string) ([]TraceableFarmsAccreditationCompany, error) {
	return _TraceableFarms.Contract.GetAccreditationCompany(&_TraceableFarms.CallOpts, _nif)
}

// GetAccreditationType is a free data retrieval call binding the contract method 0x10c74bd2.
//
// Solidity: function getAccreditationType(string _id) view returns((string,string))
func (_TraceableFarms *TraceableFarmsCaller) GetAccreditationType(opts *bind.CallOpts, _id string) (TraceableFarmsAccreditationType, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getAccreditationType", _id)

	if err != nil {
		return *new(TraceableFarmsAccreditationType), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsAccreditationType)).(*TraceableFarmsAccreditationType)

	return out0, err

}

// GetAccreditationType is a free data retrieval call binding the contract method 0x10c74bd2.
//
// Solidity: function getAccreditationType(string _id) view returns((string,string))
func (_TraceableFarms *TraceableFarmsSession) GetAccreditationType(_id string) (TraceableFarmsAccreditationType, error) {
	return _TraceableFarms.Contract.GetAccreditationType(&_TraceableFarms.CallOpts, _id)
}

// GetAccreditationType is a free data retrieval call binding the contract method 0x10c74bd2.
//
// Solidity: function getAccreditationType(string _id) view returns((string,string))
func (_TraceableFarms *TraceableFarmsCallerSession) GetAccreditationType(_id string) (TraceableFarmsAccreditationType, error) {
	return _TraceableFarms.Contract.GetAccreditationType(&_TraceableFarms.CallOpts, _id)
}

// GetBatch is a free data retrieval call binding the contract method 0xf9ca4274.
//
// Solidity: function getBatch(string _id) view returns((string,string,string,string,string,string,string,(string,string,string,string,string,(string,string))))
func (_TraceableFarms *TraceableFarmsCaller) GetBatch(opts *bind.CallOpts, _id string) (TraceableFarmsBatch, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getBatch", _id)

	if err != nil {
		return *new(TraceableFarmsBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsBatch)).(*TraceableFarmsBatch)

	return out0, err

}

// GetBatch is a free data retrieval call binding the contract method 0xf9ca4274.
//
// Solidity: function getBatch(string _id) view returns((string,string,string,string,string,string,string,(string,string,string,string,string,(string,string))))
func (_TraceableFarms *TraceableFarmsSession) GetBatch(_id string) (TraceableFarmsBatch, error) {
	return _TraceableFarms.Contract.GetBatch(&_TraceableFarms.CallOpts, _id)
}

// GetBatch is a free data retrieval call binding the contract method 0xf9ca4274.
//
// Solidity: function getBatch(string _id) view returns((string,string,string,string,string,string,string,(string,string,string,string,string,(string,string))))
func (_TraceableFarms *TraceableFarmsCallerSession) GetBatch(_id string) (TraceableFarmsBatch, error) {
	return _TraceableFarms.Contract.GetBatch(&_TraceableFarms.CallOpts, _id)
}

// GetBatchOrigin is a free data retrieval call binding the contract method 0x765344c7.
//
// Solidity: function getBatchOrigin(string _batchId) view returns((string,string,string)[])
func (_TraceableFarms *TraceableFarmsCaller) GetBatchOrigin(opts *bind.CallOpts, _batchId string) ([]TraceableFarmsBatchOrigin, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getBatchOrigin", _batchId)

	if err != nil {
		return *new([]TraceableFarmsBatchOrigin), err
	}

	out0 := *abi.ConvertType(out[0], new([]TraceableFarmsBatchOrigin)).(*[]TraceableFarmsBatchOrigin)

	return out0, err

}

// GetBatchOrigin is a free data retrieval call binding the contract method 0x765344c7.
//
// Solidity: function getBatchOrigin(string _batchId) view returns((string,string,string)[])
func (_TraceableFarms *TraceableFarmsSession) GetBatchOrigin(_batchId string) ([]TraceableFarmsBatchOrigin, error) {
	return _TraceableFarms.Contract.GetBatchOrigin(&_TraceableFarms.CallOpts, _batchId)
}

// GetBatchOrigin is a free data retrieval call binding the contract method 0x765344c7.
//
// Solidity: function getBatchOrigin(string _batchId) view returns((string,string,string)[])
func (_TraceableFarms *TraceableFarmsCallerSession) GetBatchOrigin(_batchId string) ([]TraceableFarmsBatchOrigin, error) {
	return _TraceableFarms.Contract.GetBatchOrigin(&_TraceableFarms.CallOpts, _batchId)
}

// GetBatchProcess is a free data retrieval call binding the contract method 0x3e1c21c2.
//
// Solidity: function getBatchProcess(string _batchId) view returns((string,string)[])
func (_TraceableFarms *TraceableFarmsCaller) GetBatchProcess(opts *bind.CallOpts, _batchId string) ([]TraceableFarmsBatchProcess, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getBatchProcess", _batchId)

	if err != nil {
		return *new([]TraceableFarmsBatchProcess), err
	}

	out0 := *abi.ConvertType(out[0], new([]TraceableFarmsBatchProcess)).(*[]TraceableFarmsBatchProcess)

	return out0, err

}

// GetBatchProcess is a free data retrieval call binding the contract method 0x3e1c21c2.
//
// Solidity: function getBatchProcess(string _batchId) view returns((string,string)[])
func (_TraceableFarms *TraceableFarmsSession) GetBatchProcess(_batchId string) ([]TraceableFarmsBatchProcess, error) {
	return _TraceableFarms.Contract.GetBatchProcess(&_TraceableFarms.CallOpts, _batchId)
}

// GetBatchProcess is a free data retrieval call binding the contract method 0x3e1c21c2.
//
// Solidity: function getBatchProcess(string _batchId) view returns((string,string)[])
func (_TraceableFarms *TraceableFarmsCallerSession) GetBatchProcess(_batchId string) ([]TraceableFarmsBatchProcess, error) {
	return _TraceableFarms.Contract.GetBatchProcess(&_TraceableFarms.CallOpts, _batchId)
}

// GetCompany is a free data retrieval call binding the contract method 0xb24a4e52.
//
// Solidity: function getCompany(string _nif) view returns((string,string,string,string,string,(string,string)))
func (_TraceableFarms *TraceableFarmsCaller) GetCompany(opts *bind.CallOpts, _nif string) (TraceableFarmsCompany, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getCompany", _nif)

	if err != nil {
		return *new(TraceableFarmsCompany), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsCompany)).(*TraceableFarmsCompany)

	return out0, err

}

// GetCompany is a free data retrieval call binding the contract method 0xb24a4e52.
//
// Solidity: function getCompany(string _nif) view returns((string,string,string,string,string,(string,string)))
func (_TraceableFarms *TraceableFarmsSession) GetCompany(_nif string) (TraceableFarmsCompany, error) {
	return _TraceableFarms.Contract.GetCompany(&_TraceableFarms.CallOpts, _nif)
}

// GetCompany is a free data retrieval call binding the contract method 0xb24a4e52.
//
// Solidity: function getCompany(string _nif) view returns((string,string,string,string,string,(string,string)))
func (_TraceableFarms *TraceableFarmsCallerSession) GetCompany(_nif string) (TraceableFarmsCompany, error) {
	return _TraceableFarms.Contract.GetCompany(&_TraceableFarms.CallOpts, _nif)
}

// GetConsortium is a free data retrieval call binding the contract method 0x86dae8ee.
//
// Solidity: function getConsortium(string _id) view returns((string,string))
func (_TraceableFarms *TraceableFarmsCaller) GetConsortium(opts *bind.CallOpts, _id string) (TraceableFarmsConsortium, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getConsortium", _id)

	if err != nil {
		return *new(TraceableFarmsConsortium), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsConsortium)).(*TraceableFarmsConsortium)

	return out0, err

}

// GetConsortium is a free data retrieval call binding the contract method 0x86dae8ee.
//
// Solidity: function getConsortium(string _id) view returns((string,string))
func (_TraceableFarms *TraceableFarmsSession) GetConsortium(_id string) (TraceableFarmsConsortium, error) {
	return _TraceableFarms.Contract.GetConsortium(&_TraceableFarms.CallOpts, _id)
}

// GetConsortium is a free data retrieval call binding the contract method 0x86dae8ee.
//
// Solidity: function getConsortium(string _id) view returns((string,string))
func (_TraceableFarms *TraceableFarmsCallerSession) GetConsortium(_id string) (TraceableFarmsConsortium, error) {
	return _TraceableFarms.Contract.GetConsortium(&_TraceableFarms.CallOpts, _id)
}

// GetFootprintType is a free data retrieval call binding the contract method 0x0ba33f01.
//
// Solidity: function getFootprintType(string _id) view returns((string,string,string,string))
func (_TraceableFarms *TraceableFarmsCaller) GetFootprintType(opts *bind.CallOpts, _id string) (TraceableFarmsFootprintType, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getFootprintType", _id)

	if err != nil {
		return *new(TraceableFarmsFootprintType), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsFootprintType)).(*TraceableFarmsFootprintType)

	return out0, err

}

// GetFootprintType is a free data retrieval call binding the contract method 0x0ba33f01.
//
// Solidity: function getFootprintType(string _id) view returns((string,string,string,string))
func (_TraceableFarms *TraceableFarmsSession) GetFootprintType(_id string) (TraceableFarmsFootprintType, error) {
	return _TraceableFarms.Contract.GetFootprintType(&_TraceableFarms.CallOpts, _id)
}

// GetFootprintType is a free data retrieval call binding the contract method 0x0ba33f01.
//
// Solidity: function getFootprintType(string _id) view returns((string,string,string,string))
func (_TraceableFarms *TraceableFarmsCallerSession) GetFootprintType(_id string) (TraceableFarmsFootprintType, error) {
	return _TraceableFarms.Contract.GetFootprintType(&_TraceableFarms.CallOpts, _id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraceableFarms *TraceableFarmsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraceableFarms *TraceableFarmsSession) Owner() (common.Address, error) {
	return _TraceableFarms.Contract.Owner(&_TraceableFarms.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraceableFarms *TraceableFarmsCallerSession) Owner() (common.Address, error) {
	return _TraceableFarms.Contract.Owner(&_TraceableFarms.CallOpts)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0xd2b9c9d7.
//
// Solidity: function setAccreditation(string _id, string _name, string _accreditationTypeId) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetAccreditation(opts *bind.TransactOpts, _id string, _name string, _accreditationTypeId string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setAccreditation", _id, _name, _accreditationTypeId)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0xd2b9c9d7.
//
// Solidity: function setAccreditation(string _id, string _name, string _accreditationTypeId) returns()
func (_TraceableFarms *TraceableFarmsSession) SetAccreditation(_id string, _name string, _accreditationTypeId string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetAccreditation(&_TraceableFarms.TransactOpts, _id, _name, _accreditationTypeId)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0xd2b9c9d7.
//
// Solidity: function setAccreditation(string _id, string _name, string _accreditationTypeId) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetAccreditation(_id string, _name string, _accreditationTypeId string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetAccreditation(&_TraceableFarms.TransactOpts, _id, _name, _accreditationTypeId)
}

// SetAccreditationCompany is a paid mutator transaction binding the contract method 0x838e8e78.
//
// Solidity: function setAccreditationCompany(string _nif, string _accreditationId, string _checker) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetAccreditationCompany(opts *bind.TransactOpts, _nif string, _accreditationId string, _checker string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setAccreditationCompany", _nif, _accreditationId, _checker)
}

// SetAccreditationCompany is a paid mutator transaction binding the contract method 0x838e8e78.
//
// Solidity: function setAccreditationCompany(string _nif, string _accreditationId, string _checker) returns()
func (_TraceableFarms *TraceableFarmsSession) SetAccreditationCompany(_nif string, _accreditationId string, _checker string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetAccreditationCompany(&_TraceableFarms.TransactOpts, _nif, _accreditationId, _checker)
}

// SetAccreditationCompany is a paid mutator transaction binding the contract method 0x838e8e78.
//
// Solidity: function setAccreditationCompany(string _nif, string _accreditationId, string _checker) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetAccreditationCompany(_nif string, _accreditationId string, _checker string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetAccreditationCompany(&_TraceableFarms.TransactOpts, _nif, _accreditationId, _checker)
}

// SetAccreditationType is a paid mutator transaction binding the contract method 0x09b7ba2e.
//
// Solidity: function setAccreditationType(string _id, string _name) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetAccreditationType(opts *bind.TransactOpts, _id string, _name string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setAccreditationType", _id, _name)
}

// SetAccreditationType is a paid mutator transaction binding the contract method 0x09b7ba2e.
//
// Solidity: function setAccreditationType(string _id, string _name) returns()
func (_TraceableFarms *TraceableFarmsSession) SetAccreditationType(_id string, _name string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetAccreditationType(&_TraceableFarms.TransactOpts, _id, _name)
}

// SetAccreditationType is a paid mutator transaction binding the contract method 0x09b7ba2e.
//
// Solidity: function setAccreditationType(string _id, string _name) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetAccreditationType(_id string, _name string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetAccreditationType(&_TraceableFarms.TransactOpts, _id, _name)
}

// SetBatch is a paid mutator transaction binding the contract method 0x721aaebe.
//
// Solidity: function setBatch(string _id, string _number, string _date, string _productName, string _productVariety, string _productDescription, string _productPhotoUrl, string _companyId) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetBatch(opts *bind.TransactOpts, _id string, _number string, _date string, _productName string, _productVariety string, _productDescription string, _productPhotoUrl string, _companyId string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setBatch", _id, _number, _date, _productName, _productVariety, _productDescription, _productPhotoUrl, _companyId)
}

// SetBatch is a paid mutator transaction binding the contract method 0x721aaebe.
//
// Solidity: function setBatch(string _id, string _number, string _date, string _productName, string _productVariety, string _productDescription, string _productPhotoUrl, string _companyId) returns()
func (_TraceableFarms *TraceableFarmsSession) SetBatch(_id string, _number string, _date string, _productName string, _productVariety string, _productDescription string, _productPhotoUrl string, _companyId string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatch(&_TraceableFarms.TransactOpts, _id, _number, _date, _productName, _productVariety, _productDescription, _productPhotoUrl, _companyId)
}

// SetBatch is a paid mutator transaction binding the contract method 0x721aaebe.
//
// Solidity: function setBatch(string _id, string _number, string _date, string _productName, string _productVariety, string _productDescription, string _productPhotoUrl, string _companyId) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetBatch(_id string, _number string, _date string, _productName string, _productVariety string, _productDescription string, _productPhotoUrl string, _companyId string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatch(&_TraceableFarms.TransactOpts, _id, _number, _date, _productName, _productVariety, _productDescription, _productPhotoUrl, _companyId)
}

// SetBatchOrigin is a paid mutator transaction binding the contract method 0xd03a82ee.
//
// Solidity: function setBatchOrigin(string _batchId, string _description, string _location, string _locationCoordinates) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetBatchOrigin(opts *bind.TransactOpts, _batchId string, _description string, _location string, _locationCoordinates string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setBatchOrigin", _batchId, _description, _location, _locationCoordinates)
}

// SetBatchOrigin is a paid mutator transaction binding the contract method 0xd03a82ee.
//
// Solidity: function setBatchOrigin(string _batchId, string _description, string _location, string _locationCoordinates) returns()
func (_TraceableFarms *TraceableFarmsSession) SetBatchOrigin(_batchId string, _description string, _location string, _locationCoordinates string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchOrigin(&_TraceableFarms.TransactOpts, _batchId, _description, _location, _locationCoordinates)
}

// SetBatchOrigin is a paid mutator transaction binding the contract method 0xd03a82ee.
//
// Solidity: function setBatchOrigin(string _batchId, string _description, string _location, string _locationCoordinates) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetBatchOrigin(_batchId string, _description string, _location string, _locationCoordinates string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchOrigin(&_TraceableFarms.TransactOpts, _batchId, _description, _location, _locationCoordinates)
}

// SetBatchProcess is a paid mutator transaction binding the contract method 0x4849e7a9.
//
// Solidity: function setBatchProcess(string _batchId, string _name, string _description) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetBatchProcess(opts *bind.TransactOpts, _batchId string, _name string, _description string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setBatchProcess", _batchId, _name, _description)
}

// SetBatchProcess is a paid mutator transaction binding the contract method 0x4849e7a9.
//
// Solidity: function setBatchProcess(string _batchId, string _name, string _description) returns()
func (_TraceableFarms *TraceableFarmsSession) SetBatchProcess(_batchId string, _name string, _description string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchProcess(&_TraceableFarms.TransactOpts, _batchId, _name, _description)
}

// SetBatchProcess is a paid mutator transaction binding the contract method 0x4849e7a9.
//
// Solidity: function setBatchProcess(string _batchId, string _name, string _description) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetBatchProcess(_batchId string, _name string, _description string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchProcess(&_TraceableFarms.TransactOpts, _batchId, _name, _description)
}

// SetCompany is a paid mutator transaction binding the contract method 0x97d3502e.
//
// Solidity: function setCompany(string _nif, string _bussinessName, string _location, string _locationCoordinates, string _informationalResourceUrl, string _consortiumId) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetCompany(opts *bind.TransactOpts, _nif string, _bussinessName string, _location string, _locationCoordinates string, _informationalResourceUrl string, _consortiumId string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setCompany", _nif, _bussinessName, _location, _locationCoordinates, _informationalResourceUrl, _consortiumId)
}

// SetCompany is a paid mutator transaction binding the contract method 0x97d3502e.
//
// Solidity: function setCompany(string _nif, string _bussinessName, string _location, string _locationCoordinates, string _informationalResourceUrl, string _consortiumId) returns()
func (_TraceableFarms *TraceableFarmsSession) SetCompany(_nif string, _bussinessName string, _location string, _locationCoordinates string, _informationalResourceUrl string, _consortiumId string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetCompany(&_TraceableFarms.TransactOpts, _nif, _bussinessName, _location, _locationCoordinates, _informationalResourceUrl, _consortiumId)
}

// SetCompany is a paid mutator transaction binding the contract method 0x97d3502e.
//
// Solidity: function setCompany(string _nif, string _bussinessName, string _location, string _locationCoordinates, string _informationalResourceUrl, string _consortiumId) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetCompany(_nif string, _bussinessName string, _location string, _locationCoordinates string, _informationalResourceUrl string, _consortiumId string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetCompany(&_TraceableFarms.TransactOpts, _nif, _bussinessName, _location, _locationCoordinates, _informationalResourceUrl, _consortiumId)
}

// SetConsortium is a paid mutator transaction binding the contract method 0x061d8165.
//
// Solidity: function setConsortium(string _id, string _name) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetConsortium(opts *bind.TransactOpts, _id string, _name string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setConsortium", _id, _name)
}

// SetConsortium is a paid mutator transaction binding the contract method 0x061d8165.
//
// Solidity: function setConsortium(string _id, string _name) returns()
func (_TraceableFarms *TraceableFarmsSession) SetConsortium(_id string, _name string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetConsortium(&_TraceableFarms.TransactOpts, _id, _name)
}

// SetConsortium is a paid mutator transaction binding the contract method 0x061d8165.
//
// Solidity: function setConsortium(string _id, string _name) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetConsortium(_id string, _name string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetConsortium(&_TraceableFarms.TransactOpts, _id, _name)
}

// SetFootprintType is a paid mutator transaction binding the contract method 0xfd2aeace.
//
// Solidity: function setFootprintType(string _id, string _name, string _unitMeasurementName, string _unitMeasurementSymbol) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetFootprintType(opts *bind.TransactOpts, _id string, _name string, _unitMeasurementName string, _unitMeasurementSymbol string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setFootprintType", _id, _name, _unitMeasurementName, _unitMeasurementSymbol)
}

// SetFootprintType is a paid mutator transaction binding the contract method 0xfd2aeace.
//
// Solidity: function setFootprintType(string _id, string _name, string _unitMeasurementName, string _unitMeasurementSymbol) returns()
func (_TraceableFarms *TraceableFarmsSession) SetFootprintType(_id string, _name string, _unitMeasurementName string, _unitMeasurementSymbol string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetFootprintType(&_TraceableFarms.TransactOpts, _id, _name, _unitMeasurementName, _unitMeasurementSymbol)
}

// SetFootprintType is a paid mutator transaction binding the contract method 0xfd2aeace.
//
// Solidity: function setFootprintType(string _id, string _name, string _unitMeasurementName, string _unitMeasurementSymbol) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetFootprintType(_id string, _name string, _unitMeasurementName string, _unitMeasurementSymbol string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetFootprintType(&_TraceableFarms.TransactOpts, _id, _name, _unitMeasurementName, _unitMeasurementSymbol)
}
