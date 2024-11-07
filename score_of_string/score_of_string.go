package scoreofstring

import "math"

func scoreOfString(s string) int {
	score := 0

	for i := 0; i < len(s)-1; i++ {
		score += int(math.Abs(float64(s[i]) - float64(s[i+1])))
	}

	return score
}
