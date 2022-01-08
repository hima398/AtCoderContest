package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

func SolveHonestly(n, pc int) int {
	a, b, c := make([]int, n+1), make([]int, n+1), make([]int, n+1)

	start := 9223372036854775807
	end := 0
	for i := 1; i <= n; i++ {
		a[i], b[i], c[i] = nextInt(), nextInt(), nextInt()
		start = Min(start, a[i])
		end = Max(end, b[i])
	}

	// おそらくこの容量が1GBを超えるので確保できない
	imos := make([]int, end+2)
	var mem runtime.MemStats
	// debug
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
	for i := 1; i <= n; i++ {
		//fmt.Println(i)
		imos[a[i]] += c[i]
		imos[b[i]+1] -= c[i]
	}
	ans := 0
	// ここを高速化する必要もあり
	for i := start; i <= end; i++ {
		imos[i] += imos[i-1]
		ans += Min(imos[i], pc)
	}
	return ans
}

type Event struct {
	start  int
	amount int
}

func PrintEvents(e []Event) {
	for _, v := range e {
		fmt.Println(v)
	}
}

func Solve(n, pc int) int {
	var e []Event

	for i := 0; i < n; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		e = append(e, Event{a, c})
		e = append(e, Event{b + 1, -c})
	}
	//PrintEvents(e)
	sort.Slice(e, func(i, j int) bool {
		return e[i].start < e[j].start
	})

	t := 0
	amount := 0
	ans := 0
	for _, v := range e {
		if t != v.start {
			period := v.start - t
			ans += Min(amount, pc) * period
			t = v.start
		}
		amount += v.amount
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, pc := nextInt(), nextInt()
	fmt.Println(Solve(n, pc))

}
