package queue_manager

import "github.com/otherview/gambaru/core/flowfile"

type ReadQueueItemMessage struct{}
type ReadQueueItemOKMessage struct {
	QueueItem *flowfiles.Flowfile
}
type WriteQueueItemMessage struct {
	QueueItem *flowfiles.Flowfile
}
type WriteQueueItemOKMessage struct{}
