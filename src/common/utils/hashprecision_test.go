package utils

import "testing"

func TestHashPrecision(t *testing.T) {
	tests := []struct {
		name     string
		radius   float64
		expected int
	}{
		{
			name:     "small",
			radius:   1000,
			expected: 6,
		},
		{
			name:     "medium",
			radius:   10000,
			expected: 4,
		},
		{
			name:     "large",
			radius:   100000,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := HashPrecision(tt.radius)

			if out != tt.expected {
				t.Fatalf("expected=%v, actual=%v", tt.expected, out)
			}
		})
	}
}
