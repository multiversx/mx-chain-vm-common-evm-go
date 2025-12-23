package vmcommon

var _ EVMOutputAccountHandler = (*EVMOutputAccount)(nil)

// GetCodeHash returns the evm account code hash
func (eoa *EVMOutputAccount) GetCodeHash() []byte {
	return eoa.CodeHash
}

// GetIsContractCreatedInTransaction returns is evm account is created in transaction
func (eoa *EVMOutputAccount) GetIsContractCreatedInTransaction() bool {
	return eoa.IsContractCreatedInTransaction
}

// SetCodeHash sets the evm account code hash
func (eoa *EVMOutputAccount) SetCodeHash(hash []byte) {
	eoa.CodeHash = hash
}

// SetIsContractCreatedInTransaction sets evm account is created in transaction
func (eoa *EVMOutputAccount) SetIsContractCreatedInTransaction() {
	eoa.IsContractCreatedInTransaction = true
}
