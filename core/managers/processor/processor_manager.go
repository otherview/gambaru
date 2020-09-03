package processor_manager

import (
	"fmt"
	"time"

	"github.com/otherview/gambaru/core/flowfile"

	"github.com/otherview/gambaru/core"

	queue_manager "github.com/otherview/gambaru/core/managers/queue"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type ProcessorManager struct {
	processor   *core.ProcessorInterface
	inputQueue  *actor.PID
	outputQueue *actor.PID
	stopChan    chan bool
}

func NewProcessorManager(processor *core.ProcessorInterface) *ProcessorManager {
	return &ProcessorManager{
		processor: processor,
		stopChan:  make(chan bool),
	}
}

func (state *ProcessorManager) StartProcessor() error {

	// TODO dont start the processor twice
	go func(stopChan chan bool) {
		for {
			select {
			case <-stopChan:
				return
			default:
				var queueMsg interface{}
				var responseFlowfile *flowfiles.Flowfile
				var flowfile *flowfiles.Flowfile

				if state.inputQueue != nil {
					queueMsg, _ = actor.EmptyRootContext.RequestFuture(state.inputQueue, &queue_manager.ReadQueueItemMessage{}, 5*time.Second).Result()
				}

				if queueMsg != nil && queueMsg.(queue_manager.ReadQueueItemOKMessage).QueueItem != nil {
					flowfile = queueMsg.(queue_manager.ReadQueueItemOKMessage).QueueItem
				}

				responseFlowfile, _ = (*state.processor).Execute(flowfile)

				if state.outputQueue != nil && responseFlowfile != nil {
					_, _ = actor.EmptyRootContext.RequestFuture(state.outputQueue, &queue_manager.WriteQueueItemMessage{QueueItem: responseFlowfile}, 5*time.Second).Result()
				}
			}
			//time.Sleep(time.Second)
		}

	}(state.stopChan)

	return nil
}

func (state *ProcessorManager) StopProcessor() error {

	state.stopChan <- true
	fmt.Printf("Stopping Processor\n")
	return nil
}

func (state *ProcessorManager) AddInputQueue(queuePID *actor.PID) error {

	state.inputQueue = queuePID
	return nil
}

func (state *ProcessorManager) AddOutputQueue(queuePID *actor.PID) error {

	state.outputQueue = queuePID
	return nil
}
