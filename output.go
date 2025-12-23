package vmcommon

import (
	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
)

type EVMOutputAccount struct {
	vmcommon.OutputAccount

	// CodeHash is the hash of the code of a smart contract account.
	// Like "Code", this field will be populated when a new SC must be created after the transaction.
	CodeHash []byte

	// IsContractCreatedInTransaction specifies whether the contract account has been created in the current transaction
	IsContractCreatedInTransaction bool
}
