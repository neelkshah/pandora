// Author: Neel Shah, 2020
// linkedlist_api.go provides the API for the CRUD methods for the linked-list ADT
// to be used as part of the hash table ADT.

package linkedlist

// CreateLinkedList returns a pointer to a new empty linked list instance.
func CreateLinkedList() *LinkedList {
	return createImpl()
}

// Append appends a key value pair to the linked list after checking their respective types.
func (linkedlist *LinkedList) Append(iKey interface{}, iValue interface{}) error {

	pair, conversionError := convert(iKey, iValue)

	if conversionError != nil {
		return conversionError
	}

	linkedlist.appendImpl(pair[0], pair[1])
	return nil
}

// Get gets the value array of values associated with the passed key.
func (linkedlist *LinkedList) Get(iKey interface{}) ([][]byte, bool, error) {

	key, conversionError := convert(iKey)

	if conversionError != nil {
		return nil, false, conversionError
	}

	valueArray, found := linkedlist.getImpl(key[0])

	return valueArray, found, nil
}

// Delete deletes the key-value pairs associated with the passed key. It returns the count of deleted pairs and any error.
func (linkedlist *LinkedList) Delete(iKey interface{}) (int, error) {
	key, conversionError := convert(iKey)

	if conversionError != nil {
		return 0, conversionError
	}

	deletedCount, deletionError := linkedlist.deleteImpl(key[0])

	return deletedCount, deletionError
}

// Head returns the head node of the linked list.
func (linkedlist *LinkedList) Head() *Node {
	return linkedlist.head
}

// Next returns the next node in the linked list.
func (node *Node) Next() *Node {
	return node.nextNode
}
