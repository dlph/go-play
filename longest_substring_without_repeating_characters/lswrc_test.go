package longestsubstringwithoutrepeatingcharacters

import (
	"testing"
)

func Test_memoizedLengthOfLongestSubstring(t *testing.T) {

	type input struct {
		s string
	}
	type output struct {
		i int
	}

	testCases := []struct {
		in  input
		out output
	}{
		{
			in: input{
				s: "dvdf",
			},
			out: output{
				i: 3,
			},
		},
		{
			in: input{
				s: "aab",
			},
			out: output{
				i: 2,
			},
		},
		{
			in: input{
				s: "abcabcbbe",
			},
			out: output{
				i: 3,
			},
		},
		{
			in: input{
				s: "bbbbb",
			},
			out: output{
				i: 1,
			},
		},
		{
			in: input{
				s: "pwwkew",
			},
			out: output{
				i: 3,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.in.s, func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.in.s)

			if result != tc.out.i {
				t.Errorf("not expected longest have: %d want: %d\n", result, tc.out.i)
			}
		})
	}

}
