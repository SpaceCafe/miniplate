package functions

import "errors"

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
	DecimalSymbols       = []string{",", ".", "٫"}
	ErrDivZero           = errors.New("division by zero")
	ErrInvalidArgument   = errors.New("invalid argument")
	ErrUnsupportedType   = errors.New("unsupported type")
	ErrUndefinedTemplate = errors.New("undefined template")
)
