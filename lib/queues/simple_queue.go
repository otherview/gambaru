package queues

import "fmt"

type SimpleQueue struct {
	queueItems []interface{}
}

func NewSimpleQueue() *SimpleQueue {
	return &SimpleQueue{
		queueItems: []interface{}{},
	}
}

func (state *SimpleQueue) Read() (interface{}, error) {

	var queueItem interface{}
	if len(state.queueItems) > 0 {
		queueItem = state.queueItems[0]
		state.queueItems = state.queueItems[1:]
	}

	fmt.Printf("Removed item from the Queue %v\n", queueItem)

	return queueItem, nil
}

func (state *SimpleQueue) Write(item interface{}) error {

	state.queueItems = append(state.queueItems, item)
	fmt.Printf("Added item to the Queue %v\n", item)
	return nil
}
