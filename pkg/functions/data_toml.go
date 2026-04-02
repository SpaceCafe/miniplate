//go:build toml

package functions

import (
	"github.com/BurntSushi/toml"
)

func (DataFuncs) TOML(in any) (obj any, err error) {
	switch in := in.(type) {
	case []byte:
		err = toml.Unmarshal(in, &obj)
	case string:
		err = toml.Unmarshal([]byte(in), &obj)
	}

	return
}

func (DataFuncs) ToTOML(obj any) (string, error) {
	result, err := toml.Marshal(obj)

	return string(result), err
}
