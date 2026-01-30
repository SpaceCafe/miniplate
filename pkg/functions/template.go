package functions

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"
)

type TemplateFuncs struct {
	tmpl     *template.Template
	ctx      any
	filePath string
}

func NewTemplateFuncs(
	tmpl *template.Template,
	data map[string]any,
	filePath string,
) *TemplateFuncs {
	return &TemplateFuncs{tmpl, data, filePath}
}

func (f *TemplateFuncs) Exec(name string, ctx ...any) (string, error) {
	currentCtx := f.ctx
	if len(ctx) == 1 {
		currentCtx = ctx[0]
	}

	tmpl := f.tmpl.Lookup(name)
	if tmpl == nil {
		return "", fmt.Errorf("%w: %s", ErrUndefinedTemplate, name)
	}

	return render(tmpl, currentCtx)
}

func (f *TemplateFuncs) Inline(in ...any) (string, error) {
	var (
		input string
		ctx   any
	)

	name := "<inline>"

	switch len(in) {
	case 0:
		return "", ErrInvalidArgument
	case 1:
		input = ConversionFuncs{}.ToString(in[0])
	case 2:
		switch v := in[1].(type) {
		case string:
			name = ConversionFuncs{}.ToString(in[0])
			input = v
		default:
			input = ConversionFuncs{}.ToString(in[0])
			ctx = v
		}
	case 3:
		name = ConversionFuncs{}.ToString(in[0])
		input = ConversionFuncs{}.ToString(in[1])
		ctx = in[2]
	}

	tmpl, err := f.tmpl.New(name).Parse(input)
	if err != nil {
		return "", err
	}

	return render(tmpl, ctx)
}

func (f *TemplateFuncs) Path() (string, error) {
	return f.filePath, nil
}

func (f *TemplateFuncs) PathDir() (string, error) {
	if f.filePath == "" {
		return "", nil
	}

	return filepath.Dir(f.filePath), nil
}

func render(tmpl *template.Template, ctx any) (string, error) {
	buf := &bytes.Buffer{}

	err := tmpl.Execute(buf, ctx)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
