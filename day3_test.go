package main

import "testing"

func TestFindNearestPoint(t *testing.T) {
	tests := []struct {
		input1         []string
		input2         []string
		expectedOutput int
	}{
		{[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			159},
		{[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			135},
	}
	for _, tt := range tests {
		output := findNearestPoint(tt.input1, tt.input2)
		if output != tt.expectedOutput {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", output, tt.expectedOutput)
		}
	}
}
func TestFindFewestStep(t *testing.T) {
	tests := []struct {
		input1         []string
		input2         []string
		expectedOutput int
	}{
		{[]string{"R8", "U5", "L5", "D3"},
			[]string{"U7", "R6", "D4", "L4"},
			30},
		{[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			610},
		{[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			410},
	}
	for _, tt := range tests {
		output := findFewestStep(tt.input1, tt.input2)
		if output != tt.expectedOutput {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", output, tt.expectedOutput)
		}
	}
}
