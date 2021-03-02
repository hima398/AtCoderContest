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
	approved := true
	for i := 0; i < n; i++ {
		a := nextInt()
		if a%2 == 0 {
			if a%3 != 0 && a%5 != 0 {
				approved = false
				break
			}
		}
	}
	if approved {
		fmt.Println("APPROVED")
	} else {
		fmt.Println("DENIED")
	}

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
