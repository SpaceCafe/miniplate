package functions

import (
	"fmt"
)

type NumType int

const (
	NaN NumType = iota
	Int
	Float
)

const (
	MaxSeqIterations = 100
)

var (
	DecimalSymbols     = []string{",", ".", "٫"}
	ErrDivZero         = fmt.Errorf("division by zero")
	ErrInvalidArgument = fmt.Errorf("invalid argument")
)
