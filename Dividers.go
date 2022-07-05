package main

import "fmt"

func FindDividers(x uint32) []uint32 {
	mas := make([]uint32, x)
	index := 1
	count := 0
	mas[0] = x
	for i := x / 2; i > 0; i-- {
		if x%i == 0 {
			mas[index] = i
			index++
			count++
		}
	}
	return mas[:count+1]
}

func main() {
	var x uint32
	fmt.Scan(&x)
	mas := FindDividers(x)
	fmt.Println("graph {")
	for i := 0; i < len(mas); i++ {
		fmt.Println("\t", mas[i])
	}
	/*
		index := 0
		temp := make([]int, 3)
		fl := false
		fmt.Println(mas[0], "--", mas[1])
		temp[index] = mas[1]
		index++

		k := mas[0]
		for i := 2; i < len(mas); i++ {
			if !fl {
				for j := 1; j <= i; j++ {
					fl = true
					if mas[i-1]%mas[j+1] != 0 {
						fmt.Println(k, "--", mas[j+1])
						temp[index] = mas[j+1]
					} else {
						fmt.Println(k, "--", mas[j+1])
						temp[index] = mas[j+1]
						break
					}
					index++
				}
			}
			//k = mas[i-1]
		}
		for i := 0; i < len(temp); i++ {
			march := FindDividers(temp[i])
			for j := 1; j < len(march); j++ {
				if march[j] != 1 {
					fmt.Println(temp[i], "--", march[j])
				}
			}
		}
	*/
	fl := false

	for i := 0; i < len(mas); i++ {
		for j := 0; j < len(mas); j++ {
			if (mas[i] != 0) && (mas[j] != 0) && (mas[i] != mas[j]) && (i != j) && (mas[i]%mas[j] == 0) {
				fl = true
				for d := i; d < j; d++ {
					if (d != i) && (mas[d] != 0) && (mas[i] != 0) && (mas[j] != 0) && (mas[i]%mas[d] == 0) && (mas[d]%mas[j] == 0) {
						fl = false
						break
					}
				}
				if fl {
					fmt.Println("\t", mas[i], "--", mas[j])
				}
			}
		}
	}
	fmt.Println("}")
}
