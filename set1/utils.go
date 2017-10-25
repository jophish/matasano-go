package set1

import (
	"bufio"
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

func ReadFileByLines(filename string) [][]byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buffer := [][]byte{}

	for scanner.Scan() {
		buffer = append(buffer, StrToHex(scanner.Text()))
	}

	return buffer
}
