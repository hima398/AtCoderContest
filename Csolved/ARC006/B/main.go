package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	in := bufio.NewReader(os.Stdin)
	var n, l int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &l)
	in.ReadLine()
	a := make([]string, l)
	for i := 0; i < l; i++ {
		line, _, _ := in.ReadLine()
		a[i] = string(line)
	}
	/*
		for _, v := range a {
			fmt.Println(v)
		}
	*/
	line, _, _ := in.ReadLine()
	g := string(line)
	var now int
	for i := 0; i < 2*n-1; i += 2 {
		if g[i] == 'o' {
			now = i / 2
		}
	}
	for i := l - 1; i >= 0; i-- {

		lj, rj := 2*now-1, 2*now+1
		if lj >= 0 && a[i][lj] == '-' {
			now--
		} else if rj < 2*n-1 && a[i][rj] == '-' {
			now++
		}

	}
	fmt.Println(now + 1)
}
