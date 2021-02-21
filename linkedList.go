package main

import "fmt"

// Node an element in a list
type Node struct {
	content interface{}
	next    *Node
	prev    *Node
}

// LinkedList doubly linked list
type LinkedList struct {
	head *Node
	tail *Node
}

// Push an element to the linked list
func (li *LinkedList) Push(item interface{}) {
	n := Node{content: item, next: nil, prev: nil}
	// First element in the list

	if li.head == nil {
		li.head = &n
		li.tail = &n
		return
	}

	lastNode := li.tail
	n.prev = lastNode
	lastNode.next = &n
	li.tail = &n
}

// Pop removes the last element of the list and returns it
func (li *LinkedList) Pop() interface{} {
	if li.head == nil {
		return nil
	}

	curTail := li.tail
	li.tail = curTail.prev
	li.tail.next = nil

	return curTail.content
}

// Shift removes the first element  of the list and returns it
func (li *LinkedList) Shift() interface{} {
	if li.head == nil {
		return nil
	}

	head := li.head
	li.head = head.next
	li.head.prev = nil
	return head.content
}

// Unshift adds an element to the beginning of the list
func (li *LinkedList) Unshift(item interface{}) {

	curHead := li.head
	n := Node{content: item, next: curHead, prev: nil}

	curHead.prev = &n
	li.head = &n
}

// Get returns the element at the given position
func (li *LinkedList) Get(pos int) *Node {
	if pos < 1 || li.head == nil {
		return nil
	}

	curNode := li.head
	curPos := 1

	if curPos == pos {
		return curNode
	}

	for curPos < pos {
		curNode = curNode.next
		curPos++
	}
	return curNode
}

// Traverse visits all elements of the linked list
func (li *LinkedList) Traverse() {
	curNode := li.head
	for curNode.next != nil {
		fmt.Printf("ADDRESS: '%p' NODE: '%+v'\n", curNode, curNode)
		curNode = curNode.next
	}
	fmt.Printf("ADDRESS: '%p' NODE: '%+v'\n", curNode, curNode)
}

// TraverseReverse visits all elements of the linked list in reverse order
func (li *LinkedList) TraverseReverse() {
	curNode := li.tail
	for curNode.prev != nil {
		fmt.Printf("ADDRESS: '%p' NODE: '%+v'\n", curNode, curNode)
		curNode = curNode.prev
	}
	fmt.Printf("ADDRESS: '%p' NODE: '%+v'\n", curNode, curNode)
}
