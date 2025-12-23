package mock

import (
	"errors"

	vmcommon "github.com/multiversx/mx-chain-vm-common-go"

	evmcommon "github.com/multiversx/mx-chain-vm-common-evm-go"
)

type EVMAccountsStub struct {
	GetExistingAccountCalled func(address []byte) (vmcommon.AccountHandler, error)
	LoadAccountCalled        func(address []byte) (vmcommon.AccountHandler, error)
	SaveAccountCalled        func(account vmcommon.AccountHandler) error
	RemoveAccountCalled      func(address []byte) error
	CommitCalled             func() ([]byte, error)
	JournalLenCalled         func() int
	RevertToSnapshotCalled   func(snapshot int) error
	RootHashCalled           func() ([]byte, error)
	RecreateTrieCalled       func(rootHash []byte) error
	SnapshotStateCalled      func(rootHash []byte)
	SetStateCheckpointCalled func(rootHash []byte)
	IsPruningEnabledCalled   func() bool
	GetCodeCalled            func([]byte) []byte
	SaveAliasAddressCalled   func(request *evmcommon.AliasSaveRequest) error
	RequestAddressCalled     func(request *evmcommon.AddressRequest) (*evmcommon.AddressResponse, error)
}

func (as *EVMAccountsStub) GetExistingAccount(address []byte) (vmcommon.AccountHandler, error) {
	if as.GetExistingAccountCalled != nil {
		return as.GetExistingAccountCalled(address)
	}
	return nil, errNotImplemented
}

func (as *EVMAccountsStub) LoadAccount(address []byte) (vmcommon.AccountHandler, error) {
	if as.LoadAccountCalled != nil {
		return as.LoadAccountCalled(address)
	}
	return nil, errNotImplemented
}

func (as *EVMAccountsStub) SaveAccount(account vmcommon.AccountHandler) error {
	if as.SaveAccountCalled != nil {
		return as.SaveAccountCalled(account)
	}
	return nil
}

func (as *EVMAccountsStub) RemoveAccount(address []byte) error {
	if as.RemoveAccountCalled != nil {
		return as.RemoveAccountCalled(address)
	}
	return errNotImplemented
}

func (as *EVMAccountsStub) Commit() ([]byte, error) {
	if as.CommitCalled != nil {
		return as.CommitCalled()
	}
	return nil, errNotImplemented
}

func (as *EVMAccountsStub) JournalLen() int {
	if as.JournalLenCalled != nil {
		return as.JournalLenCalled()
	}
	return 0
}

func (as *EVMAccountsStub) RevertToSnapshot(snapshot int) error {
	if as.RevertToSnapshotCalled != nil {
		return as.RevertToSnapshotCalled(snapshot)
	}
	return errNotImplemented
}

func (as *EVMAccountsStub) GetCode(codeHash []byte) []byte {
	if as.GetCodeCalled != nil {
		return as.GetCodeCalled(codeHash)
	}
	return nil
}

func (as *EVMAccountsStub) RootHash() ([]byte, error) {
	if as.RootHashCalled != nil {
		return as.RootHashCalled()
	}
	return nil, errNotImplemented
}

// SaveAliasAddress -
func (as *EVMAccountsStub) SaveAliasAddress(request *evmcommon.AliasSaveRequest) error {
	if as.SaveAliasAddressCalled != nil {
		return as.SaveAliasAddressCalled(request)
	}

	return errNotImplemented
}

// RequestAddress -
func (as *EVMAccountsStub) RequestAddress(request *evmcommon.AddressRequest) (*evmcommon.AddressResponse, error) {
	if as.RequestAddressCalled != nil {
		return as.RequestAddressCalled(request)
	}

	return nil, errNotImplemented
}

func (as *EVMAccountsStub) IsInterfaceNil() bool {
	return as == nil
}

var errNotImplemented = errors.New("not implemented")
