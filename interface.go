package vmcommon

import (
	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
)

// EVMBlockchainHook is the interface for EVM VM blockchain callbacks
type EVMBlockchainHook interface {
	vmcommon.BlockchainHook

	// SaveAliasAddress saves the given alias address
	SaveAliasAddress(request *AliasSaveRequest) error

	// RequestAddress returns the requested address
	RequestAddress(request *AddressRequest) (*AddressResponse, error)
}

// EVMAccountsAdapter is used for the structure that manages the EVM accounts on top of a trie.PatriciaMerkleTrie
// implementation
type EVMAccountsAdapter interface {
	vmcommon.AccountsAdapter
	SaveAliasAddress(request *AliasSaveRequest) error
	RequestAddress(request *AddressRequest) (*AddressResponse, error)
}

// EVMOutputAccountHandler defines methods for accessing and modifying evm account
type EVMOutputAccountHandler interface {
	vmcommon.OutputAccountHandler
	GetCodeHash() []byte
	GetIsContractCreatedInTransaction() bool
	SetCodeHash(hash []byte)
	SetIsContractCreatedInTransaction()
}

// EVMContractCallInputHandler defines the behavior expected from an evm ContractCallInput
type EVMContractCallInputHandler interface {
	vmcommon.ContractCallInputHandler
	GetRecipientAliasAddr() []byte
}

// EVMContractSameContextCallInputHandler defines the read-only behavior of a ContractSameContextCallInput.
type EVMContractSameContextCallInputHandler interface {
	EVMContractCallInputHandler
	GetDoTransfer() bool
	GetCodeAddress() []byte
}

// EVMContractCreateInputHandler defines the behavior expected from an evm ContractCreateInput
type EVMContractCreateInputHandler interface {
	vmcommon.ContractCreateInputHandler
	GetAliasAddress() []byte
}
