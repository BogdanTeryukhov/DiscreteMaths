package main

import (
	"fmt"
	"sort"
	"strings"
)

type State struct {
	index       int
	Final, flag bool
}

type DetState struct {
	index int
	name  []int
	Final bool
}

func Det(X []string, D []map[string][]*State, q *State) (Q1 []*DetState, D1 []map[string]*DetState) {
	temp := make([]*State, 1)
	temp[0] = q
	q0 := Closure(temp, D)
	arr := make([]int, len(q0))
	for i, q := range q0 {
		arr[i] = q.index
	}
	sort.Ints(arr)
	q1 := &DetState{
		index: 0,
		name:  arr,
		Final: false,
	}
	Q1 = append(make([]*DetState, 0), q1)
	D1 = make([]map[string]*DetState, 0)
	D1 = append(D1, make(map[string]*DetState))

	MyStack := make([][]*State, 0)
	MyStack = append(MyStack, q0)
	ind := make([]int, 0)
	ind = append(ind, 0)

	for len(MyStack) != 0 {
		z := MyStack[len(MyStack)-1]
		MyStack = MyStack[0 : len(MyStack)-1]
		index := ind[len(ind)-1]
		ind = ind[0 : len(ind)-1]

		q2 := Q1[index]
		for _, u := range z {
			if u.Final {
				q2.Final = true
				break
			}
		}

		for _, a := range X {
			U := make([]*State, 0)
			for _, u := range z {
				for _, w := range D[u.index][a] {
					U = append(U, w)
				}
			}

			zPr := Closure(U, D)
			arr := make([]int, len(zPr))
			for i, q := range zPr {
				arr[i] = q.index
			}
			sort.Ints(arr)

			q3 := &DetState{
				name:  arr,
				index: -1,
				Final: false,
			}

			indTemp := -1
			for i, u := range Q1 {
				flag := true
				if len(q3.name) == len(u.name) {
					for j := 0; j < len(q3.name); j++ {
						if q3.name[j] != u.name[j] {
							flag = false
							break
						}
					}
					if flag {
						indTemp = i
					}
				}
			}

			if indTemp == -1 {
				indTemp = len(Q1)
				q3.index = indTemp

				Q1 = append(Q1, q3)
				D1 = append(D1, make(map[string]*DetState))
				MyStack = append(MyStack, zPr)
				ind = append(ind, indTemp)
			} else {
				q3 = Q1[indTemp]
			}
			D1[q2.index][a] = q3
		}
	}
	return Q1, D1
}

func Closure(z []*State, D []map[string][]*State) []*State {
	C := make([]*State, 0)
	for _, q := range z {
		Dfs(q, &C, D)
	}
	for _, q := range C {
		q.flag = false
	}
	return C
}

func Dfs(q *State, C *[]*State, D []map[string][]*State) {
	if !q.flag {
		q.flag = true
		*C = append(*C, q)
		for _, w := range D[q.index]["lambda"] {
			Dfs(w, C, D)
		}
	}
}

func main() {
	var n, m, temp1, temp2, q int
	fmt.Scan(&n, &m)

	D := make([]map[string][]*State, n)
	for i := 0; i < n; i++ {
		D[i] = make(map[string][]*State)
	}

	Q := make([]*State, n)
	for i := 0; i < n; i++ {
		Q[i] = &State{
			index: i,
			Final: false,
			flag:  false,
		}
	}

	freq := make(map[string]bool)
	X := make([]string, 0)
	var s string
	for i := 0; i < m; i++ {
		fmt.Scan(&temp1, &temp2, &s)
		if s != "lambda" && !freq[s] {
			X = append(X, s)
		}
		freq[s] = true
		if _, ok := D[temp1][s]; !ok {
			D[temp1][s] = make([]*State, 0)
		}
		D[temp1][s] = append(D[temp1][s], Q[temp2])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&q)
		if q == 1 {
			Q[i].Final = true
		}
	}

	fmt.Scan(&q)
	q0 := Q[q]
	Q1, D1 := Det(X, D, q0)

	ar := make([][][]string, len(Q1))
	for i := 0; i < len(Q1); i++ {
		ar[i] = make([][]string, len(Q1))
		for j := 0; j < len(Q1); j++ {
			ar[i][j] = make([]string, 0)
		}
	}
	for _, q := range Q1 {
		for _, a := range X {
			ar[q.index][D1[q.index][a].index] = append(ar[q.index][D1[q.index][a].index], a)
		}
	}
	fmt.Printf("digraph {\n")
	fmt.Printf("\trankdir = LR\n")
	fmt.Printf("\tdummy [label = \"\", shape = none]\n")
	for i, q := range Q1 {
		if !q.Final {
			fmt.Printf("\t%d [label = \"%v\", shape = circle]\n", i, q.name)
		} else {
			fmt.Printf("\t%d [label = \"%v\", shape = doublecircle]\n", i, q.name)
		}
	}
	fmt.Printf("\tdummy -> 0\n")
	for i, line := range ar {
		for j, str := range line {
			if len(str) != 0 {
				fmt.Printf("\t%d -> %d [label = \"%s\"]\n", i, j, strings.Join(str, ", "))
			}
		}
	}
	fmt.Printf("}")
}
