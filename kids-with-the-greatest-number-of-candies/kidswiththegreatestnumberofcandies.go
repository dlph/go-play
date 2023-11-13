package kidswiththegreatestnumberofcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	if len(candies) == 0 {
		return []bool{}
	}

	var maxCandies func([]int) int
	maxCandies = func(i []int) int {
		if len(i) == 0 {
			return 0
		}

		head := i[0]
		tail := maxCandies(i[1:])
		if head >= tail {
			return head
		}

		return tail
	}

	greatestKid := make([]bool, len(candies))
	maxCandy := maxCandies(candies)
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxCandy {
			greatestKid[i] = true
		}
	}

	return greatestKid
}
