package numbertheoretic

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func GcdEuclid(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a
	}
	return GcdEuclid(b, new(big.Int).Mod(a, b))
}

func GcdExtendedEuclid(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a, big.NewInt(1), big.NewInt(0)
	}
	dTmp, xTmp, yTmp := GcdExtendedEuclid(b, new(big.Int).Mod(a, b))
	d, x, y := dTmp, yTmp, new(big.Int).Sub(xTmp, new(big.Int).Mul(yTmp, new(big.Int).Div(a, b)))
	return d, x, y
}

func ModularLinearEquationSolver(a, b, n *big.Int) ([]*big.Int, error) {
	d, x, _ := GcdExtendedEuclid(a, n)

	if new(big.Int).Mod(b, d).Cmp(big.NewInt(0)) == 0 { // x*(b/d)
		results := []*big.Int{new(big.Int).Mod(new(big.Int).Mul(x, new(big.Int).Div(b, d)), n)}

		var i int64
		for i = 1; i < d.Int64(); i++ {
			results = append(results, new(big.Int).Mod(new(big.Int).Add(results[0], new(big.Int).Mul(big.NewInt(i), new(big.Int).Div(n, d))), n))
		}

		return results, nil
	}

	return nil, errors.New("no solves")
}

func ModularExponentiation(a, b, n *big.Int) big.Int {
	c := big.NewInt(0)
	d := big.NewInt(1)

	k := b.BitLen()
	for i := k - 1; i >= 0; i-- {
		c = new(big.Int).Mul(c, big.NewInt(2))
		d = d.Mod(new(big.Int).Mul(d, d), n)

		if b.Bit(i) == 1 {
			c = new(big.Int).Add(c, big.NewInt(1))
			d = d.Mod(new(big.Int).Mul(d, a), n)
		}
	}

	return *d
}

var bOne = big.NewInt(1)

func witness(a, n *big.Int) bool {
	nn := new(big.Int).Sub(n, bOne)
	t := nn.TrailingZeroBits()
	u := new(big.Int).Rsh(n, t)

	x0 := ModularExponentiation(a, u, n)

	xprev := &x0
	var x *big.Int
	for i := 1; i <= int(t); i++ {
		x = new(big.Int).Exp(xprev, big.NewInt(2), n)
		if x.Cmp(bOne) == 0 && xprev.Cmp(bOne) != 0 && xprev.Cmp(nn) != 0 {
			return true
		}
		xprev = x
	}

	if x.Cmp(bOne) != 0 {
		return true
	}

	return false
}

// MillerRabin return:
// 1 if Simple
// -1 if not Simple
// error else
func MillerRabin(n *big.Int, s int) (int, error) {
	max := new(big.Int).Sub(n, bOne)
	for j := 0; j < s; j++ {
		a, err := rand.Int(rand.Reader, max)
		if err != nil {
			return 0, err
		}

		if witness(a, n) {
			return -1, nil
		}
	}

	return 1, nil
}

func PollardRho(n *big.Int, results chan<- *big.Int) error {
	i := 1
	x1, err := rand.Int(rand.Reader, new(big.Int).Sub(n, bOne))
	if err != nil {
		return err
	}

	y := &*x1
	k := 2
	xprev := &*x1
	for {
		i++
		x := new(big.Int).Sub(new(big.Int).Exp(xprev, big.NewInt(2), nil), bOne)
		x = new(big.Int).Mod(x, n)

		d := GcdEuclid(new(big.Int).Sub(y, x), n)
		if d.Cmp(bOne) != 0 && d.Cmp(n) != 0 {
			results <- d
		}

		if i == k {
			y = &*x
			k *= 2
		}

		xprev = &*x
	}
}
