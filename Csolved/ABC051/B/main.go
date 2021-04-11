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

	k, s := nextInt(), nextInt()
	ans := 0
	for x := 0; x <= k; x++ {
		for y := x; y <= k; y++ {
			z := s - (x + y)
			if z >= x && z >= y && z >= 0 && z <= k {
				//fmt.Println(x, y, z)
				if x != y && x != z && y != z {
					ans += 6
				} else if x == y && x == z && y == z {
					ans++
				} else {
					ans += 3
				}
			}
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
