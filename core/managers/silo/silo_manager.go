package silo_manager

import (
	"fmt"

	interface_repository "github.com/otherview/gambaru/core/interfaces/repository"

	"github.com/otherview/gambaru/core/repository"

	queue_manager "github.com/otherview/gambaru/core/managers/queue"

	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	processor_manager "github.com/otherview/gambaru/core/managers/processor"

	"github.com/google/uuid"
)

type SiloManager struct {
	processors map[uuid.UUID]*actor.PID
	queues     map[uuid.UUID]*actor.PID
	repository interface_repository.RepositoryInterface
}

func NewSiloManager() *SiloManager {
	return &SiloManager{
		processors: map[uuid.UUID]*actor.PID{},
		queues:     map[uuid.UUID]*actor.PID{},
		repository: repository.NewWALRepository(),
	}
}

func (state *SiloManager) CreateNewProcessor(msg *CreateProcessorMessage) (uuid.UUID, error) {

	newProcID := uuid.New()
	if msg.ID != uuid.Nil {
		newProcID = msg.ID
	}

	props := actor.PropsFromProducer(func() actor.Actor {
		return processor_manager.NewProcessorManager(msg.Processor, msg.Config, state.repository)
	})
	pid := actor.EmptyRootContext.Spawn(props)

	state.processors[newProcID] = pid

	return newProcID, nil
}

func (state *SiloManager) CreateNewQueue(msg *CreateQueueMessage) (uuid.UUID, error) {

	newQueueID := uuid.New()
	if msg.ID != uuid.Nil {
		newQueueID = msg.ID
	}

	props := actor.PropsFromProducer(func() actor.Actor { return queue_manager.NewQueueManager(msg.Repository, msg.Queue) })
	pid := actor.EmptyRootContext.Spawn(props)

	state.queues[newQueueID] = pid

	return newQueueID, nil
}

func (state *SiloManager) StartSilo() error {
	for processorID, processorPID := range state.processors {
		_, _ = actor.EmptyRootContext.RequestFuture(processorPID,
			&processor_manager.StartProcessorMessage{},
			5*time.Second).Result()

		fmt.Printf("Started Processor %v\n", processorID)
	}
	return nil
}

func (state *SiloManager) StopSilo() error {

	for processorID, processorPID := range state.processors {
		_, _ = actor.EmptyRootContext.RequestFuture(processorPID,
			&processor_manager.StopProcessorMessage{},
			5*time.Second).Result()

		fmt.Printf("Stopped Processor %v\n", processorID)
	}
	return nil
}

func (state *SiloManager) AddInputQueue(processorID uuid.UUID, queueID uuid.UUID) error {

	// TODO what if it doesnt exist ?
	processorPID := state.processors[processorID]
	queuePID := state.queues[queueID]

	_, _ = actor.EmptyRootContext.RequestFuture(processorPID,
		&processor_manager.AddInputQueue{QueuePID: queuePID},
		5*time.Second).Result()

	fmt.Printf("Added Input Queue %v to Processor %v\n", queueID, processorID)

	return nil
}

func (state *SiloManager) AddOutputQueue(processorID uuid.UUID, queueID uuid.UUID) error {

	// TODO what if it doesnt exist ?
	processorPID := state.processors[processorID]
	queuePID := state.queues[queueID]

	_, _ = actor.EmptyRootContext.RequestFuture(processorPID,
		&processor_manager.AddOutputQueue{QueuePID: queuePID},
		5*time.Second).Result()

	fmt.Printf("Added Output Queue %v to Processor %v\n", queueID, processorID)

	return nil

}

func (state *SiloManager) GetProcessors() map[uuid.UUID]*actor.PID {
	return state.processors
}
