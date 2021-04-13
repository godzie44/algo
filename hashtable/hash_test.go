package hashtable

import (
	"fmt"
	"hash/maphash"
	"testing"
)

func TestH(t *testing.T) {
	m := maphash.Hash{}
	m.WriteString("ONE")

	fmt.Println(m.Sum64())

	m.Reset()
	m.WriteString("ONE")

	fmt.Println(m.Sum64())

	m2 := maphash.Hash{}
	m2.SetSeed(m.Seed())

	m2.WriteString("ONE")

	fmt.Println(m2.Sum64())
}

func TestAA(t *testing.T) {
	for i := 0; i < 100; i++ {
		g2(i, i+1)
	}
}

//go:noinline
func g2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
