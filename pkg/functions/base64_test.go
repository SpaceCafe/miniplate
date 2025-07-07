package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64Funcs_Decode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "valid base64", input: "SGVsbG8sIFdvcmxkIQ==", want: "Hello, World!", wantErr: false},
		{name: "invalid base64", input: "InvalidBase64==", wantErr: true},
		{name: "empty string", input: "", want: "", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64Funcs{}.Decode(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestBase64Funcs_DecodeBytes(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []byte
		wantErr bool
	}{
		{name: "valid base64", input: "SGVsbG8sIFdvcmxkIQ==", want: []byte("Hello, World!"), wantErr: false},
		{name: "invalid base64", input: "InvalidBase64==", wantErr: true},
		{name: "empty string", input: "", want: []byte(""), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64Funcs{}.DecodeBytes(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestBase64Funcs_Encode(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{name: "String Input", input: "Hello, World!", want: "SGVsbG8sIFdvcmxkIQ==", wantErr: false},
		{name: "Byte Array Input", input: []byte("Hello, World!"), want: "SGVsbG8sIFdvcmxkIQ==", wantErr: false},
		{name: "Empty String", input: "", want: "", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64Funcs{}.Encode(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
