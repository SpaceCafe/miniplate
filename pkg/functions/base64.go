package functions

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Base64Funcs struct{}

func (Base64Funcs) Encode(in any) (text string, err error) {
	switch v := in.(type) {
	case []byte:
		text = base64.StdEncoding.EncodeToString(v)
	case string:
		text = base64.StdEncoding.EncodeToString([]byte(v))
	case fmt.Stringer:
		text = base64.StdEncoding.EncodeToString([]byte(v.String()))
	default:
		err = fmt.Errorf("unsupported type: %T", in)
	}
	return
}

func (Base64Funcs) Decode(in string) (text string, err error) {
	in = strings.TrimSpace(in)
	value, err := base64.StdEncoding.DecodeString(in)
	return string(value), err
}

func (Base64Funcs) DecodeBytes(in string) (bytes []byte, err error) {
	return base64.StdEncoding.DecodeString(in)
}
