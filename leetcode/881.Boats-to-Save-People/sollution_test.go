package leetcode881

import (
	"fmt"
	"testing"
)

type params struct {
	people []int
	limit  int
}
type testCase struct {
	params
	expected int
}

func Test_numRescueBoats(t *testing.T) {
	// nums := []int{1, 1, 1, 2, 2, 3}      // Input array
	// expectedNums := []int{1, 1, 2, 2, 3} // The expected answer with correct length

	testCases := []testCase{
		{
			params{
				[]int{1, 2},
				3,
			},
			1,
		},
		{
			params{
				[]int{3, 2, 2, 1},
				3,
			},
			3,
		},
		{
			params{
				[]int{3, 5, 3, 4},
				5,
			},
			4,
		},
		{
			params{
				[]int{3, 1, 7},
				7,
			},
			2,
		},
	}

	for _, tc := range testCases {
		result := numRescueBoats(tc.params.people, tc.params.limit)
		if result != tc.expected {
			fmt.Printf("people: %v\n", tc.params.people)
			fmt.Printf("expected: %d, result: %d\n", tc.expected, result)
			t.Error()
			break
		}
	}
}
