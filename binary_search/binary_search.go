package binary_search

func search(nums []int, target int) int {
	var low, hight int = 0, len(nums) - 1

	for low <= hight {
		mid := (low + hight) / 2
		guess := nums[mid]

		if guess == target {
			return mid
		}
		if guess > target {
			hight = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
