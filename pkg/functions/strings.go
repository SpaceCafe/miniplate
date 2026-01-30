package functions

import (
	"fmt"
	"strings"

	"github.com/gosimple/slug"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type StringsFuncs struct{}

func (StringsFuncs) Contains(substr string, in any) bool {
	return strings.Contains(ConversionFuncs{}.ToString(in), substr)
}

func (StringsFuncs) HasPrefix(prefix string, in any) bool {
	return strings.HasPrefix(ConversionFuncs{}.ToString(in), prefix)
}

func (StringsFuncs) HasSuffix(suffix string, in any) bool {
	return strings.HasSuffix(ConversionFuncs{}.ToString(in), suffix)
}

func (StringsFuncs) Quote(in any) string {
	return fmt.Sprintf("%q", ConversionFuncs{}.ToString(in))
}

func (StringsFuncs) Repeat(n int, in any) string {
	return strings.Repeat(ConversionFuncs{}.ToString(in), n)
}

func (StringsFuncs) ReplaceAll(old, replacement string, in any) string {
	return strings.ReplaceAll(ConversionFuncs{}.ToString(in), old, replacement)
}

func (StringsFuncs) ShellQuote(in any) string {
	var builder strings.Builder

	v := ConversionFuncs{}.ToStrings(in)

	for i := range v {
		builder.WriteString(shellQuote(v[i]))

		if i < len(v)-1 {
			builder.WriteRune(' ')
		}
	}

	return builder.String()
}

func (StringsFuncs) Slug(in any) string {
	return slug.Make(ConversionFuncs{}.ToString(in))
}

func (StringsFuncs) Split(sep string, in any) []string {
	return strings.Split(ConversionFuncs{}.ToString(in), sep)
}

func (StringsFuncs) SplitN(sep string, n int, in any) []string {
	return strings.SplitN(ConversionFuncs{}.ToString(in), sep, n)
}

func (StringsFuncs) Squote(in any) string {
	v := ConversionFuncs{}.ToString(in)
	v = strings.ReplaceAll(v, `'`, `''`)

	return `'` + v + `'`
}

func (StringsFuncs) Title(in any) string {
	caser := cases.Title(language.Und)

	return caser.String(ConversionFuncs{}.ToString(in))
}

func (StringsFuncs) ToLower(in any) string {
	caser := cases.Lower(language.Und)

	return caser.String(ConversionFuncs{}.ToString(in))
}

func (StringsFuncs) ToUpper(in any) string {
	caser := cases.Upper(language.Und)

	return caser.String(ConversionFuncs{}.ToString(in))
}

func (StringsFuncs) Trim(cutset string, in any) string {
	return strings.Trim(ConversionFuncs{}.ToString(in), cutset)
}

func (StringsFuncs) TrimLeft(cutset string, in any) string {
	return strings.TrimLeft(ConversionFuncs{}.ToString(in), cutset)
}

func (StringsFuncs) TrimPrefix(prefix string, in any) string {
	return strings.TrimPrefix(ConversionFuncs{}.ToString(in), prefix)
}

func (StringsFuncs) TrimRight(cutset string, in any) string {
	return strings.TrimRight(ConversionFuncs{}.ToString(in), cutset)
}

func (StringsFuncs) TrimSpace(in any) string {
	return strings.TrimSpace(ConversionFuncs{}.ToString(in))
}

func (StringsFuncs) TrimSuffix(suffix string, in any) string {
	return strings.TrimSuffix(ConversionFuncs{}.ToString(in), suffix)
}

func (StringsFuncs) Trunc(n int, in any) string {
	return fmt.Sprintf("%.*s", n, ConversionFuncs{}.ToString(in))
}

func shellQuote(in string) string {
	return `'` + strings.ReplaceAll(in, `'`, `'"'"'`) + `'`
}
