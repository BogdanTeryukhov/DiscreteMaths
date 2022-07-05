package main

import "fmt"

func add(a, b []int32, p int) []int32 {
	var (
		shrek         []int32
		lastMasCommon []int
		expr          int32
		flag          bool
		binger        int
		preorder      int
		countZeroes   int
		iter          int
		mitch         int
		temp          []int32
	)
	if len(b) >= len(a) {
		temp = a
		a = b
		b = temp
	}

	expr = 0
	flag = false

	if len(b) < len(a) {
		mitch = len(b)
	} else {
		mitch = len(a) + 10
	}

	for i := 0; i < len(a); i++ {

		if flag {
			if i >= mitch {
				expr = a[i] + int32(binger)
			} else {
				expr = a[i] + b[i] + int32(binger)
			}
			//expr = a[i] + b[i] + int32(binger)
			binger = 0
			flag = false
		} else {
			if i >= mitch {
				expr = a[i]
			} else {
				expr = a[i] + b[i]
			}
		}

		if int(expr) >= p {
			binger = 1
			//a[i+1]++
			expr -= int32(p)
			flag = true
			lastMasCommon = append(lastMasCommon, int(expr))
		} else {
			lastMasCommon = append(lastMasCommon, int(expr))
		}

		if i == len(a)-1 {
			if flag {
				preorder = 1
			}
		}
	}
	lastMasCommon = append(lastMasCommon, preorder)

	lastMas32 := make([]int32, len(lastMasCommon))
	for i := 0; i < len(lastMasCommon); i++ {
		lastMas32[i] = int32(lastMasCommon[i])
	}

	iter = len(lastMas32) - 1
	for lastMas32[iter] == 0 && iter > 0 {
		countZeroes++
		iter--
	}

	shrek = lastMas32[:len(lastMas32)-countZeroes]
	return shrek
}

func main() {

	var (
		p        int
		n1       int
		n2       int
		replyMas []int32
	)
	fmt.Scan(&n1, &n2)
	fmt.Scan(&p)
	aa := make([]int32, n1)
	bb := make([]int32, n2)
	for i := 0; i < n1; i++ {
		fmt.Scan(&aa[i])
	}

	for j := 0; j < n2; j++ {
		fmt.Scan(&bb[j])
	}

	replyMas = add(aa, bb, p)
	fmt.Println(replyMas)
}
