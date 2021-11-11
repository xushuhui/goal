package utils

func InSlice(s []interface{}, element interface{}) bool {
	for _, v := range s {
		if v == element {
			return true
		}
	}
	return false
}
