package vmcommon

var _ EVMContractCreateInputHandler = (*EVMContractCreateInput)(nil)

// GetAliasAddress return the recipient alias address
func (e *EVMContractCreateInput) GetAliasAddress() []byte {
	return e.AliasAddress
}

var _ EVMContractCallInputHandler = (*EVMContractCallInput)(nil)

// GetRecipientAliasAddr return the recipient alias address
func (e *EVMContractCallInput) GetRecipientAliasAddr() []byte {
	return e.RecipientAliasAddr
}

var _ EVMContractSameContextCallInputHandler = (*ContractSameContextCallInput)(nil)

// GetDoTransfer return true if it should do transfer
func (c *ContractSameContextCallInput) GetDoTransfer() bool {
	return c.DoTransfer
}

// GetCodeAddress return the code address
func (c *ContractSameContextCallInput) GetCodeAddress() []byte {
	return c.CodeAddress
}
