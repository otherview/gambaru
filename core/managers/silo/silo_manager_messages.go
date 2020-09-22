package silo_manager

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/google/uuid"
	interface_processor "github.com/otherview/gambaru/core/interfaces/processor"
	interface_queue "github.com/otherview/gambaru/core/interfaces/queue"
	"github.com/otherview/gambaru/core/models/flow_model"
	"github.com/otherview/gambaru/core/repository"
)

type CreateProcessorMessage struct {
	Processor interface_processor.ProcessorInterface
	ID        uuid.UUID
	Config    *flow_model.Config
}
type CreateProcessorOKMessage struct {
	PID uuid.UUID
}

type CreateQueueMessage struct {
	Queue      interface_queue.QueueInterface
	ID         uuid.UUID
	Repository *repository.WALRepository
}
type CreateQueueOKMessage struct {
	PID uuid.UUID
}

type StartSiloMessage struct{}
type StartSiloOKMessage struct{}

type StopSiloMessage struct{}
type StopSiloOKMessage struct{}

type AddOutputQueueMessage struct {
	ProcessorID uuid.UUID
	QueueID     uuid.UUID
}
type AddOutputQueueOKMessage struct{}

type AddInputQueueMessage struct {
	ProcessorID uuid.UUID
	QueueID     uuid.UUID
}
type AddInputQueueOKMessage struct{}

type GetProcessorsMessage struct{}
type GetProcessorsOKMessage struct {
	Processors map[uuid.UUID]*actor.PID
}

type GetQueuesMessage struct{}
type GetQueuesOKMessage struct {
	Queues map[uuid.UUID]*actor.PID
}
