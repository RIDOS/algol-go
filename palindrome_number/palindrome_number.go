package palindrome_number

// isPalindrome determines whether the given int is a palindrome.
// A palindrome is a number (or a text phrase) that reads the same backwards as forwards.
// For example, the number 12321 is a palindrome.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	rev := 0
	cur := x

	for cur > 0 {
		d := cur % 10
		cur /= 10
		rev = rev*10 + d
	}

	return rev == x
}
