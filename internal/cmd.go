package internal

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spacecafe/miniplate/pkg"
)

func createIOStreams(input string, output string) (reader io.ReadCloser, writer io.WriteCloser, err error) {
	if input == "-" {
		reader = os.Stdin
	} else {
		reader, err = os.Open(filepath.Clean(input))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open file '%s': %w", input, err)
		}
	}

	if output == "-" {
		writer = os.Stdout
	} else {
		err = os.MkdirAll(filepath.Dir(output), 0750)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create output directory '%s': %w", output, err)
		}
		writer, err = os.Create(filepath.Clean(output))
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

	err = pkg.RenderTemplate(reader, writer, ctx)
	if err != nil {
		return fmt.Errorf("failed to render template '%s': %w", input, err)
	}
	return
}

func Main() {
	config := ParseFlags()

	ctx, err := pkg.LoadContexts(config.Contexts)
	if err != nil {
		log.Fatal(err)
	}

	if config.InputDir != "" {
		err = filepath.WalkDir(config.InputDir, func(inputFile string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			for _, exclude := range config.Excludes {
				if match, _ := filepath.Match(exclude, filepath.Base(inputFile)); match {
					return nil
				}
			}

			inputFileRelativePath, err := filepath.Rel(config.InputDir, inputFile)
			if err != nil {
				return fmt.Errorf("failed to get relative path: %w", err)
			}

			outputFile := filepath.Join(config.OutputDir, inputFileRelativePath)
			return renderTemplate(inputFile, outputFile, ctx)
		})
	} else {
		for i, input := range config.InputFiles {
			err = renderTemplate(input, config.OutputFiles[i], ctx)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
