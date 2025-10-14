package functions

import (
	"fmt"
	"math"
	"slices"
)

type MathFuncs struct{}

func (r MathFuncs) IsInt(in any) bool {
	return r.getNumType(in) == Int
}

func (r MathFuncs) IsFloat(in any) bool {
	return r.getNumType(in) == Float
}

func (r MathFuncs) IsNum(in any) bool {
	return r.getNumType(in) != NaN
}

func (MathFuncs) getNumType(in any) NumType {
	switch in.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return Int
	case float32, float64:
		return Float
	}
	return NaN
}

func (r MathFuncs) Abs(in any) (any, error) {
	if r.IsFloat(in) {
		v, err := ConvFuncs{}.ToFloat64(in)
		if err != nil {
			return 0, err
		}
		return math.Abs(v), nil
	}

	v, err := ConvFuncs{}.ToInt64(in)
	if err != nil {
		return 0, err
	}
	if v < 0 {
		return -v, nil
	}
	return v, nil
}

func (r MathFuncs) Round(in any) (float64, error) {
	v, err := ConvFuncs{}.ToFloat64(in)
	if err != nil {
		return 0, err
	}
	return math.Round(v), nil
}

func (r MathFuncs) Ceil(in any) (float64, error) {
	v, err := ConvFuncs{}.ToFloat64(in)
	if err != nil {
		return 0, err
	}
	return math.Ceil(v), nil
}

func (r MathFuncs) Floor(in any) (float64, error) {
	v, err := ConvFuncs{}.ToFloat64(in)
	if err != nil {
		return 0, err
	}
	return math.Floor(v), nil
}

func (r MathFuncs) Min(in ...any) (any, error) {
	return r.mathOperation(in, slices.Min[[]float64], slices.Min[[]int64])
}

func (r MathFuncs) Max(in ...any) (any, error) {
	return r.mathOperation(in, slices.Max[[]float64], slices.Max[[]int64])
}

func (r MathFuncs) Add(in ...any) (any, error) {
	return r.mathOperation(in, add[float64], add[int64])
}

func (r MathFuncs) Sub(in ...any) (any, error) {
	return r.mathOperation(in, sub[float64], sub[int64])
}

func (r MathFuncs) Mul(in ...any) (any, error) {
	return r.mathOperation(in, mul[float64], mul[int64])
}

func (r MathFuncs) Div(in ...any) (float64, error) {
	if len(in) < 2 {
		return 0, ErrInvalidArgument
	}
	floats, err := ConvFuncs{}.ToFloat64s(in...)
	if err != nil {
		return 0, err
	}
	out := floats[0]
	for _, v := range floats[1:] {
		if v == 0 {
			return 0, ErrDivZero
		}
		out /= v
	}
	return out, nil
}

func (r MathFuncs) Pow(base any, exponent any) (float64, error) {
	x, err := ConvFuncs{}.ToFloat64(base)
	if err != nil {
		return 0, err
	}
	y, err := ConvFuncs{}.ToFloat64(exponent)
	if err != nil {
		return 0, err
	}
	return math.Pow(x, y), nil
}

func (r MathFuncs) Rem(divisor any, dividend any) (any, error) {
	if !r.IsInt(divisor) || !r.IsInt(dividend) {
		return 0, fmt.Errorf("invalid type: %T, %T", divisor, dividend)
	}
	x, err := ConvFuncs{}.ToInt64(divisor)
	if err != nil {
		return 0, err
	}
	y, err := ConvFuncs{}.ToInt64(dividend)
	if err != nil {
		return 0, err
	}
	if y == 0 {
		return 0, ErrDivZero
	}
	return x % y, nil
}

func (r MathFuncs) Seq(args ...any) (out []int64, err error) {
	var (
		end         int64
		start, step int64 = 1, 1
	)
	out = []int64{}

	v, err := ConvFuncs{}.ToInt64s(args...)
	if err != nil {
		return
	}

	switch len(args) {
	case 1:
		end = v[0]
	case 2:
		start, end = v[0], v[1]
	case 3:
		start, end, step = v[0], v[1], v[2]
	default:
		return out, ErrInvalidArgument
	}

	if step == 0 {
		return
	}

	return seq(start, end, step), nil
}

func seq(start, end, step int64) (out []int64) {
	out = []int64{}
	for i := start; ; i += step {
		if step > 0 && i > end {
			break
		}
		if step < 0 && i < end {
			break
		}
		out = append(out, i)
		if len(out) >= MaxSeqIterations {
			return
		}
	}
	return
}

func add[T float64 | int64](in []T) (out T) {
	for _, v := range in {
		out += v
	}
	return
}

func sub[T float64 | int64](in []T) (out T) {
	out = in[0]
	for _, v := range in[1:] {
		out -= v
	}
	return
}

func mul[T float64 | int64](in []T) (out T) {
	out = in[0]
	for _, v := range in[1:] {
		out *= v
	}
	return
}

func (r MathFuncs) mathOperation(
	in []any,
	floatOp func([]float64) float64,
	intOp func([]int64) int64,
) (v any, err error) {
	defer func() {
		if r := recover(); r != nil {
			v = nil
			err = fmt.Errorf("%v", r)
		}
	}()

	var (
		floats   []float64
		integers []int64
	)

	if len(in) == 0 {
		return 0, ErrInvalidArgument
	}

	if slices.ContainsFunc(in, r.IsFloat) {
		floats, err = ConvFuncs{}.ToFloat64s(in...)
		if err != nil {
			return 0, err
		}
		return floatOp(floats), nil
	}

	integers, err = ConvFuncs{}.ToInt64s(in...)
	if err != nil {
		return 0, err
	}
	return intOp(integers), nil
}
