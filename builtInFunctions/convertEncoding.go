package builtInFunctions

import (
	"math/big"
	"sync"

	"github.com/multiversx/mx-chain-core-evm-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
	"github.com/multiversx/mx-chain-vm-common-go/builtInFunctions"

	evmcommon "github.com/multiversx/mx-chain-vm-common-evm-go"
	"github.com/multiversx/mx-chain-vm-common-evm-go/builtInFunctions/convertEncoding"
)

const (
	minArgsCount      = 2
	signaturePosition = 0
	dataStartPosition = 1
)

type ConvertEncoding struct {
	builtInFunctions.BaseAlwaysActiveHandler
	baseOperationCost vmcommon.BaseOperationCost
	function          string
	mutExecution      sync.RWMutex
	accounts          vmcommon.AccountsAdapter
}

// NewEVMConvertEncodingFunc converts the arguments using the specified configuration for the encoding
func NewEVMConvertEncodingFunc(
	baseOperationCost vmcommon.BaseOperationCost,
	accounts vmcommon.AccountsAdapter,
	function string,
) (*ConvertEncoding, error) {
	if check.IfNil(accounts) {
		return nil, builtInFunctions.ErrNilAccountsAdapter
	}

	return &ConvertEncoding{
		baseOperationCost: baseOperationCost,
		function:          function,
		mutExecution:      sync.RWMutex{},
		accounts:          accounts,
	}, nil
}

// SetNewGasConfig is called whenever gas cost is changed
func (ce *ConvertEncoding) SetNewGasConfig(gasCost *vmcommon.GasCost) {
	if gasCost == nil {
		return
	}

	ce.mutExecution.Lock()
	ce.baseOperationCost = gasCost.BaseOperationCost
	ce.mutExecution.Unlock()
}

// ProcessBuiltinFunction resolves the convert encoding function call
func (ce *ConvertEncoding) ProcessBuiltinFunction(
	_, _ vmcommon.UserAccountHandler,
	vmInput *vmcommon.ContractCallInput,
) (*vmcommon.VMOutput, error) {
	ce.mutExecution.RLock()
	defer ce.mutExecution.RUnlock()

	err := doCommonValidation(vmInput)
	if err != nil {
		return nil, err
	}

	signature := string(vmInput.Arguments[signaturePosition])
	inputData := vmInput.Arguments[dataStartPosition:]

	gasToUse := ce.calculateGasToUse(inputData)
	if vmInput.GasProvided < gasToUse {
		return nil, builtInFunctions.ErrNotEnoughGas
	}

	evmAccounts, castOk := ce.accounts.(evmcommon.EVMAccountsAdapter)
	if !castOk {
		return nil, builtInFunctions.ErrWrongTypeAssertion
	}

	encodingHandler := convertEncoding.NewHandler(evmAccounts)
	outputData, err := ce.convertData(encodingHandler, signature, inputData)
	if err != nil {
		return nil, err
	}

	return &vmcommon.VMOutput{
		ReturnCode:   vmcommon.Ok,
		GasRemaining: vmInput.GasProvided - gasToUse,
		ReturnData:   outputData,
	}, nil
}

func (ce *ConvertEncoding) convertData(encodingHandler *convertEncoding.Handler, signature string, inputData [][]byte) ([][]byte, error) {
	switch ce.function {
	case core.BuiltInFunctionEthereumToMultiversXEncodingWithMultiversXSignature:
		return encodingHandler.EthToMvxEncodingWithMvxSignature(signature, inputData)
	case core.BuiltInFunctionEthereumToMultiversXEncodingWithEthereumSignature:
		return encodingHandler.EthToMvxEncodingWithEthSignature(signature, inputData)
	case core.BuiltInFunctionMultiversXToEthereumEncodingWithMultiversXSignature:
		return encodingHandler.MvxToEthEncodingWithMvxSignature(signature, inputData)
	case core.BuiltInFunctionMultiversXToEthereumEncodingWithEthereumSignature:
		return encodingHandler.MvxToEthEncodingWithEthSignature(signature, inputData)
	default:
		return nil, builtInFunctions.ErrInvalidArguments
	}
}

func (ce *ConvertEncoding) calculateGasToUse(inputData [][]byte) uint64 {
	totalLength := uint64(0)
	for _, arg := range inputData {
		totalLength += uint64(len(arg))
	}
	return totalLength * ce.baseOperationCost.CompilePerByte
}

func doCommonValidation(vmInput *vmcommon.ContractCallInput) error {
	if vmInput == nil {
		return builtInFunctions.ErrNilVmInput
	}
	if len(vmInput.Arguments) < minArgsCount {
		return builtInFunctions.ErrInvalidArguments
	}
	if vmInput.CallValue == nil {
		return builtInFunctions.ErrNilValue
	}
	if vmInput.CallValue.Cmp(big.NewInt(0)) != 0 {
		return builtInFunctions.ErrBuiltInFunctionCalledWithValue
	}

	return nil
}

// IsInterfaceNil returns true if underlying object in nil
func (ce *ConvertEncoding) IsInterfaceNil() bool {
	return ce == nil
}
