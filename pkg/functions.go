package pkg

import (
	"text/template"

	"github.com/SpaceCafe/miniplate/pkg/functions"
)

var FuncMap = template.FuncMap{
	"base64": func() any { return &functions.Base64Funcs{} },

	"conv":     func() any { return &functions.ConvFuncs{} },
	"bool":     functions.ConvFuncs{}.ToBool,
	"default":  functions.ConvFuncs{}.Default,
	"float":    functions.ConvFuncs{}.ToFloat64,
	"int":      functions.ConvFuncs{}.ToInt64,
	"join":     functions.ConvFuncs{}.Join,
	"string":   functions.ConvFuncs{}.ToString,
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
}
