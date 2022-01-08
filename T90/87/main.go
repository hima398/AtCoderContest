package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintField(a [][]int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func SolveHonestly(n, p, k int, a [][]int) string {
	ans := 0
	for x := 1; x <= p; x++ {
		ca := make([][]int, n)
		for i := 0; i < n; i++ {
			ca[i] = make([]int, n)
			copy(ca[i], a[i])
			for j := 0; j < n; j++ {
				if ca[i][j] == -1 {
					ca[i][j] = x
				}
			}
		}
		//O(n**3) ~ 1000くらいなのでもう少し他のループが回せるはず。
		for k := 0; k < n; k++ {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					ca[i][j] = Min(ca[i][j], ca[i][k]+ca[k][j])
				}
			}
		}
		//PrintField(ca)
		//fmt.Println()
		cnt := 0
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				if ca[i][j] <= p {
					cnt++
				}
			}
		}
		if cnt == k {
			ans++
		}
	}
	if ans == 0 {
		return "Infinity"
	} else {
		return strconv.Itoa(ans)
	}
}

//-1をxに置き換えてPスヌーク以下で到達可能なi, jの組みの数を数える
func CountPairs(x, p int, a [][]int) (ret int) {
	n := len(a)
	ca := make([][]int, n)
	for i := 0; i < n; i++ {
		ca[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if a[i][j] == -1 {
				ca[i][j] = x
			} else {
				ca[i][j] = a[i][j]
			}
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				ca[i][j] = Min(ca[i][j], ca[i][k]+ca[k][j])
			}
		}
	}
	//PrintField(ca)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if ca[i][j] <= p {
				ret++
			}
		}
	}

	return ret
}

func Search(n, p, k int, a [][]int) (ret int) {
	l, r := 1, 5*int(1e9)
	ret = 5 * int(1e9)
	for i := 0; i < 40; i++ {
		//for r-l > 1 {
		mid := (l + r) / 2
		cnt := CountPairs(mid, p, a)
		//fmt.Println(ok, ng, mid, cnt, k)
		if cnt <= k {
			r = mid
			ret = Min(ret, mid)
		} else {
			l = mid
		}
	}
	return ret
}

func Solve(n, p, k int, a [][]int) string {

	l := Search(n, p, k, a)
	r := Search(n, p, k-1, a)
	//fmt.Println(l, r)
	if r-l >= 2*int(1e9) {
		return "Infinity"
	} else {
		return strconv.Itoa(r - l)
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, p, k := nextInt(), nextInt(), nextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = nextInt()
		}
	}

	//ans := SolveHonestly(n, p, k, a)
	ans := Solve(n, p, k, a)
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

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
