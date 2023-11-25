package main

import "sort"

func NumRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	heaviest, lightest := len(people)-1, 0
	boats := 0

	for heaviest >= lightest {
		if people[heaviest]+people[lightest] <= limit {
			heaviest--
			lightest++
		} else {
			heaviest--
		}
		boats++
	}
	return boats
}
