package validation

func CheckDataForEmptiness(data [3]string) bool {
	return len(data[0]) == 0 || len(data[1]) == 0 || len(data[2]) == 0
}
