package internal

import (
	"io"
	"log"
	"os"
	"path"

	"github.com/SpaceCafe/miniplate/pkg"
)

func createIOStreams(input string, output string) (reader io.ReadCloser, writer io.WriteCloser, err error) {
	if input == "-" {
		reader = os.Stdin
	} else {
		reader, err = os.Open(path.Clean(input))
	}

	if output == "-" {
		writer = os.Stdout
	} else {
		err = os.MkdirAll(path.Dir(output), 0750)
		if err != nil {
			return nil, nil, err
		}
		writer, err = os.Create(path.Clean(output))
	}

	return
}

func renderTemplate(input string, output string, ctx map[string]any) (err error) {
	reader, writer, err := createIOStreams(input, output)
	if err != nil {
		return
	}
	defer func() { _ = reader.Close() }()
	defer func() { _ = writer.Close() }()

	return pkg.RenderTemplate(reader, writer, ctx)
}

func Main() {
	config := ParseFlags()

	ctx, err := pkg.LoadContexts(config.Contexts)
	if err != nil {
		log.Fatal(err)
	}

	for i, input := range config.InputFiles {
		err := renderTemplate(input, config.OutputFiles[i], ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
}
