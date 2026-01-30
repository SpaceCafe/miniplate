package pkg

import (
	"text/template"

	"github.com/spacecafe/miniplate/pkg/functions"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"base64": func() any { return &functions.Base64Funcs{} },

		"coll":  func() any { return &functions.CollectionFuncs{} },
		"dict":  functions.CollectionFuncs{}.Dict,
		"slice": functions.CollectionFuncs{}.Slice,

		"conv":     func() any { return &functions.ConversionFuncs{} },
		"bool":     functions.ConversionFuncs{}.ToBool,
		"default":  functions.ConversionFuncs{}.Default,
		"float":    functions.ConversionFuncs{}.ToFloat64,
		"int":      functions.ConversionFuncs{}.ToInt64,
		"join":     functions.ConversionFuncs{}.Join,
		"string":   functions.ConversionFuncs{}.ToString,
		"urlParse": functions.ConversionFuncs{}.URL,

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

		"file":  func() any { return &functions.FileFuncs{} },
		"human": func() any { return &functions.HumanFuncs{} },
		"uuid":  func() any { return &functions.UUIDFuncs{} },

		"math": func() any { return &functions.MathFuncs{} },
		"add":  functions.MathFuncs{}.Add,
		"sub":  functions.MathFuncs{}.Sub,
		"mul":  functions.MathFuncs{}.Mul,
		"div":  functions.MathFuncs{}.Div,
		"rem":  functions.MathFuncs{}.Rem,
		"pow":  functions.MathFuncs{}.Pow,
		"seq":  functions.MathFuncs{}.Seq,

		"strings":    func() any { return &functions.StringsFuncs{} },
		"quote":      functions.StringsFuncs{}.Quote,
		"replaceAll": functions.StringsFuncs{}.ReplaceAll,
		"shellQuote": functions.StringsFuncs{}.ShellQuote,
		"squote":     functions.StringsFuncs{}.Squote,
		"title":      functions.StringsFuncs{}.Title,
		"toLower":    functions.StringsFuncs{}.ToLower,
		"toUpper":    functions.StringsFuncs{}.ToUpper,
		"trimSpace":  functions.StringsFuncs{}.TrimSpace,

		"tmpl": nil,
	}
}
