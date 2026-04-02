//go:build !toml

package functions

func (DataFuncs) TOML(_ any) (any, error) {
	return nil, ErrBuiltWithoutTOML
}

func (DataFuncs) ToTOML(_ any) (string, error) {
	return "", ErrBuiltWithoutTOML
}
