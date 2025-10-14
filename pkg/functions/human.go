package functions

import (
	"github.com/dustin/go-humanize"
)

type HumanFuncs struct{}

func (HumanFuncs) Bytes(in any) (string, error) {
	v, err := ConvFuncs{}.ToInt64(in)
	if err != nil {
		return "", err
	}
	if v < 0 {
		return "", nil
	}
	return humanize.Bytes(uint64(v)), nil
}

func (r HumanFuncs) ToBytes(in any) (string, error) {
	return r.Bytes(in)
}

func (HumanFuncs) ParseBytes(in any) (int64, error) {
	v := ConvFuncs{}.ToString(in)
	result, err := humanize.ParseBytes(v)
	if err != nil {
		return 0, err
	}
	return ConvFuncs{}.ToInt64(result)
}
