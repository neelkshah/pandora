// Author: Neel Shah, 2020
// linkedlist_api.go provides the API for the CRUD methods for the linked-list ADT
// to be used as part of the hash table ADT.

package linkedlist

// Append appends a key value pair to the linked list after checking their respective types
func (linkedlist *LinkedList) Append(iKey interface{}, iValue interface{}) error {

	pair, conversionError := convert(iKey, iValue)

	if conversionError != nil {
		return conversionError
	}

	linkedlist.AppendImpl(pair[0], pair[1])
	return nil
}

// Get gets the value array of values associated with the passed key.
func (linkedlist *LinkedList) Get(iKey interface{}) ([][]byte, bool, error) {

	key, conversionError := convert(iKey)

	if conversionError != nil {
		return nil, false, conversionError
	}

	valueArray, found := linkedlist.GetImpl(key[0])

	return valueArray, found, nil
}

// Delete deletes the key-value pairs associated with the passed key. It returns the count of deleted pairs and any error.
func (linkedlist *LinkedList) Delete(iKey interface{}) (int, error) {
	key, conversionError := convert(iKey)

	if conversionError != nil {
		return 0, conversionError
	}

	deletedCount, deletionError := linkedlist.DeleteImpl(key[0])

	return deletedCount, deletionError
}
