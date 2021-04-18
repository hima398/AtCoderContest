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

	a, b, c := nextInt(), nextInt(), nextInt()
	if a+b == c && a-b == c {
		fmt.Println("?")
	} else if a+b == c && a-b != c {
		fmt.Println("+")
	} else if a+b != c && a-b == c {
		fmt.Println("-")
	} else {
		fmt.Println("!")
	}

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
