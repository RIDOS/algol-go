package remove_element

func removeElement(nums []int, val int) int {
	k := 0

	for index := 0; index < len(nums); index++ {
		if nums[index] != val {
			nums[k] = nums[index]
			k++
		}
	}

	return k
}
