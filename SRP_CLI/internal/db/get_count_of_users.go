package db

const selectCountStatement = "select count(*) from users where srp_id = $1;"

func GetCount(checkValue string) int {
	var count int
	row := dbConn.QueryRow(selectCountStatement, checkValue)
	_ = row.Scan(&count)

	return count
}
