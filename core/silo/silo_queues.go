package silo

import (
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"

	"github.com/google/uuid"
	interface_queue "github.com/otherview/gambaru/core/interfaces/queue"
	silo_manager "github.com/otherview/gambaru/core/managers/silo"
)

// CreateQueue returns an id of a created Queue
func (silo *Silo) CreateQueue(queue interface_queue.QueueInterface, queueID uuid.UUID) uuid.UUID {

	createdQueueMessage, err := silo.context.RequestFuture(silo.siloPID, &silo_manager.CreateQueueMessage{
		Queue:      queue,
		ID:         queueID,
		Repository: silo.repository,
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

func (silo *Silo) ExportQueues() (map[uuid.UUID]*actor.PID, error) {

	getQueuesMessage, err := silo.context.RequestFuture(silo.siloPID,
		&silo_manager.GetQueuesMessage{}, 5*time.Second).Result()

	if err != nil {
		// TODO: yep
		panic(err)
	}

	queues, ok := getQueuesMessage.(silo_manager.GetQueuesOKMessage)
	if !ok {
		// TODO make an error definition here between silo and service
		return nil, err
	}

	return queues.Queues, nil
}
