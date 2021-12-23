package slice

func sliceSum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return
}

func sliceMinMax(numbers []int) (min, max int) {
	if len(numbers) == 0 {
		return 0, 0
	}
	min = numbers[0]
	max = numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return
}
