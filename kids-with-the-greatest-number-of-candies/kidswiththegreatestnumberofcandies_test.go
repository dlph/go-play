package kidswiththegreatestnumberofcandies

import "testing"

func TestKidsWithCandies(t *testing.T) {
	type input struct {
		candies      []int
		extraCandies int
	}
	type want struct {
		output []bool
	}

	testCases := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "Example 1",
			input: input{
				candies:      []int{2, 3, 5, 1, 3},
				extraCandies: 3,
			},
			want: want{
				output: []bool{true, true, true, false, true},
			},
		},
		{
			name: "Example 2",
			input: input{
				candies:      []int{4, 2, 1, 1, 2},
				extraCandies: 1,
			},
			want: want{
				output: []bool{true, false, false, false, false},
			},
		},
		{
			name: "Example 3",
			input: input{
				candies:      []int{12, 1, 12},
				extraCandies: 10,
			},
			want: want{
				output: []bool{true, false, true},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			have := kidsWithCandies(tc.input.candies, tc.input.extraCandies)

			if len(have) != len(tc.want.output) {
				t.Fatalf("not same length have: %d want: %d\n", len(have), len(tc.want.output))
			}

			for i := range have {
				if have[i] != tc.want.output[i] {
					t.Errorf("not a cool kid at %d", i)
				}
			}
		})
	}
}
