package queues

import (
	"time"

	"github.com/otherview/gambaru/core/flowfiles"
)

type BufferQueue struct {
	queueItems chan *flowfiles.Flowfile
}

func NewBufferQueue() *BufferQueue {
	return &BufferQueue{
		queueItems: make(chan *flowfiles.Flowfile, 1000),
	}
}

func (state *BufferQueue) Read() (*flowfiles.Flowfile, error) {

	select {
	case flowfile := <-state.queueItems:
		return flowfile, nil
	// TODO yeah this is probably not gonna work if the file is too big ? although its just the reference to it..
	case <-time.After(time.Millisecond * 10):
		return nil, nil
	}
}

func (state *BufferQueue) Write(item *flowfiles.Flowfile) error {

	if item == nil {
		return nil
	}

	state.queueItems <- item

	return nil
}
