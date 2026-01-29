package functions

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/goccy/go-yaml"
)

type DataFuncs struct{}

func (DataFuncs) JSON(in any) (obj any, err error) {
	switch in := in.(type) {
	case []byte:
		err = json.Unmarshal(in, &obj)
	case string:
		err = json.Unmarshal([]byte(in), &obj)
	}
	return
}

func (DataFuncs) JSONArray(in any) (list []any, err error) {
	switch in := in.(type) {
	case []byte:
		err = json.Unmarshal(in, &list)
	case string:
		err = json.Unmarshal([]byte(in), &list)
	}
	return
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

func (DataFuncs) TOML(in any) (obj any, err error) {
	switch in := in.(type) {
	case []byte:
		err = toml.Unmarshal(in, &obj)
	case string:
		err = toml.Unmarshal([]byte(in), &obj)
	}
	return
}

func (DataFuncs) ToJSON(obj any) (string, error) {
	result, err := json.Marshal(obj)
	return string(result), err
}

func (DataFuncs) ToJSONPretty(indent string, obj any) (string, error) {
	result, err := json.MarshalIndent(obj, "", indent)
	return string(result), err
}

func (DataFuncs) ToYAML(obj any) (string, error) {
	result, err := yaml.Marshal(obj)
	return string(result), err
}

func (DataFuncs) ToTOML(obj any) (string, error) {
	result, err := toml.Marshal(obj)
	return string(result), err
}
