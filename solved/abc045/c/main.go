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

	s := nextString()
	n := len(s)
	bits := 1<<(n-1) - 1
	ans := 0
	for k := 0; k <= bits; k++ {
		ps := 0 // 部分的な和
		v := 0
		for i := 0; i < n; i++ {
			num := int(s[i] - '0')
			if v > 0 {
				v *= 10
			}
			v += num
			if k>>i&1 > 0 {
				ps += v
				v = 0
			}
		}
		if v > 0 {
			ps += v
		}
		ans += ps
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
