package helper

func Index(slice []string, t string) int {

	for i, v := range slice {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(slice []string, t string) bool {

	return Index(slice, t) >= 0
}

func Filter(slice []string, callback func(string) bool) []string {

	ret := make([]string, 0)
	for _, v := range slice {
		if callback(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func Map(slice []string, callback func(string) string) []string {

	ret := make([]string, len(slice))

	for i, v := range slice {

		ret[i] = callback(v)
	}
	return ret
}
