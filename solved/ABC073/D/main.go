package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintEdge(e [][]int) {
	for i := 0; i < len(e); i++ {
		fmt.Println(e[i])
	}
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = 1 << 60 //int(1e5) + 1

	n, m, r := nextInt(), nextInt(), nextInt()
	sr := nextIntSlice(r)
	//0-indexにする
	for i := range sr {
		sr[i]--
	}
	e := make([][]int, n)
	for i := 0; i < n; i++ {
		e[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				e[i][j] = 0
				continue
			}
			e[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		a--
		b--
		e[a][b] = c
		e[b][a] = c
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				e[i][j] = Min(e[i][j], e[i][k]+e[k][j])
			}
		}
	}
	sort.Ints(sr)
	//PrintEdge(e)
	ans := INF
	for {
		//fmt.Println(sr)
		cost := 0
		for i := 0; i < r-1; i++ {
			cost += e[sr[i]][sr[i+1]]
		}
		ans = Min(ans, cost)
		if !NextPermutation(sort.IntSlice(sr)) {
			break
		}
	}

	fmt.Println(ans)

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
