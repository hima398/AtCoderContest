package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstTry(n, m int) {
	s := make([]int64, n)
	res := make(map[int64]int)
	ires := make(map[int64]int)
	mask := int64(1<<m - 1)
	for i := 0; i < n; i++ {
		ts := nextString()
		si, _ := strconv.ParseInt(ts, 2, 64)
		isi := si ^ mask
		s[i] = si
		res[si]++
		ires[isi]++
	}
	fmt.Println(res)
	fmt.Println(ires)
	ans := 0
	for i := range s {
		if res[s[i]] == 1 {
			is := s[i] ^ mask
			if res[is] == 0 {
				fmt.Println(i)
				ans += (n - 1 - i)
			}
		}
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, _ := nextInt(), nextInt()
	s := make([]int64, n+1)
	even := make([]int, n+1)
	odd := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ts := nextString()
		si, _ := strconv.ParseInt(ts, 2, 64)
		cb := bits.OnesCount64(uint64(si))
		s[i] = si
		if cb%2 == 0 {
			even[i] = even[i-1] + 1
			odd[i] = odd[i-1]
		} else {
			even[i] = even[i-1]
			odd[i] = odd[i-1] + 1
		}
	}
	//fmt.Println(even)
	//fmt.Println(odd)
	ans := 0
	for i := 1; i <= n-1; i++ {
		cb := bits.OnesCount64(uint64(s[i]))
		if cb%2 == 0 {
			ans += odd[n] - odd[i]
		} else {
			ans += even[n] - even[i]
		}
	}
	fmt.Println(ans)
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
