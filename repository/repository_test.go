package repository_test

import (
	"github.com/stretchr/testify/assert"
	"proyect/repository"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		value    string
		expected bool
	}{
		{
			value:    "mk",
			expected: false,
		},
		{
			value:    "YYY",
			expected: true,
		},
		{
			value:    "ash",
			expected: false,
		},
		{
			value:    "asdffdsa",
			expected: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, repository.IsPalindrome(test.value), test.expected)
	}
}

func TestHalfPercentDiscount(t *testing.T) {
	tests := []struct {
		value    int
		expected int
	}{
		{
			value:    100,
			expected: 50,
		},
		{
			value:    40,
			expected: 20,
		},
		{
			value:    10,
			expected: 5,
		},
		{
			value:    2,
			expected: 1,
		},

	}

	for _, test := range tests {
		assert.Equal(t, repository.HalfPercentDiscount(test.value), test.expected)
	}
}
