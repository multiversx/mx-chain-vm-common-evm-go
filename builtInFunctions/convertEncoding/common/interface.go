package common

import evmcommon "github.com/multiversx/mx-chain-vm-common-evm-go"

type Argument struct {
	Type      string
	Arguments Arguments
}

type Arguments []*Argument

type EncodingContext struct {
	Accounts evmcommon.EVMAccountsAdapter
}

func BuildEncodingContext(accounts evmcommon.EVMAccountsAdapter) *EncodingContext {
	return &EncodingContext{Accounts: accounts}
}
