package mvx

import (
	"testing"

	"github.com/stretchr/testify/require"

	convertCommon "github.com/multiversx/mx-chain-vm-common-evm-go/builtInFunctions/convertEncoding/common"
)

func TestBuildMultiversXAbi(t *testing.T) {
	testBuildMultiversXAbi(t, convertCommon.MvxComplexSignature1)
	testBuildMultiversXAbi(t, convertCommon.MvxComplexSignature2)
	testBuildMultiversXAbi(t, convertCommon.MvxComplexSignature3)

	args, err := ParseMultiversXSignature(convertCommon.BuildArgName(0))
	require.NoError(t, err)
	require.NotEmpty(t, args)

	_, err = BuildMultiversXAbi(args)
	require.ErrorContains(t, err, ErrInvalidSignatureAbiType.Error())
}

func testBuildMultiversXAbi(t *testing.T, signature string) {
	args, err := ParseMultiversXSignature(signature)
	require.NoError(t, err)
	require.NotEmpty(t, args)

	abi, err := BuildMultiversXAbi(args)
	require.NoError(t, err)
	require.NotEmpty(t, abi)
}
