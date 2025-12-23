package builtInFunctions

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-vm-common-go/builtInFunctions"
	"github.com/multiversx/mx-chain-vm-common-go/mock"
	"github.com/stretchr/testify/require"

	mock2 "github.com/multiversx/mx-chain-vm-common-evm-go/mock"
)

func createMockArguments() builtInFunctions.ArgsCreateBuiltInFunctionContainer {
	gasMap := make(map[string]map[string]uint64)
	fillGasMapInternal(gasMap, 1)

	args := builtInFunctions.ArgsCreateBuiltInFunctionContainer{
		GasMap:                            gasMap,
		MapDNSAddresses:                   make(map[string]struct{}),
		MapDNSV2Addresses:                 make(map[string]struct{}),
		EnableUserNameChange:              false,
		Marshalizer:                       &mock.MarshalizerMock{},
		Accounts:                          &mock2.EVMAccountsStub{},
		ShardCoordinator:                  mock.NewMultiShardsCoordinatorMock(1),
		EnableEpochsHandler:               &mock.EnableEpochsHandlerStub{},
		GuardedAccountHandler:             &mock.GuardedAccountHandlerStub{},
		MaxNumOfAddressesForTransferRole:  100,
		MapWhiteListedCrossChainAddresses: getWhiteListedAddress(),
	}

	return args
}

func fillGasMapInternal(gasMap map[string]map[string]uint64, value uint64) map[string]map[string]uint64 {
	gasMap[core.BaseOperationCostString] = fillGasMapBaseOperationCosts(value)
	gasMap[core.BuiltInCostString] = fillGasMapBuiltInCosts(value)

	return gasMap
}

func fillGasMapBaseOperationCosts(value uint64) map[string]uint64 {
	gasMap := make(map[string]uint64)
	gasMap["StorePerByte"] = value
	gasMap["DataCopyPerByte"] = value
	gasMap["ReleasePerByte"] = value
	gasMap["PersistPerByte"] = value
	gasMap["CompilePerByte"] = value
	gasMap["AoTPreparePerByte"] = value
	gasMap["GetCode"] = value
	return gasMap
}

func fillGasMapBuiltInCosts(value uint64) map[string]uint64 {
	gasMap := make(map[string]uint64)
	gasMap["ClaimDeveloperRewards"] = value
	gasMap["ChangeOwnerAddress"] = value
	gasMap["SaveUserName"] = value
	gasMap["SaveKeyValue"] = value
	gasMap["ESDTTransfer"] = value
	gasMap["ESDTBurn"] = value
	gasMap["ChangeOwnerAddress"] = value
	gasMap["ClaimDeveloperRewards"] = value
	gasMap["SaveUserName"] = value
	gasMap["SaveKeyValue"] = value
	gasMap["ESDTTransfer"] = value
	gasMap["ESDTBurn"] = value
	gasMap["ESDTLocalMint"] = value
	gasMap["ESDTLocalBurn"] = value
	gasMap["ESDTNFTCreate"] = value
	gasMap["ESDTNFTAddQuantity"] = value
	gasMap["ESDTNFTBurn"] = value
	gasMap["ESDTNFTTransfer"] = value
	gasMap["ESDTNFTChangeCreateOwner"] = value
	gasMap["ESDTNFTAddUri"] = value
	gasMap["ESDTNFTUpdateAttributes"] = value
	gasMap["ESDTNFTMultiTransfer"] = value
	gasMap["SetGuardian"] = value
	gasMap["GuardAccount"] = value
	gasMap["UnGuardAccount"] = value
	gasMap["TrieLoadPerNode"] = value
	gasMap["TrieStorePerNode"] = value
	gasMap["ESDTModifyRoyalties"] = value
	gasMap["ESDTModifyCreator"] = value
	gasMap["ESDTNFTRecreate"] = value
	gasMap["ESDTNFTSetNewURIs"] = value
	gasMap["ESDTNFTUpdate"] = value

	return gasMap
}

func getWhiteListedAddress() map[string]struct{} {
	return map[string]struct{}{
		"whiteListedAddress": {},
	}
}

func TestEVMCreateBuiltInFunctionContainer(t *testing.T) {
	args := createMockArguments()
	f, err := NewEVMBuiltInFunctionsCreator(args)
	require.NoError(t, err)

	err = f.CreateBuiltInFunctionContainer()
	require.Nil(t, err)
	require.Equal(t, 46, f.BuiltInFunctionContainer().Len())
}
