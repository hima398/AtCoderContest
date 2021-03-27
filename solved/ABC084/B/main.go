package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	a, b := nextInt(), nextInt()
	s := nextString()
	ss := strings.Split(s, "-")
	if len(ss[0]) != a || len(ss[1]) != b {
		fmt.Println("No")
		return
	}
	_, err1 := strconv.Atoi(ss[0])
	if err1 != nil {
		fmt.Println("No")
		return
	}
	_, err2 := strconv.Atoi(ss[1])
	if err2 != nil {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
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
