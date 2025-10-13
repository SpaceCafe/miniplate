package functions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathFuncs_Abs(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    any
		wantErr bool
	}{
		{"positive integer", 5, int64(5), false},
		{"negative integer", -5, int64(5), false},
		{"zero", 0, int64(0), false},
		{"positive float", 3.14, 3.14, false},
		{"negative float", -3.14, 3.14, false},
		{"unsupported type", "string", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Abs(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_Add(t *testing.T) {
	tests := []struct {
		name    string
		input   []any
		want    any
		wantErr bool
	}{
		{"two integers", []any{2, 3}, int64(5), false},
		{"three integers", []any{1, 2, 3}, int64(6), false},
		{"two floats", []any{2.5, 3.5}, 6.0, false},
		{"mixed int and float", []any{2, 3.5}, 5.5, false},
		{"single value", []any{5}, int64(5), false},
		{"empty slice", []any{}, 0, true},
		{"unsupported type", []any{1, "string"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Add(tt.input...)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_Ceil(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    float64
		wantErr bool
	}{
		{"positive float", 3.14, 4.0, false},
		{"negative float", -3.14, -3.0, false},
		{"integer", 5, 5.0, false},
		{"already whole number", 5.0, 5.0, false},
		{"zero", 0, 0.0, false},
		{"unsupported type", "string", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Ceil(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_Div(t *testing.T) {
	tests := []struct {
		name    string
		input   []any
		want    float64
		wantErr bool
	}{
		{"two integers", []any{10, 2}, 5.0, false},
		{"three integers", []any{100, 5, 2}, 10.0, false},
		{"two floats", []any{10.0, 2.0}, 5.0, false},
		{"mixed types", []any{10, 2.0}, 5.0, false},
		{"division by zero", []any{10, 0}, 0, true},
		{"empty slice", []any{}, 0, true},
		{"single value", []any{10}, 0, true},
		{"unsupported type", []any{10, "string"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Div(tt.input...)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_Floor(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    float64
		wantErr bool
	}{
		{"positive float", 3.14, 3.0, false},
		{"negative float", -3.14, -4.0, false},
		{"integer", 5, 5.0, false},
		{"already whole number", 5.0, 5.0, false},
		{"zero", 0, 0.0, false},
		{"unsupported type", "string", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Floor(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_IsFloat(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  bool
	}{
		{"int", 5, false},
		{"int32", int32(5), false},
		{"int64", int64(5), false},
		{"uint", uint(5), false},
		{"uint32", uint32(5), false},
		{"uint64", uint64(5), false},
		{"float32", float32(3.14), true},
		{"float64", 3.14, true},
		{"string", "5", false},
		{"bool", true, false},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MathFuncs{}.IsFloat(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMathFuncs_IsInt(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  bool
	}{
		{"int", 5, true},
		{"int32", int32(5), true},
		{"int64", int64(5), true},
		{"uint", uint(5), true},
		{"uint32", uint32(5), true},
		{"uint64", uint64(5), true},
		{"float32", float32(3.14), false},
		{"float64", 3.14, false},
		{"string", "5", false},
		{"bool", true, false},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MathFuncs{}.IsInt(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMathFuncs_IsNum(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  bool
	}{
		{"int", 5, true},
		{"int32", int32(5), true},
		{"int64", int64(5), true},
		{"uint", uint(5), true},
		{"uint32", uint32(5), true},
		{"uint64", uint64(5), true},
		{"float32", float32(3.14), true},
		{"float64", 3.14, true},
		{"string", "5", false},
		{"bool", true, false},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MathFuncs{}.IsNum(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMathFuncs_Max(t *testing.T) {
	tests := []struct {
		name    string
		input   []any
		want    any
		wantErr bool
	}{
		{"two integers", []any{3, 5}, int64(5), false},
		{"three integers", []any{10, 20, 5}, int64(20), false},
		{"two floats", []any{3.5, 4.2}, 4.2, false},
		{"mixed int and float", []any{5, 3.8}, float64(5), false},
		{"single value", []any{5}, int64(5), false},
		{"empty slice", []any{}, 0, true},
		{"unsupported type", []any{"a", "b"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Max(tt.input...)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_Min(t *testing.T) {
	tests := []struct {
		name    string
		input   []any
		want    any
		wantErr bool
	}{
		{"two integers", []any{3, 5}, int64(3), false},
		{"three integers", []any{10, 20, 5}, int64(5), false},
		{"two floats", []any{3.5, 4.2}, 3.5, false},
		{"mixed int and float", []any{5, 3.8}, float64(3.8), false},
		{"single value", []any{5}, int64(5), false},
		{"empty slice", []any{}, 0, true},
		{"unsupported type", []any{"a", "b"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Min(tt.input...)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_Mul(t *testing.T) {
	tests := []struct {
		name    string
		input   []any
		want    any
		wantErr bool
	}{
		{"two integers", []any{2, 3}, int64(6), false},
		{"three integers", []any{2, 3, 4}, int64(24), false},
		{"two floats", []any{2.5, 2.0}, 5.0, false},
		{"mixed int and float", []any{2, 2.5}, 5.0, false},
		{"multiply by zero", []any{5, 0}, int64(0), false},
		{"single value", []any{5}, int64(5), false},
		{"empty slice", []any{}, 0, true},
		{"unsupported type", []any{1, "string"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Mul(tt.input...)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_Pow(t *testing.T) {
	tests := []struct {
		name      string
		inputBase any
		inputExp  any
		want      float64
		wantErr   bool
	}{
		{"integer base and exponent", 2, 3, 8.0, false},
		{"float base and integer exponent", 2.0, 3, 8.0, false},
		{"integer base and float exponent", 4, 0.5, 2.0, false},
		{"zero exponent", 5, 0, 1.0, false},
		{"negative exponent", 2, -2, 0.25, false},
		{"base zero", 0, 5, 0.0, false},
		{"unsupported base type", "string", 2, 0, true},
		{"unsupported exponent type", 2, "string", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Pow(tt.inputBase, tt.inputExp)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_Rem(t *testing.T) {
	tests := []struct {
		name          string
		inputDividend any
		inputDivisor  any
		want          int64
		wantErr       bool
	}{
		{"positive remainder", 10, 3, 1, false},
		{"no remainder", 10, 5, 0, false},
		{"negative dividend", -10, 3, -1, false},
		{"negative divisor", 10, -3, 1, false},
		{"division by zero", 10, 0, 0, true},
		{"float dividend", 10.5, 3, 0, true},
		{"float divisor", 10, 3.5, 0, true},
		{"unsupported type", "string", 3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Rem(tt.inputDividend, tt.inputDivisor)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_Round(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    float64
		wantErr bool
	}{
		{"positive float round up", 3.6, 4.0, false},
		{"positive float round down", 3.4, 3.0, false},
		{"negative float round up", -3.4, -3.0, false},
		{"negative float round down", -3.6, -4.0, false},
		{"half value", 3.5, 4.0, false},
		{"integer", 5, 5.0, false},
		{"zero", 0, 0.0, false},
		{"unsupported type", "string", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Round(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_Sub(t *testing.T) {
	tests := []struct {
		name    string
		input   []any
		want    any
		wantErr bool
	}{
		{"two integers", []any{5, 3}, int64(2), false},
		{"three integers", []any{10, 3, 2}, int64(5), false},
		{"two floats", []any{5.5, 2.5}, 3.0, false},
		{"mixed int and float", []any{10, 2.5}, 7.5, false},
		{"result negative", []any{3, 5}, int64(-2), false},
		{"single value", []any{5}, int64(5), false},
		{"empty slice", []any{}, nil, true},
		{"unsupported type", []any{10, "string"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Sub(tt.input...)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMathFuncs_Seq(t *testing.T) {
	tests := []struct {
		name    string
		input   []any
		want    []int64
		wantErr bool
	}{
		{"single arg - end only", []any{5}, []int64{1, 2, 3, 4, 5}, false},
		{"single arg - zero", []any{0}, []int64{}, false},
		{"single arg - negative", []any{-3}, []int64{}, false},
		{"two args - start and end", []any{3, 7}, []int64{3, 4, 5, 6, 7}, false},
		{"two args - equal", []any{5, 5}, []int64{5}, false},
		{"two args - start greater than end", []any{7, 3}, []int64{}, false},
		{"three args - positive step", []any{1, 10, 2}, []int64{1, 3, 5, 7, 9}, false},
		{"three args - step of 1", []any{5, 8, 1}, []int64{5, 6, 7, 8}, false},
		{"three args - negative step descending", []any{10, 1, -2}, []int64{10, 8, 6, 4, 2}, false},
		{"three args - negative step single value", []any{5, 5, -1}, []int64{5}, false},
		{"three args - large step", []any{1, 100, 50}, []int64{1, 51}, false},
		{"zero step - max iterations", []any{1, 1000, 0}, []int64{}, false},
		{"max iterations", []any{1, 1000}, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}, false},
		{"empty slice", []any{}, nil, true},
		{"too many args", []any{1, 2, 3, 4}, nil, true},
		{"unsupported type", []any{"string"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MathFuncs{}.Seq(tt.input...)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
