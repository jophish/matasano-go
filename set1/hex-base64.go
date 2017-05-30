package set1

import (
	"bytes"
	"errors"
	"strings"
)

const base64enc = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
const hexenc = "0123456789abcdef"

var hex = []byte("0123456789abcdef")

func StrToBase64(string string) []byte {
	base64Array := make([]byte, len(string))
	for i := 0; i < len(string); i++ {
		base64Array[i] = byte(strings.Index(base64enc, string[i:i+1]))
	}
	return base64Array
}
func StrToHex(string string) []byte {
	string = strings.ToLower(string)
	hexArray := make([]byte, len(string))
	for i := 0; i < len(string); i++ {
		hexArray[i] = byte(bytes.IndexByte(hex, string[i]))
	}
	return hexArray
}

// NOTE: hex chars are only FOUR bytes, not EIGHT
// assumes that each hex buffer has an even length
func HexToBase64(hex []byte) ([]byte, error) {
	extraBytes := 0
	if len(hex)%2 != 0 {
		return nil, errors.New("expected hex array of even length!")
	}
	// check if we need padding bytes in output
	if (len(hex) % 6) != 0 {
		extraBytes = 4
	}
	base64 := make([]byte, ((len(hex)/6)*4)+extraBytes)
	for i := 0; i < len(hex); i += 6 {
		offset := i / 6
		missingBytes := 0
		var bits uint32 = 0
		if i+6 > len(hex) {
			missingBytes = 6 - (len(hex) % 6)
		}

		for j := 0; j < 6-missingBytes; j++ {
			bits |= (uint32(hex[i+j]) << uint32(20-4*j))
		}
		if missingBytes == 4 {
			base64[offset*4] = byte(((bits << 8) >> 26))
			base64[offset*4+1] = byte(((bits << 14) >> 26))
			base64[offset*4+2] = 64
			base64[offset*4+3] = 64
		} else if missingBytes == 2 {
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

// must have that length of buffer is div by 4
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
	hex = make([]byte, (len(base64)/4)*6-numEquals*2)

	for i := 0; i < len(base64); i += 4 {
		offset := i / 4
		var bits uint32 = 0
		if i == len(base64)-4 && numEquals != 0 {
			for j := 0; j < 4-numEquals; j++ {
				bits |= (uint32(base64[i+j]) << uint32(18-6*j))
			}
			if numEquals == 2 {
				hex[offset*6] = byte((bits << 8) >> 28)
				hex[offset*6+1] = byte((bits << 12) >> 28)
			} else if numEquals == 1 {
				hex[offset*6] = byte((bits << 8) >> 28)
				hex[offset*6+1] = byte((bits << 12) >> 28)
				hex[offset*6+2] = byte((bits << 16) >> 28)
				hex[offset*6+3] = byte((bits << 20) >> 28)
			}

		} else {
			for j := 0; j < 4; j++ {
				bits |= (uint32(base64[i+j]) << uint32(18-6*j))
			}
			hex[offset*6] = byte((bits << 8) >> 28)
			hex[offset*6+1] = byte((bits << 12) >> 28)
			hex[offset*6+2] = byte((bits << 16) >> 28)
			hex[offset*6+3] = byte((bits << 20) >> 28)
			hex[offset*6+4] = byte((bits << 24) >> 28)
			hex[offset*6+5] = byte((bits << 28) >> 28)
		}
	}
	return hex, nil
}

func Base64ToStr(base64 []byte) string {
	var buffer bytes.Buffer
	for i := 0; i < len(base64); i++ {
		buffer.WriteString(string(base64enc[base64[i]]))
	}
	return buffer.String()
}

func HexToStr(hexStr []byte) string {
	var buffer bytes.Buffer
	for i := 0; i < len(hexStr); i++ {
		buffer.WriteString(string(hexenc[hexStr[i]]))
	}
	return buffer.String()
}
