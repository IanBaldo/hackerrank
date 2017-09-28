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
	stk := NewStack()
	var numTests int
	fmt.Scanf("%d", &numTests)

	for i:=1; i<= numTests; i++ {
     
		var test string
		fmt.Scanln(&test)
		
        result := "YES"
        
        for j := range test {
            data := string(test[j])
            if data == "(" || data == "[" || data == "{" {
                stk.Push(NewElement(data))
            } else {
                popped := stk.Pop()
                if popped != nil && (popped.data == "(" && data != ")") || (popped.data == "[" && data != "]") || (popped.data == "{" && data != "}"){
                    result = "NO"
                    
                }
            }
		}

        fmt.Println(result)
	}

}