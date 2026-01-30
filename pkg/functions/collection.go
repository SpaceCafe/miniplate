package functions

import (
	"errors"
	"fmt"
	"slices"
)

var ErrInvalidDictKey = errors.New("invalid dict key")

type CollectionFuncs struct{}

func (CollectionFuncs) Dict(in ...any) (map[string]any, error) {
	if len(in)%2 != 0 {
		in = append(in, "")
	}

	dict := make(map[string]any, len(in)/2)

	for i := 0; i < len(in); i += 2 {
		key, ok := in[i].(string)
		if !ok {
			return nil, fmt.Errorf("%w: index %d is not a string", ErrInvalidDictKey, i)
		}

		dict[key] = in[i+1]
	}

	return dict, nil
}

func (CollectionFuncs) Keys(in map[string]any) ([]string, error) {
	if len(in) == 0 {
		return nil, ErrInvalidArgument
	}

	keys := make([]string, 0, len(in))
	for k := range in {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	return keys, nil
}

func (CollectionFuncs) Slice(in ...any) []any {
	return in
}
