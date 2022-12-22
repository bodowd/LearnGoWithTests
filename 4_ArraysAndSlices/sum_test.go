package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5, 6}

		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	assertCorrectDeepEqual(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	// new technique, assign function to variable
	// useful when you want to bind a function to other local variables in "scope"
	// also reduces the surface area of your API
	checkSums := func(t testing.TB, got, want []int) {

		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)

	})

	t.Run("safely sum empty slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concat := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b"}, concat, ""), "ab")

	})
}

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25)

}

func assertCorrectDeepEqual(t testing.TB, got []int, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}

}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
