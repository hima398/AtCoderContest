package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	s := nextString()
	k := nextInt()
	//var ss []string
	ms := make(map[string]int)
	for i := 0; i < len(s); i++ {
		for j := 1; j <= Min(len(s), 5); j++ {
			ms[s[i:Min(i+j, len(s))]]++
		}
	}
	//fmt.Println(ms)
	var ss []string
	for key := range ms {
		ss = append(ss, key)
	}
	sort.Strings(ss)
	fmt.Println(ss[k-1])
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
