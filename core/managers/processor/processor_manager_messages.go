package processor_manager

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/google/uuid"
	"github.com/otherview/gambaru/core/models/queue_model"
)

type StartProcessorMessage struct{}
type StartProcessorOKMessage struct{}
type StopProcessorMessage struct{}
type StopProcessorOKMessage struct{}

type AddInputQueue struct {
	QueueID  uuid.UUID
	QueuePID *actor.PID
}
type AddInputOKQueue struct{}

type AddOutputQueue struct {
	QueueID  uuid.UUID
	QueuePID *actor.PID
}
type AddOutputOKQueue struct{}

type GetProcessorInfoMessage struct{}
type GetProcessorInfoOKMessage struct {
	ID           uuid.UUID
	Type         string
	InputQueues  []queue_model.Queue
	OutputQueues queue_model.Queue
}
