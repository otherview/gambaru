package queue_manager

type ReadQueueItemMessage struct{}
type ReadQueueItemOKMessage struct {
	Item interface{}
}
type WriteQueueItemMessage struct {
	QueueItem interface{}
}
type WriteQueueItemOKMessage struct{}
