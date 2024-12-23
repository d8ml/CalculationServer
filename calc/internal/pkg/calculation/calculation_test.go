package calculation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculaton(t *testing.T) {
	type args struct {
		input  string
		result float64
	}

	testCorrect := func(input string, expected float64) {
		res, err := Calculate(input)
		require.NoError(t, err)
		require.EqualValues(t, expected, res)
	}

	testError := func(input string, _ float64) {
		_, err := Calculate(input)
		if err == nil {
			t.Error("Expected error")
		}
	}

	tests := []struct {
		name string
		args args
		test func(string, float64)
	}{
		{
			name: "Addition 2 items",
			args: args{"2 + 3", 5},
			test: testCorrect,
		},
		{
			name: "Addition 5 items",
			args: args{"2 + 3 + 4 + 5 + 10", 24},
			test: testCorrect,
		},
		{
			name: "Subtraction",
			args: args{"2 - 4", -2},
			test: testCorrect,
		},
		{
			name: "Multiplication",
			args: args{"1 * 2", 2},
			test: testCorrect,
		},
		{
			name: "Division",
			args: args{"1 / 2", 0.5},
			test: testCorrect,
		},
		{
			name: "Division by zero",
			args: args{"1 / 0", 0},
			test: testError,
		},
		{
			name: "Brackets",
			args: args{"(1 + 2) * 3", 9},
			test: testCorrect,
		},
		{
			name: "Invalid brackets",
			args: args{"((1 + 2)", 0},
			test: testError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(tt.args.input, tt.args.result)
		})
	}
}
