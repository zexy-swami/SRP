package SRP

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/zexy-swami/SRP/SRP_CLI/pkg/big_int_handlers"
	"github.com/zexy-swami/SRP/SRP_CLI/pkg/parser"
)

type Prover struct {
	srp  *SRP
	id   string
	a    *big.Int
	A    *big.Int
	salt string
	S    *big.Int
	K    *big.Int
	M    *big.Int
}

func NewProver(pfIndex int, proverID string) *Prover {
	generator, safePrime := getGenerator(pfIndex), getSafePrime(pfIndex)

	hashArg := generator.String() + safePrime.String()
	hashResult := sha256.Sum256([]byte(hashArg))
	k := big_int_handlers.BytesToBigInt(hashResult[:])

	return &Prover{
		srp: &SRP{
			safePrime: safePrime,
			generator: generator,
			k:         k,
		},
		id: proverID,
	}
}

func (prover *Prover) GenerateClientPublicValue() *big.Int {
	a := big_int_handlers.RandBigInt(1024)
	A := big.NewInt(0).Exp(prover.srp.generator, a, prover.srp.safePrime)
	prover.a = a
	prover.A = A

	return A
}

func (prover *Prover) GenerateMValue(publicServerValue, salt string) *big.Int {
	prover.salt = salt

	// u = H(A, B)
	uAsBytes := sha256.Sum256([]byte(prover.A.String() + publicServerValue))
	u := big_int_handlers.BytesToBigInt(uAsBytes[:])

	// x = H(s, p)
	userPassword := parser.GetDataFromConfig("password")
	userPasswordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(userPassword)))
	xAsBytes := sha256.Sum256([]byte(salt + userPasswordHash))
	x := big_int_handlers.BytesToBigInt(xAsBytes[:])

	B := big_int_handlers.StringToBigInt(publicServerValue)
	prover.generateSValue(B, x, u)

	// K = H(S)
	KAsSlc := sha256.Sum256([]byte(prover.S.String()))
	prover.K = big_int_handlers.BytesToBigInt(KAsSlc[:])

	prover.generateMValue(B)

	return prover.M
}

func (prover *Prover) generateSValue(B, x, u *big.Int) {
	S1, S2, S3 := new(big.Int), new(big.Int), new(big.Int)
	S2 = S2.Exp(prover.srp.generator, x, prover.srp.safePrime)
	S1 = S1.Mul(prover.srp.k, S2)
	S1 = S1.Sub(B, S1)
	S3 = S3.Mul(u, x)
	S2 = S2.Add(prover.a, S3)
	S1 = S1.Exp(S1, S2, prover.srp.safePrime)
	prover.S = S1
}

func (prover *Prover) generateMValue(B *big.Int) {
	var MArg string
	MArg += prover.K.String()
	MArg += prover.A.String()
	MArg += B.String()
	MArg += prover.id
	MArg += prover.salt
	MArg += prover.srp.safePrime.String()
	MArg += prover.srp.generator.String()
	MAsBytes := sha256.Sum256([]byte(MArg))
	prover.M = big_int_handlers.BytesToBigInt(MAsBytes[:])
}

func (prover *Prover) CompareZValues(serverZValue string) bool {
	serverZValueAsBigInt := big_int_handlers.StringToBigInt(serverZValue)
	zValueAsSlc := sha256.Sum256([]byte(prover.M.String() + prover.K.String()))
	zValue := big_int_handlers.BytesToBigInt(zValueAsSlc[:])

	return zValue.Cmp(serverZValueAsBigInt) == 0
}
