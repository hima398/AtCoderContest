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

	m := map[int]string{0: "pon", 1: "pon", 2: "hon", 3: "bon", 4: "hon", 5: "hon", 6: "pon", 7: "hon", 8: "pon", 9: "hon"}
	n := nextInt()

	fmt.Println(m[n%10])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
