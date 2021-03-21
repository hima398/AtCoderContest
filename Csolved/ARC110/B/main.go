package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	const L = 10000000000
	n := nextInt()
	t := nextString()
	//fmt.Println(t)
	// t = "0" or "1"
	if n == 1 {
		if t == "1" {
			fmt.Println(2 * L)
			return
		} else {
			fmt.Println(L)
			return
		}
	}
	if strings.HasPrefix(t, "10") {
		t = "1" + t
	} else if strings.HasPrefix(t, "01") {
		t = "11" + t
	}
	if strings.HasSuffix(t, "11") {
		t += "0"
	} else if strings.HasSuffix(t, "01") {
		t += "10"
	}
	//fmt.Println(t)
	if t == strings.Repeat("110", len(t)/3) {
		fmt.Println(L - len(t)/3 + 1)
	} else {
		fmt.Println(0)
	}
}
