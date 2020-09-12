package linkedlist

import "fmt"

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
