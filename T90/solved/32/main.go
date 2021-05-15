package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = 1 << 60
	n := nextInt()
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = nextInt()
		}
	}
	m := nextInt()
	e := make([][]bool, n)
	for i := 0; i < n; i++ {
		e[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		x, y := nextInt(), nextInt()
		x--
		y--
		e[x][y] = true
		e[y][x] = true
	}
	//fmt.Println(e)
	ans := INF
	for i := 0; ; i++ {
		var ng bool
		sum := 0
		for i := range p {
			if i < n-1 {
				ng = e[p[i]][p[i+1]]
				if ng {
					break
				}
			}
			sum += a[p[i]][i]
		}
		// 順番として作れるものだけ答えを更新
		if !ng {
			ans = Min(ans, sum)
		}
		if !NextPermutation(sort.IntSlice(p)) {
			break
		}
	}

	//fmt.Println(pattern)
	if ans == INF {
		ans = -1
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
