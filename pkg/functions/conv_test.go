package functions_test

import (
	"testing"

	"github.com/spacecafe/miniplate/pkg/functions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvFuncs_ToBool(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input   any
		name    string
		want    bool
		wantErr bool
	}{
		// Basic type conversions
		{name: "bool true", input: true, want: true, wantErr: false},
		{name: "bool false", input: false, want: false, wantErr: false},

		{name: "int positive", input: 42, want: true, wantErr: false},
		{name: "int zero", input: 0, want: false, wantErr: false},
		{name: "int negative", input: -1, want: true, wantErr: false},

		{name: "float64 positive", input: 3.14, want: true, wantErr: false},
		{name: "float64 zero", input: float64(0), want: false, wantErr: false},
		{name: "float64 negative", input: -2.718, want: true, wantErr: false},

		// String conversions
		{name: "string 1", input: "1", want: true, wantErr: false},
		{name: "string t", input: "t", want: true, wantErr: false},
		{name: "string y", input: "y", want: true, wantErr: false},
		{name: "string true", input: "true", want: true, wantErr: false},
		{name: "string yes", input: "yes", want: true, wantErr: false},
		{name: "string on", input: "on", want: true, wantErr: false},
		{name: "string 0", input: "0", want: false, wantErr: false},
		{name: "string f", input: "f", want: false, wantErr: false},
		{name: "string n", input: "n", want: false, wantErr: false},
		{name: "string false", input: "false", want: false, wantErr: false},
		{name: "string no", input: "no", want: false, wantErr: false},
		{name: "string off", input: "off", want: false, wantErr: false},
		{name: "string invalid", input: "invalid", want: false, wantErr: true},

		// Empty string case
		{name: "empty string", input: "", want: false, wantErr: true},

		// Invalid types
		{name: "nil input", input: nil, want: false, wantErr: true},
	}

	c := functions.ConvFuncs{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := c.ToBool(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestConvFuncs_ToBools(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   []any
		want    []bool
		wantErr bool
	}{
		{name: "no args", input: []any{}, want: []bool{}},
		{name: "valid args", input: []any{"t", "off"}, want: []bool{true, false}},
		{name: "invalid args", input: []any{"t", "off", "invalid"}, wantErr: true},
	}

	c := functions.ConvFuncs{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := c.ToBools(tt.input...)
			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestConvFuncs_ToInt64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input   any
		name    string
		want    int64
		wantErr bool
	}{
		// Basic type conversions
		{name: "int zero", input: 0, want: 0, wantErr: false},
		{name: "int positive", input: 42, want: 42, wantErr: false},
		{name: "int negative", input: -17, want: -17, wantErr: false},
		{name: "uint positive", input: uint(99), want: 99, wantErr: false},
		{name: "int8 positive", input: int8(32), want: 32, wantErr: false},
		{name: "int8 negative", input: int8(-45), want: -45, wantErr: false},
		{name: "int16 positive", input: int16(1_000), want: 1_000, wantErr: false},
		{name: "int16 negative", input: int16(-3_000), want: -3_000, wantErr: false},
		{name: "int32 positive", input: int32(500_000), want: 500_000, wantErr: false},
		{name: "int32 negative", input: int32(-900_000), want: -900_000, wantErr: false},
		{name: "int64 positive", input: int64(1_000_000_000), want: 1_000_000_000, wantErr: false},
		{
			name:    "int64 negative",
			input:   int64(-2_000_000_000),
			want:    -2_000_000_000,
			wantErr: false,
		},
		{name: "uint8 positive", input: uint8(23), want: 23, wantErr: false},
		{name: "uint16 positive", input: uint16(5_000), want: 5_000, wantErr: false},
		{name: "uint32 positive", input: uint32(700_000), want: 700_000, wantErr: false},
		{name: "uint32 max", input: uint32(4_294_967_295), want: 4_294_967_295, wantErr: false},
		{
			name:    "uint64 positive",
			input:   uint64(10_000_000_000),
			want:    10_000_000_000,
			wantErr: false,
		},
		{name: "uint64 max", input: uint64(18_446_744_073_709_551_615), want: 0, wantErr: true},

		// Float conversions
		{name: "float32 positive", input: float32(3.14), want: 3, wantErr: false},
		{name: "float32 negative", input: float32(-2.718), want: -2, wantErr: false},
		{name: "float64 positive", input: 9.5, want: 9, wantErr: false},
		{name: "float64 negative", input: -8.3, want: -8, wantErr: false},

		// String conversions
		{name: "string zero", input: "0", want: 0, wantErr: false},
		{name: "string positive", input: "123", want: 123, wantErr: false},
		{name: "string negative", input: "-975", want: -975, wantErr: false},
		{name: "string invalid", input: "invalid", want: 0, wantErr: true},

		// Empty string case
		{name: "empty string", input: "", want: 0, wantErr: true},

		// Invalid types
		{name: "nil input", input: nil, want: 0, wantErr: true},
	}

	c := functions.ConvFuncs{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := c.ToInt64(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestConvFuncs_ToInt64s(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   []any
		want    []int64
		wantErr bool
	}{
		{name: "no args", input: []any{}, want: []int64{}},
		{name: "valid args", input: []any{"23", 42.0}, want: []int64{23, 42}},
		{name: "invalid args", input: []any{"23", 42.0, "invalid"}, wantErr: true},
	}

	c := functions.ConvFuncs{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := c.ToInt64s(tt.input...)
			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestConvFuncs_ToFloat64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   []any
		want    float64
		wantErr bool
	}{
		// Basic type conversions
		{name: "float32 positive", input: []any{float32(3.14)}, want: 3.14, wantErr: false},
		{name: "float32 negative", input: []any{float32(-2.718)}, want: -2.718, wantErr: false},
		{name: "float64 positive", input: []any{9.5}, want: 9.5, wantErr: false},
		{name: "float64 zero", input: []any{0.0}, want: 0.0, wantErr: false},
		{name: "float64 negative", input: []any{-8.3}, want: -8.3, wantErr: false},

		// Integer conversions
		{name: "int positive", input: []any{42}, want: 42.0, wantErr: false},
		{name: "int zero", input: []any{0}, want: 0.0, wantErr: false},
		{name: "int negative", input: []any{-17}, want: -17.0, wantErr: false},
		{name: "uint positive", input: []any{uint(99)}, want: 99.0, wantErr: false},

		// String conversions
		{name: "string decimal", input: []any{"3.14"}, want: 3.14, wantErr: false},
		{name: "string integer", input: []any{"42"}, want: 42.0, wantErr: false},
		{name: "string negative", input: []any{"-2.718"}, want: -2.718, wantErr: false},
		{name: "string separator", input: []any{",", "-2,718"}, want: -2.718, wantErr: false},
		{name: "string invalid", input: []any{"invalid"}, want: 0.0, wantErr: true},
		{name: "string invalid args", input: []any{"3,14", 23}, want: 0.0, wantErr: true},
		{name: "no args", input: []any{}, want: 0.0, wantErr: true},
		{name: "too many args", input: []any{"", "", ""}, want: 0.0, wantErr: true},

		// Empty string case
		{name: "empty string", input: []any{""}, want: 0.0, wantErr: true},

		// Invalid types
		{name: "nil input", input: []any{nil}, want: 0.0, wantErr: true},
	}

	c := functions.ConvFuncs{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := c.ToFloat64(tt.input...)
			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				assert.InDelta(t, tt.want, got, 0.000_001)
			}
		})
	}
}

func TestConvFuncs_ToFloat64s(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   []any
		want    []float64
		wantErr bool
	}{
		{name: "no args", input: []any{}, want: []float64{}},
		{name: "valid args", input: []any{"3.14", 42}, want: []float64{3.14, 42.0}},
		{name: "valid separator", input: []any{",", "3,14", 42}, want: []float64{3.14, 42.0}},
		{name: "invalid args", input: []any{"3.14", 42, "invalid"}, wantErr: true},
	}

	c := functions.ConvFuncs{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := c.ToFloat64s(tt.input...)
			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				assert.InDeltaSlice(t, tt.want, got, 0.000_001)
			}
		})
	}
}

func TestConvFuncs_ToString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input any
		want  string
	}{
		// Basic type conversions
		{name: "string", input: "hello", want: "hello"},
		{name: "int", input: 42, want: "42"},
		{name: "float32", input: float32(3.14), want: "3.14"},
		{name: "float64", input: 9.5, want: "9.5"},
		{name: "bool true", input: true, want: "true"},
		{name: "bool false", input: false, want: "false"},

		// Slices
		{name: "slice of strings", input: []string{"a", "b"}, want: "[a b]"},
		{name: "empty slice", input: []string{}, want: "[]"},

		// Maps
		{name: "map with values", input: map[string]int{"key": 42}, want: "map[key:42]"},
		{name: "empty map", input: map[string]int{}, want: "map[]"},

		// Other types
		{name: "nil input", input: nil, want: ""},
	}

	c := functions.ConvFuncs{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := c.ToString(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvFuncs_ToStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input any
		want  []string
	}{
		{name: "one string", input: "hello", want: []string{"hello"}},
		{name: "string slice", input: []string{"hello"}, want: []string{"hello"}},
		{name: "string and int", input: []any{"hello", 42}, want: []string{"hello", "42"}},
		{name: "int slice", input: []int{23, 42}, want: []string{"23", "42"}},
	}

	c := functions.ConvFuncs{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := c.ToStrings(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvFuncs_Join(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		want  string
		input []any
	}{
		{name: "no args", input: []any{}, want: ""},
		{name: "only separator", input: []any{","}, want: ""},
		{name: "one value", input: []any{"foo", ","}, want: "foo"},
		{name: "slice value", input: []any{[]any{"3.14", 42}, "|"}, want: "3.14|42"},
		{name: "values", input: []any{"3.14", 42, "|"}, want: "3.14|42"},
	}

	c := functions.ConvFuncs{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := c.Join(tt.input...)
			assert.Equal(t, tt.want, got)
		})
	}

	t.Run("string slice value", func(t *testing.T) {
		t.Parallel()

		got := c.Join([]string{"3.14", "16"}, "|")
		assert.Equal(t, "3.14|16", got)
	})
}

func TestConvFuncs_Default(t *testing.T) {
	t.Parallel()

	tests := []struct {
		def   any
		input any
		want  any
		name  string
	}{
		{name: "bool false", def: true, input: false, want: true},
		{name: "bool true", def: false, input: true, want: true},
		{name: "int zero", def: 42, input: 0, want: 42},
		{name: "int", def: 42, input: 23, want: 23},
		{name: "empty string", def: "default", input: "", want: "default"},
		{name: "string", def: "default", input: "test", want: "test"},
		{
			name:  "empty slice",
			def:   []string{"default"},
			input: []string{},
			want:  []string{"default"},
		},
		{name: "slice", def: []string{"default"}, input: []string{"test"}, want: []string{"test"}},
		{
			name:  "empty map",
			def:   map[string]int{"default": 42},
			input: map[string]int{},
			want:  map[string]int{"default": 42},
		},
		{
			name:  "map",
			def:   map[string]int{"default": 42},
			input: map[string]int{"test": 23},
			want:  map[string]int{"test": 23},
		},
		{name: "nil", def: "default", input: nil, want: "default"},
	}
	c := functions.ConvFuncs{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := c.Default(tt.def, tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
