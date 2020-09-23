package hashtable

import (
	"bytes"
	"testing"

	"github.com/neelkshah/pandora/pkg/helper"
)

func TestCreate(t *testing.T) {
	var response = *Create(10)
	if response.buckets != 10 || response.inserts != 0 || response.deletes != 0 {
		t.Fatalf("Fail at values")
	}

	for i := 0; i < int(response.buckets); i++ {
		if response.table[i] == nil {
			t.Fatalf("Fail at buckets")
		}
	}
}

func TestInsert(t *testing.T) {
	// Create and insert value in hashtable
	var key, value uint64 = 5, 10
	var hashTable = *Create(10)
	hashTable.Insert(5, 10)

	// checked linkedlist for successful insertion
	var keyBytes = helper.IntToByte(uint64(key))
	var hashKey = helper.HashKey(keyBytes, hashTable.buckets)
	valueBytes := helper.IntToByte(value)
	bucket := hashTable.table[hashKey]
	response, status := bucket.Get(keyBytes)

	// checks
	// TODO: A better way to map tests to error messages?
	if status {
		t.Fatalf("LinkedList does not contain anything insertion failed")
	}

	if len(response) != 1 {
		t.Fatalf("Linkedlist should contain only one value after single insert")
	}

	if bytes.Compare(response[0], valueBytes) != 0 {
		t.Fatalf("Hashtable did not store correct value")
	}

	if hashTable.inserts != 1 {
		t.Fatalf("Insert count failed")
	}
}

func TestGet(t *testing.T) {
	// create hashtable and insert value into underlying linkedlist
	var key, value, size uint64 = 5, 10, 10
	var hashTable = *Create(size)
	var keyBytes = helper.IntToByte(key)
	var hashKey = helper.HashKey(keyBytes, hashTable.buckets)
	valueBytes := helper.IntToByte(value)
	bucket := hashTable.table[hashKey]
	bucket.Append(keyBytes, valueBytes)

	// get value from linkedlist
	var response, status = hashTable.Get(key)

	// checks
	// TODO: A better way to map tests to error messages?
	if status {
		t.Fatalf("Get failed")
	}

	if len(response) != 1 {
		t.Fatalf("Get should return a single value. It returned %d values.", len(response))
	}

	if response[0] != value {
		t.Fatalf("Get did not return the correct value. Value should be %d is %d", value, response[0])
	}
}

func TestDelete(t *testing.T) {
	// create hashtable and insert value into underlying linkedlist
	var key, value, size uint64 = 5, 10, 10
	var hashTable = *Create(size)
	var keyBytes = helper.IntToByte(key)
	var hashKey = helper.HashKey(keyBytes, hashTable.buckets)
	valueBytes := helper.IntToByte(value)
	bucket := hashTable.table[hashKey]
	bucket.Append(keyBytes, valueBytes)

	// delete value from hash table and check underlying linkedlist
	var response, status = hashTable.Delete(key)

	// checks
	if status {
		t.Fatalf("Delete failed")
	}

	if response != 1 {
		t.Fatalf("Deleted %d values should have deleted 1 value", response)
	}

	if hashTable.deletes != 1 {
		t.Fatalf("Delete counter failed. It counted %d but should have counted 1", hashTable.deletes)
	}
}
