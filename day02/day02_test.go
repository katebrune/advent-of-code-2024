package day02

import (
	"testing"
	"reflect"
)

func TestParseStringTo2dIntArray(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected [][]int
	}{
		{
			name: "simple case",
			input: "1 2 3\n4 5 6\n7 8 9",
			expected: [][]int{{1,2,3},{4,5,6},{7,8,9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseStringTo2dIntArray(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("Length mismatch: got %d, want %d", len(result), len(tt.expected))
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Result mismatch: got %v, want %v", result, tt.expected)
				return
			}
		})
	}
}

func TestCountValidReports(t *testing.T) {
	tests := []struct{
		name string
		input [][]int
		expected int
	}{
		{
			name: "both valid",
			input: [][]int{
				{1,2,3},
				{3,2,1},
			},
			expected: 2,
		},
		{
			name: "both invalid - large intervals",
			input: [][]int{
				{1,5,6},
				{6,5,1},
			},
			expected: 0,
		},
		{
			name: "both invalid - no change",
			input: [][]int{
				{1,3,3},
				{5,3,3},
			},
			expected: 0,
		},
		{
			name: "increasing valid, decreasing invalid",
			input: [][]int{
				{1,2,3},
				{6,2,1},
			},
			expected: 1,
		},
		{
			name: "increasing invalid, decreasing valid",
			input: [][]int{
				{1,5,6},
				{3,2,1},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countValidReports(tt.input)

			if result != tt.expected {
				t.Errorf("Result mismatch: got %v, want %v", result, tt.expected)
				return
			}
		})
	}
}

func TestCheckAllIncreasing(t *testing.T) {
	tests := []struct {
		name string
		input []int
		expected bool
	}{
		{
			name: "all increasing",
			input: []int{1,2,3},
			expected: true,
		},
		{
			name: "has decrease",
			input: []int{5,6,4},
			expected: false,
		},
		{
			name: "has a no change",
			input: []int{2,4,4},
			expected: false,
		},
		{
			name: "has interval > 3",
			input: []int{1,5,6},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkAllIncreasing(tt.input)

			if result != tt.expected {
				t.Errorf("Result mismatch: got %v, want %v", result, tt.expected)
				return
			}
		})
	}
}

func TestCheckAllDecreasing(t *testing.T) {
	tests := []struct {
		name string
		input []int
		expected bool
	}{
		{
			name: "all decreasing",
			input: []int{3,2,1},
			expected: true,
		},
		{
			name: "has increase",
			input: []int{6,4,8},
			expected: false,
		},
		{
			name: "has a no change",
			input: []int{6,4,4},
			expected: false,
		},
		{
			name: "has interval < -3",
			input: []int{6,2,1},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkAllDecreasing(tt.input)

			if result != tt.expected {
				t.Errorf("Result mismatch: got %v, want %v", result, tt.expected)
				return
			}
		})
	}
}
