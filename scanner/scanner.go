package scanner

import (
	"bufio"
	"os"
	"strconv"
)

type Scanner struct {
	sc *bufio.Scanner
}

func New() *Scanner {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	return &Scanner{sc}
}

func (sc *Scanner) NextInt() int {
	sc.sc.Scan()
	i, _ := strconv.Atoi(sc.sc.Text())
	return i
}

func (sc *Scanner) NextFloat64() float64 {
	sc.sc.Scan()
	f, _ := strconv.ParseFloat(sc.sc.Text(), 64)
	return f
}

func (sc *Scanner) NextString() string {
	sc.sc.Scan()
	return sc.sc.Text()
}
