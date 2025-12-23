package vmcommon

import (
	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
)

// EVMContractCreateInput EVM VM input when creating a new contract.
// Here we have no RecipientAddr because
// the address (PK) of the created account will be provided by the vmcommon.
// We also do not need to specify a Function field,
// because on creation `init` is always called.
type EVMContractCreateInput struct {
	vmcommon.ContractCreateInput

	// AliasAddress is the alias address of the contract being created.
	AliasAddress []byte
}

// EVMContractCallInput EVM VM input when calling a function from an existing contract
type EVMContractCallInput struct {
	vmcommon.ContractCallInput

	// RecipientAliasAddr is the alias address of the smart contract.
	RecipientAliasAddr []byte
}

// ContractSameContextCallInput VM input when calling a function from an existing contract on the same context
type ContractSameContextCallInput struct {
	EVMContractCallInput

	// DoTransfer specifies whether the transfer should be executed for this same context call
	DoTransfer bool

	// CodeAddress is the address for the code
	CodeAddress []byte
}
