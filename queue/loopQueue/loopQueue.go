package loopQueue

import "fmt"

type Item interface {}

const QueueSize = 5
type LoopQueue struct {
	Items [QueueSize]Item
	Head int
	Tail int
}

//init
func (queue *LoopQueue) Init() {
	queue.Head = 0
	queue.Tail = 0
}

//enqueue
func (queue *LoopQueue) Enqueue(data Item) {
	if ((queue.Tail + 1) % QueueSize) == queue.Head {
		fmt.Println("队列满了")
	}

	queue.Items[queue.Tail] = data
	queue.Tail = (queue.Tail+1) % QueueSize
}

//dequeue
func (queue *LoopQueue) Dequeue() Item {
	if queue.Head == queue.Tail {
		fmt.Println("队列空了")
		return nil
	}

	item := queue.Items[queue.Head]
	queue.Head = (queue.Head + 1) % QueueSize

	return item
}
