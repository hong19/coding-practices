package leetcode881

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Slice(people, func(i, j int) bool {
		return people[i] > people[j]
	})
	fmt.Printf("people: %v\n", people)

	head := 0
	tail := len(people) - 1

	for head <= tail {
		remaining := limit - people[head]

		if remaining >= people[tail] {
			tail--
		}
		head++
	}

	return head
}
