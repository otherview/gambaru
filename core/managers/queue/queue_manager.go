package queue_manager

import (
	"fmt"
)

type QueueManager struct {
	queueItems []interface{}
}

func NewQueueManager() *QueueManager {
	return &QueueManager{
		queueItems: []interface{}{},
	}
}

func (state *QueueManager) ReadQueueItem() interface{} {

	var queueItem interface{}
	if len(state.queueItems) > 0 {
		queueItem = state.queueItems[0]
		state.queueItems = state.queueItems[1:]
	}

	fmt.Printf("Removed item from the Queue %v\n", queueItem)

	return queueItem
}

func (state *QueueManager) WriteQueueItem(item interface{}) error {

	state.queueItems = append(state.queueItems, item)
	fmt.Printf("Added item to the Queue %v\n", item)
	return nil
}
