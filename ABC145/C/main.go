package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	x, y := make([]float64, n), make([]float64, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextFloat64(), nextFloat64()
	}
	var is []int
	for i := 0; i < n; i++ {
		is = append(is, i)
	}
	ans := 0.0
	for {
		s := 0.0
		for i := 1; i < len(is); i++ {
			s += math.Sqrt((x[is[i]]-x[is[i-1]])*(x[is[i]]-x[is[i-1]]) + (y[is[i]]-y[is[i-1]])*(y[is[i]]-y[is[i-1]]))
		}
		ans += s
		if !NextPermutation(sort.IntSlice(is)) {
			break
		}
	}
	fn := float64(Factional(n))
	ans /= fn
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
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
