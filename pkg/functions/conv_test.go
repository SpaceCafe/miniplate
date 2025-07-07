package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvFuncs_ToBool(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    bool
		wantErr bool
	}{
		// Basic type conversions
		{"bool true", true, true, false},
		{"bool false", false, false, false},

		{"int positive", 42, true, false},
		{"int zero", 0, false, false},
		{"int negative", -1, true, false},

		{"float64 positive", 3.14, true, false},
		{"float64 zero", float64(0), false, false},
		{"float64 negative", -2.718, true, false},

		// String conversions
		{"string 1", "1", true, false},
		{"string t", "t", true, false},
		{"string y", "y", true, false},
		{"string true", "true", true, false},
		{"string yes", "yes", true, false},
		{"string on", "on", true, false},
		{"string 0", "0", false, false},
		{"string f", "f", false, false},
		{"string n", "n", false, false},
		{"string false", "false", false, false},
		{"string no", "no", false, false},
		{"string off", "off", false, false},
		{"string invalid", "invalid", false, true},

		// Empty string case
		{"empty string", "", false, true},

		// Invalid types
		{"nil input", nil, false, true},
	}

	c := ConvFuncs{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ToBool(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestConvFuncs_ToBools(t *testing.T) {
	tests := []struct {
		name    string
		input   []any
		want    []bool
		wantErr bool
	}{
		{"no args", []any{}, []bool{}, false},
		{"valid args", []any{"t", "off"}, []bool{true, false}, false},
		{"invalid args", []any{"t", "off", "invalid"}, nil, true},
	}

	c := ConvFuncs{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ToBools(tt.input...)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestConvFuncs_ToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    int64
		wantErr bool
	}{
		// Basic type conversions
		{"int zero", 0, 0, false},
		{"int positive", 42, 42, false},
		{"int negative", -17, -17, false},
		{"uint positive", uint(99), 99, false},
		{"int8 positive", int8(32), 32, false},
		{"int8 negative", int8(-45), -45, false},
		{"int16 positive", int16(1_000), 1_000, false},
		{"int16 negative", int16(-3_000), -3_000, false},
		{"int32 positive", int32(500_000), 500_000, false},
		{"int32 negative", int32(-900_000), -900_000, false},
		{"int64 positive", int64(1_000_000_000), 1_000_000_000, false},
		{"int64 negative", int64(-2_000_000_000), -2_000_000_000, false},
		{"uint8 positive", uint8(23), 23, false},
		{"uint16 positive", uint16(5_000), 5_000, false},
		{"uint32 positive", uint32(700_000), 700_000, false},
		{"uint32 max", uint32(4_294_967_295), 4_294_967_295, false},
		{"uint64 positive", uint64(10_000_000_000), 10_000_000_000, false},
		{"uint64 max", uint64(18_446_744_073_709_551_615), 0, true},

		// Float conversions
		{"float32 positive", float32(3.14), 3, false},
		{"float32 negative", float32(-2.718), -2, false},
		{"float64 positive", 9.5, 9, false},
		{"float64 negative", -8.3, -8, false},

		// String conversions
		{"string zero", "0", 0, false},
		{"string positive", "123", 123, false},
		{"string negative", "-975", -975, false},
		{"string invalid", "invalid", 0, true},

		// Empty string case
		{"empty string", "", 0, true},

		// Invalid types
		{"nil input", nil, 0, true},
	}

	c := ConvFuncs{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ToInt64(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestConvFuncs_ToInt64s(t *testing.T) {
	tests := []struct {
		name    string
		input   []any
		want    []int64
		wantErr bool
	}{
		{"no args", []any{}, []int64{}, false},
		{"valid args", []any{"23", 42.0}, []int64{23, 42}, false},
		{"invalid args", []any{"23", 42.0, "invalid"}, nil, true},
	}

	c := ConvFuncs{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ToInt64s(tt.input...)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestConvFuncs_ToFloat64(t *testing.T) {
	tests := []struct {
		name    string
		input   []any
		want    float64
		wantErr bool
	}{
		// Basic type conversions
		{"float32 positive", []any{float32(3.14)}, 3.14, false},
		{"float32 negative", []any{float32(-2.718)}, -2.718, false},
		{"float64 positive", []any{9.5}, 9.5, false},
		{"float64 zero", []any{0.0}, 0.0, false},
		{"float64 negative", []any{-8.3}, -8.3, false},

		// Integer conversions
		{"int positive", []any{42}, 42.0, false},
		{"int zero", []any{0}, 0.0, false},
		{"int negative", []any{-17}, -17.0, false},
		{"uint positive", []any{uint(99)}, 99.0, false},

		// String conversions
		{"string decimal", []any{"3.14"}, 3.14, false},
		{"string integer", []any{"42"}, 42.0, false},
		{"string negative", []any{"-2.718"}, -2.718, false},
		{"string separator", []any{",", "-2,718"}, -2.718, false},
		{"string invalid", []any{"invalid"}, 0.0, true},
		{"string invalid args", []any{"3,14", 23}, 0.0, true},
		{"no args", []any{}, 0.0, true},
		{"too many args", []any{"", "", ""}, 0.0, true},

		// Empty string case
		{"empty string", []any{""}, 0.0, true},

		// Invalid types
		{"nil input", []any{nil}, 0.0, true},
	}

	c := ConvFuncs{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ToFloat64(tt.input...)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDelta(t, tt.want, got, 0.000_001)
			}
		})
	}
}

func TestConvFuncs_ToFloat64s(t *testing.T) {
	tests := []struct {
		name    string
		input   []any
		want    []float64
		wantErr bool
	}{
		{"no args", []any{}, []float64{}, false},
		{"valid args", []any{"3.14", 42}, []float64{3.14, 42.0}, false},
		{"valid separator", []any{",", "3,14", 42}, []float64{3.14, 42.0}, false},
		{"invalid args", []any{"3.14", 42, "invalid"}, nil, true},
	}

	c := ConvFuncs{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ToFloat64s(tt.input...)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDeltaSlice(t, tt.want, got, 0.000_001)
			}
		})
	}
}

func TestConvFuncs_ToString(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		// Basic type conversions
		{"string", "hello", "hello"},
		{"int", 42, "42"},
		{"float32", float32(3.14), "3.14"},
		{"float64", 9.5, "9.5"},
		{"bool true", true, "true"},
		{"bool false", false, "false"},

		// Slices
		{"slice of strings", []string{"a", "b"}, "[a b]"},
		{"empty slice", []string{}, "[]"},

		// Maps
		{"map with values", map[string]int{"key": 42}, "map[key:42]"},
		{"empty map", map[string]int{}, "map[]"},

		// Other types
		{"nil input", nil, ""},
	}

	c := ConvFuncs{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := c.ToString(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvFuncs_Join(t *testing.T) {
	tests := []struct {
		name  string
		sep   string
		input []any
		want  string
	}{
		{"no args", ",", []any{}, ""},
		{"valid args", "|", []any{"3.14", 42}, "3.14|42"},
	}

	c := ConvFuncs{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := c.Join(tt.input, tt.sep)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvFuncs_Default(t *testing.T) {
	tests := []struct {
		name  string
		def   any
		input any
		want  any
	}{
		{"bool false", true, false, true},
		{"bool true", false, true, true},
		{"int zero", 42, 0, 42},
		{"int", 42, 23, 23},
		{"empty string", "default", "", "default"},
		{"string", "default", "test", "test"},
		{"empty slice", []string{"default"}, []string{}, []string{"default"}},
		{"slice", []string{"default"}, []string{"test"}, []string{"test"}},
		{"empty map", map[string]int{"default": 42}, map[string]int{}, map[string]int{"default": 42}},
		{"map", map[string]int{"default": 42}, map[string]int{"test": 23}, map[string]int{"test": 23}},
		{"nil", "default", nil, "default"},
	}
	c := ConvFuncs{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := c.Default(tt.def, tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
