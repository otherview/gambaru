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
	inputQueues     []*actor.PID
	outputQueue     *actor.PID
	repository      interface_repository.RepositoryInterface
	operationsQueue *operationQueue
	savePoint       bool
}

func NewSession(repository interface_repository.RepositoryInterface, inputQueues []*actor.PID, outputQueue *actor.PID, savePoint bool) *Session {

	return &Session{
		repository:      repository,
		inputQueues:     inputQueues,
		outputQueue:     outputQueue,
		savePoint:       savePoint,
		operationsQueue: newOperationQueue(),
	}
}

func (state *Session) GetFlowfile() *flowfiles.Flowfile {
	var queuedFlowfile *flowfiles.Flowfile

	//TODO add some logic around which queue to get from

	for _, queue := range state.inputQueues {
		queueMsg, _ := actor.EmptyRootContext.RequestFuture(queue, &queue_manager_messages.ReadQueueItemMessage{}, 5*time.Second).Result()
		if queueMsg != nil && queueMsg.(queue_manager_messages.ReadQueueItemOKMessage).QueueItem != nil {
			return queueMsg.(queue_manager_messages.ReadQueueItemOKMessage).QueueItem
		}
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
		// Commits the data to be available + savepointed
		err = state.repository.Commit(operation.flowfile, state.savePoint)
		if err != nil {
			operationExecuted.queueOperation(operation.flowfile, state.outputQueue)
			err = errors.Wrap(err, "failed to send message to Queue")
			break
		}

		// Notifies the queue that new events are ready + data is ready

		_, err = actor.EmptyRootContext.RequestFuture(state.outputQueue, &queue_manager_messages.WriteQueueItemMessage{QueueItem: operation.flowfile}, 5*time.Second).Result()
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

	return err
}

func (state *Session) Remove(flowfile *flowfiles.Flowfile) error {
	return state.repository.Remove(flowfile)
}
