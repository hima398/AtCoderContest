package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, x := nextInt(), nextInt(), nextInt()
	c := make([]int, n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
	}
	sc := 0
	for i := 0; i < n; i++ {
		c[i] = nextInt()
		sc += c[i]
		for j := 0; j < m; j++ {
			a[i][j] = nextInt()
		}
	}
	ans := sc + 1
	max := int(math.Pow(2.0, float64(n))) - 1
	//fmt.Printf("%b %d\n", max, ans)
	for i := 1; i <= max; i++ {
		tc := 0
		skill := make([]int, m)
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				tc += c[j]
				for k := 0; k < m; k++ {
					skill[k] += a[j][k]
				}
			}
		}
		//fmt.Println(skill)
		ok := true
		for k := 0; k < m; k++ {
			ok = ok && (skill[k] >= x)
		}
		if ok {
			ans = Min(ans, tc)
		}

	}
	if ans == sc+1 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
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
