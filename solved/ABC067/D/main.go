package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Bfs(r, n int, e [][]int) (ret map[int][]int) {
	type Node struct {
		i, d int
	}
	var q []Node
	ret = make(map[int][]int, n)
	v := make([]bool, n)
	q = append(q, Node{r, 0})
	v[r] = true
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		ret[node.d] = append(ret[node.d], node.i)
		for _, edge := range e[node.i] {
			if !v[edge] {
				q = append(q, Node{edge, node.d + 1})
				v[edge] = true
			}
		}
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}
	colors := make([]int, n)
	//1:黒、2:白
	colors[0], colors[n-1] = 1, 2
	fd, sd := Bfs(0, n, e), Bfs(n-1, n, e)
	for j := 1; j < Max(len(fd), len(sd)); j++ {
		for _, idx := range fd[j] {
			if colors[idx] == 0 {
				colors[idx] = 1
			}
		}
		for _, idx := range sd[j] {
			if colors[idx] == 0 {
				colors[idx] = 2
			}
		}
	}
	//フェネックくん、すぬけくんがそれぞれ最適に動いた場合に
	//色を塗れる可能性があるノード数
	var nf, ns int
	for i := 0; i < n; i++ {
		switch colors[i] {
		case 1:
			nf++
		case 2:
			ns++
		}
	}
	if nf > ns {
		fmt.Println("Fennec")
	} else {
		// nf <= ns
		fmt.Println("Snuke")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
