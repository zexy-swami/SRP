package big_int_handlers

import "math/big"

func StringToBigInt(value string) *big.Int {
	bigIntValue := new(big.Int)
	bigIntValue, _ = bigIntValue.SetString(value, 10)
	return bigIntValue
}
