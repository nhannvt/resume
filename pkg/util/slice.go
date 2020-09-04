package util

// ContainString checkes if given string is included in given slice.
func ContainString(slice []string, s string) bool {
	for _, v := range slice {
		if s == v {
			return true
		}
	}
	return false
}

// ContainInt checkes if given int value is included in given slice.
func ContainInt(slice []int, i int) bool {
	for _, v := range slice {
		if i == v {
			return true
		}
	}
	return false
}
