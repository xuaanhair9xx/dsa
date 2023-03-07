package main

import (
	"fmt"
	"errors"
)

type Stack struct {
    Elements []interface {}
    Size     int
}

func (s *Stack) Push(elem interface {}) {
    if s.GetLength() == s.Size && s.Size != 0 {
        fmt.Println("Overflow")
        return
    }
    s.Elements = append(s.Elements, elem)
}

func (s *Stack) Pop() interface {} {
    if s.IsEmpty() {
        fmt.Println("UnderFlow")
        return nil
    }
    element := s.Elements[s.GetLength()-1]
    if s.GetLength() == 1 {
        s.Elements = nil
        return element
    }
    s.Elements = s.Elements[:s.GetLength()-1]
    return element
}

func (s *Stack) GetLength() int {
    return len(s.Elements)
}
  
func (s *Stack) IsEmpty() bool {
    return len(s.Elements) == 0
}
  
func (s *Stack) Peek() ( interface {}, error) {
    if s.IsEmpty() {
        return nil, errors.New("Empty stack")
    }
    return s.Elements[0], nil
}

type Node struct {
	Dd int
}

func main() {
	
	stack := Stack { Size: 30 }
	stack.Push(34)
	stack.Push("122")
	stack.Push(Node {Dd: 123})

	d, okd :=  stack.Pop().(Node)
	c, okc :=  stack.Pop().(string)
	e, oke :=  stack.Pop().(int)

	fmt.Println(okd)
	fmt.Println(d)
	fmt.Println(okc)
	fmt.Println(c)
	fmt.Println(oke)
	fmt.Println(e)
}