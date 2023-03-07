package main

import (
	"fmt"
	"errors"
)

type Queue struct {
    Elements []interface {}
    Size     int
}

func (q *Queue) Enqueue(elem interface {}) {
    if q.GetLength() == q.Size && q.Size != 0 {
        fmt.Println("Overflow")
        return
    }
    q.Elements = append(q.Elements, elem)
}
  
func (q *Queue) Dequeue() interface {} {
    if q.IsEmpty() {
        fmt.Println("UnderFlow")
        return nil
    }
    element := q.Elements[0]
    if q.GetLength() == 1 {
        q.Elements = nil
        return element
    }
    q.Elements = q.Elements[1:]
    return element
}
  
func (q *Queue) GetLength() int {
    return len(q.Elements)
}
  
func (q *Queue) IsEmpty() bool {
    return len(q.Elements) == 0
}
  
func (q *Queue) Peek() ( interface {}, error) {
    if q.IsEmpty() {
        return nil, errors.New("empty queue")
    }
    return q.Elements[0], nil
}

type Node struct {
	Dd int
}

func main() {
	queue := Queue { Size: 30 }
	queue.Enqueue(34)
	queue.Enqueue(122)
	queue.Enqueue(Node {Dd: 123})
	d, ok :=  queue.Dequeue().(int)
	c, oka :=  queue.Dequeue().(string)
	e, oka :=  queue.Dequeue().(Node)
	fmt.Println(ok)
	fmt.Println(oka)
	// if (d > c) {
	// 	fmt.Println(123123)
	// }
	fmt.Printf("Deq %v \n", e.Dd)
	fmt.Printf("Deq %v \n", c)
	fmt.Printf("Deq %v \n", d)
}