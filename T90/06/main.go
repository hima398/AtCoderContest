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

	n, k := nextInt(), nextInt()
	s := nextString()

	const INF = int(1e5) + 1
	const na = 26 //小文字の種類
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, na)
		for j := 0; j < na; j++ {
			m[i][j] = INF
		}
	}
	m[n-1][int(s[n-1]-'a')] = n - 1
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < na; j++ {
			if j == int(s[i]-'a') {
				m[i][j] = i
			} else {
				m[i][j] = m[i+1][j]
			}
		}
	}

	ans := ""
	rem := k
	for i := 0; i < n; i++ {
		for j := 0; j < na; j++ {
			if n-m[i][j] > rem {
				ans += string(byte(j) + 'a')
				i = m[i][j]
				rem--
				break
			}
		}
		if rem == 0 {
			break
		}
	}
	fmt.Println(ans)
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
