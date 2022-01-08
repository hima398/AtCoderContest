package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

//愚直にやっているのでTLE
func FirstSolve() {
	n, k := nextInt(), nextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = nextInt()
		}
	}
	nn := n - k + 1
	b := make([][]int, nn)
	for i := 0; i < nn; i++ {
		b[i] = make([]int, nn)
	}
	s := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		s[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

		}
	}
	for i := 0; i < nn; i++ {
		for j := 0; j < nn; j++ {
			var s []int
			for ii := 0; ii < k; ii++ {
				for jj := 0; jj < k; jj++ {
					s = append(s, a[i+ii][j+jj])
				}
			}
			sort.Ints(s)
			//fmt.Println(s)
			b[i][j] = s[(len(s)-1)/2]
		}
	}
	/*
		for i := 0; i < nn; i++ {
			fmt.Println(b[i])
		}
	*/
	ans := int(1e9) + 1
	for i := 0; i < nn; i++ {
		for j := 0; j < nn; j++ {
			ans = Min(ans, b[i][j])
		}
	}
	fmt.Println(ans)
}

func SolveBinary(n, k int) {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	maxA := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = nextInt()
			maxA = Max(maxA, a[i][j])
		}
	}

	check := func(c int) bool {
		//fmt.Printf("c = %d\n", c)
		s := make([][]int, n+1)
		for i := 0; i < n+1; i++ {
			s[i] = make([]int, n+1)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j]
				if a[i][j] > c {
					s[i+1][j+1]++
				}
			}
		}

		/*
			for i := 0; i < n+1; i++ {
				fmt.Println(s[i])
			}
		*/

		nn := n - k + 1
		b := make([][]int, nn)
		for i := 0; i < nn; i++ {
			b[i] = make([]int, nn)
		}
		for i := 0; i < nn; i++ {
			for j := 0; j < nn; j++ {
				b[i][j] = s[i+k][j+k] - s[i][j+k] - s[i+k][j] + s[i][j]
			}
		}
		/*
			for i := 0; i < nn; i++ {
				fmt.Println(b[i])
			}
		*/

		ok := false
		for i := 0; i < nn; i++ {
			for j := 0; j < nn; j++ {
				ok = ok || b[i][j] < (k*k/2)+1
			}
		}
		return ok
	}

	ng, ok := -1, maxA
	for ok-ng > 1 {
		c := (ok + ng) / 2
		if check(c) {
			ok = c
		} else {
			ng = c
		}
	}
	fmt.Println(ok)

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	SolveBinary(n, k)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
