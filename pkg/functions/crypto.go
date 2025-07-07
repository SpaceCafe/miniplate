package functions

import (
	"crypto/md5" // #nosec CWE-327 -- Only used to create hashes of data
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
)

type CryptoFuncs struct{}

func (CryptoFuncs) Bcrypt(args ...any) (string, error) {
	var (
		in   string
		cost = 10
	)
	if len(args) == 0 || len(args) > 2 {
		return "", ErrInvalidArgument
	}

	for i, arg := range args {
		switch v := arg.(type) {
		case int:
			if i == 0 {
				cost = v
			} else {
				return "", ErrInvalidArgument
			}
		case string:
			in = v
		default:
			return "", ErrInvalidArgument
		}
	}

	result, err := bcrypt.GenerateFromPassword([]byte(in), cost)
	return string(result), err
}

func (r CryptoFuncs) SHA224(in any) string {
	return fmt.Sprintf("%x", r.SHA224Bytes(in))
}

func (r CryptoFuncs) SHA256(in any) string {
	return fmt.Sprintf("%x", r.SHA256Bytes(in))
}

func (r CryptoFuncs) SHA384(in any) string {
	return fmt.Sprintf("%x", r.SHA384Bytes(in))
}

func (r CryptoFuncs) SHA512(in any) string {
	return fmt.Sprintf("%x", r.SHA512Bytes(in))
}

func (r CryptoFuncs) MD5(in any) string {
	return fmt.Sprintf("%x", r.MD5Bytes(in))
}

func (CryptoFuncs) SHA224Bytes(in any) [28]byte {
	return sha3.Sum224(ParseBytes(in))
}

func (CryptoFuncs) SHA256Bytes(in any) [32]byte {
	return sha3.Sum256(ParseBytes(in))
}

func (CryptoFuncs) SHA384Bytes(in any) [48]byte {
	return sha3.Sum384(ParseBytes(in))
}

func (CryptoFuncs) SHA512Bytes(in any) [64]byte {
	return sha3.Sum512(ParseBytes(in))
}

func (CryptoFuncs) MD5Bytes(in any) [16]byte {
	// #nosec CWE-328 -- Only used to create hashes of data
	return md5.Sum(ParseBytes(in))
}

func ParseBytes(in any) []byte {
	switch v := in.(type) {
	case []byte:
		return v
	case string:
		return []byte(v)
	default:
		return []byte{}
	}
}
