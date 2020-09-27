package linkedlist

import (
	"fmt"

	"github.com/neelkshah/pandora/pkg/helper"
)

// Print prints the contents of a LinkedList instance.
func (linkedList *LinkedList) Print() {
	fmt.Printf("Count of elements: %v\n", linkedList.count)
	var currentNode = *linkedList.head
	for {
		for _, element := range currentNode.values {
			fmt.Printf("%v, %v\t", element.key, element.value)
		}
		if currentNode.nextNode == nil {
			return
		}
		currentNode = *currentNode.nextNode
		fmt.Println()
	}
}

// convert converts the given slice of values into a slice of byte slices.
func convert(originalValues ...interface{}) ([][]byte, error) {
	var values [][]byte

	for _, originalValue := range originalValues {
		switch keyType := originalValue.(type) {
		case int:
			values = append(values, helper.IntToByte(uint64(originalValue.(int))))
		case uint64:
			values = append(values, helper.IntToByte(originalValue.(uint64)))
		case string:
			values = append(values, []byte(originalValue.(string)))
		default:
			return nil, fmt.Errorf("%v is of invalid type %v", originalValue, keyType)
		}
	}

	return values, nil
}
