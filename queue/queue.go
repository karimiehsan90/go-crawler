package queue

type Queue struct {
	items []string
}

type EmptyQueueError struct{}

func (a *EmptyQueueError) Error() string {
	return "Queue is empty"
}

func (q *Queue) PushBack(item string) {
	q.items = append(q.items, item)
}

func (q *Queue) PopFront() (string, error) {
	if len(q.items) == 0 {
		return "", &EmptyQueueError{}
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}
