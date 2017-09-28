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

func NewElement(data string) *element {
	elem := new(element)
	elem.data = data
	elem.next = nil
	return elem
}

func main() {
	// Reads number os test cases
	var numTests int
	fmt.Scanf("%d", &numTests)

	// Runs each test case
	for i:=1; i<= numTests; i++ {
	    // Create Stack
    	stk := NewStack()
		
        // Read brackets input
		var test string
		fmt.Scanln(&test)
		
        result := "YES"
        
        for j := range test {
			data := string(test[j])
			switch data {
			case "(":
				stk.Push(NewElement(")"))
			case "[":
				stk.Push(NewElement("]"))
			case "{":
				stk.Push(NewElement("}"))
			default:
				popped := stk.Pop()
				if popped != nil && popped.data != data {
					result = "NO"
					break
				}
			}
		}

		if stk.head != nil {
			result = "NO"
		}
        fmt.Println(result)
	}

}