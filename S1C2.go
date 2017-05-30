package main

import (
	"./set1"
	"fmt"
	"os"
)

func main() {
	buffer1 := set1.StrToHex(os.Args[1])
	buffer2 := set1.StrToHex(os.Args[2])
	output, err := set1.XORBuffers(buffer1, buffer2)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	outputStr := set1.HexToStr(output)
	fmt.Println(outputStr)
}
