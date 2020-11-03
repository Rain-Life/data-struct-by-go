package main

import (
	"data-struct-by-go/queue/linkListQueue"
	"fmt"
)

func main() {
	queue := linkListQueue.Queue{}
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	queue.Enqueue("d")
	queue.Enqueue("e")
	queue.Traverse()
	queue.Dequeue()
	queue.Dequeue()
	fmt.Println()
	queue.Traverse()
}
