package main

import (
	"./set1"
	"fmt"
)

func main() {
	buffer := set1.ReadFileByLines("set1/4.txt")
	ans := set1.DetectSingleXOR(buffer)
	decrypted, _ := set1.BreakSingleByteXORCipher(ans)
	fmt.Println("Encrypted string: ", set1.HexToStr(ans))
	fmt.Println("Decrypted string: ", string(decrypted))
}
