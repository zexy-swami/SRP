package SRP

import (
	"crypto/sha256"
	"math/big"

	"github.com/zexy-swami/SRP/SRP_CLI/internal/db"
	"github.com/zexy-swami/SRP/SRP_CLI/pkg/big_int_handlers"
)

type Verifier struct {
	srp      *SRP
	proverID string
	salt     *big.Int
	v        *big.Int
	b        *big.Int
	B        *big.Int
	S        *big.Int
	K        *big.Int
	M        *big.Int
}

func NewVerifier(pfIndex int, proverID string) *Verifier {
	generator, safePrime := getGenerator(pfIndex), getSafePrime(pfIndex)

	// k = H(g, p)
	hashArg := generator.String() + safePrime.String()
	hashResult := sha256.Sum256([]byte(hashArg))
	k := big_int_handlers.BytesToBigInt(hashResult[:])

	// x = H(salt, user_password_hash)
	salt := big_int_handlers.RandBigInt(1024)
	userPasswordHash := db.GetPasswordHash(proverID)
	xAsBytes := sha256.Sum256([]byte(salt.String() + userPasswordHash))
	x := big_int_handlers.BytesToBigInt(xAsBytes[:])

	// v = g**x mod p
	v := new(big.Int)
	v = v.Exp(generator, x, safePrime)

	return &Verifier{
		srp: &SRP{
			safePrime: safePrime,
			generator: generator,
			k:         k,
		},
		proverID: proverID,
		salt:     salt,
		v:        v,
	}
}

func (verifier *Verifier) GenerateServerPublicValueAndSalt(clientPublicValue string) (*big.Int, *big.Int) {
	b := big_int_handlers.RandBigInt(1024)
	verifier.b = b

	// B = (k * v + (g**x mod p)) mod p
	verifier.generateBValue()

	// u = H(A, B)
	A := big_int_handlers.StringToBigInt(clientPublicValue)
	uAsBytes := sha256.Sum256([]byte(A.String() + verifier.B.String()))
	u := big_int_handlers.BytesToBigInt(uAsBytes[:])

	// S = (A * (v**u mod p))**b mod p
	verifier.generateSValue(u, A)

	// K = H(S)
	KAsBytes := sha256.Sum256([]byte(verifier.S.String()))
	verifier.K = big_int_handlers.BytesToBigInt(KAsBytes[:])

	verifier.generateMValue(A)

	return verifier.B, verifier.salt
}

func (verifier *Verifier) generateBValue() {
	B, B1, B2 := new(big.Int), new(big.Int), new(big.Int)
	B1 = B1.Mul(verifier.srp.k, verifier.v)
	B2 = B2.Exp(verifier.srp.generator, verifier.b, verifier.srp.safePrime)
	B = B.Add(B1, B2)
	B = B.Exp(B, big.NewInt(1), verifier.srp.safePrime)
	verifier.B = B
}

func (verifier *Verifier) generateSValue(u, A *big.Int) {
	S := new(big.Int)
	S = S.Exp(verifier.v, u, verifier.srp.safePrime)
	S = S.Mul(S, A)
	S = S.Exp(S, verifier.b, verifier.srp.safePrime)
	verifier.S = S
}

func (verifier *Verifier) generateMValue(A *big.Int) {
	var MArg string
	MArg += verifier.K.String()
	MArg += A.String()
	MArg += verifier.B.String()
	MArg += verifier.proverID
	MArg += verifier.salt.String()
	MArg += verifier.srp.safePrime.String()
	MArg += verifier.srp.generator.String()

	MAsBytes := sha256.Sum256([]byte(MArg))
	verifier.M = big_int_handlers.BytesToBigInt(MAsBytes[:])
}

func (verifier *Verifier) CompareMValues(clientMValue string) bool {
	clientMValueAsBigInt := big_int_handlers.StringToBigInt(clientMValue)
	return verifier.M.Cmp(clientMValueAsBigInt) == 0
}

func (verifier *Verifier) GenerateZValue() *big.Int {
	zValueAsSlc := sha256.Sum256([]byte(verifier.M.String() + verifier.K.String()))
	zValue := big_int_handlers.BytesToBigInt(zValueAsSlc[:])
	return zValue
}
