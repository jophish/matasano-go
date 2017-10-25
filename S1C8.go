package main

import (
	"./set1"
	"fmt"
	"os"
)

func main() {
	lines := set1.ReadFileByLines("set1/8.txt")
	// look at 16 byte chunks of each line, 1
	best, err := set1.DetectECB(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(set1.HexToStr(best))
}
