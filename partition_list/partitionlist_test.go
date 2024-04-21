package partitionlist

import (
	"testing"
)

func Test_partition(t *testing.T) {

	type input struct {
		node *ListNode
		x    int
	}
	type output struct {
		node *ListNode
	}

	testCases := []struct {
		name string
		in   input
		out  output
	}{
		{
			in:  input{},
			out: output{},
		},
		{
			in: input{
				node: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
				x: 2,
			},
			out: output{
				node: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
		},
		{
			in: input{
				node: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val:  2,
										Next: nil,
									},
								},
							},
						},
					},
				},
				x: 3,
			},
			out: output{
				node: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 3,
									Next: &ListNode{
										Val:  5,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := partition(tc.in.node, tc.in.x)

			if result == nil {
				if tc.out.node != nil {
					t.Fatal("node nil")
				}
				return // success - want node to be nil
			}

			node := result
			tcNode := tc.out.node
			for node != nil {
				if tcNode == nil {
					t.Fatal("node not nil")
					return
				}

				node = node.Next
				tcNode = tcNode.Next
			}
		})
	}
}
