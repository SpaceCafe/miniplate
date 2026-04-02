//go:build yaml

package functions

import (
	yaml "github.com/goccy/go-yaml"
)

func (DataFuncs) ToYAML(obj any) (string, error) {
	result, err := yaml.Marshal(obj)

	return string(result), err
}

func (DataFuncs) YAML(in any) (obj any, err error) {
	switch in := in.(type) {
	case []byte:
		err = yaml.Unmarshal(in, &obj)
	case string:
		err = yaml.Unmarshal([]byte(in), &obj)
	}

	return
}

func (DataFuncs) YAMLArray(in any) (list []any, err error) {
	switch in := in.(type) {
	case []byte:
		err = yaml.Unmarshal(in, &list)
	case string:
		err = yaml.Unmarshal([]byte(in), &list)
	}

	return
}
