package pkg

import (
	"io"
	"text/template"
)

func RenderTemplate(reader io.Reader, writer io.Writer, data map[string]any) (err error) {
	var (
		buf  []byte
		tmpl *template.Template
	)

	buf, err = io.ReadAll(reader)
	if err != nil {
		return err
	}

	tmpl, err = template.New("render").Funcs(FuncMap).Parse(string(buf))
	if err != nil {
		return err
	}

	return tmpl.Execute(writer, data)
}
