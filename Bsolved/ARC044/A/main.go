package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	if n == 1 {
		fmt.Println("Not Prime")
		return
	}
	if IsPrime(n) {
		fmt.Println("Prime")
		return
	}
	n1 := n % 10
	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}
	if n1%2 == 1 && n1 != 5 && s%3 != 0 {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
