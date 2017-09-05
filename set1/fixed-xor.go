package set1

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
)

// We use Score and ScoreList to sort potential keysizes from
// most-to-least likely
type Score struct {
	len   uint32
	score float64
}

type ScoreList []Score

func (s ScoreList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ScoreList) Len() int {
	return len(s)
}

func (s ScoreList) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s ScoreList) ToList() []uint32 {
	buffer := make([]uint32, s.Len())
	for i := 0; i < s.Len(); i++ {
		buffer[i] = s[i].len
	}
	return buffer
}

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

func BreakSingleByteXORCipher(buffer []byte) ([]byte, byte) {
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
	return result, best
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

// Given a buffer of byte arrays, where a single array is encrypted by
// single character XOR, returns the array which is most likely encrypted.
func DetectSingleXOR(buffer [][]byte) []byte {
	var bestChi float64 = -1
	var bestBuffer []byte

	for i := 0; i < len(buffer); i++ {
		plain, _ := BreakSingleByteXORCipher(buffer[i])
		chi := ChiSquaredPlaintext(plain)
		if bestChi < 0 {
			bestChi = chi
		}
		if chi < bestChi {
			bestChi = chi
			bestBuffer = buffer[i]
		}
	}

	return bestBuffer
}

// Given a buffer and a key, XORs the key with the buffer enough times
// to completely XOR the buffer.
func RepeatingXOR(buffer, key []byte) []byte {
	output := make([]byte, len(buffer))
	for i := 0; i < len(buffer); i++ {
		output[i] = buffer[i] ^ key[i%len(key)]
	}
	return output
}

// Given a buffer encrypted with repeated-key XOR, breaks the encryption
// and returns the decrypted buffer and the key
func BreakRepeatingXOR(buffer []byte) ([]byte, []byte) {
	keySizes := GuessKeySize(buffer)
	//for each keysize, break into keysize chunks
	var bestKey []byte
	var bestChi float64 = -1

	for keyIndex := 0; keyIndex < len(keySizes); keyIndex++ {
		chunked := ChunkBuffer(buffer, keySizes[keyIndex])
		transposed := TransposeChunks(chunked)

		keyBuf := make([]byte, len(transposed))
		for i := 0; i < len(transposed); i++ {
			_, best := BreakSingleByteXORCipher(transposed[i])
			keyBuf[i] = best
		}
		decrypted := RepeatingXOR(buffer, keyBuf)
		testChi := ChiSquaredPlaintext(decrypted)

		if testChi < bestChi || bestChi < 0 {
			bestChi = testChi
			bestKey = keyBuf
		}
	}

	return RepeatingXOR(buffer, bestKey), bestKey
}

// Given a buffer and a length , will break the buffer up into
// chunks of the given length
func ChunkBuffer(buffer []byte, chunkLen uint32) [][]byte {
	bufferArray := [][]byte{}
	for i := 0; i < len(buffer); i += int(chunkLen) {
		upperBound := i + int(chunkLen)
		if upperBound >= len(buffer) {
			upperBound = len(buffer)
		}
		chunk := buffer[i:upperBound]
		bufferArray = append(bufferArray, chunk)
	}
	return bufferArray
}

// Given a buffer of arrays, returns an array of buffers
// created by transposing the originals. For example,
// [[1,2,3],[4,5,6],[7,8]] goes to
// [[1,4,7],[2,5,8],[3,6]]
func TransposeChunks(buffer [][]byte) [][]byte {
	maxLenChunk := 0
	for i := 0; i < len(buffer); i++ {
		if len(buffer[i]) > maxLenChunk {
			maxLenChunk = len(buffer[i])
		}
	}

	bufferArray := [][]byte{}
	for i := 0; i < maxLenChunk; i++ {
		transposedChunk := []byte{}
		for j := 0; j < len(buffer); j++ {
			if i < len(buffer[j]) {
				transposedChunk = append(transposedChunk, buffer[j][i])
			}
		}
		bufferArray = append(bufferArray, transposedChunk)
	}
	return bufferArray
}

// Given two byte buffers, computes the Hamming Distance between them.
// Expects that each buffer is the same length
func HammingDistance(buffer1, buffer2 []byte) (uint32, error) {
	var count uint32 = 0
	if len(buffer1) != len(buffer2) {
		return 0, errors.New("expected both buffers to have the same length!")
	}
	for i := 0; i < len(buffer1); i++ {
		xor := buffer1[i] ^ buffer2[i]
		for j := 1; j < 8; j++ {
			count += uint32((xor << uint32(8-j)) >> 7)
		}
	}
	return count, nil
}

// Given a buffer encrypted with repeating-key XOR, returns a list of keysizes
// sorted form most-to-least likely
func GuessKeySize(buffer []byte) []uint32 {
	var minKeysize uint32 = 2
	var maxKeysize uint32 = uint32(len(buffer) / 2)
	maxKeysize = 40
	scoreBuffer := make([]Score, maxKeysize-minKeysize)
	for i := minKeysize; i < maxKeysize; i++ {
		//fmt.Println(buffer[:i], buffer[i:2*i])
		score, err := HammingDistance(buffer[:i], buffer[i:2*i])
		var normalized float64 = float64(float64(score) / float64(i))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		scoreBuffer[i-minKeysize] = Score{i, normalized}
	}
	sort.Sort(ScoreList(scoreBuffer))

	return ScoreList(scoreBuffer).ToList()
}
