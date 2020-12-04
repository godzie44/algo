package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var a = [][]int{
	{1, 3, 4, 2},
	{0, 4, 5, 1},
	{3, 6, 2, 2},
	{6, 5, 2, 6},
}

var b = [][]int{
	{1, 2, 4, 4},
	{1, 2, 2, 5},
	{4, 0, 3, 3},
	{1, 3, 2, 3},
}

var expectedC = [][]int{
	{22, 14, 26, 37},
	{25, 11, 25, 38},
	{19, 24, 34, 54},
	{25, 40, 52, 73},
}

var aSmall = [][]int{
	{1, 3},
	{0, 4},
}

var bSmall = [][]int{
	{1, 2},
	{1, 2},
}
var expectedCSmall = [][]int{
	{4, 8},
	{4, 8},
}



func TestMultiply(t *testing.T) {
	result := multiply(aSmall, bSmall)
	assert.Equal(t, expectedCSmall, result)

	result = multiply(a, b)
	assert.Equal(t, expectedC, result)
}

func TestMultiplyDivideAndRule(t *testing.T) {
	result := multiplyDivideAndRule(aSmall, bSmall)
	assert.Equal(t, expectedCSmall, result)

	result = multiplyDivideAndRule(a, b)
	assert.Equal(t, expectedC, result)
}

func TestMultiplyShtrassen(t *testing.T) {
	result := multiplyShtrassen(aSmall, bSmall)
	assert.Equal(t, expectedCSmall, result)

	result = multiplyShtrassen(a, b)
	assert.Equal(t, expectedC, result)
}
