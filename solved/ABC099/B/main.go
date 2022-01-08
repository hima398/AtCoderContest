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

	a, b := nextInt(), nextInt()
	towers := []int{1}
	for i := 2; i < 1000; i++ {
		towers = append(towers, towers[i-2]+i)
	}
	//fmt.Println(towers[len(towers)-1])
	for i := 0; i < len(towers)-1; i++ {
		da, db := towers[i]-a, towers[i+1]-b
		if da < 0 || db < 0 {
			continue
		}
		if da == db {
			fmt.Println(da)
			return
		}
	}

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
