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
	ErrBuiltWithoutTOML  = errors.New("built without TOML support")
	ErrBuiltWithoutYAML  = errors.New("built without YAML support")
	ErrDivZero           = errors.New("division by zero")
	ErrInvalidArgument   = errors.New("invalid argument")
	ErrUndefinedTemplate = errors.New("undefined template")
	ErrUnsupportedType   = errors.New("unsupported type")
)
