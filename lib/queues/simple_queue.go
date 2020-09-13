package queues

import (
	"fmt"

	"github.com/otherview/gambaru/core/flowfiles"
)

type SimpleQueue struct {
	queueItems []*flowfiles.Flowfile
}

func NewSimpleQueue() *SimpleQueue {
	return &SimpleQueue{
		queueItems: []*flowfiles.Flowfile{},
	}
}

func (state *SimpleQueue) Read() (*flowfiles.Flowfile, error) {

	var queueItem *flowfiles.Flowfile
	if len(state.queueItems) > 0 {
		queueItem = state.queueItems[0]
		state.queueItems = state.queueItems[1:]
	}

	return queueItem, nil
}

func (state *SimpleQueue) Write(item *flowfiles.Flowfile) error {

	if item == nil {
		return nil
	}

	state.queueItems = append(state.queueItems, item)

	fmt.Println("Items on the queue -> ", len(state.queueItems))
	return nil
}
