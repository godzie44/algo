package process

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActivitySelector(t *testing.T) {
	tests := []struct {
		args     []Process
		expected []Process
	}{
		{
			args: []Process{{1, 4}, {3, 5}, {0, 6}, {5, 7}, {3, 9}, {5, 9}, {6, 10},
				{8, 11}, {8, 12}, {2, 14}, {12, 16}},
			expected: []Process{{1, 4}, {5, 7}, {8, 11}, {12, 16}},
		},
		{
			args:     []Process{{1, 4}},
			expected: []Process{{1, 4}},
		},
	}
	for _, tc := range tests {
		assert.Equal(t, ActivitySelector(tc.args), tc.expected)
	}
}
