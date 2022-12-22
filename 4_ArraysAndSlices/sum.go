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

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
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
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == client {
			return currentBalance - t.Sum
		}

		if t.To == client {
			return currentBalance + t.Sum
		}
		return currentBalance
	}

	return Reduce(t, adjustBalance, 0)
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	return
}
