package hashtable

import (
	"github.com/neelkshah/pandora/config"
	"github.com/neelkshah/pandora/pkg/helper"
	ll "github.com/neelkshah/pandora/pkg/linkedlist"
)

type HashTable struct {
	table   []*ll.LinkedList
	buckets uint32
	inserts uint32
	deletes uint32
}

func CreateHashTable() *HashTable {
	var newHashTable = HashTable{table: make([]*ll.LinkedList, config.HASHTABLE_INIT_SIZE), buckets: uint32(config.HASHTABLE_INIT_SIZE), inserts: 0, deletes: 0}
	for i := 0; i < config.HASHTABLE_INIT_SIZE; i++ {
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
	ht.deletes -= uint32(count)

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
	// TODO: iterate values and convert them back to ints before returning

	return objects, status
}

func Expand() {
	// TODO: Requires method to iterate on LinkedList
}

func Contract() {
	// TODO: Requires method to iterate on LinkedList
}
