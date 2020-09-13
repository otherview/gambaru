package processor_manager

import (
	"fmt"
	"time"

	"github.com/otherview/gambaru/core/models/flow_model"

	interface_processor "github.com/otherview/gambaru/core/interfaces/processor"
	interface_repository "github.com/otherview/gambaru/core/interfaces/repository"

	"github.com/otherview/gambaru/core/sessions"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type ProcessorManager struct {
	processor     interface_processor.ProcessorInterface
	savePoint     bool
	inputQueues   []*actor.PID
	outputQueue   *actor.PID
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

func (state *ProcessorManager) AddInputQueue(queuePID *actor.PID) error {

	state.inputQueues = append(state.inputQueues, queuePID)
	return nil
}

func (state *ProcessorManager) AddOutputQueue(queuePID *actor.PID) error {

	state.outputQueue = queuePID
	return nil
}
