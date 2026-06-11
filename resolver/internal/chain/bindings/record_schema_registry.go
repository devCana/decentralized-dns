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

// RecordSchemaRegistryFieldSpec is an auto generated low-level Go binding around an user-defined struct.
type RecordSchemaRegistryFieldSpec struct {
	Name      string
	Mandatory bool
}

// RecordSchemaRegistryMetaData contains all meta data concerning the RecordSchemaRegistry contract.
var RecordSchemaRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DuplicateField\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFieldName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTypeName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyFields\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TypeAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"declarer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mandatoryCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optionalCount\",\"type\":\"uint256\"}],\"name\":\"TypeDeclared\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_FIELDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FIELD_NAME_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TYPE_NAME_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"mandatory\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"optional\",\"type\":\"string[]\"}],\"name\":\"declareType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"}],\"name\":\"getSchema\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"mandatory\",\"type\":\"bool\"}],\"internalType\":\"structRecordSchemaRegistry.FieldSpec[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listTypes\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"}],\"name\":\"mandatoryFields\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"}],\"name\":\"typeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052346103555761007560405161001a604082610375565b60078152666164647265737360c81b6020820152610036610398565b9061004082610407565b5261004a81610407565b506100536103d3565b90604051610062604082610375565b60018152604160f81b602082015261061f565b6100e4604051610086604082610375565b60078152666164647265737360c81b60208201526100a2610398565b906100ac82610407565b526100b681610407565b506100bf6103d3565b906040516100ce604082610375565b60048152634141414160e01b602082015261061f565b6040516100f2604082610375565b60048152631a1bdcdd60e21b6020820152604051610111604082610375565b60088152677072696f7269747960c01b60208201526040519060606101368184610375565b60028352601f19019260005b84811061032e57506101a2935061015883610407565b5261016282610407565b5061016c8261042a565b526101768161042a565b5061017f6103d3565b9060405161018e604082610375565b600281526109ab60f31b602082015261061f565b6102756102156040516101b6604082610375565b60068152651d185c99d95d60d21b60208201526040516101d7604082610375565b60078152667365727669636560c81b6020820152604051916101fa604084610375565b60098352681d1c985b9cdc1bdc9d60ba1b60208401526109ad565b60405190610224604083610375565b60048252631c1bdc9d60e21b602083015261023d610398565b9161024783610407565b5261025182610407565b50604051610260604082610375565b600381526253564360e81b602082015261061f565b61031f6102eb604051610289604082610375565b60088152670d2dcccde90c2e6d60c31b60208201526040516102ac604082610375565b600681526539b430991a9b60d11b6020820152604051916102ce604084610375565b600b83526a636f6e74656e745479706560a81b60208401526109ad565b6102f36103d3565b90604051610302604082610375565b600b81526a2932b9b7bab931b2a932b360a91b602082015261061f565b604051610d349081610af98239f35b806060602080938701015201610142565b634e487b7160e01b600052604160045260246000fd5b600080fd5b604081019081106001600160401b0382111761033f57604052565b601f909101601f19168101906001600160401b0382119082101761033f57604052565b604080519091906103a98382610375565b6001815291601f19018260005b8281106103c257505050565b8060606020809385010152016103b6565b604051906103e2602083610375565b600080835282815b8281106103f657505050565b8060606020809385010152016103ea565b8051156104145760200190565b634e487b7160e01b600052603260045260246000fd5b8051600110156104145760400190565b8051600210156104145760600190565b80518210156104145760209160051b010190565b90600182811c9216801561048e575b602083101461047857565b634e487b7160e01b600052602260045260246000fd5b91607f169161046d565b81519192916001600160401b03811161033f576104b5825461045e565b601f811161055b575b506020601f82116001146104f957819293946000926104ee575b50508160011b916000199060031b1c1916179055565b0151905038806104d8565b601f1982169083600052806000209160005b8181106105435750958360019596971061052a575b505050811b019055565b015160001960f88460031b161c19169055388080610520565b9192602060018192868b01518155019401920161050b565b826000526020600020601f830160051c81019160208410610599575b601f0160051c01905b81811061058d57506104be565b60008155600101610580565b9091508190610577565b80548210156104145760005260206000209060011b0190600090565b80546801000000000000000081101561033f576105e1916001820181556105a3565b6106095760016020916105f5845182610498565b01910151151560ff80198354169116179055565b634e487b7160e01b600052600060045260246000fd5b919290825180159081156109a2575b506109925760005b835181101561070957600060ff60f81b6020838701015116606160f81b81101590816106fa575b81156106d6575b81156106b2575b81156106a4575b8115610696575b50156106885750600101610636565b6228eb7360e41b8152600490fd5b605f60f81b14905038610679565b602d60f81b81149150610672565b9050600360fc1b811015806106c8575b9061066b565b50603960f81b8111156106c2565b9050604160f81b811015806106ec575b90610664565b50602d60f91b8111156106e6565b603d60f91b811115915061065d565b5090919281518151810180911161095b576010106109815783516020850120938460005260006020526040600020936001850160ff8154166109715761075183879897610498565b600160ff19825416179055600260009601955b84518110156107b2578061078361077d6001938861044a565b51610a26565b6107ac610790828861044a565b516040519061079e8261035a565b8152836020820152896105bf565b01610764565b509091929360005b845181101561080357806107d361077d6001938861044a565b6107fd6107e0828861044a565b51604051906107ee8261035a565b815260006020820152896105bf565b016107ba565b5091949390936000908054915b8281106108f1575050506001546801000000000000000081101561033f57600181018060015581101561041457600160005261086f9086907fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601610498565b519151604051906060825285519384606084015260005b8581106108db575084959650600060807faa568449c03afa7bbe1380b06c93098d804250007a14fa1f2effb173922b102e959685010152602083015260408201526080813395601f80199101168101030190a3565b80602080928a0101516080828701015201610886565b6109046108fe82846105a3565b50610a54565b602081519101206001820180831161095b575b848110610928575050600101610810565b6109356108fe82866105a3565b60208151910120821461094a57600101610917565b638516890360e01b60005260046000fd5b634e487b7160e01b600052601160045260246000fd5b62a0b14560e31b60005260046000fd5b6304d8f29f60e31b60005260046000fd5b6228eb7360e41b60005260046000fd5b60209150113861062e565b92916040519160806109bf8185610375565b60038452601f190160005b818110610a15575050908291610a1293956109e484610407565b526109ee83610407565b506109f88361042a565b52610a028261042a565b50610a0c8261043a565b5261043a565b50565b8060606020809388010152016109ca565b518015908115610a49575b50610a3857565b630ddff6f360e41b60005260046000fd5b602091501138610a31565b9060405191826000825492610a688461045e565b8084529360018116908115610ad65750600114610a8f575b50610a8d92500383610375565b565b90506000929192526020600020906000915b818310610aba575050906020610a8d9282010138610a80565b6020919350806001915483858901015201910190918492610aa1565b905060209250610a8d94915060ff191682840152151560051b82010138610a8056fe6080604052600436101561001257600080fd5b60003560e01c80632903c0f314610841578063311cec651461082557806346a508bb146106f45780638cde2d9a146102dc5780639ea7fbbb146102d7578063cbb67e581461023d578063d117e71d146101e35763f1a03f1b1461007457600080fd5b346101de5760203660031901126101de576004356001600160401b0381116101de576100a76100ae9136906004016108fd565b3691610aa9565b602081519101206000526000602052604060002060ff600182015416156101cd576002018054906100de826109c1565b916100ec60405193846109a0565b80835260208301809260005260206000206000915b83831061018557848660405191829160208301906020845251809152604083019060408160051b85010192916000905b82821061014057505050500390f35b919360019193955060208091603f19898203018552875190828061016d845160408552604085019061085c565b93015115159101529601920192018594939192610131565b6002602060019260405161019881610985565b6040516101b0816101a9818a610a26565b03826109a0565b815260ff8587015416151583820152815201920192019190610101565b630d5c6e8f60e01b60005260046000fd5b600080fd5b346101de5760203660031901126101de576004356001600160401b0381116101de576100a76102169136906004016108fd565b602081519101206000526000602052602060ff600160406000200154166040519015158152f35b346101de5760003660031901126101de5760015461025a816109c1565b9061026860405192836109a0565b8082526020820160016000527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf66000915b8383106102b257604051806102ae878261089d565b0390f35b6001602081926040516102c9816101a98189610a26565b815201920192019190610299565b610841565b346101de5760603660031901126101de576004356001600160401b0381116101de5761030c9036906004016108fd565b6024356001600160401b0381116101de5761032b90369060040161092a565b916044356001600160401b0381116101de5761036f9261035f61035561036793369060040161092a565b9790943691610aa9565b943691610aef565b933691610aef565b91815180159081156106e9575b506106d95760005b825181101561045757600060ff60f81b6020838601015116606160f81b8110159081610448575b8115610424575b8115610400575b81156103f2575b81156103e4575b50156103d65750600101610384565b6228eb7360e41b8152600490fd5b605f60f81b149050866103c7565b602d60f81b811491506103c0565b9050600360fc1b81101580610416575b906103b9565b50603960f81b811115610410565b9050604160f81b8110158061043a575b906103b2565b50602d60f91b811115610434565b603d60f91b81111591506103ab565b50918251815181018091116106a2576010106106c85781516020830120908160005260006020526040600020926001840160ff8154166106b85761049d82869596610b6a565b600160ff19825416179055600260009301925b85518110156104fe57806104cf6104c9600193896109d8565b51610cd0565b6104f86104dc82896109d8565b51604051906104ea82610985565b815283602082015286610c75565b016104b0565b5092909360005b855181101561054e578061051e6104c9600193896109d8565b61054861052b82896109d8565b516040519061053982610985565b81526000602082015286610c75565b01610505565b508385936000908054915b82811061062b5750505060015491600160401b8310156106155760018301806001558310156105ff576105d4827faa568449c03afa7bbe1380b06c93098d804250007a14fa1f2effb173922b102e9460016000527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601610b6a565b5193516105ec6040519260608452606084019061085c565b94602083015260408201528033940390a3005b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6101a961064861063b838561095a565b5060405192838092610a26565b60208151910120600182018083116106a2575b84811061066c575050600101610559565b6101a961067c61063b838761095a565b6020815191012082146106915760010161065b565b638516890360e01b60005260046000fd5b634e487b7160e01b600052601160045260246000fd5b62a0b14560e31b60005260046000fd5b6304d8f29f60e31b60005260046000fd5b6228eb7360e41b60005260046000fd5b60209150118461037c565b346101de5760203660031901126101de576004356000526000602052604060002060ff600182015416156101cd576002018054600090815b8181106107f2575061073d826109c1565b9161074b60405193846109a0565b80835261075a601f19916109c1565b0160005b8181106107e15750506000805b82811061078057604051806102ae868261089d565b8060ff600161079081948961095a565b5001541661079f575b0161076b565b6107db6107ac828861095a565b50936101a96107ca6107bd83610976565b9660405192838092610a26565b6107d482896109d8565b52866109d8565b50610799565b80606060208093870101520161075e565b60ff6001610800838761095a565b50015416610811575b60010161072c565b9161081d600191610976565b929050610809565b346101de5760003660031901126101de57602060405160108152f35b346101de5760003660031901126101de576020604051818152f35b919082519283825260005b848110610888575050826000602080949584010152601f8019910116010190565b80602080928401015182828601015201610867565b602081016020825282518091526040820191602060408360051b8301019401926000915b8383106108d057505050505090565b90919293946020806108ee600193603f19868203018752895161085c565b970193019301919392906108c1565b9181601f840112156101de578235916001600160401b0383116101de57602083818601950101116101de57565b9181601f840112156101de578235916001600160401b0383116101de576020808501948460051b0101116101de57565b80548210156105ff5760005260206000209060011b0190600090565b60001981146106a25760010190565b604081019081106001600160401b0382111761061557604052565b90601f801991011681019081106001600160401b0382111761061557604052565b6001600160401b0381116106155760051b60200190565b80518210156105ff5760209160051b010190565b90600182811c92168015610a1c575b6020831014610a0657565b634e487b7160e01b600052602260045260246000fd5b91607f16916109fb565b60009291815491610a36836109ec565b8083529260018116908115610a8c5750600114610a5257505050565b60009081526020812093945091925b838310610a72575060209250010190565b600181602092949394548385870101520191019190610a61565b915050602093945060ff929192191683830152151560051b010190565b9291926001600160401b0382116106155760405191610ad2601f8201601f1916602001846109a0565b8294818452818301116101de578281602093846000960137010152565b929192610afb826109c1565b93610b0960405195866109a0565b602085848152019260051b8201918183116101de5780935b838510610b2f575050505050565b84356001600160401b0381116101de57820183601f820112156101de57602091610b5f8583858095359101610aa9565b815201940193610b21565b91909182516001600160401b03811161061557610b8782546109ec565b601f8111610c2d575b506020601f8211600114610bcb5781929394600092610bc0575b50508160011b916000199060031b1c1916179055565b015190503880610baa565b601f1982169083600052806000209160005b818110610c1557509583600195969710610bfc575b505050811b019055565b015160001960f88460031b161c19169055388080610bf2565b9192602060018192868b015181550194019201610bdd565b826000526020600020601f830160051c81019160208410610c6b575b601f0160051c01905b818110610c5f5750610b90565b60008155600101610c52565b9091508190610c49565b8054600160401b81101561061557610c929160018201815561095a565b610cba576001602091610ca6845182610b6a565b01910151151560ff80198354169116179055565b634e487b7160e01b600052600060045260246000fd5b518015908115610cf3575b50610ce257565b630ddff6f360e41b60005260046000fd5b602091501138610cdb56fea2646970667358221220c66085b22b9a8266145835398f2f9f7ae7c31782968a3735a86c119f19e535eb64736f6c634300081c0033",
}

// RecordSchemaRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RecordSchemaRegistryMetaData.ABI instead.
var RecordSchemaRegistryABI = RecordSchemaRegistryMetaData.ABI

// RecordSchemaRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RecordSchemaRegistryMetaData.Bin instead.
var RecordSchemaRegistryBin = RecordSchemaRegistryMetaData.Bin

// DeployRecordSchemaRegistry deploys a new Ethereum contract, binding an instance of RecordSchemaRegistry to it.
func DeployRecordSchemaRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RecordSchemaRegistry, error) {
	parsed, err := RecordSchemaRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecordSchemaRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RecordSchemaRegistry{RecordSchemaRegistryCaller: RecordSchemaRegistryCaller{contract: contract}, RecordSchemaRegistryTransactor: RecordSchemaRegistryTransactor{contract: contract}, RecordSchemaRegistryFilterer: RecordSchemaRegistryFilterer{contract: contract}}, nil
}

// RecordSchemaRegistry is an auto generated Go binding around an Ethereum contract.
type RecordSchemaRegistry struct {
	RecordSchemaRegistryCaller     // Read-only binding to the contract
	RecordSchemaRegistryTransactor // Write-only binding to the contract
	RecordSchemaRegistryFilterer   // Log filterer for contract events
}

// RecordSchemaRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecordSchemaRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordSchemaRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecordSchemaRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordSchemaRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecordSchemaRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordSchemaRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecordSchemaRegistrySession struct {
	Contract     *RecordSchemaRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RecordSchemaRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecordSchemaRegistryCallerSession struct {
	Contract *RecordSchemaRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// RecordSchemaRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecordSchemaRegistryTransactorSession struct {
	Contract     *RecordSchemaRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// RecordSchemaRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecordSchemaRegistryRaw struct {
	Contract *RecordSchemaRegistry // Generic contract binding to access the raw methods on
}

// RecordSchemaRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecordSchemaRegistryCallerRaw struct {
	Contract *RecordSchemaRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RecordSchemaRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecordSchemaRegistryTransactorRaw struct {
	Contract *RecordSchemaRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecordSchemaRegistry creates a new instance of RecordSchemaRegistry, bound to a specific deployed contract.
func NewRecordSchemaRegistry(address common.Address, backend bind.ContractBackend) (*RecordSchemaRegistry, error) {
	contract, err := bindRecordSchemaRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RecordSchemaRegistry{RecordSchemaRegistryCaller: RecordSchemaRegistryCaller{contract: contract}, RecordSchemaRegistryTransactor: RecordSchemaRegistryTransactor{contract: contract}, RecordSchemaRegistryFilterer: RecordSchemaRegistryFilterer{contract: contract}}, nil
}

// NewRecordSchemaRegistryCaller creates a new read-only instance of RecordSchemaRegistry, bound to a specific deployed contract.
func NewRecordSchemaRegistryCaller(address common.Address, caller bind.ContractCaller) (*RecordSchemaRegistryCaller, error) {
	contract, err := bindRecordSchemaRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecordSchemaRegistryCaller{contract: contract}, nil
}

// NewRecordSchemaRegistryTransactor creates a new write-only instance of RecordSchemaRegistry, bound to a specific deployed contract.
func NewRecordSchemaRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RecordSchemaRegistryTransactor, error) {
	contract, err := bindRecordSchemaRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecordSchemaRegistryTransactor{contract: contract}, nil
}

// NewRecordSchemaRegistryFilterer creates a new log filterer instance of RecordSchemaRegistry, bound to a specific deployed contract.
func NewRecordSchemaRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RecordSchemaRegistryFilterer, error) {
	contract, err := bindRecordSchemaRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecordSchemaRegistryFilterer{contract: contract}, nil
}

// bindRecordSchemaRegistry binds a generic wrapper to an already deployed contract.
func bindRecordSchemaRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RecordSchemaRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecordSchemaRegistry *RecordSchemaRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecordSchemaRegistry.Contract.RecordSchemaRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecordSchemaRegistry *RecordSchemaRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecordSchemaRegistry.Contract.RecordSchemaRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecordSchemaRegistry *RecordSchemaRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecordSchemaRegistry.Contract.RecordSchemaRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecordSchemaRegistry *RecordSchemaRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecordSchemaRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecordSchemaRegistry *RecordSchemaRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecordSchemaRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecordSchemaRegistry *RecordSchemaRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecordSchemaRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXFIELDS is a free data retrieval call binding the contract method 0x311cec65.
//
// Solidity: function MAX_FIELDS() view returns(uint256)
func (_RecordSchemaRegistry *RecordSchemaRegistryCaller) MAXFIELDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RecordSchemaRegistry.contract.Call(opts, &out, "MAX_FIELDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFIELDS is a free data retrieval call binding the contract method 0x311cec65.
//
// Solidity: function MAX_FIELDS() view returns(uint256)
func (_RecordSchemaRegistry *RecordSchemaRegistrySession) MAXFIELDS() (*big.Int, error) {
	return _RecordSchemaRegistry.Contract.MAXFIELDS(&_RecordSchemaRegistry.CallOpts)
}

// MAXFIELDS is a free data retrieval call binding the contract method 0x311cec65.
//
// Solidity: function MAX_FIELDS() view returns(uint256)
func (_RecordSchemaRegistry *RecordSchemaRegistryCallerSession) MAXFIELDS() (*big.Int, error) {
	return _RecordSchemaRegistry.Contract.MAXFIELDS(&_RecordSchemaRegistry.CallOpts)
}

// MAXFIELDNAMELENGTH is a free data retrieval call binding the contract method 0x2903c0f3.
//
// Solidity: function MAX_FIELD_NAME_LENGTH() view returns(uint256)
func (_RecordSchemaRegistry *RecordSchemaRegistryCaller) MAXFIELDNAMELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RecordSchemaRegistry.contract.Call(opts, &out, "MAX_FIELD_NAME_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFIELDNAMELENGTH is a free data retrieval call binding the contract method 0x2903c0f3.
//
// Solidity: function MAX_FIELD_NAME_LENGTH() view returns(uint256)
func (_RecordSchemaRegistry *RecordSchemaRegistrySession) MAXFIELDNAMELENGTH() (*big.Int, error) {
	return _RecordSchemaRegistry.Contract.MAXFIELDNAMELENGTH(&_RecordSchemaRegistry.CallOpts)
}

// MAXFIELDNAMELENGTH is a free data retrieval call binding the contract method 0x2903c0f3.
//
// Solidity: function MAX_FIELD_NAME_LENGTH() view returns(uint256)
func (_RecordSchemaRegistry *RecordSchemaRegistryCallerSession) MAXFIELDNAMELENGTH() (*big.Int, error) {
	return _RecordSchemaRegistry.Contract.MAXFIELDNAMELENGTH(&_RecordSchemaRegistry.CallOpts)
}

// MAXTYPENAMELENGTH is a free data retrieval call binding the contract method 0x9ea7fbbb.
//
// Solidity: function MAX_TYPE_NAME_LENGTH() view returns(uint256)
func (_RecordSchemaRegistry *RecordSchemaRegistryCaller) MAXTYPENAMELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RecordSchemaRegistry.contract.Call(opts, &out, "MAX_TYPE_NAME_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTYPENAMELENGTH is a free data retrieval call binding the contract method 0x9ea7fbbb.
//
// Solidity: function MAX_TYPE_NAME_LENGTH() view returns(uint256)
func (_RecordSchemaRegistry *RecordSchemaRegistrySession) MAXTYPENAMELENGTH() (*big.Int, error) {
	return _RecordSchemaRegistry.Contract.MAXTYPENAMELENGTH(&_RecordSchemaRegistry.CallOpts)
}

// MAXTYPENAMELENGTH is a free data retrieval call binding the contract method 0x9ea7fbbb.
//
// Solidity: function MAX_TYPE_NAME_LENGTH() view returns(uint256)
func (_RecordSchemaRegistry *RecordSchemaRegistryCallerSession) MAXTYPENAMELENGTH() (*big.Int, error) {
	return _RecordSchemaRegistry.Contract.MAXTYPENAMELENGTH(&_RecordSchemaRegistry.CallOpts)
}

// GetSchema is a free data retrieval call binding the contract method 0xf1a03f1b.
//
// Solidity: function getSchema(string typeName) view returns((string,bool)[])
func (_RecordSchemaRegistry *RecordSchemaRegistryCaller) GetSchema(opts *bind.CallOpts, typeName string) ([]RecordSchemaRegistryFieldSpec, error) {
	var out []interface{}
	err := _RecordSchemaRegistry.contract.Call(opts, &out, "getSchema", typeName)

	if err != nil {
		return *new([]RecordSchemaRegistryFieldSpec), err
	}

	out0 := *abi.ConvertType(out[0], new([]RecordSchemaRegistryFieldSpec)).(*[]RecordSchemaRegistryFieldSpec)

	return out0, err

}

// GetSchema is a free data retrieval call binding the contract method 0xf1a03f1b.
//
// Solidity: function getSchema(string typeName) view returns((string,bool)[])
func (_RecordSchemaRegistry *RecordSchemaRegistrySession) GetSchema(typeName string) ([]RecordSchemaRegistryFieldSpec, error) {
	return _RecordSchemaRegistry.Contract.GetSchema(&_RecordSchemaRegistry.CallOpts, typeName)
}

// GetSchema is a free data retrieval call binding the contract method 0xf1a03f1b.
//
// Solidity: function getSchema(string typeName) view returns((string,bool)[])
func (_RecordSchemaRegistry *RecordSchemaRegistryCallerSession) GetSchema(typeName string) ([]RecordSchemaRegistryFieldSpec, error) {
	return _RecordSchemaRegistry.Contract.GetSchema(&_RecordSchemaRegistry.CallOpts, typeName)
}

// ListTypes is a free data retrieval call binding the contract method 0xcbb67e58.
//
// Solidity: function listTypes() view returns(string[])
func (_RecordSchemaRegistry *RecordSchemaRegistryCaller) ListTypes(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _RecordSchemaRegistry.contract.Call(opts, &out, "listTypes")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ListTypes is a free data retrieval call binding the contract method 0xcbb67e58.
//
// Solidity: function listTypes() view returns(string[])
func (_RecordSchemaRegistry *RecordSchemaRegistrySession) ListTypes() ([]string, error) {
	return _RecordSchemaRegistry.Contract.ListTypes(&_RecordSchemaRegistry.CallOpts)
}

// ListTypes is a free data retrieval call binding the contract method 0xcbb67e58.
//
// Solidity: function listTypes() view returns(string[])
func (_RecordSchemaRegistry *RecordSchemaRegistryCallerSession) ListTypes() ([]string, error) {
	return _RecordSchemaRegistry.Contract.ListTypes(&_RecordSchemaRegistry.CallOpts)
}

// MandatoryFields is a free data retrieval call binding the contract method 0x46a508bb.
//
// Solidity: function mandatoryFields(bytes32 typeHash) view returns(string[])
func (_RecordSchemaRegistry *RecordSchemaRegistryCaller) MandatoryFields(opts *bind.CallOpts, typeHash [32]byte) ([]string, error) {
	var out []interface{}
	err := _RecordSchemaRegistry.contract.Call(opts, &out, "mandatoryFields", typeHash)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// MandatoryFields is a free data retrieval call binding the contract method 0x46a508bb.
//
// Solidity: function mandatoryFields(bytes32 typeHash) view returns(string[])
func (_RecordSchemaRegistry *RecordSchemaRegistrySession) MandatoryFields(typeHash [32]byte) ([]string, error) {
	return _RecordSchemaRegistry.Contract.MandatoryFields(&_RecordSchemaRegistry.CallOpts, typeHash)
}

// MandatoryFields is a free data retrieval call binding the contract method 0x46a508bb.
//
// Solidity: function mandatoryFields(bytes32 typeHash) view returns(string[])
func (_RecordSchemaRegistry *RecordSchemaRegistryCallerSession) MandatoryFields(typeHash [32]byte) ([]string, error) {
	return _RecordSchemaRegistry.Contract.MandatoryFields(&_RecordSchemaRegistry.CallOpts, typeHash)
}

// TypeExists is a free data retrieval call binding the contract method 0xd117e71d.
//
// Solidity: function typeExists(string typeName) view returns(bool)
func (_RecordSchemaRegistry *RecordSchemaRegistryCaller) TypeExists(opts *bind.CallOpts, typeName string) (bool, error) {
	var out []interface{}
	err := _RecordSchemaRegistry.contract.Call(opts, &out, "typeExists", typeName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TypeExists is a free data retrieval call binding the contract method 0xd117e71d.
//
// Solidity: function typeExists(string typeName) view returns(bool)
func (_RecordSchemaRegistry *RecordSchemaRegistrySession) TypeExists(typeName string) (bool, error) {
	return _RecordSchemaRegistry.Contract.TypeExists(&_RecordSchemaRegistry.CallOpts, typeName)
}

// TypeExists is a free data retrieval call binding the contract method 0xd117e71d.
//
// Solidity: function typeExists(string typeName) view returns(bool)
func (_RecordSchemaRegistry *RecordSchemaRegistryCallerSession) TypeExists(typeName string) (bool, error) {
	return _RecordSchemaRegistry.Contract.TypeExists(&_RecordSchemaRegistry.CallOpts, typeName)
}

// DeclareType is a paid mutator transaction binding the contract method 0x8cde2d9a.
//
// Solidity: function declareType(string name, string[] mandatory, string[] optional) returns()
func (_RecordSchemaRegistry *RecordSchemaRegistryTransactor) DeclareType(opts *bind.TransactOpts, name string, mandatory []string, optional []string) (*types.Transaction, error) {
	return _RecordSchemaRegistry.contract.Transact(opts, "declareType", name, mandatory, optional)
}

// DeclareType is a paid mutator transaction binding the contract method 0x8cde2d9a.
//
// Solidity: function declareType(string name, string[] mandatory, string[] optional) returns()
func (_RecordSchemaRegistry *RecordSchemaRegistrySession) DeclareType(name string, mandatory []string, optional []string) (*types.Transaction, error) {
	return _RecordSchemaRegistry.Contract.DeclareType(&_RecordSchemaRegistry.TransactOpts, name, mandatory, optional)
}

// DeclareType is a paid mutator transaction binding the contract method 0x8cde2d9a.
//
// Solidity: function declareType(string name, string[] mandatory, string[] optional) returns()
func (_RecordSchemaRegistry *RecordSchemaRegistryTransactorSession) DeclareType(name string, mandatory []string, optional []string) (*types.Transaction, error) {
	return _RecordSchemaRegistry.Contract.DeclareType(&_RecordSchemaRegistry.TransactOpts, name, mandatory, optional)
}

// RecordSchemaRegistryTypeDeclaredIterator is returned from FilterTypeDeclared and is used to iterate over the raw logs and unpacked data for TypeDeclared events raised by the RecordSchemaRegistry contract.
type RecordSchemaRegistryTypeDeclaredIterator struct {
	Event *RecordSchemaRegistryTypeDeclared // Event containing the contract specifics and raw log

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
func (it *RecordSchemaRegistryTypeDeclaredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordSchemaRegistryTypeDeclared)
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
		it.Event = new(RecordSchemaRegistryTypeDeclared)
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
func (it *RecordSchemaRegistryTypeDeclaredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordSchemaRegistryTypeDeclaredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordSchemaRegistryTypeDeclared represents a TypeDeclared event raised by the RecordSchemaRegistry contract.
type RecordSchemaRegistryTypeDeclared struct {
	TypeHash       [32]byte
	Name           string
	Declarer       common.Address
	MandatoryCount *big.Int
	OptionalCount  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTypeDeclared is a free log retrieval operation binding the contract event 0xaa568449c03afa7bbe1380b06c93098d804250007a14fa1f2effb173922b102e.
//
// Solidity: event TypeDeclared(bytes32 indexed typeHash, string name, address indexed declarer, uint256 mandatoryCount, uint256 optionalCount)
func (_RecordSchemaRegistry *RecordSchemaRegistryFilterer) FilterTypeDeclared(opts *bind.FilterOpts, typeHash [][32]byte, declarer []common.Address) (*RecordSchemaRegistryTypeDeclaredIterator, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}

	var declarerRule []interface{}
	for _, declarerItem := range declarer {
		declarerRule = append(declarerRule, declarerItem)
	}

	logs, sub, err := _RecordSchemaRegistry.contract.FilterLogs(opts, "TypeDeclared", typeHashRule, declarerRule)
	if err != nil {
		return nil, err
	}
	return &RecordSchemaRegistryTypeDeclaredIterator{contract: _RecordSchemaRegistry.contract, event: "TypeDeclared", logs: logs, sub: sub}, nil
}

// WatchTypeDeclared is a free log subscription operation binding the contract event 0xaa568449c03afa7bbe1380b06c93098d804250007a14fa1f2effb173922b102e.
//
// Solidity: event TypeDeclared(bytes32 indexed typeHash, string name, address indexed declarer, uint256 mandatoryCount, uint256 optionalCount)
func (_RecordSchemaRegistry *RecordSchemaRegistryFilterer) WatchTypeDeclared(opts *bind.WatchOpts, sink chan<- *RecordSchemaRegistryTypeDeclared, typeHash [][32]byte, declarer []common.Address) (event.Subscription, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}

	var declarerRule []interface{}
	for _, declarerItem := range declarer {
		declarerRule = append(declarerRule, declarerItem)
	}

	logs, sub, err := _RecordSchemaRegistry.contract.WatchLogs(opts, "TypeDeclared", typeHashRule, declarerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordSchemaRegistryTypeDeclared)
				if err := _RecordSchemaRegistry.contract.UnpackLog(event, "TypeDeclared", log); err != nil {
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

// ParseTypeDeclared is a log parse operation binding the contract event 0xaa568449c03afa7bbe1380b06c93098d804250007a14fa1f2effb173922b102e.
//
// Solidity: event TypeDeclared(bytes32 indexed typeHash, string name, address indexed declarer, uint256 mandatoryCount, uint256 optionalCount)
func (_RecordSchemaRegistry *RecordSchemaRegistryFilterer) ParseTypeDeclared(log types.Log) (*RecordSchemaRegistryTypeDeclared, error) {
	event := new(RecordSchemaRegistryTypeDeclared)
	if err := _RecordSchemaRegistry.contract.UnpackLog(event, "TypeDeclared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
