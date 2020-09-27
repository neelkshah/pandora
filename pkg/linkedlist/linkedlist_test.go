// Author: Neel Shah, 2020
// linkedlist_test.go contains unit tests for the CRUD methods for the linked-list ADT
// to be used as part of the hash table ADT.

package linkedlist

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/neelkshah/pandora/pkg/helper"

	"github.com/neelkshah/pandora/config"
)

// Test create function.
func TestCreateLinkedList(t *testing.T) {
	var response = *CreateLinkedList()
	if response.head != nil || response.tail != nil || response.count != 0 {
		t.Fail()
	}
}

// Test append function for single value append.
func TestAppend(t *testing.T) {
	var response = *CreateLinkedList()
	appendError := response.Append("a", 5)
	if appendError != nil ||
		response.head == nil ||
		response.tail == nil ||
		response.count != 1 ||
		response.head != response.tail ||
		helper.ByteToInt(response.head.values[0].value) != 5 ||
		len(response.head.values) != 1 {
		t.Fatalf("Failed to append single key-value pair to empty linked list.")
	}
}

// Test whether fatness is maintained during append.
func TestFatness(t *testing.T) {
	var list = CreateLinkedList()
	var N = 3 * config.NODEFATNESS
	for i := 1; i <= N; i++ {
		var key = rand.Uint64()
		var value = rand.Uint64()
		if appendError := list.Append(key, value); appendError != nil {
			t.Fatalf("Error in appending key-value pair %v, %v to linkedlist.\nInner error: %v", key, value, appendError)
		}
	}
	var responselist = *list
	if (responselist).count != 3*config.NODEFATNESS ||
		responselist.head == responselist.tail ||
		responselist.head == nil ||
		responselist.tail == nil {
		t.Fatalf("Error in appending to fat nodes.")
	}
	var currentNode = *responselist.head
	for {
		if len(currentNode.values) != config.NODEFATNESS {
			t.Fatalf("Error in maintaining fatness.")
		}
		if currentNode.nextNode == nil {
			return
		}
		currentNode = *currentNode.nextNode
	}
}

// Test get
func TestGet(t *testing.T) {
	var list = CreateLinkedList()
	if appendError := list.Append(5, 7); appendError != nil {
		t.Fatalf("Error in appending key-value pair to linkedlist.\nInner error: %v", appendError)
	}
	if result, found, _ := list.Get(5); !found || !bytes.Equal(result[0], helper.IntToByte(7)) {
		t.Fatalf("Value present in linkedlist, but not found.")
	}
	if result, found, _ := list.Get(helper.IntToByte(7)); found || len(result) != 0 {
		t.Fatalf("Value not present in linkedlist, returned as found.")
	}
}

// Test get behaviour for an empty LinkedList instance
func TestGetNilValidation(t *testing.T) {
	var response = CreateLinkedList()
	if result, found, _ := response.Get(helper.IntToByte(5)); found || result != nil {
		t.Fatalf("Getting values from empty linkedlist.")
	}
}

// Test delete behaviour for an empty LinkedList instance
func TestDeleteNilValidation(t *testing.T) {
	var response = CreateLinkedList()
	if count, errorValue := response.Delete(helper.IntToByte(5)); errorValue == nil || count != 0 {
		t.Fatalf("Deleting values from empty linkedlist.")
	}
}

// Test deletion
func TestDelete(t *testing.T) {
	var response = CreateLinkedList()
	for i := 1; i <= 5; i++ {
		var key = rand.Uint64()
		var value = rand.Uint64()
		if appendResponse := response.Append(key, value); appendResponse != nil {
			t.Fatalf("Error in appending key-value pair to linkedlist.")
		}
	}
	response.Append(5, 7)
	for i := 1; i <= 5; i++ {
		var key = rand.Uint64()
		var value = rand.Uint64()
		if appendResponse := response.Append(key, value); appendResponse != nil {
			t.Fatalf("Error in appending key-value pair to linkedlist.")
		}
	}

	if count, errorValue := response.Delete(5); count < 1 || errorValue != nil {
		t.Fatalf("Error in deleting key-value pair.")
	}
}
