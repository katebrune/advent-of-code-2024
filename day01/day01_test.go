package day01

import (
	"testing"
	"reflect"
)

func TestParseStringToTwoLists(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected TwoLists
	}{
		{ name: "simple case",
		  input: "1  2\n3  4\n5  6",
		  expected: TwoLists{
			ListA: []int{1,3,5},
			ListB: []int{2,4,6},
		  },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseStringToTwoLists(tt.input)

			if len(result.ListA) != len(tt.expected.ListA) {
				t.Errorf("ListA Length mismatch: got %d, want %d", len(result.ListA), len(tt.expected.ListA))
				return
			}

			if len(result.ListB) != len(tt.expected.ListB) {
				t.Errorf("ListB Length mismatch: got %d, want %d", len(result.ListB), len(tt.expected.ListB))
				return
			}

			if !reflect.DeepEqual(result.ListA, tt.expected.ListA) {
				t.Errorf("ListA Result mismatch: got %v, want %v", result.ListA, tt.expected.ListA)
				return
			}

			if !reflect.DeepEqual(result.ListB, tt.expected.ListB) {
				t.Errorf("ListB Result mismatch: got %v, want %v", result.ListB, tt.expected.ListB)
				return
			}
		})
	}
}

func TestOrderTwoLists(t *testing.T) {
	tests := []struct{
		name string
		input TwoLists
		expected TwoLists
	}{
		{
			name: "simple case",
			input: TwoLists{
				ListA: []int{3,9,1},
				ListB: []int{9,2,5},
			},
			expected: TwoLists{
				ListA: []int{1,3,9},
				ListB: []int{2,5,9},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := orderTwoLists(tt.input)

			if len(result.ListA) != len(tt.expected.ListA) {
				t.Errorf("ListA Length mismatch: got %d, want %d", len(result.ListA), len(tt.expected.ListA))
				return
			}

			if len(result.ListB) != len(tt.expected.ListB) {
				t.Errorf("ListB Length mismatch: got %d, want %d", len(result.ListB), len(tt.expected.ListB))
				return
			}

			if !reflect.DeepEqual(result.ListA, tt.expected.ListA) {
				t.Errorf("ListA Result mismatch: got %v, want %v", result.ListA, tt.expected.ListA)
				return
			}

			if !reflect.DeepEqual(result.ListB, tt.expected.ListB) {
				t.Errorf("ListB Result mismatch: got %v, want %v", result.ListB, tt.expected.ListB)
				return
			}
		})
	}
}

func TestGetDifferencesOfTwoLists(t *testing.T) {
	tests := []struct{
		name string
		input TwoLists
		expected []int
	}{
		{
			name: "simple case",
			input: TwoLists{
				ListA: []int{1,3,5},
				ListB: []int{2,3,4},
			},
			expected: []int{1,0,1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getDifferencesOfTwoLists(tt.input)

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

func TestGetSimilaritiesOfTwoLists(t *testing.T) {
	tests := []struct{
		name string
		input TwoLists
		expected []int
	}{
		{
			name: "simple case",
			input: TwoLists{
				ListA: []int{1,2,3},
				ListB: []int{1,2,2},
			},
			expected: []int{1,4,0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getSimilaritiesOfTwoLists(tt.input)

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


func TestSumList(t *testing.T) {
	tests := []struct{
		name string
		input []int
		expected int
	}{
		{
			name: "simple case",
			input: []int{1,2,3},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumList(tt.input)

			if result != tt.expected {
				t.Errorf("Result mismatch: got %v, want %v", result, tt.expected)
				return
			}
		})
	}
}