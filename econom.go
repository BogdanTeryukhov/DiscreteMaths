package main

import (
	"bufio"
	"fmt"
	"os"
)

var arr [][2]int
var helpArr []int

func ParseExpression(i int, F string) {
	l := helpArr[len(helpArr)-1]
	helpArr = helpArr[:len(helpArr)-1]
	for _, pair := range arr {
		if F[l:i] == F[pair[0]:pair[1]] {
			return
		}
	}
	arr = append(arr, [2]int{l, i})
}

func main() {
	var F string
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	F = myscanner.Text()

	for i, symb := range F {
		if symb == '(' {
			helpArr = append(helpArr, i)
		} else if symb == ')' {
			ParseExpression(i, F)
		}
	}
	fmt.Println(len(arr))
}
