package stringmatch

import (
	"algorithms/numbertheoretic"
	"crypto/rand"
	"math"
	"math/big"
	"strings"
)

func NaiveMatching(str, needle string) (matches []int) {
	n := len(str)
	m := len(needle)

	for s := 0; s <= n-m; s++ {
		if strings.EqualFold(needle, str[s:s+m]) {
			matches = append(matches, s)
		}
	}

	return matches
}

func RabinKarpMatcher(str, needle string) (matches []int, err error) {
	const DictionaryLen = 128 // only 7-bit ASCII symbols

	d := big.NewInt(DictionaryLen)
	q, err := rand.Prime(rand.Reader, 128)
	if err != nil {
		return nil, err
	}

	n := len(str)
	m := len(needle)
	h := numbertheoretic.ModularExponentiation(d, big.NewInt(int64(m)-1), q)
	p := big.NewInt(0)
	tPrev := big.NewInt(0)

	for i := 0; i < m; i++ {
		pSum := new(big.Int).Add(new(big.Int).Mul(p, d), big.NewInt(int64(needle[i])))
		p = new(big.Int).Mod(pSum, q)

		t0Sum := new(big.Int).Add(new(big.Int).Mul(tPrev, d), big.NewInt(int64(str[i])))
		tPrev = new(big.Int).Mod(t0Sum, q)
	}

	tNext := &*tPrev
	for s := 0; s <= n-m; s++ {
		if p.Cmp(tNext) == 0 {
			if strings.EqualFold(needle, str[s:s+m]) {
				matches = append(matches, s)
			}
		}
		tPrev = tNext
		if s < n-m {
			tNext = new(big.Int).Mul(big.NewInt(int64(str[s])), &h)
			tNext = new(big.Int).Sub(tPrev, tNext)
			tNext = new(big.Int).Mul(d, tNext)
			tNext = new(big.Int).Add(tNext, big.NewInt(int64(str[s+m])))
			tNext = new(big.Int).Mod(tNext, q)
		}
	}

	return matches, nil
}

func computeTransitionFunc(str string, dict []string) map[int]map[string]int {
	m := len(str)

	result := make(map[int]map[string]int)

	for q := 0; q <= m; q++ {
		for _, a := range dict {
			//fmt.Println("next dict symb ", a)
			k := int(math.Min(float64(m), float64(q+1)))

			for {
				p := str[:k]
				//fmt.Println(k)
				//fmt.Println(p)
				//fmt.Println(str[:q] + a)

				if strings.HasSuffix(str[:q]+a, p) {
					break
				} else {
					k--
				}
			}

			if result[q] == nil {
				result[q] = make(map[string]int)
			}
			result[q][a] = k
		}
	}

	return result
}

func finiteAutomationMatcher(sigm map[int]map[string]int, str string, m int) (matches []int) {
	q := 0
	for pos, char := range str {
		q = sigm[q][string(char)]
		if q == m {
			matches = append(matches, pos-m+1)
		}
	}
	return matches
}

func FiniteAutomationMatcher(str, needle string, maxCodeInDict int) (matches []int) {
	characters := make([]string, maxCodeInDict)
	for i := 0; i < maxCodeInDict; i++ {
		characters[i] = string(rune(i))
	}

	sign := computeTransitionFunc(needle, characters)

	return finiteAutomationMatcher(sign, str, len(needle))
}

func getCharCode(str string, ind int) uint8 {
	return str[ind-1]
}

func KMPMatcher(str, needle string) (matches []int) {
	n := len(str)
	m := len(needle)
	p := computePrefixFunc(needle)
	q := 0

	for i := 1; i <= n; i++ {
		for q > 0 && getCharCode(needle, q+1) != getCharCode(str, i) {
			q = p[q]
		}

		if getCharCode(needle, q+1) == getCharCode(str, i) {
			q++
		}

		if q == m {
			matches = append(matches, i-m)
			q = p[q]
		}
	}
	return matches
}

func computePrefixFunc(needle string) map[int]int {
	m := len(needle)
	p := map[int]int{}
	p[1] = 0
	k := 0

	for q := 2; q <= m; q++ {
		for k > 0 && getCharCode(needle, k+1) != getCharCode(needle, q) {
			k = p[k]
		}

		if getCharCode(needle, k+1) == getCharCode(needle, q) {
			k++
		}
		p[q] = k
	}
	return p
}
