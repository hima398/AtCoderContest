package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

//しゃくとり法で解いてみる
func SolveMeasuring(n int, a []int) (ans int) {
	l, r := 0, 0
	m := make(map[int]int)
	for {
		for r < n && m[a[r]] == 0 {
			m[a[r]]++
			//fmt.Println(l, r)
			if m[a[r]] <= 1 {
				ans = Max(ans, r-l+1)
				r++
			} else {
				//l-rまでの間に重複している味があるのでlを進めるためループを抜ける
				break
			}
		}
		// rがnを指していたらループ全体から抜ける
		if r >= n {
			break
		}
		tar := a[r]
		// rが指している味が左側からなくなるまでlを進める
		for m[tar] > 0 && l < r {
			//fmt.Println(m[tar], m)
			m[a[l]]--
			l++
		}
		// rがnを指していたらループ全体から抜ける
		if r >= n {
			break
		}
	}
	return ans
}

func SolveCommentary(n int, a []int) (ans int) {
	l := 0
	m := make(map[int]bool)
	for r := 0; r < n; r++ {
		tar := a[r]
		for m[tar] {
			m[a[l]] = false
			l++
		}
		m[tar] = true
		ans = Max(ans, r-l+1)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	//ans := SolveMeasuring(n, a)
	ans := SolveCommentary(n, a)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
