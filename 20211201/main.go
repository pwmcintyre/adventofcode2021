package main

func part1(numbers []int) int {
	increasedCount := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			increasedCount++
		}
	}
	return increasedCount
}
