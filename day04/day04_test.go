package day04

import (
	"testing"
	"reflect"
)

func TestCountXmasOccurances(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected int
	}{
		{
			name: "simple case",
			input: "MMMSXXMASM",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countXmasOccurances(tt.input)

			if result != tt.expected {
				t.Errorf("Result mismatch: got %v, want %v", result, tt.expected)
				return
			}
		})
	}
}

func TestParseHorizontalL2R(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected []string
	}{
		{
			name: "simple case",
			input: "123\n456\n789",
			expected: []string{"123","456","789"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseHorizontalL2R(tt.input)

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

func TestParseHorizontalR2L(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected []string
	}{
		{
			name: "simple case",
			input: "123\n456\n789",
			expected: []string{"321","654","987"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseHorizontalR2L(tt.input)

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

func TestParseVerticalT2B(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected []string
	}{
		{
			name: "simple case",
			input: "123\n456\n789",
			expected: []string{"147","258","369"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseVerticalT2B(tt.input)

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

func TestParseVerticalB2T(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected []string
	}{
		{
			name: "simple case",
			input: "123\n456\n789",
			expected: []string{"741","852","963"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseVerticalB2T(tt.input)

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

func TestParseDiagonalTL2BR(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected []string
	}{
		{
			name: "3x3",
			input: "123\n456\n789",
			expected: []string{"159","26","3","48","7"},
		},
		{
			name: "4x4",
			input: "ABCD\nEFGH\nIJKL\nMNOP",
			expected: []string{"AFKP","BGL","CH","D","EJO","IN","M"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseDiagonalTL2BR(tt.input)

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

func TestParseDiagonalBR2TL(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected []string
	}{
		{
			name: "simple case",
			input: "123\n456\n789",
			expected: []string{"951","62","3","84","7"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseDiagonalBR2TL(tt.input)

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

func TestParseDiagonalTR2BL(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected []string
	}{
		{
			name: "simple case",
			input: "123\n456\n789",
			expected: []string{"357","24","1","68","9"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseDiagonalTR2BL(tt.input)

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

func TestParseDiagonalBL2TR(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected []string
	}{
		{
			name: "simple case",
			input: "123\n456\n789",
			expected: []string{"753","42","1","86","9"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseDiagonalBL2TR(tt.input)

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

func TestReverseString(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected string
	}{
		{
			name: "simple case",
			input: "hello",
			expected: "olleh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseString(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Result mismatch: got %v, want %v", result, tt.expected)
				return
			}
		})
	}
}

func TestParseStringToLetterAs(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected []LetterA
	}{
		{
			name: "simple case",
			input: "123\n4A6\n789",
			expected: []LetterA{
				LetterA{
					tl: '1',
					tr: '3',
					bl: '7',
					br: '9',
				},
			},
		},
		{
			name: "a on top border",
			input: "1A3\n456\n789",
			expected: []LetterA{},
		},
		{
			name: "a on left border",
			input: "123\nA56\n789",
			expected: []LetterA{},
		},
		{
			name: "a on right border",
			input: "123\n45A\n789",
			expected: []LetterA{},
		},
		{
			name: "a on bottom border",
			input: "123,456,7A9",
			expected: []LetterA{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseStringToLetterAs(tt.input)

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

func TestCountMasInLetterA(t *testing.T) {
	tests := []struct{
		name string
		input LetterA
		expected int
	}{
		{
			name: "MASxMAS",
			input: LetterA{
				tl: 'M',
				tr: 'M',
				bl: 'S',
				br: 'S',
			},
			expected: 2,
		},
		{
			name: "MASxSAM",
			input: LetterA{
				tl: 'M',
				tr: 'S',
				bl: 'M',
				br: 'S',
			},
			expected: 2,
		},
		{
			name: "SAMxSAM",
			input: LetterA{
				tl: 'S',
				tr: 'S',
				bl: 'M',
				br: 'M',
			},
			expected: 2,
		},
		{
			name: "SAMxMAS",
			input: LetterA {
				tl: 'S',
				tr: 'M',
				bl: 'S',
				br: 'M',
			},
			expected: 2,
		},
		{
			name: "MASxFAF",
			input: LetterA{
				tl: 'M',
				tr: 'F',
				bl: 'F',
				br: 'S',
			},
			expected: 1,
		},
		{
			name: "SAMxFAF",
			input: LetterA {
				tl: 'S',
				tr: 'F',
				bl: 'F',
				br: 'M',
			},
			expected: 1,
		},
		{
			name: "FAFxMAS",
			input: LetterA {
				tl: 'F',
				tr: 'M',
				bl: 'S',
				br: 'F',
			},
			expected: 1,
		},
		{
			name: "FAFxSAM",
			input: LetterA {
				tl: 'F',
				tr: 'S',
				bl: 'M',
				br: 'F',
			},
			expected: 1,
		},
		{
			name: "FAFxFAF",
			input: LetterA{
				tl: 'F',
				tr: 'F',
				bl: 'F',
				br: 'F',
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countMasInLetterA(tt.input)

			if result != tt.expected {
				t.Errorf("Result mismatch: got %v, want %v", result, tt.expected)
				return
			}	
		})
	}
}