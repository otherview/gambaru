package processor_manager

import (
	"fmt"
	"time"

	queue_manager "github.com/otherview/gambaru/core/managers/queue"

	"github.com/AsynkronIT/protoactor-go/actor"

	"github.com/otherview/gambaru/core/processors"
)

type ProcessorManager struct {
	processor   *processors.ProcessorInterface
	inputQueue  *actor.PID
	outputQueue *actor.PID
	stopChan    chan bool
}

func NewProcessorManager(processor *processors.ProcessorInterface) *ProcessorManager {
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
				var queueItem interface{}

				if state.inputQueue != nil {
					queueItem, _ = actor.EmptyRootContext.RequestFuture(state.inputQueue, &queue_manager.ReadQueueItemMessage{}, 5*time.Second).Result()
				}

				responseData, _ := (*state.processor).Execute(queueItem)

				if state.outputQueue != nil {
					_, _ = actor.EmptyRootContext.RequestFuture(state.outputQueue, &queue_manager.WriteQueueItemMessage{QueueItem: responseData}, 5*time.Second).Result()
				}
			}
			time.Sleep(time.Second)
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
