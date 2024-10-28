package two_sum

// Indices of two numbers such that key add up to target
//
// Complexity: O(n)
func twoSum(nums []int, target int) []int {
	exist := make(map[int]int)
	for index, num := range nums {
		x := target - num
		if value, ok := exist[x]; ok {
			return []int{index, value}
		}
		exist[num] = index
	}
	return []int{}
}
