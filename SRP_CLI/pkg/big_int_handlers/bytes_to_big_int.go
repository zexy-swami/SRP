package big_int_handlers

import "math/big"

func BytesToBigInt(byteSlc []byte) *big.Int {
	bigIntValue := big.NewInt(0).SetBytes(byteSlc)
	return bigIntValue
}
