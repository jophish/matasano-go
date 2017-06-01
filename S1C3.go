package main

import (
	"./set1"
	"fmt"
	"os"
)

func main() {
	var hexBuffer []byte
	if len(os.Args) != 2 {
		hexBuffer = set1.StrToHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	} else {
		hexBuffer = set1.StrToHex(os.Args[1])
	}

	result := set1.DecryptSingleByteXORCipher(hexBuffer)
	fmt.Printf("%+q\n", result)

}
