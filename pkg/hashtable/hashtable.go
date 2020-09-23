package hashtable

import (
	"github.com/neelkshah/pandora/config"
	"github.com/neelkshah/pandora/pkg/helper"
	ll "github.com/neelkshah/pandora/pkg/linkedlist"
)

type HashTable struct {
	table   []*ll.LinkedList
	buckets uint64
	inserts uint64
	deletes uint64
}

func CreateHashTable(size uint64) *HashTable {
	var newHashTable = HashTable{table: make([]*ll.LinkedList, size), buckets: size, inserts: 0, deletes: 0}
	for i := 0; i < int(size); i++ {
		newHashTable.table[i] = ll.CreateLinkedList()
	}

	return &newHashTable
}

func (ht *HashTable) Insert(key uint64, value uint64) {
	keyBytes := helper.IntToByte(key)
	hashKey := helper.HashKey(keyBytes, ht.buckets)
	valueBytes := helper.IntToByte(value)
	bucket := ht.table[hashKey]

	bucket.Append(keyBytes, valueBytes)
	ht.inserts += 1
}

func (ht *HashTable) Delete(key uint64) (int, bool) {
	keyBytes := helper.IntToByte(key)
	hashKey := helper.HashKey(keyBytes, ht.buckets)
	bucket := ht.table[hashKey]
	count, status := bucket.Delete(keyBytes)
	ht.deletes -= uint64(count)

	return count, status
}

func (ht *HashTable) Get(key uint64) ([]uint64, bool) {
	keyBytes := helper.IntToByte(key)
	hashKey := helper.HashKey(keyBytes, ht.buckets)
	bucket := ht.table[hashKey]

	values, status := bucket.Get(keyBytes)
	n := len(values)
	objects := make([]uint64, n)
	for i := 0; i < n; i++ {
		objects[i] = helper.ByteToInt(values[i])
	}

	return objects, status
}

func (ht *HashTable) Grow() *HashTable {
	var newBuckets = uint64(ht.buckets * 2)

	// Do note shrink below limit
	if newBuckets <= config.HASHTABLE_GROW_LIMIT {
		return ht
	} else {
		return ht.Resize(newBuckets)
	}
}

func (ht *HashTable) Shrink() *HashTable {
	var newBuckets = uint64(ht.buckets / 2)

	// Do note shrink below limit
	if newBuckets <= config.HASHTABLE_SHRINK_LIMIT {
		return ht
	} else {
		return ht.Resize(newBuckets)
	}
}

func (ht *HashTable) Resize(size uint64) *HashTable {
	var newHt = CreateHashTable(size)

	// iterate over all values and insert into new hash table

	return newHt
}
