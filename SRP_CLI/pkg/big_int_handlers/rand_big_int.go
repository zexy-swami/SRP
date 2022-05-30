package big_int_handlers

import (
	"crypto/rand"
	"io"
	"math/big"
)

func RandBigInt(bitsCount int) *big.Int {
	bytesBuffer := randomBytes(bitsCount)
	randomBigInt := big.NewInt(0).SetBytes(bytesBuffer)
	return randomBigInt
}

func randomBytes(bitsCount int) []byte {
	bytesBuffer := make([]byte, bitsCount/8)
	_, _ = io.ReadFull(rand.Reader, bytesBuffer)
	return bytesBuffer
}
