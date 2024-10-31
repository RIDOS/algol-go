package romantoint

// romanToInt converts a Roman numeral string to an integer.
// The function iterates over each character in the string, using a map to
// determine the integer value of each Roman numeral. It handles subtractive
// notation by checking if a numeral is smaller than the one following it and
// adjusting the result accordingly.
func romanToInt(s string) int {
	romanNums := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	for index := range s {
		if nextValue, ok := romanNums[s[index]]; ok {
			if index == 0 {
				result += nextValue
			} else if reduxValue, ok := romanNums[s[index-1]]; ok {
				if nextValue > reduxValue {
					result += (nextValue - reduxValue) - reduxValue
				} else {
					result += nextValue
				}
			}
		}
	}

	return result
}
