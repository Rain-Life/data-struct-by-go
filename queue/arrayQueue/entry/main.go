package main

import (
	"data-struct-by-go/queue/arrayQueue"
	"fmt"
)

func main() {
	queue := arrayQueue.Queue{Length: 10}
	queue.Init()
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	queue.Enqueue("d")
	queue.Enqueue("e")

	for _, value := range queue.Queue {
		fmt.Printf("%v\t", value)
	}
	fmt.Println()
	queue.Dequeue()
	queue.Dequeue()
	for _, value := range queue.Queue {
		fmt.Printf("%v\t", value)
	}
}
