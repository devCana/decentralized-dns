// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ResolverIncentivesMetaData contains all meta data concerning the ResolverIncentives contract.
var ResolverIncentivesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BadVoucher\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChannel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotClient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotResolver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToClaim\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroResolver\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refunded\",\"type\":\"uint256\"}],\"name\":\"ChannelClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"name\":\"ChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulative\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_DURATION\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"channels\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimed\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cumulative\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"clientSig\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"closeChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"openChannel\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cumulative\",\"type\":\"uint256\"}],\"name\":\"voucherDigest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557610715908161001b8239f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c90816330b7596c1461041f575080634c2ee09d146103195780637964ea871461011c5780637a7ebd7b146100ac578063b6a6d1771461008f5763e55aee261461006157600080fd5b3461008a57604036600319011261008a576020610082602435600435610671565b604051908152f35b600080fd5b3461008a57600036600319011261008a576020604051610e108152f35b3461008a57602036600319011261008a57600435600052600060205260a06040600020600180831b0381541690600180841b0360018201541690600281015467ffffffffffffffff6004600384015493015416926040519485526020850152604084015260608301526080820152f35b3461008a57606036600319011261008a5760043560443560243567ffffffffffffffff821161008a573660238301121561008a5781600401359067ffffffffffffffff821161008a57366024838501011161008a5783600052600060205260406000209260018060a01b038454169081156103085760018501546001600160a01b031693338590036102f7576041036102b257606481013560001a601b81106102cf575b6080600091604460209460ff6101d6898d610671565b936040519485521686840152602481013560408401520135606082015282805260015afa156102c3576000516001600160a01b0316036102b2576002830154600391818111156102aa5750925b0180549081841115610299576000808093928695610242839588610624565b968792555af1610250610631565b5015610288577f41628d0ba42442e4aa4fc514eeb97bb7154969e70e6678229c836f3b9732ba909160409182519182526020820152a2005b6312171d8360e31b60005260046000fd5b6312d37ee560e31b60005260046000fd5b905092610223565b630de54e5d60e31b60005260046000fd5b6040513d6000823e3d90fd5b601b019060ff82116102e157906101c0565b634e487b7160e01b600052601160045260246000fd5b635d154fe160e11b60005260046000fd5b63078226dd60e51b60005260046000fd5b3461008a57602036600319011261008a57600435600081815260208190526040902080546001600160a01b0316801561030857330361040e5767ffffffffffffffff60048201541642106103fd57806003600261037b93015491015490610624565b81600052600060205260006004604082208281558260018201558260028201558260038201550155806103d6575b60207f74e9aa18d6bb2c4887e76896296ce0a296a2e8315bb319b08b7607ff92fbef7991604051908152a2005b600080808084335af16103e7610631565b506103a9576312171d8360e31b60005260046000fd5b63d0404f8560e01b60005260046000fd5b630836f21d60e21b60005260046000fd5b604036600319011261008a576004356001600160a01b038116919082900361008a576024359067ffffffffffffffff82169182810361008a579183156105f15734156105e057610e10116105d6575b60015460001981146102e15767ffffffffffffffff916001820160015560208101913383528560408301526060820152606081526104ad608082610602565b519020911667ffffffffffffffff4216019167ffffffffffffffff83116102e15760405160a081019381851067ffffffffffffffff8611176105c05760409485523380835260208381018581523485890181815260006060880181815267ffffffffffffffff97881660808a018181528c8452838852928d902099518a546001600160a01b03199081166001600160a01b03928316178c55965160018c018054909816911617909555915160028901559051600388015551600496909601805467ffffffffffffffff191696909516959095179093558651938452838101929092529094909184917f9872b10740b75c20e0eb3eebab184398d737141b8ba28f48c11db6632c60856291a4604051908152f35b634e487b7160e01b600052604160045260246000fd5b610e10915061046e565b633a6a68b160e01b60005260046000fd5b6309e7fc4760e01b60005260046000fd5b90601f8019910116810190811067ffffffffffffffff8211176105c057604052565b919082039182116102e157565b3d1561066c573d9067ffffffffffffffff82116105c05760405191610660601f8201601f191660200184610602565b82523d6000602084013e565b606090565b906040519060208201923084526040830152606082015260608152610697608082610602565b51902060405160208101917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008352603c820152603c81526106d9605c82610602565b5190209056fea2646970667358221220de08872b5058448152900fe7e412755e328d920824db227a100b9ea6a174b72664736f6c634300081c0033",
}

// ResolverIncentivesABI is the input ABI used to generate the binding from.
// Deprecated: Use ResolverIncentivesMetaData.ABI instead.
var ResolverIncentivesABI = ResolverIncentivesMetaData.ABI

// ResolverIncentivesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ResolverIncentivesMetaData.Bin instead.
var ResolverIncentivesBin = ResolverIncentivesMetaData.Bin

// DeployResolverIncentives deploys a new Ethereum contract, binding an instance of ResolverIncentives to it.
func DeployResolverIncentives(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ResolverIncentives, error) {
	parsed, err := ResolverIncentivesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ResolverIncentivesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ResolverIncentives{ResolverIncentivesCaller: ResolverIncentivesCaller{contract: contract}, ResolverIncentivesTransactor: ResolverIncentivesTransactor{contract: contract}, ResolverIncentivesFilterer: ResolverIncentivesFilterer{contract: contract}}, nil
}

// ResolverIncentives is an auto generated Go binding around an Ethereum contract.
type ResolverIncentives struct {
	ResolverIncentivesCaller     // Read-only binding to the contract
	ResolverIncentivesTransactor // Write-only binding to the contract
	ResolverIncentivesFilterer   // Log filterer for contract events
}

// ResolverIncentivesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverIncentivesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverIncentivesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverIncentivesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverIncentivesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverIncentivesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverIncentivesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverIncentivesSession struct {
	Contract     *ResolverIncentives // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ResolverIncentivesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverIncentivesCallerSession struct {
	Contract *ResolverIncentivesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ResolverIncentivesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverIncentivesTransactorSession struct {
	Contract     *ResolverIncentivesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ResolverIncentivesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverIncentivesRaw struct {
	Contract *ResolverIncentives // Generic contract binding to access the raw methods on
}

// ResolverIncentivesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverIncentivesCallerRaw struct {
	Contract *ResolverIncentivesCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverIncentivesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverIncentivesTransactorRaw struct {
	Contract *ResolverIncentivesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolverIncentives creates a new instance of ResolverIncentives, bound to a specific deployed contract.
func NewResolverIncentives(address common.Address, backend bind.ContractBackend) (*ResolverIncentives, error) {
	contract, err := bindResolverIncentives(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResolverIncentives{ResolverIncentivesCaller: ResolverIncentivesCaller{contract: contract}, ResolverIncentivesTransactor: ResolverIncentivesTransactor{contract: contract}, ResolverIncentivesFilterer: ResolverIncentivesFilterer{contract: contract}}, nil
}

// NewResolverIncentivesCaller creates a new read-only instance of ResolverIncentives, bound to a specific deployed contract.
func NewResolverIncentivesCaller(address common.Address, caller bind.ContractCaller) (*ResolverIncentivesCaller, error) {
	contract, err := bindResolverIncentives(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverIncentivesCaller{contract: contract}, nil
}

// NewResolverIncentivesTransactor creates a new write-only instance of ResolverIncentives, bound to a specific deployed contract.
func NewResolverIncentivesTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverIncentivesTransactor, error) {
	contract, err := bindResolverIncentives(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverIncentivesTransactor{contract: contract}, nil
}

// NewResolverIncentivesFilterer creates a new log filterer instance of ResolverIncentives, bound to a specific deployed contract.
func NewResolverIncentivesFilterer(address common.Address, filterer bind.ContractFilterer) (*ResolverIncentivesFilterer, error) {
	contract, err := bindResolverIncentives(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverIncentivesFilterer{contract: contract}, nil
}

// bindResolverIncentives binds a generic wrapper to an already deployed contract.
func bindResolverIncentives(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResolverIncentivesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverIncentives *ResolverIncentivesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolverIncentives.Contract.ResolverIncentivesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverIncentives *ResolverIncentivesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverIncentives.Contract.ResolverIncentivesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverIncentives *ResolverIncentivesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverIncentives.Contract.ResolverIncentivesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverIncentives *ResolverIncentivesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolverIncentives.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverIncentives *ResolverIncentivesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverIncentives.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverIncentives *ResolverIncentivesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverIncentives.Contract.contract.Transact(opts, method, params...)
}

// MINDURATION is a free data retrieval call binding the contract method 0xb6a6d177.
//
// Solidity: function MIN_DURATION() view returns(uint64)
func (_ResolverIncentives *ResolverIncentivesCaller) MINDURATION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ResolverIncentives.contract.Call(opts, &out, "MIN_DURATION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINDURATION is a free data retrieval call binding the contract method 0xb6a6d177.
//
// Solidity: function MIN_DURATION() view returns(uint64)
func (_ResolverIncentives *ResolverIncentivesSession) MINDURATION() (uint64, error) {
	return _ResolverIncentives.Contract.MINDURATION(&_ResolverIncentives.CallOpts)
}

// MINDURATION is a free data retrieval call binding the contract method 0xb6a6d177.
//
// Solidity: function MIN_DURATION() view returns(uint64)
func (_ResolverIncentives *ResolverIncentivesCallerSession) MINDURATION() (uint64, error) {
	return _ResolverIncentives.Contract.MINDURATION(&_ResolverIncentives.CallOpts)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address client, address resolver, uint256 deposit, uint256 claimed, uint64 expiresAt)
func (_ResolverIncentives *ResolverIncentivesCaller) Channels(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Client    common.Address
	Resolver  common.Address
	Deposit   *big.Int
	Claimed   *big.Int
	ExpiresAt uint64
}, error) {
	var out []interface{}
	err := _ResolverIncentives.contract.Call(opts, &out, "channels", arg0)

	outstruct := new(struct {
		Client    common.Address
		Resolver  common.Address
		Deposit   *big.Int
		Claimed   *big.Int
		ExpiresAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Client = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Resolver = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Deposit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ExpiresAt = *abi.ConvertType(out[4], new(uint64)).(*uint64)

	return *outstruct, err

}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address client, address resolver, uint256 deposit, uint256 claimed, uint64 expiresAt)
func (_ResolverIncentives *ResolverIncentivesSession) Channels(arg0 [32]byte) (struct {
	Client    common.Address
	Resolver  common.Address
	Deposit   *big.Int
	Claimed   *big.Int
	ExpiresAt uint64
}, error) {
	return _ResolverIncentives.Contract.Channels(&_ResolverIncentives.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address client, address resolver, uint256 deposit, uint256 claimed, uint64 expiresAt)
func (_ResolverIncentives *ResolverIncentivesCallerSession) Channels(arg0 [32]byte) (struct {
	Client    common.Address
	Resolver  common.Address
	Deposit   *big.Int
	Claimed   *big.Int
	ExpiresAt uint64
}, error) {
	return _ResolverIncentives.Contract.Channels(&_ResolverIncentives.CallOpts, arg0)
}

// VoucherDigest is a free data retrieval call binding the contract method 0xe55aee26.
//
// Solidity: function voucherDigest(bytes32 id, uint256 cumulative) view returns(bytes32)
func (_ResolverIncentives *ResolverIncentivesCaller) VoucherDigest(opts *bind.CallOpts, id [32]byte, cumulative *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ResolverIncentives.contract.Call(opts, &out, "voucherDigest", id, cumulative)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VoucherDigest is a free data retrieval call binding the contract method 0xe55aee26.
//
// Solidity: function voucherDigest(bytes32 id, uint256 cumulative) view returns(bytes32)
func (_ResolverIncentives *ResolverIncentivesSession) VoucherDigest(id [32]byte, cumulative *big.Int) ([32]byte, error) {
	return _ResolverIncentives.Contract.VoucherDigest(&_ResolverIncentives.CallOpts, id, cumulative)
}

// VoucherDigest is a free data retrieval call binding the contract method 0xe55aee26.
//
// Solidity: function voucherDigest(bytes32 id, uint256 cumulative) view returns(bytes32)
func (_ResolverIncentives *ResolverIncentivesCallerSession) VoucherDigest(id [32]byte, cumulative *big.Int) ([32]byte, error) {
	return _ResolverIncentives.Contract.VoucherDigest(&_ResolverIncentives.CallOpts, id, cumulative)
}

// Claim is a paid mutator transaction binding the contract method 0x7964ea87.
//
// Solidity: function claim(bytes32 id, uint256 cumulative, bytes clientSig) returns()
func (_ResolverIncentives *ResolverIncentivesTransactor) Claim(opts *bind.TransactOpts, id [32]byte, cumulative *big.Int, clientSig []byte) (*types.Transaction, error) {
	return _ResolverIncentives.contract.Transact(opts, "claim", id, cumulative, clientSig)
}

// Claim is a paid mutator transaction binding the contract method 0x7964ea87.
//
// Solidity: function claim(bytes32 id, uint256 cumulative, bytes clientSig) returns()
func (_ResolverIncentives *ResolverIncentivesSession) Claim(id [32]byte, cumulative *big.Int, clientSig []byte) (*types.Transaction, error) {
	return _ResolverIncentives.Contract.Claim(&_ResolverIncentives.TransactOpts, id, cumulative, clientSig)
}

// Claim is a paid mutator transaction binding the contract method 0x7964ea87.
//
// Solidity: function claim(bytes32 id, uint256 cumulative, bytes clientSig) returns()
func (_ResolverIncentives *ResolverIncentivesTransactorSession) Claim(id [32]byte, cumulative *big.Int, clientSig []byte) (*types.Transaction, error) {
	return _ResolverIncentives.Contract.Claim(&_ResolverIncentives.TransactOpts, id, cumulative, clientSig)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 id) returns()
func (_ResolverIncentives *ResolverIncentivesTransactor) CloseChannel(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _ResolverIncentives.contract.Transact(opts, "closeChannel", id)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 id) returns()
func (_ResolverIncentives *ResolverIncentivesSession) CloseChannel(id [32]byte) (*types.Transaction, error) {
	return _ResolverIncentives.Contract.CloseChannel(&_ResolverIncentives.TransactOpts, id)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 id) returns()
func (_ResolverIncentives *ResolverIncentivesTransactorSession) CloseChannel(id [32]byte) (*types.Transaction, error) {
	return _ResolverIncentives.Contract.CloseChannel(&_ResolverIncentives.TransactOpts, id)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x30b7596c.
//
// Solidity: function openChannel(address resolver, uint64 duration) payable returns(bytes32 id)
func (_ResolverIncentives *ResolverIncentivesTransactor) OpenChannel(opts *bind.TransactOpts, resolver common.Address, duration uint64) (*types.Transaction, error) {
	return _ResolverIncentives.contract.Transact(opts, "openChannel", resolver, duration)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x30b7596c.
//
// Solidity: function openChannel(address resolver, uint64 duration) payable returns(bytes32 id)
func (_ResolverIncentives *ResolverIncentivesSession) OpenChannel(resolver common.Address, duration uint64) (*types.Transaction, error) {
	return _ResolverIncentives.Contract.OpenChannel(&_ResolverIncentives.TransactOpts, resolver, duration)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x30b7596c.
//
// Solidity: function openChannel(address resolver, uint64 duration) payable returns(bytes32 id)
func (_ResolverIncentives *ResolverIncentivesTransactorSession) OpenChannel(resolver common.Address, duration uint64) (*types.Transaction, error) {
	return _ResolverIncentives.Contract.OpenChannel(&_ResolverIncentives.TransactOpts, resolver, duration)
}

// ResolverIncentivesChannelClosedIterator is returned from FilterChannelClosed and is used to iterate over the raw logs and unpacked data for ChannelClosed events raised by the ResolverIncentives contract.
type ResolverIncentivesChannelClosedIterator struct {
	Event *ResolverIncentivesChannelClosed // Event containing the contract specifics and raw log

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
func (it *ResolverIncentivesChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverIncentivesChannelClosed)
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
		it.Event = new(ResolverIncentivesChannelClosed)
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
func (it *ResolverIncentivesChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverIncentivesChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverIncentivesChannelClosed represents a ChannelClosed event raised by the ResolverIncentives contract.
type ResolverIncentivesChannelClosed struct {
	Id       [32]byte
	Refunded *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelClosed is a free log retrieval operation binding the contract event 0x74e9aa18d6bb2c4887e76896296ce0a296a2e8315bb319b08b7607ff92fbef79.
//
// Solidity: event ChannelClosed(bytes32 indexed id, uint256 refunded)
func (_ResolverIncentives *ResolverIncentivesFilterer) FilterChannelClosed(opts *bind.FilterOpts, id [][32]byte) (*ResolverIncentivesChannelClosedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ResolverIncentives.contract.FilterLogs(opts, "ChannelClosed", idRule)
	if err != nil {
		return nil, err
	}
	return &ResolverIncentivesChannelClosedIterator{contract: _ResolverIncentives.contract, event: "ChannelClosed", logs: logs, sub: sub}, nil
}

// WatchChannelClosed is a free log subscription operation binding the contract event 0x74e9aa18d6bb2c4887e76896296ce0a296a2e8315bb319b08b7607ff92fbef79.
//
// Solidity: event ChannelClosed(bytes32 indexed id, uint256 refunded)
func (_ResolverIncentives *ResolverIncentivesFilterer) WatchChannelClosed(opts *bind.WatchOpts, sink chan<- *ResolverIncentivesChannelClosed, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ResolverIncentives.contract.WatchLogs(opts, "ChannelClosed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverIncentivesChannelClosed)
				if err := _ResolverIncentives.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
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

// ParseChannelClosed is a log parse operation binding the contract event 0x74e9aa18d6bb2c4887e76896296ce0a296a2e8315bb319b08b7607ff92fbef79.
//
// Solidity: event ChannelClosed(bytes32 indexed id, uint256 refunded)
func (_ResolverIncentives *ResolverIncentivesFilterer) ParseChannelClosed(log types.Log) (*ResolverIncentivesChannelClosed, error) {
	event := new(ResolverIncentivesChannelClosed)
	if err := _ResolverIncentives.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverIncentivesChannelOpenedIterator is returned from FilterChannelOpened and is used to iterate over the raw logs and unpacked data for ChannelOpened events raised by the ResolverIncentives contract.
type ResolverIncentivesChannelOpenedIterator struct {
	Event *ResolverIncentivesChannelOpened // Event containing the contract specifics and raw log

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
func (it *ResolverIncentivesChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverIncentivesChannelOpened)
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
		it.Event = new(ResolverIncentivesChannelOpened)
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
func (it *ResolverIncentivesChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverIncentivesChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverIncentivesChannelOpened represents a ChannelOpened event raised by the ResolverIncentives contract.
type ResolverIncentivesChannelOpened struct {
	Id        [32]byte
	Client    common.Address
	Resolver  common.Address
	Deposit   *big.Int
	ExpiresAt uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelOpened is a free log retrieval operation binding the contract event 0x9872b10740b75c20e0eb3eebab184398d737141b8ba28f48c11db6632c608562.
//
// Solidity: event ChannelOpened(bytes32 indexed id, address indexed client, address indexed resolver, uint256 deposit, uint64 expiresAt)
func (_ResolverIncentives *ResolverIncentivesFilterer) FilterChannelOpened(opts *bind.FilterOpts, id [][32]byte, client []common.Address, resolver []common.Address) (*ResolverIncentivesChannelOpenedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var resolverRule []interface{}
	for _, resolverItem := range resolver {
		resolverRule = append(resolverRule, resolverItem)
	}

	logs, sub, err := _ResolverIncentives.contract.FilterLogs(opts, "ChannelOpened", idRule, clientRule, resolverRule)
	if err != nil {
		return nil, err
	}
	return &ResolverIncentivesChannelOpenedIterator{contract: _ResolverIncentives.contract, event: "ChannelOpened", logs: logs, sub: sub}, nil
}

// WatchChannelOpened is a free log subscription operation binding the contract event 0x9872b10740b75c20e0eb3eebab184398d737141b8ba28f48c11db6632c608562.
//
// Solidity: event ChannelOpened(bytes32 indexed id, address indexed client, address indexed resolver, uint256 deposit, uint64 expiresAt)
func (_ResolverIncentives *ResolverIncentivesFilterer) WatchChannelOpened(opts *bind.WatchOpts, sink chan<- *ResolverIncentivesChannelOpened, id [][32]byte, client []common.Address, resolver []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var resolverRule []interface{}
	for _, resolverItem := range resolver {
		resolverRule = append(resolverRule, resolverItem)
	}

	logs, sub, err := _ResolverIncentives.contract.WatchLogs(opts, "ChannelOpened", idRule, clientRule, resolverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverIncentivesChannelOpened)
				if err := _ResolverIncentives.contract.UnpackLog(event, "ChannelOpened", log); err != nil {
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

// ParseChannelOpened is a log parse operation binding the contract event 0x9872b10740b75c20e0eb3eebab184398d737141b8ba28f48c11db6632c608562.
//
// Solidity: event ChannelOpened(bytes32 indexed id, address indexed client, address indexed resolver, uint256 deposit, uint64 expiresAt)
func (_ResolverIncentives *ResolverIncentivesFilterer) ParseChannelOpened(log types.Log) (*ResolverIncentivesChannelOpened, error) {
	event := new(ResolverIncentivesChannelOpened)
	if err := _ResolverIncentives.contract.UnpackLog(event, "ChannelOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverIncentivesClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the ResolverIncentives contract.
type ResolverIncentivesClaimedIterator struct {
	Event *ResolverIncentivesClaimed // Event containing the contract specifics and raw log

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
func (it *ResolverIncentivesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverIncentivesClaimed)
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
		it.Event = new(ResolverIncentivesClaimed)
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
func (it *ResolverIncentivesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverIncentivesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverIncentivesClaimed represents a Claimed event raised by the ResolverIncentives contract.
type ResolverIncentivesClaimed struct {
	Id         [32]byte
	Cumulative *big.Int
	Paid       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x41628d0ba42442e4aa4fc514eeb97bb7154969e70e6678229c836f3b9732ba90.
//
// Solidity: event Claimed(bytes32 indexed id, uint256 cumulative, uint256 paid)
func (_ResolverIncentives *ResolverIncentivesFilterer) FilterClaimed(opts *bind.FilterOpts, id [][32]byte) (*ResolverIncentivesClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ResolverIncentives.contract.FilterLogs(opts, "Claimed", idRule)
	if err != nil {
		return nil, err
	}
	return &ResolverIncentivesClaimedIterator{contract: _ResolverIncentives.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x41628d0ba42442e4aa4fc514eeb97bb7154969e70e6678229c836f3b9732ba90.
//
// Solidity: event Claimed(bytes32 indexed id, uint256 cumulative, uint256 paid)
func (_ResolverIncentives *ResolverIncentivesFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ResolverIncentivesClaimed, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ResolverIncentives.contract.WatchLogs(opts, "Claimed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverIncentivesClaimed)
				if err := _ResolverIncentives.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x41628d0ba42442e4aa4fc514eeb97bb7154969e70e6678229c836f3b9732ba90.
//
// Solidity: event Claimed(bytes32 indexed id, uint256 cumulative, uint256 paid)
func (_ResolverIncentives *ResolverIncentivesFilterer) ParseClaimed(log types.Log) (*ResolverIncentivesClaimed, error) {
	event := new(ResolverIncentivesClaimed)
	if err := _ResolverIncentives.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
