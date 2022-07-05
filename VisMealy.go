package main

/*
import (
	"fmt"
)

type MiliAutomate struct {
	D       [][]int
	F       [][]string
	n, m, q int
}

func (a *MiliAutomate) initialization() {

	a.D = make([][]int, a.n)
	a.F = make([][]string, a.n)

	for i := 0; i < a.n; i++ {
		a.D[i] = make([]int, a.m)
		a.F[i] = make([]string, a.m)
	}
}

func (a *MiliAutomate) creation() {
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			fmt.Scan(&a.D[i][j])
		}
	}

	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			fmt.Scan(&a.F[i][j])
		}
	}
}

func (a *MiliAutomate) print() {
	fmt.Println("digraph {\n\trankdir = LR\t")

	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			fmt.Printf(
				"\t%d -> %d  [label = \"%s(%s)\"]\n",
				i,
				a.D[i][j],
				string(rune(97+j)),
				a.F[i][j],
			)
		}
	}
	fmt.Print("}\n")
}

func main() {

	var n, m, q int

	fmt.Scan(&n, &m, &q)

	var X = MiliAutomate{n: n, m: m, q: q}

	X.initialization()

	X.creation()

	X.print()
}*/
