package silo_manager

import (
	"fmt"

	"github.com/otherview/gambaru/core/queues"

	queue_manager "github.com/otherview/gambaru/core/managers/queue"

	"github.com/teamwork/deskapi/util/time"

	"github.com/AsynkronIT/protoactor-go/actor"
	processor_manager "github.com/otherview/gambaru/core/managers/processor"

	"github.com/google/uuid"
)

type SiloManager struct {
	processors map[uuid.UUID]*actor.PID
	queues     map[uuid.UUID]*actor.PID
}

func NewSiloManager() *SiloManager {
	return &SiloManager{
		processors: map[uuid.UUID]*actor.PID{},
		queues:     map[uuid.UUID]*actor.PID{},
	}
}

func (state *SiloManager) CreateNewProcessor(msg *CreateProcessorMessage) (uuid.UUID, error) {

	newProcID := uuid.New()
	if msg.ID != uuid.Nil {
		newProcID = msg.ID
	}

	props := actor.PropsFromProducer(func() actor.Actor { return processor_manager.NewProcessorManager(msg.Processor) })
	pid := actor.EmptyRootContext.Spawn(props)

	state.processors[newProcID] = pid

	return newProcID, nil
}

func (state *SiloManager) CreateNewQueue(queue *queues.QueueInterface, queueID uuid.UUID) (uuid.UUID, error) {

	newQueueID := uuid.New()
	if queueID != uuid.Nil {
		newQueueID = queueID
	}

	props := actor.PropsFromProducer(func() actor.Actor { return queue_manager.NewQueueManager() })
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

	fmt.Printf("Started Processor %v\n", processorID)

	return nil
}

func (state *SiloManager) AddOutputQueue(processorID uuid.UUID, queueID uuid.UUID) error {

	// TODO what if it doesnt exist ?
	processorPID := state.processors[processorID]
	queuePID := state.queues[queueID]

	_, _ = actor.EmptyRootContext.RequestFuture(processorPID,
		&processor_manager.AddOutputQueue{QueuePID: queuePID},
		5*time.Second).Result()

	fmt.Printf("Started Processor %v\n", processorID)

	return nil

}
