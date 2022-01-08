package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func EtoD(s string) int {
	//ss := strings.Split(s, "")
	ns := len(s)
	w := 1
	ret := 0
	for i := 0; i < ns; i++ {
		ni := int(s[ns-i-1] - '0')
		ret += ni * w
		w *= 8
	}
	return ret
}

func DtoN(n int) []string {
	var s []int
	for n != 0 {
		ni := n % 9
		s = append(s, ni)
		n = (n - ni) / 9
	}
	ret := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		ret[i] = strconv.Itoa(s[len(s)-i-1])
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextString(), nextInt()
	if n == "0" {
		fmt.Println(0)
		return
	}
	//sn := strings.Split(n, "")
	for i := 0; i < k; i++ {
		d := EtoD(n)
		nine := DtoN(d)
		for i := range nine {
			if nine[i] == "8" {
				nine[i] = "5"
			}
		}
		n = strings.Join(nine, "")
	}
	fmt.Println(n)
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
