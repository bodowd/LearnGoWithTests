package main

func Sum(numbers []int) int {
	add := func(acc, x int) int {
		return acc + x
	}
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(res, x []int) []int {
		if len(x) == 0 {
			return append(res, 0)
		} else {
			tail := x[1:]
			return append(res, Sum(tail))
		}
	}

	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(t []Transaction, client string) float64 {
	var balance float64
	for _, t := range t {
		if t.From == client {
			balance -= t.Sum
		}
		if t.To == client {
			balance += t.Sum
		}
	}
	return balance
}
