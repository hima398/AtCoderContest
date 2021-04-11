package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	sx, sy, tx, ty := nextInt(), nextInt(), nextInt(), nextInt()
	ans := ""
	ans += strings.Repeat("R", tx-sx)
	ans += strings.Repeat("U", ty-sy)
	ans += strings.Repeat("L", tx-sx)
	ans += strings.Repeat("D", ty-sy)
	ans += "D"
	ans += strings.Repeat("R", tx-sx+1)
	ans += strings.Repeat("U", ty-sy+1)
	ans += "LU"
	ans += strings.Repeat("L", tx-sx+1)
	ans += strings.Repeat("D", ty-sy+1)
	ans += "R"
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
