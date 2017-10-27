package buffstreams

import (
	"encoding/binary"
	"math"
)

func byteArrayToUInt32(bytes []byte) (result int64, bytesRead int) {
	return binary.Varint(bytes)
}

func intToByteArray(value int64, bufferSize int) []byte {
	toWriteLen := make([]byte, 0, 4)
	for value >= 1<<7 {
		toWriteLen = append(toWriteLen, uint8(value&0x7f|0x80))
		value >>= 7
	}
	toWriteLen = append(toWriteLen, uint8(value))
	return toWriteLen
}

// Formula for taking size in bytes and calculating # of bits to express that size
// http://www.exploringbinary.com/number-of-bits-in-a-decimal-integer/
func messageSizeToBitLength(messageSize int) int {
	bytes := float64(messageSize)
	header := math.Ceil(math.Floor(math.Log2(bytes)+1)/8.0) + 1
	return int(header)
}
