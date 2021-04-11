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

	const p = int(1e9) + 7

	n := nextInt()
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		a := nextInt()
		m[a]++
	}
	if n%2 == 1 {
		m[0]++
	}
	//fmt.Println(m)
	if n%2 == 1 {
		for k, v := range m {
			if k%2 == 1 {
				fmt.Println(0)
				return
			}
			if v != 2 {
				fmt.Println(0)
				return
			}
		}
	} else {
		for k, v := range m {
			if k%2 == 0 {
				fmt.Println(0)
				return
			}
			if v != 2 {
				fmt.Println(0)
				return
			}
		}
	}
	ans := 1
	if n%2 == 0 {
		for i := 0; i < len(m); i++ {
			ans *= 2
			ans %= p
		}
	} else {
		for i := 0; i < len(m)-1; i++ {
			ans *= 2
			ans %= p
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
