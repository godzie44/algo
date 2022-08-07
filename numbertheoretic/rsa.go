package numbertheoretic

import (
	"crypto/rand"
	"math/big"
)

var e = big.NewInt(35537)

type PubKey struct {
	N []byte
	E []byte
}

type SecretKey struct {
	N []byte
	D []byte
}

func RsaKeys(bits int) (*PubKey, *SecretKey, error) {
	p, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	q, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	n := new(big.Int).Mul(p, q)
	fn := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	d, err := ModularLinearEquationSolver(e, big.NewInt(1), fn)
	if err != nil {
		return nil, nil, err
	}

	return &PubKey{E: e.Bytes(), N: n.Bytes()}, &SecretKey{D: d[0].Bytes(), N: n.Bytes()}, nil
}

func RsaEncode(msg []byte, pubKey *PubKey) []byte {
	m := new(big.Int).SetBytes(msg)
	n := new(big.Int).SetBytes(pubKey.N)

	pm := ModularExponentiation(m, new(big.Int).SetBytes(pubKey.E), n)

	return pm.Bytes()
}

func RsaDecode(cypher []byte, secretKey *SecretKey) []byte {
	c := new(big.Int).SetBytes(cypher)
	d := new(big.Int).SetBytes(secretKey.D)
	n := new(big.Int).SetBytes(secretKey.N)

	sc := ModularExponentiation(c, d, n)

	return sc.Bytes()
}
