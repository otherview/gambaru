package sessions

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/otherview/gambaru/core/flowfiles"
)

type operationQueue struct {
	queue []*operation
}

func newOperationQueue() *operationQueue {
	return &operationQueue{
		queue: []*operation{},
	}
}

func (q *operationQueue) queueOperation(flowfile *flowfiles.Flowfile, boundForQueue *actor.PID) {

	q.queue = append(q.queue, newOperation(flowfile, boundForQueue))
}

func (q *operationQueue) failed() bool {

	return len(q.queue) > 0
}
