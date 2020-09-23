package utils

func InArray(s []interface{}, element interface{}) bool {
	for _, v := range s {
		if v == element {
			return true
		}
	}
	return false
}
func ArrayRemove(s []interface{}, element interface{}) (otherArr []interface{}) {
	for _, v := range s {
		if v == element {
			otherArr = append(otherArr, v)
		}
	}
	return
}
