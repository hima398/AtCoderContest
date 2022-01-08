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
	ans := make([]int, 46)
	ans[0] = 1
	ans[1] = 1
	//n2 := n / 2
	for i := 2; i <= 45; i++ {
		ans[i] = ans[i-1] + ans[i-2]
	}
	fmt.Println(ans[n])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
