package main

import "testing"

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		input          int
		expectedOutput bool
	}{
		{111111, true},
		{223450, false},
		{123789, false},
	}
	for _, tt := range tests {
		output := isValidPassword(tt.input)
		if output != tt.expectedOutput {
			t.Errorf("Sum was incorrect, got: %t, want: %t.", output, tt.expectedOutput)
		}
	}
}

func TestIsValidPasswordExtraRule(t *testing.T) {
	tests := []struct {
		input          int
		expectedOutput bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},
		{111234, false},
	}
	for _, tt := range tests {
		output := isValidPasswordExtraRule(tt.input)
		if output != tt.expectedOutput {
			t.Errorf("%d, got: %t, want: %t.", tt.input, output, tt.expectedOutput)
		}
	}
}
