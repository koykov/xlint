package xlint

import "errors"

var (
	ErrLengthUUID = errors.New("unsupported UUID length")
	ErrPrefixUUID = errors.New("invalid urn prefix")
	ErrFormatUUID = errors.New("invalid UUID format")
)
