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
	return strings.Contains(ConvFuncs{}.ToString(in), substr)
}

func (StringsFuncs) HasPrefix(prefix string, in any) bool {
	return strings.HasPrefix(ConvFuncs{}.ToString(in), prefix)
}

func (StringsFuncs) HasSuffix(suffix string, in any) bool {
	return strings.HasSuffix(ConvFuncs{}.ToString(in), suffix)
}

func (StringsFuncs) Split(sep string, in any) []string {
	return strings.Split(ConvFuncs{}.ToString(in), sep)
}

func (StringsFuncs) SplitN(sep string, n int, in any) []string {
	return strings.SplitN(ConvFuncs{}.ToString(in), sep, n)
}

func (StringsFuncs) Quote(in any) string {
	return fmt.Sprintf("%q", ConvFuncs{}.ToString(in))
}

func (StringsFuncs) Repeat(n int, in any) string {
	return strings.Repeat(ConvFuncs{}.ToString(in), n)
}

func (StringsFuncs) ReplaceAll(old, new string, in any) string {
	return strings.ReplaceAll(ConvFuncs{}.ToString(in), old, new)
}

func (StringsFuncs) Slug(in any) string {
	return slug.Make(ConvFuncs{}.ToString(in))
}

func (StringsFuncs) ShellQuote(in any) string {
	var (
		builder strings.Builder
	)
	v := ConvFuncs{}.ToStrings(in)

	for i := range v {
		builder.WriteString(shellQuote(v[i]))
		if i < len(v)-1 {
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}

func (StringsFuncs) Squote(in any) string {
	v := ConvFuncs{}.ToString(in)
	v = strings.ReplaceAll(v, `'`, `''`)
	return `'` + v + `'`
}

func (StringsFuncs) Title(in any) string {
	caser := cases.Title(language.Und)
	return caser.String(ConvFuncs{}.ToString(in))
}

func (StringsFuncs) ToLower(in any) string {
	caser := cases.Lower(language.Und)
	return caser.String(ConvFuncs{}.ToString(in))
}

func (StringsFuncs) ToUpper(in any) string {
	caser := cases.Upper(language.Und)
	return caser.String(ConvFuncs{}.ToString(in))
}

func (StringsFuncs) Trim(cutset string, in any) string {
	return strings.Trim(ConvFuncs{}.ToString(in), cutset)
}

func (StringsFuncs) TrimLeft(cutset string, in any) string {
	return strings.TrimLeft(ConvFuncs{}.ToString(in), cutset)
}

func (StringsFuncs) TrimPrefix(prefix string, in any) string {
	return strings.TrimPrefix(ConvFuncs{}.ToString(in), prefix)
}

func (StringsFuncs) TrimRight(cutset string, in any) string {
	return strings.TrimRight(ConvFuncs{}.ToString(in), cutset)
}

func (StringsFuncs) TrimSpace(in any) string {
	return strings.TrimSpace(ConvFuncs{}.ToString(in))
}

func (StringsFuncs) TrimSuffix(suffix string, in any) string {
	return strings.TrimSuffix(ConvFuncs{}.ToString(in), suffix)
}

func (StringsFuncs) Trunc(n int, in any) string {
	return fmt.Sprintf("%.*s", n, ConvFuncs{}.ToString(in))
}

func shellQuote(in string) string {
	return `'` + strings.ReplaceAll(in, `'`, `'"'"'`) + `'`
}
