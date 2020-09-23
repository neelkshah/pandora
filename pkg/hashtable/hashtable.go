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

// Create hashtable with number of buckets given by size
func Create(size uint64) *HashTable {
	var newHashTable = HashTable{table: make([]*ll.LinkedList, size), buckets: size, inserts: 0, deletes: 0}
	for i := 0; i < int(size); i++ {
		newHashTable.table[i] = ll.CreateLinkedList()
	}

	return &newHashTable
}

// Insert value into hashtable
func (ht *HashTable) Insert(key uint64, value uint64) {
	keyBytes := helper.IntToByte(key)
	hashKey := helper.HashKey(keyBytes, ht.buckets)
	valueBytes := helper.IntToByte(value)
	bucket := ht.table[hashKey]

	bucket.Append(keyBytes, valueBytes)
	ht.inserts++
}

// Delete all values for the given key from
// hashtable. Return number of values deleted
// and true if no values were deleted
func (ht *HashTable) Delete(key uint64) (int, bool) {
	keyBytes := helper.IntToByte(key)
	hashKey := helper.HashKey(keyBytes, ht.buckets)
	bucket := ht.table[hashKey]
	count, status := bucket.Delete(keyBytes)
	ht.deletes += uint64(count)

	return count, status
}

// Get all the values for the given key stored
// in the table. Return array of values and status
// true if there are no value to return
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

// Grow grows the current hashtable by a constant factor
// it returns reference to new hashtable
func (ht *HashTable) Grow() *HashTable {
	var newBuckets uint64 = uint64(float64(ht.buckets) * config.HASHTABLE_GROW_FACTOR)

	// Do note shrink below limit
	if newBuckets <= config.HASHTABLE_GROW_LIMIT {
		return ht
	} else {
		return ht.Resize(newBuckets)
	}
}

// Shrink shrinks the current hashtable by a constant factor
// it returns reference to new hashtable
func (ht *HashTable) Shrink() *HashTable {
	var newBuckets uint64 = uint64(float64(ht.buckets) * config.HASHTABLE_SHRINK_FACTOR)

	// Do note shrink below limit
	if newBuckets <= config.HASHTABLE_SHRINK_LIMIT {
		return ht
	} else {
		return ht.Resize(newBuckets)
	}
}

// Resize given hashtable to new hashtable of given size
// copies all values from existing hash table to new hashtable
// returns reference to new hashtable
func (ht *HashTable) Resize(size uint64) *HashTable {
	var newHt = Create(size)

	// iterate over all values and insert into new hash table

	return newHt
}
