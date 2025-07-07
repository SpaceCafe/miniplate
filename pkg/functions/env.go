package functions

import (
	"os"
)

type EnvFuncs struct{}

func (EnvFuncs) Getenv(key string, def ...string) string {
	fileFuncs := &FileFuncs{}
	value := os.Getenv(key)
	if value == "" {
		valueFile := os.Getenv(key + "_FILE")
		if valueFile != "" && fileFuncs.IsFile(valueFile) {
			value, _ = fileFuncs.Read(valueFile)
		}
	}
	if value == "" && len(def) > 0 {
		return def[0]
	}
	return value
}
