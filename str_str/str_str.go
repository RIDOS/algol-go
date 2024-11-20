package str_str

func strStr(haystack, needle string) int {
	for i := 0; i < len(haystack); i++ {
		k := len(needle)
		if haystack[i] == needle[0] && len(needle) <= len(haystack)-i {
			for j := 0; j < len(needle); j++ {
				if haystack[i+j] == needle[j] {
					k--
				}
				if k == 0 {
					return i
				}
			}
		}
	}
	return -1
}
