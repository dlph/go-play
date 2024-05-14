package rotatearray

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/
func rotate(nums []int, k int) {
	if nums == nil {
		return
	}

	if len(nums) == 0 {
		return
	}

	if k <= 0 {
		return
	}

	for i := 0; i < k; i++ {
		swp := nums[len(nums)-1]
		copy(nums[1:], nums[:len(nums)-1])
		nums[0] = swp
	}
}
