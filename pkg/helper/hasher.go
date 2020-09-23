package helper

func HashKey(key []byte, buckets uint64) uint64 {
	return uint64(key[0]) % buckets
}
