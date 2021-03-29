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

	s := nextString()
	var a [4]int
	for i := 0; i < len(s); i++ {
		a[i] = int(s[i] - '0')
	}
	bits := 0
	for bits < 1<<3 {
		sa := a[0]
		for i := 1; i < 4; i++ {
			if bits&(1<<(i-1)) > 0 {
				sa += a[i]
			} else {
				sa -= a[i]
			}
		}
		if sa == 7 {
			break
		}
		bits++
	}
	ans := fmt.Sprintf("%d", a[0])
	for i := 1; i < 4; i++ {
		if bits&(1<<(i-1)) > 0 {
			ans += "+" + fmt.Sprintf("%d", a[i])
		} else {
			ans += "-" + fmt.Sprintf("%d", a[i])
		}
	}
	ans += "=7"
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
