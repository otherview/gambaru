package sessions

import (
	"time"

	"github.com/google/uuid"

	interface_repository "github.com/otherview/gambaru/core/interfaces/repository"

	"github.com/pkg/errors"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/otherview/gambaru/core/flowfiles"
	queue_manager_messages "github.com/otherview/gambaru/core/managers/queue/messages"
)

type Session struct {
	inputQueue      *actor.PID
	outputQueue     *actor.PID
	repository      interface_repository.RepositoryInterface
	operationsQueue *operationQueue
}

func NewSession(repository interface_repository.RepositoryInterface, inputQueue *actor.PID, outputQueue *actor.PID) *Session {

	return &Session{
		repository:      repository,
		inputQueue:      inputQueue,
		outputQueue:     outputQueue,
		operationsQueue: newOperationQueue(),
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
		state.operationsQueue.queueOperation(flowfile, state.outputQueue)
	}
	return nil
}

func (state *Session) ReadFlowfileData(flowfile *flowfiles.Flowfile) (interface{}, error) {

	return state.repository.Read(flowfile)
}

func (state *Session) WriteFlowfileData(flowfile *flowfiles.Flowfile, phrase string) error {

	return state.repository.Write(flowfile, uuid.New(), phrase)
}

func (state *Session) Commit() error {

	var err error
	operationExecuted := newOperationQueue()
	for _, operation := range state.operationsQueue.queue {
		_, err := actor.EmptyRootContext.RequestFuture(state.outputQueue, &queue_manager_messages.WriteQueueItemMessage{QueueItem: operation.flowfile}, 5*time.Second).Result()
		if err != nil {
			operationExecuted.queueOperation(operation.flowfile, state.outputQueue)
			err = errors.Wrap(err, "failed to send message to Queue")
			break
		}
	}

	if operationExecuted.failed() {
		// TODO this should be reversed
		for _, operation := range operationExecuted.queue {
			_, err := actor.EmptyRootContext.RequestFuture(state.outputQueue, &queue_manager_messages.RemoveQueueItemMessage{QueueItem: operation.flowfile}, 5*time.Second).Result()
			if err != nil {
				// TODO yep...
				panic(err)
			}

			state.repository.Rollback(operation.flowfile)
		}
		return err
	}

	for _, operation := range state.operationsQueue.queue {
		state.repository.Commit(operation.flowfile)
	}
	return err
}

func (state *Session) Remove(flowfile *flowfiles.Flowfile) error {
	return state.repository.Remove(flowfile)
}
