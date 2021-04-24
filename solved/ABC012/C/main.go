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

	sum := 0
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			sum += i * j
		}
	}
	//fmt.Println(sum)
	n := nextInt()
	diff := sum - n
	var ans []string
	for i := 1; i <= 9; i++ {
		if diff%i == 0 && diff/i <= 9 {
			ans = append(ans, fmt.Sprintf("%d x %d", i, diff/i))
		}
	}
	if len(ans) > 0 {
		for _, v := range ans {
			fmt.Println(v)
		}
	} else {
		fmt.Println("1 x 1")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
