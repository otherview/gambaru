package silo

import (
	"time"

	"github.com/otherview/gambaru/core"

	"github.com/google/uuid"
	silo_manager "github.com/otherview/gambaru/core/managers/silo"
)

// CreateQueue returns an id of a created Queue
func (silo *Silo) CreateQueue(queue *core.QueueInterface, queueID uuid.UUID) uuid.UUID {

	createdQueueMessage, err := silo.context.RequestFuture(silo.siloPID, &silo_manager.CreateQueueMessage{
		Queue: queue,
		ID:    queueID,
	}, 5*time.Second).Result()
	if err != nil {
		// TODO: yep
		panic(err)
	}

	createdQueue, ok := createdQueueMessage.(silo_manager.CreateQueueOKMessage)
	if !ok {
		return uuid.Nil
	}

	return createdQueue.PID
}