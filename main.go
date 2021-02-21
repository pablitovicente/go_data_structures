package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	myLinkedList := LinkedList{}

	for i := 0; i < 5; i++ {
		myLinkedList.Push([]int{i, rnd()})
	}

	myLinkedList.Traverse()
	fmt.Println("")
	myLinkedList.TraverseReverse()
	fmt.Println("")

	lastElement := myLinkedList.Pop()
	fmt.Printf("Pop last element: %+v\n address %p", lastElement, &lastElement)
	fmt.Println("Element type:", reflect.TypeOf(lastElement))

	myLinkedList.Traverse()
	fmt.Println("")

	firstElement := myLinkedList.Shift()
	fmt.Println("Shift first element:", firstElement)
	fmt.Println("Element type:", reflect.TypeOf(firstElement))
	myLinkedList.Traverse()
	fmt.Println("")

	c := []float64{0, 3.14, 2.16, 0.9}
	fmt.Println("Unshift:", c)
	myLinkedList.Unshift(c)

	fmt.Println("")
	myLinkedList.Traverse()

	fmt.Println("")
	idx := 2
	found := myLinkedList.Get(idx)
	fmt.Printf("Node at position %v is %+v\n", idx, found)
	fmt.Println("Element type:", reflect.TypeOf(found))

	fmt.Println("")
	myLinkedList.TraverseReverse()
}

func rnd() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 1000
	return rand.Intn(max-min+1) + min
}
