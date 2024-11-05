package valid_parentheses

func isValid(s string) bool {
	rev := map[byte]byte{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	arr := make([]byte, 0, 0)

	for key := range s {
		value := s[key]

		if value == '{' || value == '[' || value == '(' {
			arr = append(arr, rev[value])
		} else {
			if len(arr) > 0 && value == arr[len(arr)-1] {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		}
	}

	if len(arr) > 0 {
		return false
	}

	return true
}
