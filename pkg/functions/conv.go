package functions

import (
	"fmt"
	"net/url"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

type ConvFuncs struct{}

func (f ConvFuncs) Atoi(in any) (int64, error) {
	return f.ToInt64(in)
}

func (f ConvFuncs) Bool(in any) (bool, error) {
	return f.ToBool(in)
}

func (f ConvFuncs) Default(def, in any) any {
	if in == nil {
		return def
	}

	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		if v.Len() == 0 {
			return def
		}
	default:
		if v.IsZero() {
			return def
		}
	}

	return in
}

func (f ConvFuncs) Float(args ...any) (float64, error) {
	return f.ToFloat64(args...)
}

func (f ConvFuncs) Int(in any) (int64, error) {
	return f.ToInt64(in)
}

func (f ConvFuncs) Join(args ...any) string {
	var (
		sep string
		ok  bool
	)

	if len(args) < 2 {
		return ""
	}

	if sep, ok = args[len(args)-1].(string); !ok {
		return ""
	}

	v := f.ToStrings(args[:len(args)-1]...)

	var builder strings.Builder

	for i, s := range v {
		if i > 0 {
			builder.WriteString(sep)
		}

		builder.WriteString(s)
	}

	return builder.String()
}

func (f ConvFuncs) ParseFloat(args ...any) (float64, error) {
	return f.ToFloat64(args...)
}

func (f ConvFuncs) ParseInt(in any) (int64, error) {
	return f.ToInt64(in)
}

func (f ConvFuncs) String(in any) string {
	return f.ToString(in)
}

func (f ConvFuncs) ToBool(in any) (bool, error) {
	switch v := in.(type) {
	case bool:
		return v, nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v != 0, nil
	case float32, float64:
		return v != 0.0, nil
	case string:
		return stringToBool(v)
	case fmt.Stringer:
		return stringToBool(v.String())
	}

	return false, fmt.Errorf("%w: not a boolean value: %v", ErrInvalidArgument, in)
}

func (f ConvFuncs) ToBools(in ...any) (list []bool, err error) {
	list = make([]bool, len(in))
	for i, v := range in {
		list[i], err = f.ToBool(v)
		if err != nil {
			return
		}
	}

	return
}

func (f ConvFuncs) ToFloat(args ...any) (float64, error) {
	return f.ToFloat64(args...)
}

func (f ConvFuncs) ToFloat64(args ...any) (float64, error) {
	var (
		in  any
		sep = "."
		ok  bool
	)

	if len(args) == 0 {
		return 0.0, ErrInvalidArgument
	}

	if len(args) == 2 {
		if sep, ok = decimalSymbol(args[0]); ok {
			in = args[1]
		} else {
			return 0.0, ErrInvalidArgument
		}
	} else {
		in = args[0]
	}

	switch v := in.(type) {
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case string:
		return stringToFloat(sep, v)
	case fmt.Stringer:
		return stringToFloat(sep, v.String())
	}

	return 0, fmt.Errorf("%w: not a float value: %v", ErrInvalidArgument, args[0])
}

func (f ConvFuncs) ToFloat64s(args ...any) (list []float64, err error) {
	if len(args) > 1 {
		if sep, ok := decimalSymbol(args[0]); ok {
			list = make([]float64, len(args[1:]))
			for i, v := range args[1:] {
				list[i], err = f.ToFloat64(sep, v)
				if err != nil {
					return
				}
			}

			return
		}
	}

	list = make([]float64, len(args))
	for i, v := range args {
		list[i], err = f.ToFloat64(v)
		if err != nil {
			return
		}
	}

	return
}

func (f ConvFuncs) ToInt(in any) (int64, error) {
	return f.ToInt64(in)
}

func (f ConvFuncs) ToInt64(in any) (int64, error) {
	switch v := in.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint:
		if v > 1<<63-1 {
			return 0, fmt.Errorf("%w: value out of range for int64: %d", ErrInvalidArgument, v)
		}

		return int64(v), nil
	case uint64:
		if v > 1<<63-1 {
			return 0, fmt.Errorf("%w: value out of range for int64: %d", ErrInvalidArgument, v)
		}

		return int64(v), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case string:
		return stringToInt(v)
	case fmt.Stringer:
		return stringToInt(v.String())
	}

	return 0, fmt.Errorf("%w: invalid integer value: %v", ErrInvalidArgument, in)
}

func (f ConvFuncs) ToInt64s(in ...any) (list []int64, err error) {
	list = make([]int64, len(in))
	for i, v := range in {
		list[i], err = f.ToInt64(v)
		if err != nil {
			return
		}
	}

	return
}

func (f ConvFuncs) ToInts(in ...any) ([]int64, error) {
	return f.ToInt64s(in...)
}

func (f ConvFuncs) ToString(in any) string {
	if in == nil {
		return ""
	}

	switch v := in.(type) {
	case bool:
		return strconv.FormatBool(v)
	case int, int8, int16, int32, int64:
		return intToString(v)
	case uint, uint8, uint16, uint32, uint64:
		return uintToString(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (f ConvFuncs) ToStrings(in ...any) (list []string) {
	list = make([]string, 0)

	for _, v := range in {
		val := reflect.ValueOf(v)
		switch val.Kind() {
		case reflect.Slice, reflect.Array:
			for j := range val.Len() {
				list = append(list, f.ToString(val.Index(j).Interface()))
			}
		default:
			list = append(list, f.ToString(v))
		}
	}

	return
}

func (f ConvFuncs) URL(in string) (*url.URL, error) {
	return url.Parse(in)
}

func intToString(in any) string {
	var value int64

	switch v := in.(type) {
	case int:
		value = int64(v)
	case int8:
		value = int64(v)
	case int16:
		value = int64(v)
	case int32:
		value = int64(v)
	case int64:
		value = v
	}

	return strconv.FormatInt(value, 10)
}

func uintToString(in any) string {
	var value uint64

	switch v := in.(type) {
	case uint:
		value = uint64(v)
	case uint8:
		value = uint64(v)
	case uint16:
		value = uint64(v)
	case uint32:
		value = uint64(v)
	case uint64:
		value = v
	}

	return strconv.FormatUint(value, 10)
}

func decimalSymbol(in any) (sep string, ok bool) {
	switch v := in.(type) {
	case fmt.Stringer:
		sep = v.String()
	case string:
		sep = v
	case rune:
		sep = string(v)
	default:
		return "", false
	}

	return sep, slices.Contains(DecimalSymbols, sep)
}

func stringToBool(in string) (bool, error) {
	in = strings.ToLower(strings.TrimSpace(in))
	switch in {
	case "1", "t", "y", "true", "yes", "on":
		return true, nil
	case "0", "f", "n", "false", "no", "off":
		return false, nil
	}

	return false, fmt.Errorf("%w: not a boolean value: %v", ErrInvalidArgument, in)
}

func stringToInt(in string) (int64, error) {
	in = strings.ToLower(strings.TrimSpace(in))

	return strconv.ParseInt(in, 10, 64)
}

func stringToFloat(sep, in string) (float64, error) {
	in = strings.ToLower(strings.TrimSpace(in))
	in = strings.Replace(in, sep, ".", 1)

	return strconv.ParseFloat(in, 64)
}
