package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ComputeDeg(d int) string {
	if d >= 113 && d <= 337 {
		return "NNE"
	} else if d >= 338 && d <= 562 {
		return "NE"
	} else if d >= 563 && d <= 787 {
		return "ENE"
	} else if d >= 788 && d <= 1012 {
		return "E"
	} else if d >= 1013 && d <= 1237 {
		return "ESE"
	} else if d >= 1238 && d <= 1462 {
		return "SE"
	} else if d >= 1463 && d <= 1687 {
		return "SSE"
	} else if d >= 1688 && d <= 1912 {
		return "S"
	} else if d >= 1913 && d <= 2137 {
		return "SSW"
	} else if d >= 2138 && d <= 2362 {
		return "SW"
	} else if d >= 2363 && d <= 2587 {
		return "WSW"
	} else if d >= 2588 && d <= 2812 {
		return "W"
	} else if d >= 2813 && d <= 3037 {
		return "WNW"
	} else if d >= 3038 && d <= 3262 {
		return "NW"
	} else if d >= 3263 && d <= 3487 {
		return "NNW"
	} else {
		return "N"
	}
}

func ComputeDis(d int) int {
	d1 := (10 * d / 6) % 10
	if d1 >= 0 && d1 < 5 {
		d = Floor(d, 6)
	} else {
		d = Ceil(d, 6)
	}
	if d >= 0 && d <= 2 {
		return 0
	} else if d >= 3 && d <= 15 {
		return 1
	} else if d >= 16 && d <= 33 {
		return 2
	} else if d >= 34 && d <= 54 {
		return 3
	} else if d >= 55 && d <= 79 {
		return 4
	} else if d >= 80 && d <= 107 {
		return 5
	} else if d >= 108 && d <= 138 {
		return 6
	} else if d >= 139 && d <= 171 {
		return 7
	} else if d >= 172 && d <= 207 {
		return 8
	} else if d >= 208 && d <= 244 {
		return 9
	} else if d >= 245 && d <= 284 {
		return 10
	} else if d >= 285 && d <= 326 {
		return 11
	} else {
		return 12
	}

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	deg, dis := nextInt(), nextInt()

	ansDeg := ComputeDeg(deg)
	ansDis := ComputeDis(dis)
	if ansDis == 0 {
		fmt.Println("C", ansDis)
	} else {
		fmt.Println(ansDeg, ansDis)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
