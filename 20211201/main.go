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

const windowSize = 3

func part2(numbers []int) int {
	sums := []int{}
	for i := windowSize - 1; i < len(numbers); i++ {
		sum := 0
		for n := 0; n < windowSize; n++ {
			sum += numbers[i-n]
		}
		sums = append(sums, sum)
	}
	return part1(sums)
}
