package main

import (
	"bufio"
	"fmt"
	"os"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n := nextString()
	idx := len(n) - 1
	for idx >= 0 && n[idx] == '0' {
		idx--
	}
	if idx < 0 {
		fmt.Println("Yes")
		return
	}
	nn := n[:idx+1]
	//fmt.Println(nn)
	ok := true
	for i := 0; i <= len(nn)/2; i++ {
		ok = ok && nn[i] == nn[len(nn)-1-i]
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
