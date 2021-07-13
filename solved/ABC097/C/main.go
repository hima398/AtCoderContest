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
	m := make(map[string]bool)
	for i := 1; i < 6; i++ {
		for j := 0; j < len(s)-i+1; j++ {
			m[s[j:j+i]] = true
		}
	}
	var ps []string
	for k := range m {
		ps = append(ps, k)
	}
	sort.Strings(ps)
	//fmt.Println(k, ps)
	fmt.Println(ps[k-1])
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
