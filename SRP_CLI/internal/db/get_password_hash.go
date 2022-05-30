package db

var selectPasswordHashStatement = "select user_password_hash from users where srp_id = $1;"

func GetPasswordHash(userID string) string {
	var passwordHash string
	row := dbConn.QueryRow(selectPasswordHashStatement, userID)
	_ = row.Scan(&passwordHash)

	return passwordHash
}
