package consumer

type WorkerQueue struct {
	items []Worker
}

func (queue *WorkerQueue) Enqueue(w Worker) {
	queue.items = append(queue.items, w)

}

func (queue *WorkerQueue) Dequeue() Worker {
	var w Worker
	if len(queue.items) > 0 {
		w = queue.items[0]
		queue.items = queue.items[1:]
	}
	return w
}

func (queue *WorkerQueue) Len() int {
	return len(queue.items)
}

func (queue *WorkerQueue) Peek(index int) Worker {
	return queue.items[index]
}
