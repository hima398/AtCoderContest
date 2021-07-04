package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = int(1e9) + 7
	x, y := nextInt(), nextInt()
	if (x+y)%3 != 0 {
		fmt.Println(0)
		return
	}
	m := (2*x - y) / 3
	n := (2*y - x) / 3
	if m < 0 || n < 0 {
		fmt.Println(0)
		return
	}
	mc := NewModCombination()
	mc.Init(1000001, p)
	ans := mc.Comb(m+n, n)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type ModCombination struct {
	fac, finv, inv []int
	p              int
}

func NewModCombination() *ModCombination {
	return new(ModCombination)
}

func (mc *ModCombination) Init(n, p int) {
	mc.fac = make([]int, n)
	mc.finv = make([]int, n)
	mc.inv = make([]int, n)
	mc.p = p
	mc.fac[0], mc.fac[1] = 1, 1
	mc.finv[0], mc.finv[1] = 1, 1
	mc.inv[1] = 1
	for i := 2; i < n; i++ {
		mc.fac[i] = mc.fac[i-1] * i % p
		mc.inv[i] = p - mc.inv[p%i]*(p/i)%p
		mc.finv[i] = mc.finv[i-1] * mc.inv[i] % p
	}
}

func (mc *ModCombination) Comb(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return mc.fac[n] * (mc.finv[k] * mc.finv[n-k] % mc.p) % mc.p
}
