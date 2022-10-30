package main

import "fmt"

func main() {
	q := queue{}
	q.push(1, 2, 3, 4)
	fmt.Println(q, q.size(), q.peek())
	fmt.Println(q.pop(), q.size(), q)
}

/*
A queue ds implementation, a FIFO implementation
grow as the data is pushed
size will give the size of the queue
will squash the underlying queue - ?

major functions:
- push => pushes the data to the queue
- pop  => pops the
- peek

	H - - - - T

	push pushes to tail of queue
*/
type queue struct {
	data []int
}

func (q *queue) push(ele ...int) {
	(*q).data = append((*q).data, ele...)
}

func (q *queue) pop() int {
	ele := q.data[0]
	q.data = q.data[1:]
	return ele
}

func (q *queue) size() int {
	return len(q.data)
}

func (q *queue) peek() int {
	return q.data[0]
}
