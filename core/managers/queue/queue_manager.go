package queue_manager

import (
	"github.com/otherview/gambaru/core"
)

type QueueManager struct {
	queue *core.QueueInterface
}

func NewQueueManager(queue *core.QueueInterface) *QueueManager {
	return &QueueManager{
		queue: queue,
	}
}

func (state *QueueManager) ReadQueueItem() interface{} {

	queueItem, err := (*state.queue).Read()
	if err != nil {
		// TODO yep
		panic(err)
	}

	return queueItem
}

func (state *QueueManager) WriteQueueItem(item interface{}) error {

	err := (*state.queue).Write(item)
	if err != nil {
		// TODO yep
		panic(err)
	}

	return err
}
