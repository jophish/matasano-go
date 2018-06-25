package main

import (
	"./set1"
	"fmt"
)

func main() {
	plainText := "YELLOW SUBMARINE"
	plainTextBuf := set1.StrToASCII(plainText)
	padded, _ := set1.PadPKCS7(plainTextBuf, 20)
	outStr := set1.ASCIIToStr(padded)
	fmt.Println(outStr)
}
