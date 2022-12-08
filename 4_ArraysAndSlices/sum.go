package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	numberOfSlices := len(numbersToSum)
	// create a slice with a starting capacity of the len of the nuberOfSlices we need to work through
	sums := make([]int, numberOfSlices)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums

}
