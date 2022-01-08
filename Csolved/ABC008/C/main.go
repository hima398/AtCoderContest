package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type coins []int

func (cs coins) Len() int {
	return len(cs)
}
func (cs coins) Less(i, j int) bool {
	return cs[i] <= cs[j]
}
func (cs coins) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

// n<=8の制約化で解く
func SolveHonestly(n int, c []int) float64 {

	sort.Ints(c)
	ret := 0
	for {
		// 初期化を省略するため 0:表 1:裏で扱う
		pattern := make([]int, n)
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				if c[j]%c[i] == 0 {
					//表裏反転
					pattern[j] ^= 1
				}
			}
		}
		s := 0
		for i := 0; i < n; i++ {
			s += (pattern[i] ^ 1)
		}
		ret += s
		if !NextPermutation(coins(c)) {
			break
		}
	}
	fmt.Println(ret, Factional(n))
	return float64(ret) / float64(Factional(n))
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	c := nextIntSlice(n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if c[i]%c[j] == 0 {
				s[i]++
			}
		}
	}
	ans := 0.0
	for i := 0; i < n; i++ {
		if s[i]%2 == 0 {
			ans += float64(s[i]/2+1) / float64(s[i]+1)
		} else {
			ans += 0.5
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
