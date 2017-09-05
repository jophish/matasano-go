package main

import (
	"./set1"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./set1/4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buffer := make([][]byte, 327)
	i := 0
	for scanner.Scan() {

		buffer[i] = set1.StrToHex(scanner.Text())
		i += 1
	}

	ans := set1.DetectSingleXOR(buffer)
	decrypted, _ := set1.BreakSingleByteXORCipher(ans)
	fmt.Println("Encrypted string: ", set1.HexToStr(ans))
	fmt.Println("Decrypted string: ", string(decrypted))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
