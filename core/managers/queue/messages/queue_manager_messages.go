package queue_manager_actor

import "github.com/otherview/gambaru/core/flowfiles"

type ReadQueueItemMessage struct{}
type ReadQueueItemOKMessage struct {
	QueueItem *flowfiles.Flowfile
}
type WriteQueueItemMessage struct {
	QueueItem *flowfiles.Flowfile
}
type WriteQueueItemOKMessage struct{}

type RemoveQueueItemMessage struct {
	QueueItem *flowfiles.Flowfile
}
type RemoveQueueItemOKMessage struct{}
