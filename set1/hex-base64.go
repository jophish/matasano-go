package set1

import (
	"bytes"
	"errors"
	"strings"
)

const ascii = " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
const base64enc = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
const hexenc = "0123456789abcdef"

var hex = []byte("0123456789abcdef")

// Given an ASCII string, returns a byte array representing the ASCII encoding.
func StrToASCII(string string) []byte {
	return []byte(string)
}

// Given a string containing base64 characters, returns a buffer of bytes representing the base64 encoding of the string.
// each character is represented by the first six bits of of a byte.
func StrToBase64(string string) []byte {
	base64Array := make([]byte, len(string))
	for i := 0; i < len(string); i++ {
		base64Array[i] = byte(strings.Index(base64enc, string[i:i+1]))
	}
	return base64Array

}

// Given a string containing hex characters, returns a buffer of bytes half the length of the given string
// containing the hex encoding of the string. Each byte in the buffer corresponds to two characters of the string.
// The length of the input string must be a multiple of two.
func StrToHex(string string) []byte {
	string = strings.ToLower(string)
	hexArray := make([]byte, len(string)/2)
	for i := 0; i < len(string); i += 2 {
		hexArray[i/2] = byte((bytes.IndexByte(hex, string[i]) << 4)) | byte(bytes.IndexByte(hex, string[i+1]))
	}
	return hexArray
}

// Given a buffer containing the hex encoding of a string, returns a buffer containing the base64 representation
// of the same string.
func HexToBase64(hex []byte) ([]byte, error) {
	extraBytes := 0
	// check if we need padding bytes in output
	if (len(hex) % 3) != 0 {
		extraBytes = 4
	}
	base64 := make([]byte, ((len(hex)/3)*4)+extraBytes)
	for i := 0; i < len(hex); i += 3 {
		offset := i / 3
		missingBytes := 0
		var bits uint32 = 0
		if i+3 > len(hex) {
			missingBytes = 3 - (len(hex) % 3)
		}

		for j := 0; j < 3-missingBytes; j++ {
			bits |= (uint32(hex[i+j]) << uint32(16-8*j))
		}

		if missingBytes == 2 {
			base64[offset*4] = byte(((bits << 8) >> 26))
			base64[offset*4+1] = byte(((bits << 14) >> 26))
			base64[offset*4+2] = 64
			base64[offset*4+3] = 64
		} else if missingBytes == 1 {
			base64[offset*4] = byte(((bits << 8) >> 26))
			base64[offset*4+1] = byte(((bits << 14) >> 26))
			base64[offset*4+2] = byte(((bits << 20) >> 26))
			base64[offset*4+3] = 64
		} else {
			base64[offset*4] = byte(((bits << 8) >> 26))
			base64[offset*4+1] = byte(((bits << 14) >> 26))
			base64[offset*4+2] = byte(((bits << 20) >> 26))
			base64[offset*4+3] = byte(((bits << 26) >> 26))
		}
	}

	return base64, nil
}

// Give a byte array containing the base64 encoding of a string, returns a byte
// array containing the hex representation of the same string.
func Base64ToHex(base64 []byte) ([]byte, error) {

	if len(base64)%4 != 0 {
		return nil, errors.New("expected length of buffer to a multiple of 4!")
	}
	numEquals := 0
	if base64[len(base64)-2] == 64 {
		numEquals = 2
	} else if base64[len(base64)-1] == 64 {
		numEquals = 1
	}
	hex = make([]byte, (len(base64)/4)*3-numEquals)
	for i := 0; i < len(base64); i += 4 {
		offset := i / 4
		var bits uint32 = 0
		if i == len(base64)-4 && numEquals != 0 {
			for j := 0; j < 4-numEquals; j++ {
				bits |= (uint32(base64[i+j]) << uint32(18-6*j))
			}
			if numEquals == 2 {
				hex[offset*3] = byte((bits << 8) >> 24)
			} else if numEquals == 1 {
				hex[offset*3] = byte((bits << 8) >> 24)
				hex[offset*3+1] = byte((bits << 16) >> 24)
			}

		} else {
			for j := 0; j < 4; j++ {
				bits |= (uint32(base64[i+j]) << uint32(18-6*j))
			}
			hex[offset*3] = byte((bits << 8) >> 24)
			hex[offset*3+1] = byte((bits << 16) >> 24)
			hex[offset*3+2] = byte((bits << 24) >> 24)
		}
	}
	return hex, nil
}

// Given a byte array representing the base64 encoding of a string, returns
// the actual human-readable string.
func Base64ToStr(base64 []byte) string {
	var buffer bytes.Buffer
	for i := 0; i < len(base64); i++ {
		buffer.WriteString(string(base64enc[base64[i]]))
	}
	return buffer.String()
}

// Given a byte array representing the hex encoding of a string, returns
// the actual human-readable string.
func HexToStr(hexStr []byte) string {
	var buffer bytes.Buffer
	for i := 0; i < len(hexStr); i++ {
		buffer.WriteString(string(hexenc[hexStr[i]>>4]))
		buffer.WriteString(string(hexenc[(hexStr[i]<<4)>>4]))
	}
	return buffer.String()
}

// Given a byte array representing the ASCII encoding of a string, returns
// the actual human-readable string
func ASCIIToStr(str []byte) string {
	var result string = ""
	for i := 0; i < len(str); i++ {
		result += string(ascii[str[i]-32])
	}
	return result
}
