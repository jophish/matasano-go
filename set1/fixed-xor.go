package set1

import (
	"errors"
)

func XORBuffers(buffer1, buffer2 []byte) ([]byte, error) {
	if len(buffer1) != len(buffer2) {
		return nil, errors.New("expected both buffers to have same length!")
	}

	output := make([]byte, len(buffer1))
	for i := 0; i < len(buffer1); i++ {
		output[i] = buffer1[i] ^ buffer2[i]
	}
	return output, nil
}

func DecryptSingleByteXORCipher(buffer []byte) []byte {
	// try XORing with each character, maintain a dictionary mapping chars to freq scores,
	// return the decryption using the character with the best score
	//chars := "0123456789abcdef"
	return nil

}

func XORBufferWithChar(buffer []byte, char byte) ([]byte, error) {
	buffer2 := make([]byte, len(buffer))
	for i := 0; i < len(buffer); i++ {
		buffer2[i] = char
	}
	buffer, err := XORBuffers(buffer, buffer2)
	return buffer, err
}
