package vmcommon

import "errors"

// ErrNilRequest signals that the provided request is nil
var ErrNilRequest = errors.New("nil request")

// ErrInvalidSourceIdentifier signals that the source identifier is invalid
var ErrInvalidSourceIdentifier = errors.New("invalid source identifier")

// ErrInvalidRequestedIdentifier signals that the requested identifier is invalid
var ErrInvalidRequestedIdentifier = errors.New("invalid requested identifier")

// ErrSourceIdentifierMatchesRequestedIdentifier signals that the source identifier matches the requested identifier
var ErrSourceIdentifierMatchesRequestedIdentifier = errors.New("source identifier matches requested identifier")

// ErrIdentifierNotHandledForBlankAddress signals that the identifier is not handled for blank address
var ErrIdentifierNotHandledForBlankAddress = errors.New("address identifier not handled for blank address")
