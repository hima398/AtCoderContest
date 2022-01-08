package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func F(k int) int {
	sk := strconv.Itoa(k)
	n := len(sk)
	res := 0
	w := 1
	for i := n - 1; i >= 0; i-- {
		x := int(sk[i] - '0')
		res += x * w
		w *= k
	}
	return res
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a := nextInt()

	ma := make(map[int]int)
	max := int(1e16)
	k := 10
	aa := F(k)
	for aa <= max {
		ma[aa] = k
		k++
		aa = F(k)
	}
	if k, ok := ma[a]; ok {
		fmt.Println(k)
	} else {
		fmt.Println(-1)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
