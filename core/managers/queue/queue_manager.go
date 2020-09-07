package queue_manager

import (
	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/flowfiles"
	interface_queue "github.com/otherview/gambaru/core/interfaces/queue"
	interface_repository "github.com/otherview/gambaru/core/interfaces/repository"
)

type QueueManager struct {
	queue      interface_queue.QueueInterface
	repository interface_repository.RepositoryInterface
}

func NewQueueManager(repository interface_repository.RepositoryInterface, queue interface_queue.QueueInterface) *QueueManager {
	return &QueueManager{
		queue:      queue,
		repository: repository,
	}
}

func (state *QueueManager) ReadQueueItem() *flowfiles.Flowfile {

	queueItem, err := state.queue.Read()
	if err != nil {
		// TODO yep
		panic(err)
	}

	return queueItem
}

func (state *QueueManager) WriteQueueItem(item *flowfiles.Flowfile) error {

	if item.ID == uuid.Nil {
		return nil
	}

	err := state.queue.Write(item)
	if err != nil {
		// TODO yep
		panic(err)
	}

	return err
}
