package framework

type RequestQueue struct {
	items []Request
}

func (queue *RequestQueue) Enqueue(request Request) {
	queue.items = append(queue.items, request)

}

func (queue *RequestQueue) Dequeue() Request {
	var request Request
	if len(queue.items) > 0 {
		request = queue.items[0]
		queue.items = queue.items[1:]
	}
	return request
}

func (queue *RequestQueue) Len() int {
	return len(queue.items)
}

func (queue *RequestQueue) Peek(index int) Request {
	return queue.items[index]
}
