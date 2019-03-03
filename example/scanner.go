package main

import (
	"fmt"

	"github.com/syumai/proconutil/scanner"
)

func main() {
	sc := scanner.New()
	fmt.Println(sc.NextString())
	fmt.Println(sc.NextInt())
	fmt.Println(sc.NextFloat64())
}
