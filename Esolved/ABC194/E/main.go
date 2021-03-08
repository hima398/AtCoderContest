package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstTry(n, m int) int {
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}

	mex := 0
	d := make(map[int]int)
	for i := 1; i < 1+m; i++ {
		d[a[i]]++
	}
	//fmt.Println(d)
	for i := 0; i < 1500001; i++ {
		if d[i] == 0 {
			mex = i
			break
		}
	}
	//fmt.Println(mex)
	ans := mex
	for l := 1; l <= n-m+1; l++ {
		r := l + m - 1
		if d[a[l]] == 1 && a[l] != a[r] {
			mex = d[l]
		} else if a[r] == mex {
			mex++
		}
		ans = Min(ans, mex)
		d[a[r]]++
		d[a[l]]--
	}
	return ans
}

func SolveCommentary(n, m int, a []int) int {
	p := make(map[int][]int)
	max := 0
	for i := 0; i < n; i++ {
		max = Max(max, a[i])
		p[a[i]] = append(p[a[i]], i)
	}
	//fmt.Println(p)
	// Ai < Nなのでi < nでも良いはず
	for i := 0; i <= max+1; i++ {
		pre := -1
		p[i] = append(p[i], n)
		for _, v := range p[i] {
			if v-pre > m {
				return i
			}
			pre = v
		}
	}
	return n
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	fmt.Println(SolveCommentary(n, m, a))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	if x < y {
		x, y = y, x
	}
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}
