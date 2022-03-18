package main

import "fmt"

// Each linked list item maintains a pointer to the next item in the list
type list struct {
	item int
	next *list
}

func main() {
	l := list{1, nil}
	l.Add(2)
	fmt.Printf("%v", l)
}

// Linked lists support the following operations:
// - Searching the list for an item
// - Add a new item to the list
// - Delete an item from the list
func (l *list) Add(item int) {
	new := list{item, l}
	l = new
}

func (l *list) Search(i int) bool {
	if l == nil {
		return false
	}

	if l.item == i {
		return true
	} else {
		return l.next.Search(i)
	}
}

func (l *list) Delete(item int) {

}
