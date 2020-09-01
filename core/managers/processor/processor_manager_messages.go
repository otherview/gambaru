package processor_manager

import "github.com/AsynkronIT/protoactor-go/actor"

type StartProcessorMessage struct{}
type StartProcessorOKMessage struct{}
type StopProcessorMessage struct{}
type StopProcessorOKMessage struct{}

type AddInputQueue struct {
	QueuePID *actor.PID
}
type AddInputOKQueue struct{}

type AddOutputQueue struct {
	QueuePID *actor.PID
}
type AddOutputOKQueue struct{}
