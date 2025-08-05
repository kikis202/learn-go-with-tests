package arrays

func Sum(arr []int) (sum int) {
	add := func(acc, v int) int {
		return acc + v
	}
	return Reduce(arr, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc, v []int) []int {
		if len(v) == 0 {
			acc = append(acc, 0)
		} else {
			tail := v[1:]
			acc = append(acc, Sum(tail))
		}

		return acc
	}

	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[T any](arr []T, f func(sum T, val T) T, initialVal T) T {
	sums := initialVal
	for _, val := range arr {
		sums = f(sums, val)
	}
	return sums
}
