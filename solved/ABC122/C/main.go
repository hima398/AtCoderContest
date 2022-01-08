package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	s := nextString()

	//0-iまでに出現するACの数
	m := make([]int, n)
	for i := 1; i < n; i++ {
		m[i] = m[i-1]
		if string(s[i-1])+string(s[i]) == "AC" {
			m[i]++
		}
	}
	//fmt.Println(m)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		l, r := nextInt(), nextInt()
		l--
		r--
		fmt.Fprintln(out, m[r]-m[l])
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
