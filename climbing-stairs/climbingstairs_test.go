package climbstairs

import "testing"

func Test_reverseVowels(t *testing.T) {
	type input struct {
		n int
	}
	type want struct {
		n int
	}

	testCases := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "test-a",
			input: input{
				n: 4,
			},
			want: want{
				n: 5,
			},
		},
		{
			name: "test-b",
			input: input{
				n: 5,
			},
			want: want{
				n: 8,
			},
		},
		{
			name: "Example 2",
			input: input{
				n: 3,
			},
			want: want{
				n: 3,
			},
		},
		{
			name: "Example 1",
			input: input{
				n: 2,
			},
			want: want{
				n: 2,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			have := climbStairs(tc.input.n)

			if have != tc.want.n {
				t.Errorf("have: %v want: %v\n", have, tc.want.n)
			}
		})
	}
}
