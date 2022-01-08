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

	s := nextString()
	//5,000,000,000,000,000
	k := nextInt()
	no := 0
	for _, r := range s {
		if r == '1' {
			no++
		} else {
			break
		}
	}
	if k <= no {
		fmt.Println(1)
	} else {
		fmt.Println(int(s[no] - '0'))
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
