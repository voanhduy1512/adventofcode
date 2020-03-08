package main

import (
	"testing"
)

func TestCalcFuel(t *testing.T) {
	tests := []struct {
		input          int64
		expectedOutput int64
	}{
		{int64(14), int64(2)},
		{int64(1969), int64(966)},
		{int64(100756), int64(50346)},
	}
	for _, tt := range tests {
		output := calcFuel(tt.input)
		if output != tt.expectedOutput {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", output, tt.expectedOutput)
		}
	}
}
