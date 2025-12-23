package builtInFunctions

import (
	"fmt"

	"github.com/multiversx/mx-chain-core-evm-go/core"
	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
	"github.com/multiversx/mx-chain-vm-common-go/builtInFunctions"

	evmcommon "github.com/multiversx/mx-chain-vm-common-evm-go"
)

type evmBuiltInFuncCreator struct {
	vmcommon.BuiltInFunctionFactory
}

// NewEVMBuiltInFunctionsCreator creates a component which will instantiate the built in functions contracts
func NewEVMBuiltInFunctionsCreator(args builtInFunctions.ArgsCreateBuiltInFunctionContainer) (*evmBuiltInFuncCreator, error) {
	builtInFuncCreator, err := builtInFunctions.NewBuiltInFunctionsCreator(args)
	if err != nil {
		return nil, err
	}

	return &evmBuiltInFuncCreator{builtInFuncCreator}, nil
}

// CreateBuiltInFunctionContainer will create the list of built-in functions
func (b *evmBuiltInFuncCreator) CreateBuiltInFunctionContainer() error {
	err := b.BuiltInFunctionFactory.CreateBuiltInFunctionContainer()
	if err != nil {
		return err
	}

	evmAccounts, castOk := b.Accounts().(evmcommon.EVMAccountsAdapter)
	if !castOk {
		return fmt.Errorf("%w when casting accounts to evm accounts", builtInFunctions.ErrWrongTypeAssertion)
	}

	newFunc, err := NewEVMConvertEncodingFunc(b.GasConfig().BaseOperationCost, evmAccounts, core.BuiltInFunctionEthereumToMultiversXEncodingWithMultiversXSignature)
	if err != nil {
		return err
	}
	err = b.BuiltInFunctionContainer().Add(core.BuiltInFunctionEthereumToMultiversXEncodingWithMultiversXSignature, newFunc)
	if err != nil {
		return err
	}

	newFunc, err = NewEVMConvertEncodingFunc(b.GasConfig().BaseOperationCost, evmAccounts, core.BuiltInFunctionEthereumToMultiversXEncodingWithEthereumSignature)
	if err != nil {
		return err
	}
	err = b.BuiltInFunctionContainer().Add(core.BuiltInFunctionEthereumToMultiversXEncodingWithEthereumSignature, newFunc)
	if err != nil {
		return err
	}

	newFunc, err = NewEVMConvertEncodingFunc(b.GasConfig().BaseOperationCost, evmAccounts, core.BuiltInFunctionMultiversXToEthereumEncodingWithMultiversXSignature)
	if err != nil {
		return err
	}
	err = b.BuiltInFunctionContainer().Add(core.BuiltInFunctionMultiversXToEthereumEncodingWithMultiversXSignature, newFunc)
	if err != nil {
		return err
	}

	newFunc, err = NewEVMConvertEncodingFunc(b.GasConfig().BaseOperationCost, evmAccounts, core.BuiltInFunctionMultiversXToEthereumEncodingWithEthereumSignature)
	if err != nil {
		return err
	}
	err = b.BuiltInFunctionContainer().Add(core.BuiltInFunctionMultiversXToEthereumEncodingWithEthereumSignature, newFunc)
	if err != nil {
		return err
	}

	return nil
}
