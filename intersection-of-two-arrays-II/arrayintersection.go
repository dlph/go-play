package palindromelinkedlist

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/674/

func intersect(nums1 []int, nums2 []int) []int {
	if nums1 == nil {
		return nil
	}
	if nums2 == nil {
		return nil
	}

	if len(nums1) == 0 {
		return nil
	}
	if len(nums2) == 0 {
		return nil
	}

	freq1 := frequency(nums1)
	freq2 := frequency(nums2)

	res := make([]int, 0, max(len(freq1), len(freq2)))
	for k, v1 := range freq1 {
		if v2, exists := freq2[k]; exists {
			for i := 0; i < min(v1, v2); i++ {
				res = append(res, k)
			}
		}
	}

	return res
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}

func max[T int](left, right T) T {
	if left > right {
		return left
	}
	return right
}

func frequency[T int](slice []T) map[T]int {
	freq := make(map[T]int)
	for i := range slice {
		if _, exists := freq[slice[i]]; !exists {
			freq[slice[i]] = 1
			continue
		}
		freq[slice[i]]++
	}
	return freq
}
