package main

import (
	"./set1"
	"fmt"
)

func main() {

	buf := set1.ReadFile("set1/7.txt")
	hexBuf := set1.StrToBase64(string(buf))
	cipher, _ := set1.Base64ToHex(hexBuf)
	key := set1.StrToASCII("YELLOW SUBMARINE")

	plaintext := set1.AES_ECB_Decrypt(cipher, key)
	fmt.Println(string(plaintext))
}
