package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	abc := make([]int, 3)
	abc[0], abc[1], abc[2] = nextInt(), nextInt(), nextInt()
	sort.Ints(abc)
	fmt.Println(10*abc[2] + abc[1] + abc[0])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
