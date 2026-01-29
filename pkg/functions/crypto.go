package functions

import (
	"crypto/md5" // #nosec CWE-327 -- Only used to create hashes of data
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

type CryptoFuncs struct{}

func (f CryptoFuncs) Bcrypt(args ...any) (string, error) {
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

func (f CryptoFuncs) MD5(in any) string {
	return fmt.Sprintf("%x", f.MD5Bytes(in))
}

func (CryptoFuncs) MD5Bytes(in any) [16]byte {
	// #nosec CWE-328 -- Only used to create hashes of data
	return md5.Sum(ParseBytes(in))
}

func (f CryptoFuncs) SHA224(in any) string {
	return fmt.Sprintf("%x", f.SHA224Bytes(in))
}

func (CryptoFuncs) SHA224Bytes(in any) [28]byte {
	return sha3.Sum224(ParseBytes(in))
}

func (f CryptoFuncs) SHA256(in any) string {
	return fmt.Sprintf("%x", f.SHA256Bytes(in))
}

func (CryptoFuncs) SHA256Bytes(in any) [32]byte {
	return sha3.Sum256(ParseBytes(in))
}

func (f CryptoFuncs) SHA384(in any) string {
	return fmt.Sprintf("%x", f.SHA384Bytes(in))
}

func (CryptoFuncs) SHA384Bytes(in any) [48]byte {
	return sha3.Sum384(ParseBytes(in))
}

func (f CryptoFuncs) SHA512(in any) string {
	return fmt.Sprintf("%x", f.SHA512Bytes(in))
}

func (CryptoFuncs) SHA512Bytes(in any) [64]byte {
	return sha3.Sum512(ParseBytes(in))
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
