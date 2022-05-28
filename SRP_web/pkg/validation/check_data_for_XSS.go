package validation

import "github.com/microcosm-cc/bluemonday"

func CheckDataForXSS(data [3]string) (bool, string) {
	policy := bluemonday.StrictPolicy()
	for i := 0; i < 3; i++ {
		if policy.Sanitize(data[i]) != data[i] {
			return true, data[i]
		}
	}
	return false, ""
}
