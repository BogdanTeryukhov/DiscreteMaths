package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
// DataSum - Функция, выполняющая разбор переданного скобочного подвыражения
func DataSum(ch string) int {
	count := 0
	countNumbers := 0
	for range ch {
		count++
	}
	mas := make([]int, count)
	i := 0
	for _, b := range ch {
		if b > 47 && b < 58 {
			mas[i] = int(b) - 48
			//fmt.Println(mas[i])
			countNumbers++
			i++
		}
	}
	expr := 1
	for _, b := range ch {
		switch string(b) {
		case "*":
			for check := 0; check < len(mas); check++ {
				if mas[check] > 0 && mas[check] < 10 {
					expr *= mas[check]
				}
			}
			return expr
		case "+":
			for check := 0; check < len(mas); check++ {
				if mas[check] > 0 && mas[check] < 10 {
					expr += mas[check]
				}
			}
			return expr - 1
		case "-":
			iter := 0
			for check := 0; check < len(mas); check++ {
				if mas[check] > 0 && mas[check] < 10 {
					if iter == 0 {
						expr = mas[check]
					} else {
						expr -= mas[check]
					}
					//fmt.Println(expr, " ", mas[check])
					iter++
				}
			}
			//fmt.Println(expr)
			return expr
		}
	}
	return 0
}

func main() {
	for {
		var (
			F string
		)
		//F = "(* 2 (* 2 2 (- 5 3)(* 5 (- 4 2))(+ 5 1) 3))"
		//fmt.Scan(&F)
		myscanner := bufio.NewScanner(os.Stdin)
		myscanner.Scan()
		F = myscanner.Text()

		var (
			head     int
			tail     int
			lastRep  int
			i        int
			iterator int
		)
		iterator = 0
		mas := make([]int, len(F))
		masOfOperands := make([]int, len(F))
		p := 0

		//Заполнение массива операций
		flOfOperateons := false
		for _, b := range F {
			if string(b) == "*" || string(b) == "-" || string(b) == "+" {
				masOfOperands[p] = int(b)
				p++
				flOfOperateons = true
			}
		}
		if !flOfOperateons {
			fmt.Println(F)
		} else {
			//Полный разбор выражения
			for F != "" {
				head = 0
				tail = 0
				count := 0
				//Динамически высчитываем длину оставшегося массива
				for range F {
					count++
				}
				//Основной цикл для проверки скобок
				for ind, b := range F {
					if b == 40 {
						head = ind
					} else if b == 41 {
						tail = ind
						lastRep = DataSum(F[head+1 : tail])
						//fmt.Println(lastRep)
						mas[i] = lastRep
						i++
						F = F[0:head] + F[tail+1:count]
						//fmt.Println(F[head : tail+1])
						break
					}
				}
				iterator++
			}
			totalAnswer := 1

			//Финальное высчитывание выражения
			switch masOfOperands[0] {
			case 42:
				for h := 0; h < iterator; h++ {
					totalAnswer *= mas[h]
				}
				fmt.Println(totalAnswer)
			case 43:
				for h := 0; h < iterator; h++ {
					totalAnswer += mas[h]
				}
				fmt.Println(totalAnswer - 1)
			case 45:
				iter := 0
				for h := 0; h < iterator; h++ {
					if iter == 0 {
						totalAnswer = mas[h]
					} else {
						totalAnswer -= mas[h]
					}
					iter++
					//fmt.Println(mas[h])
				}
				fmt.Println(totalAnswer)
			}
		}
	}
}*/

//Рекурсия oneLove.....
var i int

func TotalCounter(F string) int {
	i++

	if i >= len(F) {
		return 0
	}

	if F[i] == '(' || F[i] == ')' || F[i] == ' ' {
		return TotalCounter(F)
	}

	switch F[i] {
	case '+':
		return TotalCounter(F) + TotalCounter(F)
	case '-':
		return TotalCounter(F) - TotalCounter(F)
	case '*':
		return TotalCounter(F) * TotalCounter(F)
	}

	return int(F[i] - 48)
}

func main() {
	var (
		F string
		//sum  int
	)
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	F = myscanner.Text()

	if len(F) == 1 {
		fmt.Println(F)
	} else {
		fmt.Printf("%d\n", TotalCounter(F))
	}
}
