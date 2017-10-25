package main

import (
	"./set1"
	"fmt"
)

func main() {

	buf := set1.ReadFile("set1/6.txt")
	hexBuf := set1.StrToBase64(string(buf))
	cipher, _ := set1.Base64ToHex(hexBuf)
	str1 := "this is a test"
	str2 := "wokka wokka!!!"

	buf1 := set1.StrToASCII(str1)
	buf2 := set1.StrToASCII(str2)
	ans, _ := set1.HammingDistance(buf1, buf2)
	fmt.Println(ans)

	decrypted, key := set1.BreakRepeatingXOR(cipher)
	fmt.Println(string(decrypted))
	fmt.Println(string(key))
}
