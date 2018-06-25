package main

import (
	"./set1"
	"fmt"
)

func main() {
	buf := set1.ReadFile("set2/2.txt")
	hexBuf := set1.StrToBase64(string(buf))
	ciphertext, _ := set1.Base64ToHex(hexBuf)
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	key := set1.StrToASCII("YELLOW SUBMARINE")
	plaintext := set1.AES_ECB_CBC_Decrypt(ciphertext, key, iv)
	fmt.Println(string(plaintext))
}
