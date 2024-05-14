package climbstairs

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/569/

func climbStairs(n int) int {
	return memoizedClimbStairs()(n)
}

func memoizedClimbStairs() func(int) int {
	cache := make(map[int]int)

	var fn func(n int) int
	fn = func(n int) int {
		if x, found := cache[n]; found {
			return x
		}

		if n < 0 {
			return 0
		}

		if n == 0 {
			return 1
		}

		if n == 1 {
			return 1
		}

		x := fn(n-2) + fn(n-1)
		cache[n] = x
		return x
	}

	return fn
}

func climbStairsStep(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	if n == 1 {
		return 1
	}

	return climbStairsStep(n-2) + climbStairsStep(n-1)
}
