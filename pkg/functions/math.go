package functions

import (
	"fmt"
	"math"
	"slices"
)

type MathFuncs struct{}

func (f MathFuncs) Abs(in any) (any, error) {
	if f.IsFloat(in) {
		v, err := ConversionFuncs{}.ToFloat64(in)
		if err != nil {
			return 0, err
		}

		return math.Abs(v), nil
	}

	v, err := ConversionFuncs{}.ToInt64(in)
	if err != nil {
		return 0, err
	}

	if v < 0 {
		return -v, nil
	}

	return v, nil
}

func (f MathFuncs) Add(in ...any) (any, error) {
	return f.mathOperation(in, add[float64], add[int64])
}

func (f MathFuncs) Ceil(in any) (float64, error) {
	v, err := ConversionFuncs{}.ToFloat64(in)
	if err != nil {
		return 0, err
	}

	return math.Ceil(v), nil
}

func (f MathFuncs) Div(in ...any) (float64, error) {
	if len(in) < 2 {
		return 0, ErrInvalidArgument
	}

	floats, err := ConversionFuncs{}.ToFloat64s(in...)
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

func (f MathFuncs) Floor(in any) (float64, error) {
	v, err := ConversionFuncs{}.ToFloat64(in)
	if err != nil {
		return 0, err
	}

	return math.Floor(v), nil
}

func (f MathFuncs) IsFloat(in any) bool {
	return f.getNumType(in) == Float
}

func (f MathFuncs) IsInt(in any) bool {
	return f.getNumType(in) == Int
}

func (f MathFuncs) IsNum(in any) bool {
	return f.getNumType(in) != NaN
}

func (f MathFuncs) Max(in ...any) (any, error) {
	return f.mathOperation(in, slices.Max[[]float64], slices.Max[[]int64])
}

func (f MathFuncs) Min(in ...any) (any, error) {
	return f.mathOperation(in, slices.Min[[]float64], slices.Min[[]int64])
}

func (f MathFuncs) Mul(in ...any) (any, error) {
	return f.mathOperation(in, mul[float64], mul[int64])
}

func (f MathFuncs) Pow(base, exponent any) (float64, error) {
	x, err := ConversionFuncs{}.ToFloat64(base)
	if err != nil {
		return 0, err
	}

	y, err := ConversionFuncs{}.ToFloat64(exponent)
	if err != nil {
		return 0, err
	}

	return math.Pow(x, y), nil
}

func (f MathFuncs) Rem(divisor, dividend any) (any, error) {
	if !f.IsInt(divisor) || !f.IsInt(dividend) {
		return 0, fmt.Errorf("%w: invalid type: %T, %T", ErrInvalidArgument, divisor, dividend)
	}

	x, err := ConversionFuncs{}.ToInt64(divisor)
	if err != nil {
		return 0, err
	}

	y, err := ConversionFuncs{}.ToInt64(dividend)
	if err != nil {
		return 0, err
	}

	if y == 0 {
		return 0, ErrDivZero
	}

	return x % y, nil
}

func (f MathFuncs) Round(in any) (float64, error) {
	v, err := ConversionFuncs{}.ToFloat64(in)
	if err != nil {
		return 0, err
	}

	return math.Round(v), nil
}

func (f MathFuncs) Seq(args ...any) (out []int64, err error) {
	var (
		end         int64
		start, step int64 = 1, 1
	)

	out = []int64{}

	v, err := ConversionFuncs{}.ToInt64s(args...)
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

func (f MathFuncs) Sub(in ...any) (any, error) {
	return f.mathOperation(in, sub[float64], sub[int64])
}

func add[T float64 | int64](in []T) (out T) {
	for _, v := range in {
		out += v
	}

	return
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

func (f MathFuncs) mathOperation(
	in []any,
	floatOp func([]float64) float64,
	intOp func([]int64) int64,
) (v any, err error) {
	defer func() {
		if r := recover(); r != nil {
			v = nil
			err = fmt.Errorf("%w: %v", ErrInvalidArgument, r)
		}
	}()

	var (
		floats   []float64
		integers []int64
	)

	if len(in) == 0 {
		return 0, ErrInvalidArgument
	}

	if slices.ContainsFunc(in, f.IsFloat) {
		floats, err = ConversionFuncs{}.ToFloat64s(in...)
		if err != nil {
			return 0, err
		}

		return floatOp(floats), nil
	}

	integers, err = ConversionFuncs{}.ToInt64s(in...)
	if err != nil {
		return 0, err
	}

	return intOp(integers), nil
}

func mul[T float64 | int64](in []T) (out T) {
	out = in[0]
	for _, v := range in[1:] {
		out *= v
	}

	return
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

func sub[T float64 | int64](in []T) (out T) {
	out = in[0]
	for _, v := range in[1:] {
		out -= v
	}

	return
}
