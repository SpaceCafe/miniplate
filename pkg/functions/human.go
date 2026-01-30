package functions

import (
	"github.com/dustin/go-humanize"
)

type HumanFuncs struct{}

func (f HumanFuncs) Bytes(in any) (string, error) {
	v, err := ConversionFuncs{}.ToInt64(in)
	if err != nil {
		return "", err
	}

	if v < 0 {
		return "", nil
	}

	return humanize.Bytes(uint64(v)), nil
}

func (HumanFuncs) ParseBytes(in any) (int64, error) {
	v := ConversionFuncs{}.ToString(in)

	result, err := humanize.ParseBytes(v)
	if err != nil {
		return 0, err
	}

	return ConversionFuncs{}.ToInt64(result)
}

func (f HumanFuncs) ToBytes(in any) (string, error) {
	return f.Bytes(in)
}
