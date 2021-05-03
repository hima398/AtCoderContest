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

	n := nextInt()
	s := nextString()
	i := 0
	t := "b"
	for len(t) <= n {
		if t == s {
			fmt.Println(i)
			return
		}
		i++
		switch i % 3 {
		case 0:
			t = "b" + t + "b"
		case 1:
			t = "a" + t + "c"
		case 2:
			t = "c" + t + "a"
		default:
		}
	}
	fmt.Println(-1)
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
