package db

func CloseDB() {
	_ = dbConn.Close()
}
