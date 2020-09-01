package silo_manager

import (
	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/processors"
	"github.com/otherview/gambaru/core/queues"
)

type CreateProcessorMessage struct {
	Processor *processors.ProcessorInterface
	ID        uuid.UUID
}
type CreateProcessorOKMessage struct {
	PID uuid.UUID
}

type CreateQueueMessage struct {
	Queue *queues.QueueInterface
	ID    uuid.UUID
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
