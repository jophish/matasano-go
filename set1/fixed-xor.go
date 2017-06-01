package set1

import (
	"errors"
	"math"
)

// statistics are from here: fitaly.com/board/domper3/posts/b136.html
var frequencies = map[byte]float64{
	32:  17.1662,
	33:  0.0072,
	34:  0.2442,
	35:  0.0179,
	36:  0.0561,
	37:  0.0160,
	38:  0.0226,
	39:  0.2447,
	40:  0.2178,
	41:  0.2233,
	42:  0.0628,
	43:  0.0215,
	44:  0.7384,
	45:  1.3734,
	46:  1.5124,
	47:  0.1549,
	48:  0.5516,
	49:  0.4594,
	50:  0.3322,
	51:  0.1847,
	52:  0.1348,
	53:  0.1663,
	54:  0.1153,
	55:  0.1030,
	56:  0.1054,
	57:  0.1024,
	58:  0.4354,
	59:  0.1214,
	60:  0.1225,
	61:  0.0227,
	62:  0.1242,
	63:  0.1474,
	64:  0.0073,
	65:  0.3132,
	66:  0.2163,
	67:  0.3906,
	68:  0.3151,
	69:  0.2673,
	70:  0.1416,
	71:  0.1876,
	72:  0.2321,
	73:  0.3211,
	74:  0.1726,
	75:  0.0687,
	76:  0.1884,
	77:  0.3529,
	78:  0.2085,
	79:  0.1842,
	80:  0.2614,
	81:  0.0316,
	82:  0.2519,
	83:  0.4003,
	84:  0.3322,
	85:  0.0814,
	86:  0.0892,
	87:  0.2527,
	88:  0.0343,
	89:  0.0304,
	90:  0.0076,
	91:  0.0086,
	92:  0.0016,
	93:  0.0088,
	94:  0.0003,
	95:  0.1159,
	96:  0.0009,
	97:  5.1880,
	98:  1.0195,
	99:  2.1129,
	100: 2.5071,
	101: 8.5771,
	102: 1.3725,
	103: 1.5597,
	104: 2.7444,
	105: 4.9019,
	106: 0.0867,
	107: 0.6753,
	108: 3.1750,
	109: 1.6437,
	110: 4.9701,
	111: 5.7701,
	112: 1.5482,
	113: 0.0747,
	114: 4.2586,
	115: 4.3686,
	116: 6.3700,
	117: 2.0999,
	118: 0.8462,
	119: 1.3034,
	120: 0.1950,
	121: 1.1330,
	122: 0.0596,
	123: 0.0026,
	124: 0.0007,
	125: 0.0026,
	126: 0.0003,
}

// Given two byte buffers of equal length, returns a byte array produced by
// XORing repsective bytes of the input array
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
	var best byte
	var min float64 = -1
	for i := 32; i <= 126; i++ {
		XOR, _ := XORBufferWithChar(buffer, byte(i))
		chi := ChiSquaredPlaintext(XOR)
		if min < 0 {
			min = chi
		}
		if chi <= min {
			best = byte(i)
			min = chi
		}

	}
	result, _ := XORBufferWithChar(buffer, best)
	return result
}

// Given a byte buffer and a single byte, returns the byte array
// resulting from XORing each byte in the buffer with the single input byte
func XORBufferWithChar(buffer []byte, char byte) ([]byte, error) {
	buffer2 := make([]byte, len(buffer))
	for i := 0; i < len(buffer); i++ {
		buffer2[i] = char
	}
	buffer, err := XORBuffers(buffer, buffer2)
	return buffer, err
}

// Given a byte buffer of ASCII encoded text, perform a chi squared
// test on character frequency and return the result
func ChiSquaredPlaintext(buffer []byte) float64 {
	var chi float64 = 0
	for k, v := range frequencies {

		var Oi float64 = GetByteFrequency(buffer, k)
		var Ei float64 = v
		chi += math.Pow((Oi-Ei), 2) / Ei

	}
	return chi
}

// Given a byte buffer and a single byte, returns the frequency of the single
// byte's occurence in the given buffer
func GetByteFrequency(buffer []byte, char byte) float64 {
	var count float64 = 0
	for i := 0; i < len(buffer); i++ {
		if buffer[i] == char {

			count += 1
		}
	}
	return float64(count / float64(len(buffer)))
}
