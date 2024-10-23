package yetanothercalc

import (
	"errors"
)

// Internal errors
var ErrRegexp = errors.New("failed to set regexp")
var ErrParseFloat = errors.New("failed to parse float")

// expression errors
var ErrExtraCharacters = errors.New("expression has extra characters")
var ErrUnpairedBracket = errors.New("expression has unpaired brackets")
var ErrWrongBracketOrder = errors.New("expression has wrong sequence of brackets")
var ErrMultipleOperands = errors.New("expression has multiple sequential operands")
var ErrMultipleNumbers = errors.New("expression has multiple sequential operands")
var ErrZeroByDivision = errors.New("expression has zero by division")
