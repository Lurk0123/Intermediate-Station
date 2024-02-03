package utils

import "sync"

type Queue struct {
	items []string
	lock  sync.Mutex
}

func (q *Queue) Enqueue(item string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() string {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.items) == 0 {
		return ""
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}
func (q *Queue) Clear() {
	q.lock.Lock()
	defer q.lock.Lock()
	q.items = make([]string, 0)
}
