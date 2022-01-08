package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ComputeDivisor(n int) []int {
	var ret []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			if i != n/i {
				ret = append(ret, n/i)
			}
		}
	}
	sort.Ints(ret)
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	k := nextInt()
	if k == 1 {
		fmt.Println(1)
		return
	}
	ds := ComputeDivisor(k)
	//fmt.Println(ds)
	ans := 0
	for i := 0; i < len(ds); i++ {
		for j := i; j < len(ds); j++ {
			a, b := ds[i], ds[j]
			if k/a < b {
				continue
			}
			if k%(a*b) != 0 {
				continue
			}
			c := k / (a * b)
			if b <= c {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
