package SRP

import (
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/zexy-swami/SRP/SRP_CLI/pkg/big_int_handlers"
)

type primeField struct {
	generator *big.Int
	safePrime *big.Int
}

const countOfPrimeFields = 990

var primeFields [countOfPrimeFields]*primeField

func init() {
	generatorsAndSafePrimesAsStringSlice := strings.Split(generatorsAndSafePrimes, "\n")
	for i, generatorAndSafePrime := range generatorsAndSafePrimesAsStringSlice {
		primeFields[i] = &primeField{
			generator: big_int_handlers.StringToBigInt(strings.Split(generatorAndSafePrime, " ")[0]),
			safePrime: big_int_handlers.StringToBigInt(strings.Split(generatorAndSafePrime, " ")[1]),
		}
	}
}

func GetPrimeFieldIndex() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(countOfPrimeFields)
}

func getGenerator(index int) *big.Int {
	return primeFields[index].generator
}

func getSafePrime(index int) *big.Int {
	return primeFields[index].safePrime
}
