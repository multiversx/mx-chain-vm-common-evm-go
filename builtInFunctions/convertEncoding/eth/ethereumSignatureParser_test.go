package eth

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	convertCommon "github.com/multiversx/mx-chain-vm-common-evm-go/builtInFunctions/convertEncoding/common"
)

func TestParseEthereumSignature(t *testing.T) {
	args, err := ParseEthereumSignature(convertCommon.EthComplexSignature)
	require.NoError(t, err)
	validateArguments(t, args)

	_, err = ParseEthereumSignature("")
	require.Equal(t, convertCommon.ErrBlankExpression, err)

	_, err = ParseEthereumSignature(convertCommon.Comma)
	require.Equal(t, convertCommon.ErrExpressionStartsWithDelimiter, err)

	_, err = ParseEthereumSignature(BeginTuple)
	require.Equal(t, convertCommon.ErrBlankExpression, err)

	_, err = ParseEthereumSignature(BeginTuple + Address)
	require.Equal(t, ErrExpectedTupleEnd, err)

	_, err = ParseEthereumSignature(Address + convertCommon.Comma)
	require.Equal(t, convertCommon.ErrExpectedExpressionAfterComma, err)

	_, err = ParseEthereumSignature(Address + EndTuple)
	require.Equal(t, convertCommon.ErrExpectedBlankRemainder, err)
}

func validateArguments(t *testing.T, args convertCommon.Arguments) {
	require.NotEmpty(t, args)
	for _, arg := range args {
		if strings.HasPrefix(arg.Type, Tuple) {
			validateArguments(t, arg.Arguments)
		}
	}
}
