package silo_manager

import (
	"github.com/google/uuid"
	"github.com/otherview/gambaru/core"
	"github.com/otherview/gambaru/core/repository"
)

type CreateProcessorMessage struct {
	Processor *core.ProcessorInterface
	ID        uuid.UUID
}
type CreateProcessorOKMessage struct {
	PID uuid.UUID
}

type CreateQueueMessage struct {
	Queue      *core.QueueInterface
	ID         uuid.UUID
	Repository *repository.Repository
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
