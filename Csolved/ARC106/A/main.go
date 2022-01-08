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

	for a := 1; a < 38; a++ {
		for b := 1; b < 26; b++ {
			a3, b5 := 1, 1
			for i := 0; i < a; i++ {
				a3 *= 3
			}
			for i := 0; i < b; i++ {
				b5 *= 5
			}
			if a3+b5 == n {
				fmt.Println(a, b)
				return
			}
		}
	}
	fmt.Println(-1)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}
