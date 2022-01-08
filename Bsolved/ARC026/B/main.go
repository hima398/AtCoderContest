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
	if n == 1 {
		fmt.Println("Deficient")
		return
	}
	d := make(map[int]int)
	d[1]++
	//d := []int{1}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			d[i]++
			d[n/i]++
		}
	}
	//fmt.Println(d)
	s := 0
	for k := range d {
		s += k
	}
	if n < s {
		fmt.Println("Abundant")
	} else if n > s {
		fmt.Println("Deficient")
	} else {
		fmt.Println("Perfect")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
