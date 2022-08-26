package stringmatch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNaiveMatch(t *testing.T) {
	str := "addfassaddfdadd"
	assert.Equal(t, []int{0, 7, 12}, NaiveMatching(str, "add"))
}

func TestRabinKarpMatcher(t *testing.T) {
	str := "addfassaddfdadd"
	matches, err := RabinKarpMatcher(str, "add")
	assert.NoError(t, err)

	assert.Equal(t, []int{0, 7, 12}, matches)

	str = "123112233112"
	matches, err = RabinKarpMatcher(str, "11")
	assert.NoError(t, err)

	assert.Equal(t, []int{3, 9}, matches)
}

func TestFiniteAutomationMatcher(t *testing.T) {
	str := "abababacaba"
	matches := FiniteAutomationMatcher(str, "ababaca", 128)
	assert.Equal(t, []int{2}, matches)

	str = "addfassaddfdadd"
	matches = FiniteAutomationMatcher(str, "add", 128)
	assert.Equal(t, []int{0, 7, 12}, matches)

	str = "123112233112"
	matches = FiniteAutomationMatcher(str, "11", 128)

	assert.Equal(t, []int{3, 9}, matches)
}

func TestKMPMatcher(t *testing.T) {
	str := "abababacaba"
	matches := KMPMatcher(str, "ababaca")
	assert.Equal(t, []int{2}, matches)

	str = "addfassaddfdadd"
	matches = FiniteAutomationMatcher(str, "add", 128)
	assert.Equal(t, []int{0, 7, 12}, matches)

	str3 := "123112233112"
	matches3 := KMPMatcher(str3, "11")
	assert.Equal(t, []int{3, 9}, matches3)
}
