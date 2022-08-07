package numbertheoretic

import (
	"crypto/rand"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"runtime"
	"testing"
	"time"
)

func TestGcdEuclid(t *testing.T) {
	assert.Equal(t, int64(3), GcdEuclid(big.NewInt(30), big.NewInt(21)).Int64())
}

func TestGcdExtendedEuclid(t *testing.T) {
	d, x, y := GcdExtendedEuclid(big.NewInt(99), big.NewInt(78))
	assert.Equal(t, int64(3), d.Int64())
	assert.Equal(t, int64(-11), x.Int64())
	assert.Equal(t, int64(14), y.Int64())
}

func TestModularLinearEquationSolver(t *testing.T) {
	result, err := ModularLinearEquationSolver(big.NewInt(14), big.NewInt(30), big.NewInt(100))
	assert.NoError(t, err)
	assert.Equal(t, []*big.Int{big.NewInt(95), big.NewInt(45)}, result)
}

func TestModularLinearEquationSolver2(t *testing.T) {
	result, err := ModularLinearEquationSolver(big.NewInt(65537), big.NewInt(1), big.NewInt(2873905464))
	assert.NoError(t, err)
	assert.Equal(t, []*big.Int{big.NewInt(155410241)}, result)
}

func TestModularExponentiation(t *testing.T) {
	exp := ModularExponentiation(big.NewInt(7), big.NewInt(560), big.NewInt(561))
	assert.Equal(t, int64(1), exp.Int64())
}

func TestMillerRabin(t *testing.T) {
	prime, err := rand.Prime(rand.Reader, 64)
	assert.NoError(t, err)

	res, err := MillerRabin(prime, 32)
	assert.NoError(t, err)
	assert.Equal(t, 1, res)

	notPrime := big.NewInt(9 * 101)
	res, err = MillerRabin(notPrime, 32)
	assert.NoError(t, err)
	assert.Equal(t, -1, res)
}

func TestPollardRho(t *testing.T) {
	resChan := make(chan *big.Int, 1000)
	go func() {
		val, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		fmt.Printf("Number: %d \n", val.Int64())

		PollardRho(val, resChan)
	}()

	timer := time.After(time.Millisecond)
	for {
		select {
		case v := <-resChan:
			fmt.Println(v)
			runtime.Gosched()

			select {
			case <-timer:
				return
			default:
			}
		}
	}
}
