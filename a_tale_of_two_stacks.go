// https://www.hackerrank.com/challenges/ctci-queue-using-two-stacks

package main

import "fmt"

type element struct {
	data int
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
	if stk.head == nil {
		return nil
	}
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

func NewElement(data int) *element {
	elem := new(element)
	elem.data = data
	elem.next = nil
	return elem
}

//##################  PROGRAM STARTS HERE  ######################
func main() {
	mainStack := new(stack)

	var numQuerys int
	fmt.Scanf("%d", &numQuerys)

	for i:=0; i < numQuerys; i++ {
		op, _ := myScanln()
		
		switch op[0] {
		case 1:
			// Enqueue
			mainStack.Push(NewElement(op[1]))
		case 2:
			//Dequeue
			mainStack.Pop()
		case 3:
			// Print
			printFirstElem(mainStack)
		}
	}

}

// TODO - function to read integers
func myScanln() ([]int, error) {
	vals := make([]int, 3)
	y := make([]interface{}, 3)

	for i:= range vals {
		y[i] = &vals[i]
	}

	n, err := fmt.Scanln(y...)
	vals = vals[:n]

	return vals, err
}

// TODO- function to print
func printFirstElem(stk *stack) {
	aux := new(stack)

	elem := stk.head
	for elem != nil {
		aux.Push(NewElement(elem.data))
		elem = elem.next
	}

	fmt.Println(aux.Pop().data)
}