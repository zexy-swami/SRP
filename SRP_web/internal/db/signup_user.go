package db

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

const insertStatement = "insert into users (srp_id, user_password_hash) values ($1, $2, $3, $4)"

func SignupUser(firstname, lastname, password string) (srpID string, retError error) {
	defer func() {
		if err := recover(); err != nil {
			retError = errors.New("user already signed up")
		}
	}()

	srpID = fmt.Sprintf("%x", sha256.Sum256([]byte(firstname+lastname+password)))
	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	tx := dbConn.MustBegin()
	tx.MustExec(insertStatement, srpID, firstname, lastname, passwordHash)
	if retError = tx.Commit(); retError != nil {
		return "", retError
	}

	return srpID, nil
}
