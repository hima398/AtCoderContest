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

	dir := []string{"R", "RU", "U", "LU", "L", "LD", "D", "RD"}
	d := map[string][3]int{"R": {0, 0, 1}, "RU": {1, -1, 1}, "U": {2, -1, 0}, "LU": {3, -1, -1}, "L": {4, 0, -1}, "LD": {5, 1, -1}, "D": {6, 1, 0}, "RD": {7, 1, 1}}

	x, y, w := nextInt(), nextInt(), nextString()
	x--
	y--
	var c [9]string

	for i := 0; i < 9; i++ {
		c[i] = nextString()
	}
	var ans string
	for k := 0; k < 4; k++ {
		ans += string(c[y][x])

		//角度を変える処理が必要ならば以下の条件式内で行われる
		if len(w) == 1 {
			// w = "R", "D", "L", "U"
			if (x == 0 && w == "L") || (y == 0 && w == "U") || (x == 8 && w == "R") || (y == 8 && w == "D") {
				//180度回転
				w = dir[(d[w][0]+4)%8]
			}
		} else {
			if (x == 0 && y == 0 && w == "LU") || (x == 0 && y == 8 && w == "LD") || (x == 8 && y == 0 && w == "RU") || (x == 8 && y == 8 && w == "RD") {
				//180度回転
				w = dir[(d[w][0]+4)%8]
			} else if (x == 0 && w == "LD") || (y == 0 && w == "LU") || (x == 8 && w == "RU") || (y == 8 && w == "RD") {
				//45度回転
				w = dir[(d[w][0]+2)%8]
			} else if (x == 0 && w == "LU") || (y == 0 && w == "RU") || (x == 8 && w == "RD") || (y == 8 && w == "LD") {
				//-45度回転
				w = dir[(d[w][0]+6)%8]
			}
		}

		//位置を更新
		x, y = x+d[w][2], y+d[w][1]
		//fmt.Println(y, x)
	}
	fmt.Println(ans)
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
