package main

import (
	"./set1"
	"fmt"
	"os"
)

func main() {
	buffer1 := set1.StrToHex("1c0111001f010100061a024b53535009181c")
	buffer2 := set1.StrToHex("686974207468652062756c6c277320657965")
	output, err := set1.XORBuffers(buffer1, buffer2)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	outputStr := set1.HexToStr(output)
	fmt.Println(outputStr)
}
