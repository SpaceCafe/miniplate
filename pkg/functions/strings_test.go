package functions_test

import (
	"testing"

	"github.com/spacecafe/miniplate/pkg/functions"
	"github.com/stretchr/testify/assert"
)

func TestStringsFuncs_ShellQuote(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"Empty string", "", "''"},
		{"Regular string", "hello", "'hello'"},
		{"String with spaces", "hello world", "'hello world'"},
		{"Special characters", "$@&*", "'$@&*'"},
		{"String slice", []string{"hello", "world"}, "'hello' 'world'"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := functions.StringsFuncs{}.ShellQuote(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStringsFuncs_Squote(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"Empty string", "", "''"},
		{"Regular string", "hello", "'hello'"},
		{"String with spaces", "hello world", "'hello world'"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := functions.StringsFuncs{}.Squote(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
