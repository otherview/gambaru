package queue_manager

import (
	"github.com/google/uuid"
	"github.com/otherview/gambaru/core"
	"github.com/otherview/gambaru/core/flowfiles"
	"github.com/otherview/gambaru/core/repository"
)

type QueueManager struct {
	queue      *core.QueueInterface
	repository *repository.Repository
}

func NewQueueManager(repository *repository.Repository, queue *core.QueueInterface) *QueueManager {
	return &QueueManager{
		queue:      queue,
		repository: repository,
	}
}

func (state *QueueManager) ReadQueueItem() *flowfiles.Flowfile {

	queueItem, err := (*state.queue).Read()
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

	err := (*state.queue).Write(item)
	if err != nil {
		// TODO yep
		panic(err)
	}

	return err
}
