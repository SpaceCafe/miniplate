package functions

import (
	"encoding/json"
	"encoding/xml"
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

func (DataFuncs) ToJSON(obj any) (string, error) {
	result, err := json.Marshal(obj)

	return string(result), err
}

func (DataFuncs) ToJSONPretty(indent string, obj any) (string, error) {
	result, err := json.MarshalIndent(obj, "", indent)

	return string(result), err
}

func (DataFuncs) ToXML(obj any) (string, error) {
	result, err := xml.Marshal(obj)

	return string(result), err
}

func (DataFuncs) ToXMLPretty(indent string, obj any) (string, error) {
	result, err := xml.MarshalIndent(obj, "", indent)

	return string(result), err
}

func (DataFuncs) XML(in any) (obj any, err error) {
	switch in := in.(type) {
	case []byte:
		err = xml.Unmarshal(in, &obj)
	case string:
		err = xml.Unmarshal([]byte(in), &obj)
	}

	return
}
