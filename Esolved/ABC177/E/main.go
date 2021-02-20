package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var cnts [1000001]int

func SieveOfEratosthenes(n int) []int {
	ret := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 2 * i; j <= n; j += i {
			if ret[j] == 0 {
				ret[j] = i
			}
		}
	}
	return ret
}

func SliceGcd(s []int) int {
	if len(s) < 2 {
		return -1
	}
	gcd := Gcd(s[0], s[1])
	for i := 2; i < len(s); i++ {
		gcd = Gcd(gcd, s[i])
	}
	return gcd
}

func Solve() string {
	n := nextInt()
	a := make([]int, n+1)
	maxa := 0
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
		maxa = Max(maxa, a[i])
	}

	// エラトステネスの篩(Sieve of Eratosthenes)
	soe := SieveOfEratosthenes(maxa)
	//fmt.Println(a)
	//fmt.Println(soe)

	m := make(map[int][]int)
	// 素因数分解
	for i := 1; i <= n; i++ {
		v := a[i]
		for {
			if soe[v] == 0 {
				m[v] = append(m[v], i)
				break
			}
			m[soe[v]] = append(m[soe[v]], i)
			v = v / soe[v]
		}
	}
	IsPC := true // is pairwise coprime
	for k, v := range m {
		//fmt.Println(k, v)
		if k == 1 {
			continue
		}
		if len(v) > 1 {
			pre := v[0]
			for i := 1; i < len(v); i++ {
				if pre != v[i] {
					IsPC = false
					break
				}
				pre = v[i]
			}
		}
	}
	if IsPC {
		return "pairwise coprime"
	}

	IsSC := SliceGcd(a[1:]) == 1 // is setwise coprime
	if IsSC {
		return "setwise coprime"
	} else {
		return "not coprime"
	}
}

func SolveCommentary() string {
	n := nextInt()
	a := make([]int, n)

	//ma := make(map[int]int)
	maxa := 0
	a[0] = nextInt()
	gcd := a[0]
	cnts[a[0]]++
	for i := 1; i < n; i++ {
		a[i] = nextInt()
		gcd = Gcd(gcd, a[i])
		maxa = Max(maxa, a[i])
		cnts[a[i]]++
	}
	if gcd > 1 {
		return "not coprime"
	}
	// エラトステネスの篩(Sieve of Eratosthenes)
	for i := 2; i <= maxa; i++ {
		total := 0
		for j := i; j <= maxa; j += i {
			total += cnts[j]
		}

		if total >= 2 {
			return "setwise coprime"
		}
	}

	return "pairwise coprime"
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	fmt.Println(SolveCommentary())
	//fmt.Println(Solve())

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
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
