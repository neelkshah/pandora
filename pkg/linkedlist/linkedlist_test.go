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
	response.Append([]byte("a"), helper.IntToByte(5))
	if response.head == nil ||
		response.tail == nil ||
		response.count != 1 ||
		response.head != response.tail ||
		helper.ByteToInt(response.head.values[0].value) != 5 ||
		len(response.head.values) != 1 {
		t.Fail()
	}
}

// Test whether fatness is maintained during append.
func TestFatness(t *testing.T) {
	var response = CreateLinkedList()
	var N = 3 * config.NODEFATNESS
	for i := 1; i <= N; i++ {
		var key = helper.IntToByte(rand.Uint64())
		var value = helper.IntToByte(rand.Uint64())
		response.Append(key, value)
	}
	var responselist = *response
	if (responselist).count != 3*config.NODEFATNESS ||
		responselist.head == responselist.tail ||
		responselist.head == nil ||
		responselist.tail == nil {
		t.Fail()
	}
	var currentNode = *responselist.head
	for {
		if len(currentNode.values) != config.NODEFATNESS {
			t.Fail()
		}
		if currentNode.nextNode == nil {
			return
		}
		currentNode = *currentNode.nextNode
	}
}

// Test get
func TestGet(t *testing.T) {
	var response = CreateLinkedList()
	response.Append(helper.IntToByte(5), helper.IntToByte(7))
	if result, isEmpty := response.Get(helper.IntToByte(5)); isEmpty == true || !bytes.Equal(result[0], helper.IntToByte(7)) {
		t.Fail()
	}
	if result, isEmpty := response.Get(helper.IntToByte(7)); isEmpty == true || len(result) != 0 {
		t.Fail()
	}
}

// Test get behaviour for an empty LinkedList instance
func TestGetNilValidation(t *testing.T) {
	var response = CreateLinkedList()
	if result, isEmpty := response.Get(helper.IntToByte(5)); isEmpty == false || result != nil {
		t.Fail()
	}
}

// Test delete behaviour for an empty LinkedList instance
func TestDeleteNilValidation(t *testing.T) {
	var response = CreateLinkedList()
	if count, isEmpty := response.Delete(helper.IntToByte(5)); isEmpty == false || count != 0 {
		t.Fail()
	}
}

// Test deletion
func TestDelete(t *testing.T) {
	var response = CreateLinkedList()
	for i := 1; i <= 5; i++ {
		var key = helper.IntToByte(rand.Uint64())
		var value = helper.IntToByte(rand.Uint64())
		response.Append(key, value)
	}
	response.Append(helper.IntToByte(5), helper.IntToByte(7))
	for i := 1; i <= 5; i++ {
		var key = helper.IntToByte(rand.Uint64())
		var value = helper.IntToByte(rand.Uint64())
		response.Append(key, value)
	}

	if count, status := response.Delete(helper.IntToByte(5)); count < 1 || status == true {
		t.Fail()
	}
}
