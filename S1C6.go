package main

import (
	"./set1"
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(file)
	bytes := make([]byte, 0, 1024)

	for scan.Scan() {
		input := []byte(scan.Text())
		bytes = append(bytes, input...)
	}
	return bytes
}
func main() {

	buf := ReadFile("set1/6.txt")
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
