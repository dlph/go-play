package rotatearray

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	type input struct {
		nums []int
		k    int
	}
	type output struct {
		nums []int
	}

	testCases := []struct {
		name string
		in   input
		out  output
	}{
		{
			in: input{
				nums: []int{1, 2},
				k:    3,
			},
			out: output{
				nums: []int{2, 1},
			},
		},
		{
			in: input{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			out: output{
				nums: []int{5, 6, 7, 1, 2, 3, 4},
			},
		},
		{
			in: input{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			out: output{
				nums: []int{3, 99, -1, -100},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var res []int = make([]int, len(tc.in.nums))
			copy(res, tc.in.nums)
			rotate(res, tc.in.k)

			if len(res) != len(tc.out.nums) {
				t.Fatal("not same result length")
			}

			for i := 0; i < len(res); i++ {
				if res[i] != tc.out.nums[i] {
					t.Errorf("%d) %d != %d\n", i, res[i], tc.out.nums[i])
				}
			}
		})
	}
}
