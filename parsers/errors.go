package parsers

import (
	"errors"
)

// ErrUnexpectedInputSize signals that the input does not meet the required size
var ErrUnexpectedInputSize = errors.New("unexpected input size")
