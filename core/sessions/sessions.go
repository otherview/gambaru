package sessions

import (
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/otherview/gambaru/core/flowfiles"
	queue_manager_messages "github.com/otherview/gambaru/core/managers/queue/messages"
	"github.com/otherview/gambaru/core/repository"
)

type Session struct {
	inputQueue  *actor.PID
	outputQueue *actor.PID
	repository  *repository.Repository
}

func NewSession(repository *repository.Repository, inputQueue *actor.PID, outputQueue *actor.PID) *Session {

	return &Session{
		repository:  repository,
		inputQueue:  inputQueue,
		outputQueue: outputQueue,
	}
}

func (state *Session) GetFlowfile() *flowfiles.Flowfile {
	var queuedFlowfile *flowfiles.Flowfile

	queueMsg, _ := actor.EmptyRootContext.RequestFuture(state.inputQueue, &queue_manager_messages.ReadQueueItemMessage{}, 5*time.Second).Result()
	if queueMsg != nil && queueMsg.(queue_manager_messages.ReadQueueItemOKMessage).QueueItem != nil {
		queuedFlowfile = queueMsg.(queue_manager_messages.ReadQueueItemOKMessage).QueueItem
	}

	return queuedFlowfile
}

func (state *Session) TransferFlowfile(flowfile *flowfiles.Flowfile) error {

	if state.outputQueue != nil && flowfile != nil {
		_, _ = actor.EmptyRootContext.RequestFuture(state.outputQueue, &queue_manager_messages.WriteQueueItemMessage{QueueItem: flowfile}, 5*time.Second).Result()
	}
	return nil
}

func (state *Session) ReadFlowfileData(flowfile *flowfiles.Flowfile) (interface{}, interface{}) {

	return state.repository.Read(flowfile)
}

func (state *Session) WriteFlowfileData(flowfile *flowfiles.Flowfile, phrase string) error {

	return state.repository.Write(flowfile, phrase)
}
