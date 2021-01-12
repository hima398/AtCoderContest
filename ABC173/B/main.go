package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		s := nextString()
		m[s]++
	}
	fmt.Printf("AC x %d\n", m["AC"])
	fmt.Printf("WA x %d\n", m["WA"])
	fmt.Printf("TLE x %d\n", m["TLE"])
	fmt.Printf("RE x %d\n", m["RE"])
}
