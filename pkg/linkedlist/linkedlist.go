// Author: Neel Shah, 2020
// linkedlist.go contains the CRUD methods for the linked-list ADT
// to be used as part of the hash table ADT.

package linkedlist

import (
	"bytes"

	"github.com/neelkshah/pandora/config"
)

// LinkedList is the data structure used for the hash table.
type LinkedList struct {
	head  *linkedListNode
	tail  *linkedListNode
	count int
}

// linkedListNode is the data structure that forms the LinkedList.
type linkedListNode struct {
	values   []valueNode
	nextNode *linkedListNode
}

// valueNode contains a single value to be stored in the LinkedList.
type valueNode struct {
	key   []byte
	value []byte
}

// CreateLinkedList returns a pointer to a new empty linked list instance.
func CreateLinkedList() *LinkedList {
	linkedList := LinkedList{head: nil, tail: nil, count: 0}
	return &linkedList
}

// Append appends a given value to the end of the referenced LinkedList instance.
func (linkedList *LinkedList) Append(key []byte, value []byte) {
	var newValue = valueNode{key: key, value: value}
	if linkedList.count == 0 {
		var newNode = linkedListNode{values: []valueNode{newValue}, nextNode: nil}
		linkedList.head = &newNode
		linkedList.tail = &newNode
		linkedList.count = 1
		return
	}
	var tailNode = linkedList.tail
	if len(tailNode.values) < config.NODEFATNESS {
		tailNode.values = append(tailNode.values, newValue)
	} else {
		var newNode = linkedListNode{values: []valueNode{newValue}, nextNode: nil}
		tailNode.nextNode = &newNode
		linkedList.tail = &newNode
	}
	linkedList.count++
}

// Get returns the values associate with the key.
func (linkedList *LinkedList) Get(key []byte) ([][]byte, bool) {
	if linkedList == nil || linkedList.head == nil {
		return nil, true
	}
	var currentNode = linkedList.head
	var result = make([][]byte, 0)
	for {
		if currentNode == nil {
			break
		}
		for _, vnode := range (*currentNode).values {
			if bytes.Equal(vnode.key, key) {
				result = append(result, vnode.value)
			}
		}
		currentNode = currentNode.nextNode
	}
	return result, false
}

// Delete deletes all key-value pairs having the key passed as parameter.
// It returns the number of deleted pairs and a bool indicating occurrence of an error.
func (linkedList *LinkedList) Delete(key []byte) (int, bool) {
	if linkedList == nil || linkedList.head == nil {
		return 0, true
	}
	var currentNode = linkedList.head
	var count = 0
	var k = 0
	for {
		if currentNode == nil {
			break
		}
		for _, vnode := range (*currentNode).values {
			if !bytes.Equal(vnode.key, key) {
				(*currentNode).values[k] = vnode
				k++
				continue
			}
			count++
		}
		(*currentNode).values = (*currentNode).values[:k]
		currentNode = currentNode.nextNode
	}
	return count, false
}
