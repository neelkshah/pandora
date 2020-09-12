package helper

func HashKey(key []byte, buckets uint32) uint32 {
	return uint32(key[0]) % buckets
}
