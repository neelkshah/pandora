// Author: Shubhankar Ranade
// Implements hash functions to be used in hashtable

package hashfunc

import (
	"hash"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

// FNV64a implements FNV-1a hash
func FNV64a(key []byte, tableSize uint64) uint64 {
	h64a := fnv.New64a()
	return uint64Hash(h64a, key) % tableSize
}

// FNV64 implements FNV-1 hash
func FNV64(key []byte, tableSize uint64) uint64 {
	h64 := fnv.New64()
	return uint64Hash(h64, key) % tableSize
}

// uint64Hash calculates prehash for a byte slice
func uint64Hash(hasher hash.Hash64, key []byte) uint64 {
	hasher.Write(key)
	return hasher.Sum64()
}

// DJB2a implements djb2a hash
func DJB2a(key []byte, tableSize uint64) uint64 {
	var hash uint64 = 5382
	for _, num := range key {
		hash = ((hash << 5) + hash) ^ uint64(num)
	}

	return hash % tableSize
}

// DJB2 implements djb2 hash
func DJB2(key []byte, tableSize uint64) uint64 {
	var hash uint64 = 5382
	for _, num := range key {
		hash = (hash * 33) + uint64(num)
	}
	return hash % tableSize
}

// SDBM implements SDBM hash
func SDBM(key []byte, tableSize uint64) uint64 {
	var hash uint64 = 0
	for _, num := range key {
		hash = (hash << 6) + (hash << 16) - hash + uint64(num)
	}
	return hash % tableSize
}

// PJW implements PJW hash
func PJW(key []byte, tableSize uint64) uint64 {
	var h uint64 = 0
	var high uint64
	for _, num := range key {
		h = (h << 4) + uint64(num)
		high = h & 0xF0000000
		if high != 0 {
			h ^= high >> 24
		}
		h &= (high ^ 0xFFFFFFFF)
	}
	return h % tableSize
}

// Murmur returns murmur3 hash
func Murmur(key []byte, tableSize uint64) uint64 {
	return murmur3.Sum64(key) % tableSize
}
