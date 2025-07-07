package pkg

import (
	"text/template"

	"github.com/SpaceCafe/miniplate/pkg/functions"
)

var FuncMap = template.FuncMap{
	"base64": func() any { return &functions.Base64Funcs{} },

	"conv":     func() any { return &functions.ConvFuncs{} },
	"bool":     functions.ConvFuncs{}.Bool,
	"default":  functions.ConvFuncs{}.Default,
	"join":     functions.ConvFuncs{}.Join,
	"urlParse": functions.ConvFuncs{}.URL,

	"crypto": func() any { return &functions.CryptoFuncs{} },

	"data":         func() any { return &functions.DataFuncs{} },
	"json":         functions.DataFuncs{}.JSON,
	"jsonArray":    functions.DataFuncs{}.JSONArray,
	"yaml":         functions.DataFuncs{}.YAML,
	"yamlArray":    functions.DataFuncs{}.YAMLArray,
	"toml":         functions.DataFuncs{}.TOML,
	"toJSON":       functions.DataFuncs{}.ToJSON,
	"toJSONPretty": functions.DataFuncs{}.ToJSONPretty,
	"toYAML":       functions.DataFuncs{}.ToYAML,
	"toTOML":       functions.DataFuncs{}.ToTOML,

	"env":    func() any { return &functions.EnvFuncs{} },
	"getenv": functions.EnvFuncs{}.Getenv,

	"file": func() any { return &functions.FileFuncs{} },
	"uuid": func() any { return &functions.UUIDFuncs{} },
}
