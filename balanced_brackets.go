// https://www.hackerrank.com/challenges/ctci-balanced-brackets
package main

import "fmt"

type element struct {
	data string
	next *element
}

type stack struct {
	head *element
}

func (stk *stack) Push(elem *element) {
	
	elem.next = stk.head
	stk.head = elem
}

func (stk *stack) Pop() *element {
	tmp := stk.head
	stk.head = stk.head.next
	return tmp
}

func (stk *stack) Print() {
	e := stk.head
	for e != nil {
		fmt.Println(e.data)
		e = e.next
	}
}


func NewStack() *stack {
	stk := new(stack)
	stk.head = nil
	return stk
}

func NewElement(data string) *element {
	elem := new(element)
	elem.data = data
	elem.next = nil
	return elem
}

func main() {
	stk := NewStack()
	stk.Push(NewElement("a"))
	stk.Push(NewElement("b"))
	stk.Push(NewElement("c"))
	stk.Push(NewElement("d"))
	stk.Push(NewElement("e"))
	stk.Print()
}