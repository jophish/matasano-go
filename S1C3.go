package main

import (
	"./set1"
	"bufio"
	"fmt"
	"os"
)

func main() {
	hexBuffer := set1.StrToHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	for i := 0; i < 256; i++ {
		XOR, err := set1.XORBufferWithChar(hexBuffer, byte('X'))
		fmt.Println(XOR)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(set1.HexToStr(XOR))
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
