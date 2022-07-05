package main

import (
	"fmt"
)

var mas []int

func less(i, j int) bool {
	if mas[i] < mas[j] {
		return true
	} else {
		return false
	}
}
func swap(i, j int) {
	mas[i], mas[j] = mas[j], mas[i]
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	qsortrec(less, swap, 0, n-1)
}

func qsortrec(less func(i, j int) bool, swap func(i, j int), low int, high int) {
	if low < high {
		q := Partition(less, swap, low, high)
		qsortrec(less, swap, low, q-1)
		qsortrec(less, swap, q+1, high)
	}
}

func Partition(less func(i, j int) bool, swap func(i, j int), low int, high int) int {
	i := low
	j := low
	for j < high {
		if less(j, high) {
			swap(i, j)
			i++
		}
		j++
	}
	swap(i, high)
	return i
}

func main() {
	var N int
	fmt.Scan(&N)
	mas = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&mas[i])
	}

	qsort(N, less, swap)

	for i := 0; i < N; i++ {
		fmt.Printf("%d ", mas[i])
	}
}
