package main

import (
	"fmt"
)

type Graph struct {
	mark, equal bool
	Link        []int
	dest        int
}

func (g *Graph) initRootTop() {
	g.mark = true
	g.dest = 0
	g.equal = false
}

func (g *Graph) makeVisit(p *Graph) {
	g.mark = true
	g.equal = g.equal && (g.dest == -2 || g.dest == (*p).dest+1)
	g.dest = (*p).dest + 1
}

func (g *Graph) CheckMarkAfter() {
	if !g.mark {
		g.equal = false
	}
}

type Queue struct {
	data                   []int
	cap, Count, head, tail int
}

func (p *Queue) Enqueue(x int) {
	p.data[p.tail] = x
	p.tail++
	p.Count++
}

func (p *Queue) Dequeue() int {
	var x = p.data[p.head]
	p.head++
	p.Count--
	return x
}

func main() {
	var N, M int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)

	arr := make([]Graph, N)
	for i := 0; i < N; i++ {
		arr[i] = Graph{equal: true, dest: -2, Link: make([]int, 0)}
	}

	for i := 0; i < M; i++ {
		var u, v int
		_, err := fmt.Scanf("%d %d", &u, &v)
		if err != nil {
			return
		}
		arr[u].Link = append(arr[u].Link, v)
		arr[v].Link = append(arr[v].Link, u)
	}
	var K int
	fmt.Scanf("%d", &K)
	for i := 0; i < K; i++ {
		var top int
		fmt.Scanf("%d", &top)
		BFS(&arr, top)
	}
	PrintAnswer(&arr, N)
}

func BFS(graph *[]Graph, top int) {
	for i := 0; i < len(*graph); i++ {
		if i == top {
			continue
		}
		(*graph)[i].mark = false
	}
	(*graph)[top].initRootTop()
	var Q = Queue{data: make([]int, len(*graph))}
	Q.Enqueue(top)
	for Q.Count > 0 {
		var v = Q.Dequeue()
		for i := 0; i < len((*graph)[v].Link); i++ {
			var u = (*graph)[v].Link[i]
			if !(*graph)[u].mark {
				(*graph)[u].makeVisit(&(*graph)[v])
				Q.Enqueue(u)
			}
		}
	}
	for i := 0; i < len(*graph); i++ {
		(*graph)[i].CheckMarkAfter()
	}
}
func PrintAnswer(graph *[]Graph, N int) {
	var ans = false
	for i := 0; i < N; i++ {
		if (*graph)[i].equal {
			fmt.Print(i)
			fmt.Print(" ")
			ans = true
		}
	}
	if !ans {
		fmt.Println("-")
	}
}
