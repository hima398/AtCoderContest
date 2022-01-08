package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := "indeednow"
	mt := make(map[byte]int)
	for i := range t {
		mt[t[i]]++
	}

	n := nextInt()
	for i := 0; i < n; i++ {
		s := nextString()
		ms := make(map[byte]int)
		for i := range s {
			ms[s[i]]++
		}
		ok := true
		for k, v := range ms {
			ok = ok && mt[k] == v
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
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
