package set1

import (
	"crypto/aes"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
)

// Returns a buffer of len random bytes
func RandomBytes(len int) []byte {
	buf := make([]byte, len)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return buf
}

// Given a plaintext, this function first pads it at the beginning and end with
// 5-10 random bytes, (the SAME bytes, at the beginning and end), and encrypts it
// with a random 16 byte key in ECB or CBC mode. If CBC mode is chosen, the IV is also
// randomly initialized. The ciphertext is returned.
func AES_EncryptionOracle(plaintext []byte) []byte {
	// Frist, generate a 5-10 byte prefix/suffix
	pre_len, _ := rand.Int(rand.Reader, big.NewInt(6))
	prefix := make([]byte, pre_len.Int64()+5)
	_, err := rand.Read(prefix)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Now, append and prepend
	plaintext = append(plaintext, prefix...)
	plaintext = append(prefix, plaintext...)

	// AES requires a 16 byte aligned plaintext, so we have to pad
	// to the closest 16 byte boundary
	padlen := len(plaintext) + (16 - (len(plaintext) % 16))
	plaintext, err = PadPKCS7(plaintext, padlen)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Generate a random key
	key := RandomKey()
	var ciphertext []byte

	mode, _ := rand.Int(rand.Reader, big.NewInt(2))
	if mode.Int64() == 0 {
		// ECB mode
		fmt.Println("We're encrypting with ECB mode")
		ciphertext = AES_ECB_Encrypt(plaintext, key)
	} else {
		// CBC mode
		fmt.Println("We're encrypting with CBC mode")
		iv := RandomKey()
		ciphertext = AES_CBC_Encrypt(plaintext, key, iv)
	}

	return ciphertext
}

// Returns a random 16 byte AES key
func RandomKey() []byte {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return key
}

// Given a plaintext which is a multiple of 16 bytes,
// 16 byte initialization vector, and a 16 byte key,
// returns the AES-128-ECB encrypted ciphertext, operating in
// CBC mode.
func AES_CBC_Encrypt(plaintext, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bs := block.BlockSize()
	ciphertext := []byte{}
	prevVec := iv
	for len(plaintext) > 0 {
		// Before encrypting, each plaintext block should be XOR'd with the output
		// of the previous encryption round. The very first plaintext block should
		// instead be XOR'd with the given initialization vector.
		chunk := make([]byte, block.BlockSize())
		XORblock, _ := XORBuffers(plaintext[:bs], prevVec)
		block.Encrypt(chunk, XORblock)
		prevVec = chunk
		ciphertext = append(ciphertext, chunk...)
		plaintext = plaintext[bs:]
	}
	return ciphertext
}

// Given a ciphertext which is a multiple of 16 bytes, a
// 16 byte key, and a 16 byte initialization vector, returns
// the AES-128-ECB decrypted plaintext, operating in CBC mode
func AES_CBC_Decrypt(ciphertext, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bs := block.BlockSize()
	plaintext := []byte{}
	prevVec := iv
	for len(ciphertext) > 0 {
		chunk := make([]byte, bs)
		block.Decrypt(chunk, ciphertext)
		XORblock, err := XORBuffers(chunk, prevVec)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		plaintext = append(plaintext, XORblock...)
		prevVec = ciphertext[:bs]
		ciphertext = ciphertext[bs:]
	}
	return DepadPKCS(plaintext)
}

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
