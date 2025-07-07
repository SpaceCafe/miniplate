package pkg

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/SpaceCafe/miniplate/pkg/functions"
)

func LoadContexts(files []string) (ctx map[string]any, err error) {
	ctx = make(map[string]any)
	dataFuncs := &functions.DataFuncs{}
	for _, file := range files {
		var (
			rawData []byte
			data    map[string]any
		)

		parts := strings.SplitN(file, "=", 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return ctx, fmt.Errorf("invalid file format: %s", file)
		}

		filePath := path.Clean(parts[1])
		rawData, err = os.ReadFile(path.Clean(filePath))
		if err != nil {
			return
		}
		switch path.Ext(filePath) {
		case ".json":
			data, err = dataFuncs.JSON(rawData)
		case ".toml":
			data, err = dataFuncs.TOML(rawData)
		case ".yaml":
			data, err = dataFuncs.YAML(rawData)
		}
		if err != nil {
			return
		}
		ctx[parts[0]] = data
	}
	return
}
