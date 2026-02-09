package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "add 2 + 2",
			a:        2,
			b:        2,
			expected: 4,
		},
		{
			name:     "add 23 + 23",
			a:        23,
			b:        23,
			expected: 46,
		},
		{
			name:     "add positive numbers",
			a:        5,
			b:        3,
			expected: 8,
		},
		{
			name:     "add negative numbers",
			a:        -5,
			b:        -3,
			expected: -8,
		},
		{
			name:     "add positive and negative",
			a:        10,
			b:        -3,
			expected: 7,
		},
		{
			name:     "add zero",
			a:        0,
			b:        0,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, result)
		})
	}
}
