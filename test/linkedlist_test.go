// Author: Neel Shah, 2020
// linkedlist_test.go contains E2E tests for the CRUD methods for the linked-list ADT
// to be used as part of the hash table ADT.

package test

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/neelkshah/pandora/pkg/helper"
	"github.com/neelkshah/pandora/pkg/linkedlist"
)

func TestCreateLinkedList(t *testing.T) {
	var response = linkedlist.CreateLinkedList()
	if response == nil {
		t.Fatalf("Create is failing")
	}
}

func TestNext(t *testing.T) {
	var list = linkedlist.CreateLinkedList()
	list.Append(5, 7)
	list.Append(7, 5)
	var v1 = helper.IntToByte(5)
	var v2 = helper.IntToByte(7)
	var headValues = list.Head().Values
	if !bytes.Equal(headValues[0].Key, v1) ||
		!bytes.Equal(headValues[0].Value, v2) ||
		!bytes.Equal(headValues[1].Key, v2) ||
		!bytes.Equal(headValues[1].Value, v1) {
		t.Fatalf("Next is failing")
	}
}

func BenchmarkAppend(b *testing.B) {
	var response = linkedlist.CreateLinkedList()
	for i := 0; i < b.N; i++ {
		var key = helper.IntToByte(rand.Uint64())
		var value = helper.IntToByte(rand.Uint64())
		response.Append(key, value)
	}
}

func BenchmarkGet(b *testing.B) {
	var response = populatedList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		response.Get(helper.IntToByte(rand.Uint64()))
	}
}

func populatedList() *(linkedlist.LinkedList) {
	var response = linkedlist.CreateLinkedList()
	for i := 0; i < 10; i++ {
		var key = helper.IntToByte(rand.Uint64())
		var value = helper.IntToByte(rand.Uint64())
		response.Append(key, value)
	}
	return response
}
