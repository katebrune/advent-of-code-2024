package day03

import (
	"testing"
	"reflect"
)

func TestParseJumbledStringToArithmeticList(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected []Arithmetic
	}{
		{
			name: "simple case",
			input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			expected: []Arithmetic{
				Arithmetic{
					numbers: []int{2,4},
					command: "mul",
				},
				Arithmetic{
					numbers: []int{5,5},
					command: "mul",
				},
				Arithmetic{
					numbers: []int{11,8},
					command: "mul",
				},
				Arithmetic{
					numbers: []int{8,5},
					command: "mul",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := parseJumbledStringToArithmeticList(tt.input)

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

func TestParseJumbledStringToArithmeticListWithDoDont(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected []Arithmetic
	}{
		{
			name: "simple case",
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expected: []Arithmetic{
				Arithmetic{
					numbers: []int{2,4},
					command: "mul",
				},
				Arithmetic{
					numbers: []int{8,5},
					command: "mul",
				},
			},
		},
		{
			name: "no do pattern",
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)und()?mul(8,5))",
			expected: []Arithmetic{
				Arithmetic{
					numbers: []int{2,4},
					command: "mul",
				},
			},
		},
		{
			name: "multiple do in a row",
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))do()?mul(2,3)[]",
			expected: []Arithmetic{
				Arithmetic{
					numbers: []int{2,4},
					command: "mul",
				},
				Arithmetic{
					numbers: []int{8,5},
					command: "mul",
				},
				Arithmetic{
					numbers: []int{2,3},
					command: "mul",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := parseJumbledStringToArithmeticListWithDoDont(tt.input)

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

func TestComputeArithmeticList(t *testing.T) {
	tests := []struct{
		name string
		input []Arithmetic
		expected []int
	}{
		{
			name: "simple case",
			input: []Arithmetic{
				Arithmetic{
					numbers: []int{2,4},
					command: "mul",
				},
				Arithmetic{
					numbers: []int{5,5},
					command: "mul",
				},
				Arithmetic{
					numbers: []int{11,8},
					command: "mul",
				},
				Arithmetic{
					numbers: []int{8,5},
					command: "mul",
				},
			},
			expected: []int{8,25,88,40},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeArthmeticList(tt.input)

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
			input: []int{8,25,88,40},
			expected: 161,
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