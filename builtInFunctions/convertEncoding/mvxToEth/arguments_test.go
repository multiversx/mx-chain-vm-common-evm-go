package mvxToEth

import (
	"testing"

	"github.com/stretchr/testify/require"

	convertCommon "github.com/multiversx/mx-chain-vm-common-evm-go/builtInFunctions/convertEncoding/common"
	"github.com/multiversx/mx-chain-vm-common-evm-go/builtInFunctions/convertEncoding/mvx"
)

func TestMultiversXToEthereumArguments(t *testing.T) {
	testMultiversXToEthereumArguments(t, convertCommon.MvxComplexSignature1)
	testMultiversXToEthereumArguments(t, convertCommon.MvxComplexSignature2)
	testMultiversXToEthereumArguments(t, convertCommon.MvxComplexSignature3)
}

func testMultiversXToEthereumArguments(t *testing.T, signature string) {
	args, err := mvx.ParseMultiversXSignature(signature)
	require.NoError(t, err)
	require.NotEmpty(t, args)

	convertedArgs, err := MultiversXToEthereumArguments(args)
	require.NoError(t, err)
	require.NotEmpty(t, convertedArgs)
}
