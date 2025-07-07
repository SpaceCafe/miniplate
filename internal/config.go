package internal

import (
	"flag"
	"fmt"
	"log"
)

type Config struct {
	InputFiles  flagStringArray
	OutputFiles flagStringArray
	Excludes    flagStringArray
	Contexts    flagStringArray
	InputDir    string
	OutputDir   string
}

func (r *Config) ApplyDefaults() {
	if len(r.InputFiles) == 0 {
		r.InputFiles = []string{"-"}
	}
	if len(r.OutputFiles) == 0 {
		r.OutputFiles = []string{"-"}
	}
}

func (r *Config) Validate() error {
	if len(r.InputFiles) != len(r.OutputFiles) {
		return fmt.Errorf("input files and output files have different length")
	}
	return nil
}

type flagStringArray []string

func (r *flagStringArray) String() string {
	return fmt.Sprintf("%v", *r)
}

func (r *flagStringArray) Set(value string) error {
	*r = append(*r, value)
	return nil
}

func registerFlagAliases(name string, aliases ...string) {
	for _, v := range aliases {
		flagSet := flag.Lookup(name)
		flag.Var(flagSet.Value, v, fmt.Sprintf("Alias to -%s", flagSet.Name))
	}
}

func ParseFlags() *Config {
	config := &Config{}

	flag.Var(&config.InputFiles, "in", "Path to a specific input template file. The special value '-' means 'Stdin'. (default \"-\")")
	registerFlagAliases("in", "i", "file", "f")

	flag.Var(&config.OutputFiles, "out", "Path to save output to file. The special value '-' means 'Stdout'. (default \"-\")")
	registerFlagAliases("out", "o")

	flag.StringVar(&config.InputDir, "input-dir", "", "Path to input directory, where all files will be processed recursively as templates.")
	flag.StringVar(&config.OutputDir, "output-dir", "", "Path to output directory where all resulting files will be stored.")
	flag.Var(&config.Excludes, "exclude", "When using the -input-dir argument, it can be useful to filter which files are processed.")

	flag.Var(&config.Contexts, "context", "Add a data source in 'name=URL' form, and make it available in the default context as '.<name>'.")

	flag.Parse()
	config.ApplyDefaults()
	if err := config.Validate(); err != nil {
		log.Fatal(err)
	}
	return config
}
