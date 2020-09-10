package lab0028

func StrStr(haystack string, needle string) int {
	if needle == "" || needle == haystack {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
