package functions_test

import (
	"testing"

	"github.com/spacecafe/miniplate/pkg/functions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMathFuncs_Abs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input   any
		want    any
		name    string
		wantErr bool
	}{
		{name: "positive integer", input: 5, want: int64(5), wantErr: false},
		{name: "negative integer", input: -5, want: int64(5), wantErr: false},
		{name: "zero", input: 0, want: int64(0), wantErr: false},
		{name: "positive float", input: 3.14, want: 3.14, wantErr: false},
		{name: "negative float", input: -3.14, want: 3.14, wantErr: false},
		{name: "unsupported type", input: "string", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Abs(tt.input)

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

func TestMathFuncs_Add(t *testing.T) {
	t.Parallel()

	tests := []struct {
		want    any
		name    string
		input   []any
		wantErr bool
	}{
		{name: "two integers", input: []any{2, 3}, want: int64(5), wantErr: false},
		{name: "three integers", input: []any{1, 2, 3}, want: int64(6), wantErr: false},
		{name: "two floats", input: []any{2.5, 3.5}, want: 6.0, wantErr: false},
		{name: "mixed int and float", input: []any{2, 3.5}, want: 5.5, wantErr: false},
		{name: "single value", input: []any{5}, want: int64(5), wantErr: false},
		{name: "empty slice", input: []any{}, want: 0, wantErr: true},
		{name: "unsupported type", input: []any{1, "string"}, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Add(tt.input...)

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

func TestMathFuncs_Ceil(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input   any
		name    string
		want    float64
		wantErr bool
	}{
		{name: "positive float", input: 3.14, want: 4.0, wantErr: false},
		{name: "negative float", input: -3.14, want: -3.0, wantErr: false},
		{name: "integer", input: 5, want: 5.0, wantErr: false},
		{name: "already whole number", input: 5.0, want: 5.0, wantErr: false},
		{name: "zero", input: 0, want: 0.0, wantErr: false},
		{name: "unsupported type", input: "string", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Ceil(tt.input)

			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				assert.InDelta(t, tt.want, got, 0.000001)
			}
		})
	}
}

func TestMathFuncs_Div(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   []any
		want    float64
		wantErr bool
	}{
		{name: "two integers", input: []any{10, 2}, want: 5.0, wantErr: false},
		{name: "three integers", input: []any{100, 5, 2}, want: 10.0, wantErr: false},
		{name: "two floats", input: []any{10.0, 2.0}, want: 5.0, wantErr: false},
		{name: "mixed types", input: []any{10, 2.0}, want: 5.0, wantErr: false},
		{name: "division by zero", input: []any{10, 0}, want: 0, wantErr: true},
		{name: "empty slice", input: []any{}, want: 0, wantErr: true},
		{name: "single value", input: []any{10}, want: 0, wantErr: true},
		{name: "unsupported type", input: []any{10, "string"}, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Div(tt.input...)

			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				assert.InDelta(t, tt.want, got, 0.000001)
			}
		})
	}
}

func TestMathFuncs_Floor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input   any
		name    string
		want    float64
		wantErr bool
	}{
		{name: "positive float", input: 3.14, want: 3.0, wantErr: false},
		{name: "negative float", input: -3.14, want: -4.0, wantErr: false},
		{name: "integer", input: 5, want: 5.0, wantErr: false},
		{name: "already whole number", input: 5.0, want: 5.0, wantErr: false},
		{name: "zero", input: 0, want: 0.0, wantErr: false},
		{name: "unsupported type", input: "string", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Floor(tt.input)

			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				assert.InDelta(t, tt.want, got, 0.000001)
			}
		})
	}
}

func TestMathFuncs_IsFloat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input any
		name  string
		want  bool
	}{
		{name: "int", input: 5, want: false},
		{name: "int32", input: int32(5), want: false},
		{name: "int64", input: int64(5), want: false},
		{name: "uint", input: uint(5), want: false},
		{name: "uint32", input: uint32(5), want: false},
		{name: "uint64", input: uint64(5), want: false},
		{name: "float32", input: float32(3.14), want: true},
		{name: "float64", input: 3.14, want: true},
		{name: "string", input: "5", want: false},
		{name: "bool", input: true, want: false},
		{name: "nil", input: nil, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := functions.MathFuncs{}.IsFloat(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMathFuncs_IsInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input any
		name  string
		want  bool
	}{
		{name: "int", input: 5, want: true},
		{name: "int32", input: int32(5), want: true},
		{name: "int64", input: int64(5), want: true},
		{name: "uint", input: uint(5), want: true},
		{name: "uint32", input: uint32(5), want: true},
		{name: "uint64", input: uint64(5), want: true},
		{name: "float32", input: float32(3.14), want: false},
		{name: "float64", input: 3.14, want: false},
		{name: "string", input: "5", want: false},
		{name: "bool", input: true, want: false},
		{name: "nil", input: nil, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := functions.MathFuncs{}.IsInt(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMathFuncs_IsNum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input any
		name  string
		want  bool
	}{
		{name: "int", input: 5, want: true},
		{name: "int32", input: int32(5), want: true},
		{name: "int64", input: int64(5), want: true},
		{name: "uint", input: uint(5), want: true},
		{name: "uint32", input: uint32(5), want: true},
		{name: "uint64", input: uint64(5), want: true},
		{name: "float32", input: float32(3.14), want: true},
		{name: "float64", input: 3.14, want: true},
		{name: "string", input: "5", want: false},
		{name: "bool", input: true, want: false},
		{name: "nil", input: nil, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := functions.MathFuncs{}.IsNum(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMathFuncs_Max(t *testing.T) {
	t.Parallel()

	tests := []struct {
		want    any
		name    string
		input   []any
		wantErr bool
	}{
		{name: "two integers", input: []any{3, 5}, want: int64(5), wantErr: false},
		{name: "three integers", input: []any{10, 20, 5}, want: int64(20), wantErr: false},
		{name: "two floats", input: []any{3.5, 4.2}, want: 4.2, wantErr: false},
		{name: "mixed int and float", input: []any{5, 3.8}, want: float64(5), wantErr: false},
		{name: "single value", input: []any{5}, want: int64(5), wantErr: false},
		{name: "empty slice", input: []any{}, want: 0, wantErr: true},
		{name: "unsupported type", input: []any{"a", "b"}, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Max(tt.input...)

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

func TestMathFuncs_Min(t *testing.T) {
	t.Parallel()

	tests := []struct {
		want    any
		name    string
		input   []any
		wantErr bool
	}{
		{name: "two integers", input: []any{3, 5}, want: int64(3), wantErr: false},
		{name: "three integers", input: []any{10, 20, 5}, want: int64(5), wantErr: false},
		{name: "two floats", input: []any{3.5, 4.2}, want: 3.5, wantErr: false},
		{name: "mixed int and float", input: []any{5, 3.8}, want: float64(3.8), wantErr: false},
		{name: "single value", input: []any{5}, want: int64(5), wantErr: false},
		{name: "empty slice", input: []any{}, want: 0, wantErr: true},
		{name: "unsupported type", input: []any{"a", "b"}, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Min(tt.input...)

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

func TestMathFuncs_Mul(t *testing.T) {
	t.Parallel()

	tests := []struct {
		want    any
		name    string
		input   []any
		wantErr bool
	}{
		{name: "two integers", input: []any{2, 3}, want: int64(6), wantErr: false},
		{name: "three integers", input: []any{2, 3, 4}, want: int64(24), wantErr: false},
		{name: "two floats", input: []any{2.5, 2.0}, want: 5.0, wantErr: false},
		{name: "mixed int and float", input: []any{2, 2.5}, want: 5.0, wantErr: false},
		{name: "multiply by zero", input: []any{5, 0}, want: int64(0), wantErr: false},
		{name: "single value", input: []any{5}, want: int64(5), wantErr: false},
		{name: "empty slice", input: []any{}, want: 0, wantErr: true},
		{name: "unsupported type", input: []any{1, "string"}, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Mul(tt.input...)

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

func TestMathFuncs_Pow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		inputBase any
		inputExp  any
		name      string
		want      float64
		wantErr   bool
	}{
		{name: "integer base and exponent", inputBase: 2, inputExp: 3, want: 8.0, wantErr: false},
		{
			name:      "float base and integer exponent",
			inputBase: 2.0,
			inputExp:  3,
			want:      8.0,
			wantErr:   false,
		},
		{
			name:      "integer base and float exponent",
			inputBase: 4,
			inputExp:  0.5,
			want:      2.0,
			wantErr:   false,
		},
		{name: "zero exponent", inputBase: 5, inputExp: 0, want: 1.0, wantErr: false},
		{name: "negative exponent", inputBase: 2, inputExp: -2, want: 0.25, wantErr: false},
		{name: "base zero", inputBase: 0, inputExp: 5, want: 0.0, wantErr: false},
		{name: "unsupported base type", inputBase: "string", inputExp: 2, want: 0, wantErr: true},
		{
			name:      "unsupported exponent type",
			inputBase: 2,
			inputExp:  "string",
			want:      0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Pow(tt.inputBase, tt.inputExp)

			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				assert.InDelta(t, tt.want, got, 0.000001)
			}
		})
	}
}

func TestMathFuncs_Rem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		inputDividend any
		inputDivisor  any
		name          string
		want          int64
		wantErr       bool
	}{
		{name: "positive remainder", inputDividend: 10, inputDivisor: 3, want: 1, wantErr: false},
		{name: "no remainder", inputDividend: 10, inputDivisor: 5, want: 0, wantErr: false},
		{name: "negative dividend", inputDividend: -10, inputDivisor: 3, want: -1, wantErr: false},
		{name: "negative divisor", inputDividend: 10, inputDivisor: -3, want: 1, wantErr: false},
		{name: "division by zero", inputDividend: 10, inputDivisor: 0, want: 0, wantErr: true},
		{name: "float dividend", inputDividend: 10.5, inputDivisor: 3, want: 0, wantErr: true},
		{name: "float divisor", inputDividend: 10, inputDivisor: 3.5, want: 0, wantErr: true},
		{
			name:          "unsupported type",
			inputDividend: "string",
			inputDivisor:  3,
			want:          0,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Rem(tt.inputDividend, tt.inputDivisor)

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

func TestMathFuncs_Round(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input   any
		name    string
		want    float64
		wantErr bool
	}{
		{name: "positive float round up", input: 3.6, want: 4.0, wantErr: false},
		{name: "positive float round down", input: 3.4, want: 3.0, wantErr: false},
		{name: "negative float round up", input: -3.4, want: -3.0, wantErr: false},
		{name: "negative float round down", input: -3.6, want: -4.0, wantErr: false},
		{name: "half value", input: 3.5, want: 4.0, wantErr: false},
		{name: "integer", input: 5, want: 5.0, wantErr: false},
		{name: "zero", input: 0, want: 0.0, wantErr: false},
		{name: "unsupported type", input: "string", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Round(tt.input)

			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				assert.InDelta(t, tt.want, got, 0.000001)
			}
		})
	}
}

func TestMathFuncs_Sub(t *testing.T) {
	t.Parallel()

	tests := []struct {
		want    any
		name    string
		input   []any
		wantErr bool
	}{
		{name: "two integers", input: []any{5, 3}, want: int64(2), wantErr: false},
		{name: "three integers", input: []any{10, 3, 2}, want: int64(5), wantErr: false},
		{name: "two floats", input: []any{5.5, 2.5}, want: 3.0, wantErr: false},
		{name: "mixed int and float", input: []any{10, 2.5}, want: 7.5, wantErr: false},
		{name: "result negative", input: []any{3, 5}, want: int64(-2), wantErr: false},
		{name: "single value", input: []any{5}, want: int64(5), wantErr: false},
		{name: "empty slice", input: []any{}, want: nil, wantErr: true},
		{name: "unsupported type", input: []any{10, "string"}, want: nil, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Sub(tt.input...)

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

func TestMathFuncs_Seq(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   []any
		want    []int64
		wantErr bool
	}{
		{
			name:    "single arg - end only",
			input:   []any{5},
			want:    []int64{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{name: "single arg - zero", input: []any{0}, want: []int64{}, wantErr: false},
		{name: "single arg - negative", input: []any{-3}, want: []int64{}, wantErr: false},
		{
			name:    "two args - start and end",
			input:   []any{3, 7},
			want:    []int64{3, 4, 5, 6, 7},
			wantErr: false,
		},
		{name: "two args - equal", input: []any{5, 5}, want: []int64{5}, wantErr: false},
		{
			name:    "two args - start greater than end",
			input:   []any{7, 3},
			want:    []int64{},
			wantErr: false,
		},
		{
			name:    "three args - positive step",
			input:   []any{1, 10, 2},
			want:    []int64{1, 3, 5, 7, 9},
			wantErr: false,
		},
		{
			name:    "three args - step of 1",
			input:   []any{5, 8, 1},
			want:    []int64{5, 6, 7, 8},
			wantErr: false,
		},
		{
			name:    "three args - negative step descending",
			input:   []any{10, 1, -2},
			want:    []int64{10, 8, 6, 4, 2},
			wantErr: false,
		},
		{
			name:    "three args - negative step single value",
			input:   []any{5, 5, -1},
			want:    []int64{5},
			wantErr: false,
		},
		{
			name:    "three args - large step",
			input:   []any{1, 100, 50},
			want:    []int64{1, 51},
			wantErr: false,
		},
		{
			name:    "zero step - max iterations",
			input:   []any{1, 1000, 0},
			want:    []int64{},
			wantErr: false,
		},
		{
			name:  "max iterations",
			input: []any{1, 1000},
			want: []int64{
				1,
				2,
				3,
				4,
				5,
				6,
				7,
				8,
				9,
				10,
				11,
				12,
				13,
				14,
				15,
				16,
				17,
				18,
				19,
				20,
				21,
				22,
				23,
				24,
				25,
				26,
				27,
				28,
				29,
				30,
				31,
				32,
				33,
				34,
				35,
				36,
				37,
				38,
				39,
				40,
				41,
				42,
				43,
				44,
				45,
				46,
				47,
				48,
				49,
				50,
				51,
				52,
				53,
				54,
				55,
				56,
				57,
				58,
				59,
				60,
				61,
				62,
				63,
				64,
				65,
				66,
				67,
				68,
				69,
				70,
				71,
				72,
				73,
				74,
				75,
				76,
				77,
				78,
				79,
				80,
				81,
				82,
				83,
				84,
				85,
				86,
				87,
				88,
				89,
				90,
				91,
				92,
				93,
				94,
				95,
				96,
				97,
				98,
				99,
				100,
			},
			wantErr: false,
		},
		{name: "empty slice", input: []any{}, want: nil, wantErr: true},
		{name: "too many args", input: []any{1, 2, 3, 4}, want: nil, wantErr: true},
		{name: "unsupported type", input: []any{"string"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := functions.MathFuncs{}.Seq(tt.input...)

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
