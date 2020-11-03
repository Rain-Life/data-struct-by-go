package arrayQueue

import "fmt"

type Item interface {}

type Queue struct {
	Queue []Item
	Length int
}

func (queue *Queue) Init() {
	queue.Queue = []Item{}
}

func (queue *Queue) Enqueue(data Item) {
	if len(queue.Queue) > queue.Length {
		fmt.Println("队列满了")
		return
	}
	queue.Queue = append(queue.Queue, data)
}

func (queue *Queue) Dequeue() Item {
	if len(queue.Queue) == 0 {
		fmt.Println("队列空了")
		return nil
	}
	item := queue.Queue[0]
	queue.Queue = queue.Queue[1:]

	return item
}


