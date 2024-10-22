package ordinarycalc

import (
	"errors"
)

// Internal errors
var ErrRegexp = errors.New("failed to set regexp")

// expression errors
var ErrExtraCharacters = errors.New("expression have extra characters")
var ErrUnpairedBracket = errors.New("expression have unpaired brackets")
