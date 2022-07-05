package main

import "fmt"

func main() {
	var kin, kout, n, i, j int
	fmt.Scan(&kin)
	X := make([]string, kin)
	for i = 0; i < kin; i++ {
		fmt.Scan(&X[i])
	}
	fmt.Scan(&kout)
	Y := make([]string, kout)
	for i = 0; i < kout; i++ {
		fmt.Scan(&Y[i])
	}
	fmt.Scan(&n)
	D := make([][]int, n)
	for i = 0; i < n; i++ {
		D[i] = make([]int, kin)
		for j = 0; j < kin; j++ {
			fmt.Scan(&D[i][j])
		}
	}
	F := make([][]string, n)
	for i = 0; i < n; i++ {
		F[i] = make([]string, kin)
		for j = 0; j < kin; j++ {
			fmt.Scan(&F[i][j])
		}
	}
	type pair struct {
		ind  int
		alph string
	}
	vertex := make([]pair, n*kin)
	lenfreq := 0
	for i = 0; i < n; i++ {
		for j = 0; j < kin; j++ {
			repeat := 0
			for k := 0; k < lenfreq; k++ {
				if vertex[k].ind == D[i][j] && vertex[k].alph == F[i][j] {
					repeat++
				}
			}
			if repeat == 0 {
				vertex[lenfreq].ind = D[i][j]
				vertex[lenfreq].alph = F[i][j]
				lenfreq++
			}
		}
	}

	fmt.Printf("digraph {\n")
	fmt.Printf("\trankdir = LR\n")
	t := vertex[0].ind
	for i = 0; i < lenfreq; i++ {
		fmt.Printf("\t%d [label = \"(%d,%s)\"]\n", i, vertex[i].ind, vertex[i].alph)
		for k := 0; k < kin; k++ {
			t = t % n
			for x := 0; x < len(vertex); x++ {
				if vertex[x].ind == D[t][k] && vertex[x].alph == F[t][k] {
					fmt.Printf("\t%d -> %d [label = \"%s\"]\n", i, x, X[k])
					break
				}
			}
		}
		if i+1 < lenfreq && vertex[i].ind != vertex[i+1].ind {
			t++
		}
	}
	fmt.Printf("}")
}
