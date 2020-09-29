package hashtable

import (
	"github.com/neelkshah/pandora/config"
	"github.com/neelkshah/pandora/pkg/helper"
	ll "github.com/neelkshah/pandora/pkg/linkedlist"
)

type HashTable struct {
	table     []*ll.LinkedList
	hashFunc  func([]byte, uint64) uint64
	buckets   uint64
	occupancy uint64
}

// Create hashtable with number of buckets given by size
func Create(size uint64) *HashTable {
	var newHashTable = HashTable{table: make([]*ll.LinkedList, size), buckets: size, occupancy: 0}
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
	ht.occupancy++

	// if hash table has exceeded grow ratio grow it
	if ht.growCheck() {
		ht = ht.Grow()
	}
}

func (ht *HashTable) growCheck() bool {
	return float64(ht.occupancy)/float64(ht.buckets) > config.HASHTABLE_GROW_RATIO
}

// Delete all values for the given key from
// hashtable. Return number of values deleted
// and true if no values were deleted
func (ht *HashTable) Delete(key uint64) (int, bool) {
	keyBytes := helper.IntToByte(key)
	hashKey := helper.HashKey(keyBytes, ht.buckets)
	bucket := ht.table[hashKey]
	count, status := bucket.Delete(keyBytes)
	ht.occupancy -= uint64(count)

	// if hash table has falls below shrink ratio shrink it
	if ht.shrinkCheck() {
		ht.Shrink()
	}

	return count, status
}

func (ht *HashTable) shrinkCheck() bool {
	return float64(ht.occupancy)/float64(ht.buckets) < config.HASHTABLE_SHRINK_RATIO
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

	// Do not shrink below limit
	if newBuckets > config.HASHTABLE_GROW_LIMIT {
		ht.Resize(newBuckets)
	}
}

// Shrink shrinks the current hashtable by a constant factor
// it returns reference to new hashtable
func (ht *HashTable) Shrink() {
	var newBuckets uint64 = uint64(float64(ht.buckets) * config.HASHTABLE_SHRINK_FACTOR)

	// Do not shrink below limit
	if newBuckets > config.HASHTABLE_SHRINK_LIMIT {
		ht.Resize(newBuckets)
	}
}

// Resize given hashtable to new hashtable of given size
// copies all values from existing hash table to new hashtable
// returns reference to new hashtable
func (ht *HashTable) Resize(size uint64) {
	// var newHt = Create(size)
	var oldTable = ht.table
	var oldBuckets = ht.buckets

	var newTable = make([]*ll.LinkedList, size)
	for i := 0; i < int(size); i++ {
		newTable[i] = ll.CreateLinkedList()
	}

	ht.buckets = size
	ht.occupancy = 0
	ht.table = newTable

	// // iterate over all values and insert into new hash table
	// for i := 0; i < oldBuckets; i++ {
	// 	iter = oldTable[i].iterator()
	// 	while iter.hasNext() {
	// 		key, value = iter.next()
	// 		ht.Insert(key, value)
	// 	}
	// }
}
