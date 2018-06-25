package set1

import (
	"crypto/aes"
	"errors"
	"fmt"
	"os"
)

// Given a plaintext which is a multiple of 16 bytes and a
// 16 byte key, returns the AES-128-ECB encrypted ciphertext
func AES_ECB_Encrypt(plaintext, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bs := block.BlockSize()
	ciphertext := []byte{}
	for len(plaintext) > 0 {
		chunk := make([]byte, block.BlockSize())
		block.Encrypt(chunk, plaintext)
		ciphertext = append(ciphertext, chunk...)
		plaintext = plaintext[bs:]
	}
	return ciphertext
}

// Given a ciphertext which is a multiple of 16 bytes and a
// 16 byte key, returns the AES-128-ECB decrypted plaintext
func AES_ECB_Decrypt(ciphertext, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bs := block.BlockSize()
	plaintext := []byte{}

	for len(ciphertext) > 0 {
		chunk := make([]byte, block.BlockSize())
		block.Decrypt(chunk, ciphertext)
		plaintext = append(plaintext, chunk...)
		ciphertext = ciphertext[bs:]
	}
	return DepadPKCS(plaintext)
}

// Given a buffer which may be PKCS#07 padded,
// return the de-padded version
func DepadPKCS(buffer []byte) []byte {
	for i := 0; i < len(buffer); i++ {
		if buffer[i] == 4 {
			return buffer[:i]
		}
	}
	return buffer
}

// Given a buffer of byte arrays, returns the array which
// is most probably ecnrypted with AES-ECB-128
func DetectECB(buffer [][]byte) ([]byte, error) {
	for i := 0; i < len(buffer); i++ {
		chunked := ChunkBuffer(buffer[i], 16)
		for j := 0; j < len(chunked); j++ {
			for k := j + 1; k < len(chunked); k++ {
				score, err := HammingDistance(chunked[j], chunked[k])
				if err != nil {
					return nil, errors.New("Expected all arrays to have length divisible by 16 bytes!")
				}
				// The Hamming Distance is 0 if two encrypted chunks are the same. This
				// should never happen in any sane encryption scheme
				if score == 0 {
					return buffer[i], nil
				}
			}
		}
	}

	return nil, errors.New("No AES-ECB-128 array detected!")
}

// Given a buffer (a "block") and a length, pads the buffer
// to the given length with \x04 bytes. `length` must be greater
// than the size of the given buffer
func PadPKCS7(buffer []byte, length int) ([]byte, error) {
	if len(buffer) > length {
		return nil, errors.New("Expected buffer to be less than or equal to given length!")
	}
	newBuf := make([]byte, length)
	for i := 0; i < len(buffer); i++ {
		newBuf[i] = buffer[i]
	}
	for i := len(buffer); i < len(newBuf); i++ {
		newBuf[i] = byte(length - len(buffer))
	}
	return newBuf, nil
}
