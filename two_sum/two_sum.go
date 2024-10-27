package two_sum

// Indices of two numbers such that key add up to target
//
// Complexity: O(n)
func twoSum(nums []int, target int) []int {
	var exist = make(map[int]int)
	for index, num := range nums {
		x := target - num
		if value, ok := exist[x]; !ok {
			exist[num] = index
		} else {
			return []int{index, value}
		}
	}
	return []int{}
}
