package buffstreams

import (
	"encoding/binary"
	"math"
)

func byteArrayToUInt32(bytes []byte) (result int64, bytesRead int) {
	return binary.Varint(bytes)
}

func intToByteArray(value int64, bufferSize int) []byte {
	toWriteLen := make([]byte, 4)
	binary.LittleEndian.PutUint32(toWriteLen, uint32(value))
	if value < 255 {
		toWriteLen = []byte{toWriteLen[0]}
	} else if value < 65536 {
		toWriteLen = toWriteLen[0:2]
	} else if value < 16777215 {
		toWriteLen = toWriteLen[0:3]
	}
	return toWriteLen
}

// Formula for taking size in bytes and calculating # of bits to express that size
// http://www.exploringbinary.com/number-of-bits-in-a-decimal-integer/
func messageSizeToBitLength(messageSize int) int {
	bytes := float64(messageSize)
	header := math.Ceil(math.Floor(math.Log2(bytes)+1)/8.0) + 1
	return int(header)
}
