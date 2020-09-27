// Author: Neel Shah, 2020
// linkedlist.go contains the CRUD methods for the linked-list ADT
// to be used as part of the hash table ADT.

package linkedlist

import (
	"bytes"
	"fmt"

	"github.com/neelkshah/pandora/config"
)

// LinkedList is the data structure used for the hash table.
type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

// Node is the data structure that forms the LinkedList.
type Node struct {
	Values   []valueNode
	nextNode *Node
}

// ValueNode contains a single value to be stored in the LinkedList.
type valueNode struct {
	Key   []byte
	Value []byte
}

// createImpl returns a pointer to a new empty linked list instance.
func createImpl() *LinkedList {
	linkedList := LinkedList{head: nil, tail: nil, count: 0}
	return &linkedList
}

// AppendImpl appends a given value to the end of the referenced LinkedList instance.
func (linkedList *LinkedList) appendImpl(key []byte, value []byte) {
	var newValue = valueNode{Key: key, Value: value}
	if linkedList.count == 0 {
		var newNode = Node{Values: []valueNode{newValue}, nextNode: nil}
		linkedList.head = &newNode
		linkedList.tail = &newNode
		linkedList.count = 1
		return
	}
	var tailNode = linkedList.tail
	if len(tailNode.Values) < config.NODEFATNESS {
		tailNode.Values = append(tailNode.Values, newValue)
	} else {
		var newNode = Node{Values: []valueNode{newValue}, nextNode: nil}
		tailNode.nextNode = &newNode
		linkedList.tail = &newNode
	}
	linkedList.count++
}

// GetImpl returns the values associate with the key.
func (linkedList *LinkedList) getImpl(key []byte) ([][]byte, bool) {
	if linkedList == nil || linkedList.head == nil {
		return nil, true
	}
	var currentNode = linkedList.head
	var result = make([][]byte, 0)
	for {
		if currentNode == nil {
			return result, len(result) != 0
		}
		for _, vnode := range (*currentNode).Values {
			if bytes.Equal(vnode.Key, key) {
				result = append(result, vnode.Value)
			}
		}
		currentNode = currentNode.nextNode
	}
}

// DeleteImpl deletes all key-value pairs having the key passed as parameter.
// It returns the number of deleted pairs and any error.
func (linkedList *LinkedList) deleteImpl(key []byte) (int, error) {
	if linkedList == nil || linkedList.head == nil {
		return 0, fmt.Errorf("The linked list is empty")
	}
	var currentNode = linkedList.head
	var count = 0
	var k = 0
	for {
		if currentNode == nil {
			break
		}
		for _, vnode := range (*currentNode).Values {
			if !bytes.Equal(vnode.Key, key) {
				(*currentNode).Values[k] = vnode
				k++
				continue
			}
			count++
		}
		(*currentNode).Values = (*currentNode).Values[:k]
		currentNode = currentNode.nextNode
	}
	return count, nil
}
