package processor_manager

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/otherview/gambaru/core/models/queue_model"

	"github.com/otherview/gambaru/core/models/flow_model"

	interface_processor "github.com/otherview/gambaru/core/interfaces/processor"
	interface_repository "github.com/otherview/gambaru/core/interfaces/repository"

	"github.com/otherview/gambaru/core/sessions"
)

type ProcessorManager struct {
	processorID   uuid.UUID
	procType      string
	processor     interface_processor.ProcessorInterface
	savePoint     bool
	inputQueues   []queue_model.Queue
	outputQueue   queue_model.Queue
	stopChan      chan bool
	repository    interface_repository.RepositoryInterface
	executionTime time.Duration
}

func NewProcessorManager(processor interface_processor.ProcessorInterface, config *flow_model.Config, repository interface_repository.RepositoryInterface) *ProcessorManager {

	executionTime, err := time.ParseDuration(config.ExecutionTime)
	if err != nil {
		//TODO yep
		panic(err)
	}

	return &ProcessorManager{
		processor:     processor,
		savePoint:     config.SavePoint,
		executionTime: executionTime,
		repository:    repository,
		stopChan:      make(chan bool),
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
				newSession := sessions.NewSession(state.repository, state.inputQueues, state.outputQueue, state.savePoint)
				_ = state.processor.Execute(newSession)
				//TODO ensure that all flowfiles read/written are pushed or removed
				//newSession.VerifyCleanSession()
			}
			time.Sleep(state.executionTime)
		}

	}(state.stopChan)

	return nil
}

func (state *ProcessorManager) StopProcessor() error {

	state.stopChan <- true
	fmt.Printf("Stopping Processor\n")
	return nil
}

func (state *ProcessorManager) AddInputQueue(msg *AddInputQueue) error {

	state.inputQueues = append(state.inputQueues, queue_model.Queue{
		ID:      msg.QueueID,
		ActorID: msg.QueuePID,
	})
	return nil
}

func (state *ProcessorManager) AddOutputQueue(msg *AddOutputQueue) error {

	state.outputQueue = queue_model.Queue{
		ID:      msg.QueueID,
		ActorID: msg.QueuePID,
	}
	return nil
}

func (state *ProcessorManager) GetProcessorInfo() (uuid.UUID, string, []queue_model.Queue, queue_model.Queue) {

	return state.processorID, state.procType, state.inputQueues, state.outputQueue
}
