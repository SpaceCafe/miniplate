package pkg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spacecafe/miniplate/pkg/functions"
)

var (
	ErrOpenFile  = errors.New("failed to open file")
	ErrCreateDir = errors.New("failed to create output directory")
)

type Renderer struct {
	tmpl   *template.Template
	input  string
	output string
}

func (r *Renderer) Render(input, output string, ctx map[string]any) error {
	r.init(input, output, ctx)

	reader, writer, cleanup, err := r.createIOStreams()
	defer cleanup()

	if err != nil {
		return err
	}

	var buf []byte

	buf, err = io.ReadAll(reader)
	if err != nil {
		return err
	}

	_, err = r.tmpl.Parse(string(buf))
	if err != nil {
		return err
	}

	return r.tmpl.Execute(writer, ctx)
}

func (r *Renderer) createIOStreams() (io.ReadCloser, io.WriteCloser, func(), error) {
	var (
		reader io.ReadCloser
		writer io.WriteCloser
		err    error
	)

	if r.input == "-" {
		reader = os.Stdin
	} else {
		reader, err = os.Open(filepath.Clean(r.input))
		if err != nil {
			return nil, nil, nil, fmt.Errorf("%w: %s: %w", ErrOpenFile, r.input, err)
		}
	}

	if r.output == "-" {
		writer = os.Stdout
	} else {
		err = os.MkdirAll(filepath.Dir(r.output), 0o750)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("%w: %s: %w", ErrCreateDir, r.output, err)
		}

		writer, err = os.Create(filepath.Clean(r.output))
		if err != nil {
			return nil, nil, nil, fmt.Errorf("%w: %s: %w", ErrOpenFile, r.output, err)
		}
	}

	return reader, writer, func() { _ = reader.Close(); _ = writer.Close() }, nil
}

func (r *Renderer) init(input, output string, ctx map[string]any) {
	r.tmpl = template.New(input)
	r.input = input
	r.output = output

	funcMap := FuncMap()
	funcMap["tmpl"] = func() any { return functions.NewTemplateFuncs(r.tmpl, ctx, input) }

	r.tmpl.Funcs(funcMap)
}
