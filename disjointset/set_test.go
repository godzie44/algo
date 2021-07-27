package disjointset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDisjointSet(t *testing.T) {
	type testCase struct {
		v1, v2 rune
		same   bool
	}
	cases := []testCase{
		{
			v1:   'a',
			v2:   'b',
			same: true,
		},
		{
			v1:   'a',
			v2:   'd',
			same: true,
		},
		{
			v1:   'c',
			v2:   'd',
			same: true,
		},
		{
			v1:   'e',
			v2:   'f',
			same: true,
		},
		{
			v1:   'g',
			v2:   'f',
			same: true,
		},
		{
			v1:   'h',
			v2:   'i',
			same: true,
		},
		{
			v1:   'j',
			v2:   'i',
			same: false,
		},
		{
			v1:   'f',
			v2:   'a',
			same: false,
		},
		{
			v1:   'a',
			v2:   'h',
			same: false,
		},
		{
			v1:   'e',
			v2:   'j',
			same: false,
		},
		{
			v1:   'b',
			v2:   'g',
			same: false,
		},
	}

	vertexes := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	edges := [][]rune{{'b', 'd'}, {'e', 'g'}, {'a', 'c'}, {'h', 'i'}, {'a', 'b'}, {'e', 'f'}, {'b', 'c'}}

	vertexSets := make(map[rune]*Set, len(vertexes))

	for _, v := range vertexes {
		vertexSets[v] = NewSet(v)
	}

	for _, e := range edges {
		if FindSet(vertexSets[e[0]]) != FindSet(vertexSets[e[1]]) {
			Union(vertexSets[e[0]], vertexSets[e[1]])
		}
	}

	for _, tc := range cases {
		assert.Equal(t, tc.same, FindSet(vertexSets[tc.v1]) == FindSet(vertexSets[tc.v2]))
	}
}
