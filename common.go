package vmcommon

import (
	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
)

func GetEVMContractCodeMetadata() vmcommon.CodeMetadata {
	return vmcommon.CodeMetadata{Payable: false, PayableBySC: false, Upgradeable: false, Readable: false, Guarded: false}
}
