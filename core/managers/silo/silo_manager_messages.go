package silo_manager

import (
	"github.com/google/uuid"
	interface_processor "github.com/otherview/gambaru/core/interfaces/processor"
	interface_queue "github.com/otherview/gambaru/core/interfaces/queue"
	"github.com/otherview/gambaru/core/repository"
)

type CreateProcessorMessage struct {
	Processor interface_processor.ProcessorInterface
	ID        uuid.UUID
}
type CreateProcessorOKMessage struct {
	PID uuid.UUID
}

type CreateQueueMessage struct {
	Queue      interface_queue.QueueInterface
	ID         uuid.UUID
	Repository *repository.MemoryRepository
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
