package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextUint() uint64 {
	sc.Scan()
	u, _ := strconv.ParseUint(sc.Text(), 10, 64)
	return u
}

func Pow(x, y int) int {
	ret := 1
	for i := 1; i <= y; i++ {
		ret *= x
	}
	return ret
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextUint()
	k := Pow(2, int(n)) - 1
	max := Pow(2, int(n))
	fmt.Println(k)
	for i := 1; i <= k; i++ {
		s := ""
		for j := 1; j <= max; j++ {
			if bits.OnesCount64(uint64(i&j))%2 == 1 {
				s += "A"
			} else {
				s += "B"
			}
		}
		fmt.Println(s)
	}

}
