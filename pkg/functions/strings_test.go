package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsFuncs_ShellQuote(t *testing.T) {
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
			got := StringsFuncs{}.ShellQuote(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStringsFuncs_Squote(t *testing.T) {
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
			got := StringsFuncs{}.Squote(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
