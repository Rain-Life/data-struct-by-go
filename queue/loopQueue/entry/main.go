package main

import (
	loopQueue2 "data-struct-by-go/queue/loopQueue"
	"fmt"
)

func main() {
	loopQueue := loopQueue2.LoopQueue{}
	loopQueue.Init()

	loopQueue.Enqueue("1")
	loopQueue.Enqueue("2")
	loopQueue.Enqueue("3")
	loopQueue.Enqueue("4")
	loopQueue.Enqueue("5")

	fmt.Printf("%v\t", loopQueue.Dequeue())
	fmt.Printf("%v\t", loopQueue.Dequeue())
	fmt.Printf("%v\t", loopQueue.Dequeue())
	fmt.Printf("%v\t", loopQueue.Dequeue())
	fmt.Printf("%v\t", loopQueue.Dequeue())
}
