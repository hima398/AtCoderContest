package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func DeleteS(s string, idx int) string {
	var result string
	result = s[:idx]
	result += s[idx+1:]
	return result
}

func CreateRoute(s string) []string {
	var ss []string
	if len(s) == 1 {
		return []string{s + "1"}
	}
	for i := 0; i < len(s); i++ {
		sss := CreateRoute(DeleteS(s, i))
		for j := 0; j < len(sss); j++ {
			sss[j] = string(s[i]) + sss[j]
		}
		ss = append(ss, sss...)
	}
	return ss
}

func Find(s []int, t int) bool {
	for _, v := range s {
		if v == t {
			return true
		}
	}
	return false
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = 1 << 60
	n := nextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = nextInt()
		}
	}
	m := nextInt()
	e := make(map[int][]int)
	for i := 0; i < m; i++ {
		x, y := nextInt(), nextInt()
		x--
		y--
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}
	//fmt.Println(e)
	var F func(in, used []int)
	var pattern [][]int
	F = func(in, used []int) {
		if len(in) == 0 {
			pattern = append(pattern, used)
		}
		for i, v := range in {
			nu := make([]int, len(used))
			copy(nu, used)
			nu = append(nu, v)
			nin := make([]int, len(in))
			copy(nin, in)
			nin = append(nin[:i], nin[i+1:]...)
			F(nin, nu)
		}
	}
	var ns []int
	for i := 0; i < n; i++ {
		ns = append(ns, i)
	}
	F(ns, []int{})
	ans := INF
	//fmt.Println(pattern)
	for _, v := range pattern {
		pre := v[0]
		time := a[pre][0]
		ok := true
		for j := 1; j < n; j++ {
			if Find(e[pre], v[j]) {
				ok = false
				break
			}
			time += a[v[j]][j]
			pre = v[j]
		}
		if ok {
			ans = Min(ans, time)
			//fmt.Println(ans, v)
		}
	}
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

func Permutation(N, K int) int {
	v := 1
	if 0 < K && K <= N {
		for i := 0; i < K; i++ {
			v *= (N - i)
		}
	} else if K > N {
		v = 0
	}
	return v
}

func Factional(N int) int {
	return Permutation(N, N-1)
}

func Combination(N, K int) int {
	if K == 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}
