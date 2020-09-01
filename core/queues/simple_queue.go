package queues

type SimpleQueue struct {
	queueData []byte
}

func NewSimpleQueue() *SimpleQueue {
	return &SimpleQueue{}
}

func (queue *SimpleQueue) Read() ([]byte, error) {

	return queue.queueData, nil
}

func (queue *SimpleQueue) Write(data []byte) error {

	queue.queueData = data
	return nil
}
