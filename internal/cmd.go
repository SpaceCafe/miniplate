package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spacecafe/miniplate/pkg"
)

func Main() {
	config := ParseFlags()

	ctx, err := pkg.LoadContexts(config.Contexts)
	if err != nil {
		log.Fatal(err)
	}

	if config.InputDir == "" {
		for i, input := range config.InputFiles {
			err = (&pkg.Renderer{}).Render(input, config.OutputFiles[i], ctx)
			if err != nil {
				log.Fatal(err)

				return
			}
		}

		return
	}

	err = filepath.WalkDir(
		config.InputDir,
		func(inputFile string, d os.DirEntry, err error) error {
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

			return (&pkg.Renderer{}).Render(inputFile, outputFile, ctx)
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
