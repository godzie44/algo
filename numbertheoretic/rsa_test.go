package numbertheoretic

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestRsaKeys(t *testing.T) {
	_, _, err := RsaKeys(1024)
	assert.NoError(t, err)
}

func TestRsaDecodeEncode(t *testing.T) {
	pub, secret, err := RsaKeys(1024)
	assert.NoError(t, err)

	msg := make([]byte, 8)
	binary.BigEndian.PutUint64(msg, 80087)
	cypher := RsaEncode(msg, pub)
	res := RsaDecode(cypher, secret)
	assert.Equal(t, 0, new(big.Int).SetBytes(res).Cmp(big.NewInt(80087)))

	pub, secret, err = RsaKeys(1024)
	assert.NoError(t, err)

	cypher = RsaEncode([]byte("hello world"), pub)
	res = RsaDecode(cypher, secret)
	assert.Equal(t, []byte("hello world"), res)
}
