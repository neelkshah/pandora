package helper

import "encoding/binary"

// IntToByte takes a uint64 value and converts it into a little endian byte array.
func IntToByte(value uint64) []byte {
	arr := make([]byte, 8)
	binary.LittleEndian.PutUint64(arr, value)
	return arr
}

// ByteToInt takes a little endian byte array and converts it into a uint64 value.
func ByteToInt(arr []byte) uint64 {
	value := binary.LittleEndian.Uint64(arr[0:])
	return value
}
