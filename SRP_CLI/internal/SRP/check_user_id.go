package SRP

import "github.com/zexy-swami/SRP/SRP_CLI/internal/db"

func CheckUserID(userID string) bool {
	count := db.GetCount(userID)
	return count == 1
}
