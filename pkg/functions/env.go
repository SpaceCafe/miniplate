package functions

import (
	"fmt"
	"os"
	"strings"
)

type EnvFuncs struct{}

func (EnvFuncs) Getenv(key string, def ...string) (out string) {
	out, _ = EnvFuncs{}.MustGetenv(key, def...)
	return
}

func (EnvFuncs) MustGetenv(key string, def ...string) (string, error) {
	v, ok := os.LookupEnv(key)
	if ok {
		return strings.TrimSpace(v), nil
	}

	v, ok = os.LookupEnv(key + "_FILE")
	if ok {
		if (FileFuncs{}.IsFile(strings.TrimSpace(v))) {
			out, err := FileFuncs{}.Read(v)
			return strings.TrimSpace(out), err
		}
		return "", fmt.Errorf("environment variable %s_FILE is not a file", key)
	}

	if len(def) > 0 {
		return def[0], nil
	} else {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
}
