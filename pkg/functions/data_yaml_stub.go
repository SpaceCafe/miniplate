//go:build !yaml

package functions

func (DataFuncs) ToYAML(_ any) (string, error) {
	return "", ErrBuiltWithoutYAML
}

func (DataFuncs) YAML(_ any) (any, error) {
	return nil, ErrBuiltWithoutYAML
}

func (DataFuncs) YAMLArray(_ any) ([]any, error) {
	return nil, ErrBuiltWithoutYAML
}
