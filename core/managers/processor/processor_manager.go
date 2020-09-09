package processor_manager

import (
	"fmt"
	"time"

	interface_processor "github.com/otherview/gambaru/core/interfaces/processor"
	interface_repository "github.com/otherview/gambaru/core/interfaces/repository"

	"github.com/otherview/gambaru/core/sessions"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type ProcessorManager struct {
	processor   interface_processor.ProcessorInterface
	inputQueue  *actor.PID
	outputQueue *actor.PID
	stopChan    chan bool
	repository  interface_repository.RepositoryInterface
}

func NewProcessorManager(processor interface_processor.ProcessorInterface, repository interface_repository.RepositoryInterface) *ProcessorManager {
	return &ProcessorManager{
		processor:  processor,
		repository: repository,
		stopChan:   make(chan bool),
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

				_ = state.processor.Execute(sessions.NewSession(state.repository, state.inputQueue, state.outputQueue))
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
