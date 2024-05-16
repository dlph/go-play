package palindromelinkedlist

import "testing"

func Test_fn(t *testing.T) {
	type input struct {
		nums1 []int
		nums2 []int
	}
	type want struct {
		output []int
	}

	testCases := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "Eaxmple 1",
			input: input{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: want{
				output: []int{2, 2},
			},
		},
		{
			name: "Eaxmple 2",
			input: input{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: want{
				output: []int{9, 4},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			have := intersect(tc.input.nums1, tc.input.nums2)

			if have == nil {
				t.Fatalf("have is nil\n")
			}
			if len(have) != len(tc.want.output) {
				t.Fatalf("not matching slice size, have: %d, want: %d\n", len(have), len(tc.want.output))
			}
		})
	}
}
