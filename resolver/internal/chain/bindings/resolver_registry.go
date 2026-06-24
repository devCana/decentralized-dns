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

// ResolverRegistryMetaData contains all meta data concerning the ResolverRegistry contract.
var ResolverRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EmptyEndpoint\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyPubKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EndpointTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegistered\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"ResolverAnnounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ResolverRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ENDPOINT_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeResolvers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubKeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"string[]\",\"name\":\"endpoints\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"announce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getResolver\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"updatedAt\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"operatorAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557610973908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c8063261423351461072d578063319a5b07146104885780635392c157146104035780637c6f3158146103e5578063842929d514610102578063b6549f751461008b5763c85d3eb01461006957600080fd5b346100865760003660031901126100865760206040516101008152f35b600080fd5b34610086576000366003190112610086573360005260006020526002604060002001805460ff8160401c16156100f15768ff000000000000000019169055337f6cc76de175eb7c3679f8d4c3529e23883784a40e1f9d7f54e8a9f97115aab347600080a2005b63aba4733960e01b60005260046000fd5b346100865760403660031901126100865760243560043567ffffffffffffffff821161008657366023830112156100865781600401359067ffffffffffffffff821161008657602483019136602482860101116100865781156103d457600092601f8201601f1981169460405161017c60208801826107fb565b8481528484602083013781602086830101528051156103c5576101009051116103b657338152806020526040812091600283019267ffffffffffffffff84541615610363575b8087600192550191506101d5825461085f565b90601f8211610323575b50506000601f851160011461028f5792849283926000966060967f2af369faa85a14ab1ea1ae248aca2b45493d2e3e197e4ce1adbb32208f79a1c39a9b8992610281575b50508460011b9088198660031b1c19161790555b600160401b67ffffffffffffffff421668ffffffffffffffffff198354161717905560405195865260406020870152816040870152838601378301015260608133948101030190a2005b602492500101358b80610223565b8181526020812090601f198616815b818110610308575092606095927f2af369faa85a14ab1ea1ae248aca2b45493d2e3e197e4ce1adbb32208f79a1c3999a8896938760009a97106102ec575b505050600184811b019055610237565b01602401358819600387901b60f8161c191690558a80806102dc565b91926020600181926024878f0101358155019401920161029e565b8260005260206000209060051c81019160208710610359575b601f0160051c01905b818110156101df5760008155600101610345565b909150819061033c565b600154600160401b8110156103a2579061038382600180940184556107a5565b81549060031b9033821b91858060a01b03901b191617905590506101c2565b634e487b7160e01b84526041600452602484fd5b639049092560e01b8152600490fd5b63221418ad60e21b8252600482fd5b633f26f9c760e21b60005260046000fd5b34610086576000366003190112610086576020600154604051908152f35b34610086576020366003190112610086576004356001600160a01b0381169081900361008657600052600060205260406000208054610449600160028401549301610899565b9160ff6104686040519485948552608060208601526080850190610764565b9167ffffffffffffffff8116604085015260401c16151560608301520390f35b3461008657600036600319011261008657600080600154905b8181106106dd5750906104b381610833565b906104c160405192836107fb565b8082526104cd81610833565b602083019390601f19013685376104e382610833565b936104f160405195866107fb565b8285526104fd83610833565b602086019290601f190136843761051384610833565b9361052160405195866107fb565b808552610530601f1991610833565b0160005b8181106106ca5750506000805b82811061062c575050506040519485946060860190606087525180915260808601929060005b81811061060a5750505060209085830382870152519182815201919060005b8181106105f1575050508281036040840152815180825260208201916020808360051b8301019401926000915b8383106105c05786860387f35b9193955091936020806105df600193601f198682030187528951610764565b970193019301909286959492936105b3565b8251845286955060209384019390920191600101610586565b82516001600160a01b0316855288975060209485019490920191600101610567565b61063b819894959697986107a5565b60018060a01b0391549060031b1c168060005260006020526040600020908960ff600284015460401c1661067b575b505050600101969594939296610541565b60016106c19386936106928397986106a69561084b565b52805461069f858b61084b565b5201610899565b6106b0828b61084b565b526106bb818a61084b565b506107d6565b9190898961066a565b6060602082880181019190915201610534565b6106e6816107a5565b60018060a01b0391549060031b1c16600052600060205260ff60026040600020015460401c16610719575b6001016104a1565b916107256001916107d6565b929050610711565b3461008657602036600319011261008657602061074b6004356107a5565b905460405160039290921b1c6001600160a01b03168152f35b919082519283825260005b848110610790575050826000602080949584010152601f8019910116010190565b8060208092840101518282860101520161076f565b6001548110156107c057600160005260206000200190600090565b634e487b7160e01b600052603260045260246000fd5b60001981146107e55760010190565b634e487b7160e01b600052601160045260246000fd5b90601f8019910116810190811067ffffffffffffffff82111761081d57604052565b634e487b7160e01b600052604160045260246000fd5b67ffffffffffffffff811161081d5760051b60200190565b80518210156107c05760209160051b010190565b90600182811c9216801561088f575b602083101461087957565b634e487b7160e01b600052602260045260246000fd5b91607f169161086e565b90604051918260008254926108ad8461085f565b808452936001811690811561091b57506001146108d4575b506108d2925003836107fb565b565b90506000929192526020600020906000915b8183106108ff5750509060206108d292820101386108c5565b60209193508060019154838589010152019101909184926108e6565b9050602092506108d294915060ff191682840152151560051b820101386108c556fea2646970667358221220d825a2eeb6001adfb22974236381d05c66e50b05b45a1722ac492de8cc5a1a2764736f6c634300081c0033",
}

// ResolverRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ResolverRegistryMetaData.ABI instead.
var ResolverRegistryABI = ResolverRegistryMetaData.ABI

// ResolverRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ResolverRegistryMetaData.Bin instead.
var ResolverRegistryBin = ResolverRegistryMetaData.Bin

// DeployResolverRegistry deploys a new Ethereum contract, binding an instance of ResolverRegistry to it.
func DeployResolverRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ResolverRegistry, error) {
	parsed, err := ResolverRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ResolverRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ResolverRegistry{ResolverRegistryCaller: ResolverRegistryCaller{contract: contract}, ResolverRegistryTransactor: ResolverRegistryTransactor{contract: contract}, ResolverRegistryFilterer: ResolverRegistryFilterer{contract: contract}}, nil
}

// ResolverRegistry is an auto generated Go binding around an Ethereum contract.
type ResolverRegistry struct {
	ResolverRegistryCaller     // Read-only binding to the contract
	ResolverRegistryTransactor // Write-only binding to the contract
	ResolverRegistryFilterer   // Log filterer for contract events
}

// ResolverRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverRegistrySession struct {
	Contract     *ResolverRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverRegistryCallerSession struct {
	Contract *ResolverRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ResolverRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverRegistryTransactorSession struct {
	Contract     *ResolverRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ResolverRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverRegistryRaw struct {
	Contract *ResolverRegistry // Generic contract binding to access the raw methods on
}

// ResolverRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverRegistryCallerRaw struct {
	Contract *ResolverRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverRegistryTransactorRaw struct {
	Contract *ResolverRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolverRegistry creates a new instance of ResolverRegistry, bound to a specific deployed contract.
func NewResolverRegistry(address common.Address, backend bind.ContractBackend) (*ResolverRegistry, error) {
	contract, err := bindResolverRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResolverRegistry{ResolverRegistryCaller: ResolverRegistryCaller{contract: contract}, ResolverRegistryTransactor: ResolverRegistryTransactor{contract: contract}, ResolverRegistryFilterer: ResolverRegistryFilterer{contract: contract}}, nil
}

// NewResolverRegistryCaller creates a new read-only instance of ResolverRegistry, bound to a specific deployed contract.
func NewResolverRegistryCaller(address common.Address, caller bind.ContractCaller) (*ResolverRegistryCaller, error) {
	contract, err := bindResolverRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverRegistryCaller{contract: contract}, nil
}

// NewResolverRegistryTransactor creates a new write-only instance of ResolverRegistry, bound to a specific deployed contract.
func NewResolverRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverRegistryTransactor, error) {
	contract, err := bindResolverRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverRegistryTransactor{contract: contract}, nil
}

// NewResolverRegistryFilterer creates a new log filterer instance of ResolverRegistry, bound to a specific deployed contract.
func NewResolverRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ResolverRegistryFilterer, error) {
	contract, err := bindResolverRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverRegistryFilterer{contract: contract}, nil
}

// bindResolverRegistry binds a generic wrapper to an already deployed contract.
func bindResolverRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResolverRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverRegistry *ResolverRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolverRegistry.Contract.ResolverRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverRegistry *ResolverRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverRegistry.Contract.ResolverRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverRegistry *ResolverRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverRegistry.Contract.ResolverRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverRegistry *ResolverRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolverRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverRegistry *ResolverRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverRegistry *ResolverRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXENDPOINTLENGTH is a free data retrieval call binding the contract method 0xc85d3eb0.
//
// Solidity: function MAX_ENDPOINT_LENGTH() view returns(uint256)
func (_ResolverRegistry *ResolverRegistryCaller) MAXENDPOINTLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ResolverRegistry.contract.Call(opts, &out, "MAX_ENDPOINT_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXENDPOINTLENGTH is a free data retrieval call binding the contract method 0xc85d3eb0.
//
// Solidity: function MAX_ENDPOINT_LENGTH() view returns(uint256)
func (_ResolverRegistry *ResolverRegistrySession) MAXENDPOINTLENGTH() (*big.Int, error) {
	return _ResolverRegistry.Contract.MAXENDPOINTLENGTH(&_ResolverRegistry.CallOpts)
}

// MAXENDPOINTLENGTH is a free data retrieval call binding the contract method 0xc85d3eb0.
//
// Solidity: function MAX_ENDPOINT_LENGTH() view returns(uint256)
func (_ResolverRegistry *ResolverRegistryCallerSession) MAXENDPOINTLENGTH() (*big.Int, error) {
	return _ResolverRegistry.Contract.MAXENDPOINTLENGTH(&_ResolverRegistry.CallOpts)
}

// ActiveResolvers is a free data retrieval call binding the contract method 0x319a5b07.
//
// Solidity: function activeResolvers() view returns(address[] operators, bytes32[] pubKeys, string[] endpoints)
func (_ResolverRegistry *ResolverRegistryCaller) ActiveResolvers(opts *bind.CallOpts) (struct {
	Operators []common.Address
	PubKeys   [][32]byte
	Endpoints []string
}, error) {
	var out []interface{}
	err := _ResolverRegistry.contract.Call(opts, &out, "activeResolvers")

	outstruct := new(struct {
		Operators []common.Address
		PubKeys   [][32]byte
		Endpoints []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.PubKeys = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)
	outstruct.Endpoints = *abi.ConvertType(out[2], new([]string)).(*[]string)

	return *outstruct, err

}

// ActiveResolvers is a free data retrieval call binding the contract method 0x319a5b07.
//
// Solidity: function activeResolvers() view returns(address[] operators, bytes32[] pubKeys, string[] endpoints)
func (_ResolverRegistry *ResolverRegistrySession) ActiveResolvers() (struct {
	Operators []common.Address
	PubKeys   [][32]byte
	Endpoints []string
}, error) {
	return _ResolverRegistry.Contract.ActiveResolvers(&_ResolverRegistry.CallOpts)
}

// ActiveResolvers is a free data retrieval call binding the contract method 0x319a5b07.
//
// Solidity: function activeResolvers() view returns(address[] operators, bytes32[] pubKeys, string[] endpoints)
func (_ResolverRegistry *ResolverRegistryCallerSession) ActiveResolvers() (struct {
	Operators []common.Address
	PubKeys   [][32]byte
	Endpoints []string
}, error) {
	return _ResolverRegistry.Contract.ActiveResolvers(&_ResolverRegistry.CallOpts)
}

// GetResolver is a free data retrieval call binding the contract method 0x5392c157.
//
// Solidity: function getResolver(address operator) view returns(bytes32 pubKey, string endpoint, uint64 updatedAt, bool active)
func (_ResolverRegistry *ResolverRegistryCaller) GetResolver(opts *bind.CallOpts, operator common.Address) (struct {
	PubKey    [32]byte
	Endpoint  string
	UpdatedAt uint64
	Active    bool
}, error) {
	var out []interface{}
	err := _ResolverRegistry.contract.Call(opts, &out, "getResolver", operator)

	outstruct := new(struct {
		PubKey    [32]byte
		Endpoint  string
		UpdatedAt uint64
		Active    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PubKey = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Endpoint = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.UpdatedAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Active = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetResolver is a free data retrieval call binding the contract method 0x5392c157.
//
// Solidity: function getResolver(address operator) view returns(bytes32 pubKey, string endpoint, uint64 updatedAt, bool active)
func (_ResolverRegistry *ResolverRegistrySession) GetResolver(operator common.Address) (struct {
	PubKey    [32]byte
	Endpoint  string
	UpdatedAt uint64
	Active    bool
}, error) {
	return _ResolverRegistry.Contract.GetResolver(&_ResolverRegistry.CallOpts, operator)
}

// GetResolver is a free data retrieval call binding the contract method 0x5392c157.
//
// Solidity: function getResolver(address operator) view returns(bytes32 pubKey, string endpoint, uint64 updatedAt, bool active)
func (_ResolverRegistry *ResolverRegistryCallerSession) GetResolver(operator common.Address) (struct {
	PubKey    [32]byte
	Endpoint  string
	UpdatedAt uint64
	Active    bool
}, error) {
	return _ResolverRegistry.Contract.GetResolver(&_ResolverRegistry.CallOpts, operator)
}

// OperatorAt is a free data retrieval call binding the contract method 0x26142335.
//
// Solidity: function operatorAt(uint256 i) view returns(address)
func (_ResolverRegistry *ResolverRegistryCaller) OperatorAt(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ResolverRegistry.contract.Call(opts, &out, "operatorAt", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorAt is a free data retrieval call binding the contract method 0x26142335.
//
// Solidity: function operatorAt(uint256 i) view returns(address)
func (_ResolverRegistry *ResolverRegistrySession) OperatorAt(i *big.Int) (common.Address, error) {
	return _ResolverRegistry.Contract.OperatorAt(&_ResolverRegistry.CallOpts, i)
}

// OperatorAt is a free data retrieval call binding the contract method 0x26142335.
//
// Solidity: function operatorAt(uint256 i) view returns(address)
func (_ResolverRegistry *ResolverRegistryCallerSession) OperatorAt(i *big.Int) (common.Address, error) {
	return _ResolverRegistry.Contract.OperatorAt(&_ResolverRegistry.CallOpts, i)
}

// OperatorCount is a free data retrieval call binding the contract method 0x7c6f3158.
//
// Solidity: function operatorCount() view returns(uint256)
func (_ResolverRegistry *ResolverRegistryCaller) OperatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ResolverRegistry.contract.Call(opts, &out, "operatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorCount is a free data retrieval call binding the contract method 0x7c6f3158.
//
// Solidity: function operatorCount() view returns(uint256)
func (_ResolverRegistry *ResolverRegistrySession) OperatorCount() (*big.Int, error) {
	return _ResolverRegistry.Contract.OperatorCount(&_ResolverRegistry.CallOpts)
}

// OperatorCount is a free data retrieval call binding the contract method 0x7c6f3158.
//
// Solidity: function operatorCount() view returns(uint256)
func (_ResolverRegistry *ResolverRegistryCallerSession) OperatorCount() (*big.Int, error) {
	return _ResolverRegistry.Contract.OperatorCount(&_ResolverRegistry.CallOpts)
}

// Announce is a paid mutator transaction binding the contract method 0x842929d5.
//
// Solidity: function announce(bytes32 pubKey, string endpoint) returns()
func (_ResolverRegistry *ResolverRegistryTransactor) Announce(opts *bind.TransactOpts, pubKey [32]byte, endpoint string) (*types.Transaction, error) {
	return _ResolverRegistry.contract.Transact(opts, "announce", pubKey, endpoint)
}

// Announce is a paid mutator transaction binding the contract method 0x842929d5.
//
// Solidity: function announce(bytes32 pubKey, string endpoint) returns()
func (_ResolverRegistry *ResolverRegistrySession) Announce(pubKey [32]byte, endpoint string) (*types.Transaction, error) {
	return _ResolverRegistry.Contract.Announce(&_ResolverRegistry.TransactOpts, pubKey, endpoint)
}

// Announce is a paid mutator transaction binding the contract method 0x842929d5.
//
// Solidity: function announce(bytes32 pubKey, string endpoint) returns()
func (_ResolverRegistry *ResolverRegistryTransactorSession) Announce(pubKey [32]byte, endpoint string) (*types.Transaction, error) {
	return _ResolverRegistry.Contract.Announce(&_ResolverRegistry.TransactOpts, pubKey, endpoint)
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_ResolverRegistry *ResolverRegistryTransactor) Revoke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverRegistry.contract.Transact(opts, "revoke")
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_ResolverRegistry *ResolverRegistrySession) Revoke() (*types.Transaction, error) {
	return _ResolverRegistry.Contract.Revoke(&_ResolverRegistry.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_ResolverRegistry *ResolverRegistryTransactorSession) Revoke() (*types.Transaction, error) {
	return _ResolverRegistry.Contract.Revoke(&_ResolverRegistry.TransactOpts)
}

// ResolverRegistryResolverAnnouncedIterator is returned from FilterResolverAnnounced and is used to iterate over the raw logs and unpacked data for ResolverAnnounced events raised by the ResolverRegistry contract.
type ResolverRegistryResolverAnnouncedIterator struct {
	Event *ResolverRegistryResolverAnnounced // Event containing the contract specifics and raw log

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
func (it *ResolverRegistryResolverAnnouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverRegistryResolverAnnounced)
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
		it.Event = new(ResolverRegistryResolverAnnounced)
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
func (it *ResolverRegistryResolverAnnouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverRegistryResolverAnnouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverRegistryResolverAnnounced represents a ResolverAnnounced event raised by the ResolverRegistry contract.
type ResolverRegistryResolverAnnounced struct {
	Operator common.Address
	PubKey   [32]byte
	Endpoint string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterResolverAnnounced is a free log retrieval operation binding the contract event 0x2af369faa85a14ab1ea1ae248aca2b45493d2e3e197e4ce1adbb32208f79a1c3.
//
// Solidity: event ResolverAnnounced(address indexed operator, bytes32 pubKey, string endpoint)
func (_ResolverRegistry *ResolverRegistryFilterer) FilterResolverAnnounced(opts *bind.FilterOpts, operator []common.Address) (*ResolverRegistryResolverAnnouncedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ResolverRegistry.contract.FilterLogs(opts, "ResolverAnnounced", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ResolverRegistryResolverAnnouncedIterator{contract: _ResolverRegistry.contract, event: "ResolverAnnounced", logs: logs, sub: sub}, nil
}

// WatchResolverAnnounced is a free log subscription operation binding the contract event 0x2af369faa85a14ab1ea1ae248aca2b45493d2e3e197e4ce1adbb32208f79a1c3.
//
// Solidity: event ResolverAnnounced(address indexed operator, bytes32 pubKey, string endpoint)
func (_ResolverRegistry *ResolverRegistryFilterer) WatchResolverAnnounced(opts *bind.WatchOpts, sink chan<- *ResolverRegistryResolverAnnounced, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ResolverRegistry.contract.WatchLogs(opts, "ResolverAnnounced", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverRegistryResolverAnnounced)
				if err := _ResolverRegistry.contract.UnpackLog(event, "ResolverAnnounced", log); err != nil {
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

// ParseResolverAnnounced is a log parse operation binding the contract event 0x2af369faa85a14ab1ea1ae248aca2b45493d2e3e197e4ce1adbb32208f79a1c3.
//
// Solidity: event ResolverAnnounced(address indexed operator, bytes32 pubKey, string endpoint)
func (_ResolverRegistry *ResolverRegistryFilterer) ParseResolverAnnounced(log types.Log) (*ResolverRegistryResolverAnnounced, error) {
	event := new(ResolverRegistryResolverAnnounced)
	if err := _ResolverRegistry.contract.UnpackLog(event, "ResolverAnnounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverRegistryResolverRevokedIterator is returned from FilterResolverRevoked and is used to iterate over the raw logs and unpacked data for ResolverRevoked events raised by the ResolverRegistry contract.
type ResolverRegistryResolverRevokedIterator struct {
	Event *ResolverRegistryResolverRevoked // Event containing the contract specifics and raw log

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
func (it *ResolverRegistryResolverRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverRegistryResolverRevoked)
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
		it.Event = new(ResolverRegistryResolverRevoked)
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
func (it *ResolverRegistryResolverRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverRegistryResolverRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverRegistryResolverRevoked represents a ResolverRevoked event raised by the ResolverRegistry contract.
type ResolverRegistryResolverRevoked struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterResolverRevoked is a free log retrieval operation binding the contract event 0x6cc76de175eb7c3679f8d4c3529e23883784a40e1f9d7f54e8a9f97115aab347.
//
// Solidity: event ResolverRevoked(address indexed operator)
func (_ResolverRegistry *ResolverRegistryFilterer) FilterResolverRevoked(opts *bind.FilterOpts, operator []common.Address) (*ResolverRegistryResolverRevokedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ResolverRegistry.contract.FilterLogs(opts, "ResolverRevoked", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ResolverRegistryResolverRevokedIterator{contract: _ResolverRegistry.contract, event: "ResolverRevoked", logs: logs, sub: sub}, nil
}

// WatchResolverRevoked is a free log subscription operation binding the contract event 0x6cc76de175eb7c3679f8d4c3529e23883784a40e1f9d7f54e8a9f97115aab347.
//
// Solidity: event ResolverRevoked(address indexed operator)
func (_ResolverRegistry *ResolverRegistryFilterer) WatchResolverRevoked(opts *bind.WatchOpts, sink chan<- *ResolverRegistryResolverRevoked, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ResolverRegistry.contract.WatchLogs(opts, "ResolverRevoked", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverRegistryResolverRevoked)
				if err := _ResolverRegistry.contract.UnpackLog(event, "ResolverRevoked", log); err != nil {
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

// ParseResolverRevoked is a log parse operation binding the contract event 0x6cc76de175eb7c3679f8d4c3529e23883784a40e1f9d7f54e8a9f97115aab347.
//
// Solidity: event ResolverRevoked(address indexed operator)
func (_ResolverRegistry *ResolverRegistryFilterer) ParseResolverRevoked(log types.Log) (*ResolverRegistryResolverRevoked, error) {
	event := new(ResolverRegistryResolverRevoked)
	if err := _ResolverRegistry.contract.UnpackLog(event, "ResolverRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
