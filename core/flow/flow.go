package flow

import (
	"encoding/json"
	"fmt"

	"github.com/otherview/gambaru/core"

	"github.com/pkg/errors"

	"github.com/otherview/gambaru/lib/queues"

	"github.com/otherview/gambaru/core/silo"
)

type Flow struct {
	flowStructure Structure
}

// NewFlow creates a Flow model from a string
func NewFlow(flowJson string) *Flow {

	var flow Structure
	if err := json.Unmarshal([]byte(flowJson), &flow); err != nil {
		// TODO yep
		panic(err)
	}

	return &Flow{flowStructure: flow}
}

// CreateFlow creates the flow in the silo
func (flow *Flow) CreateFlow(silo *silo.Silo) error {

	if err := flow.validateComponentTypes(silo); err != nil {
		return errors.Wrap(err, "failed to validate the flow")
	}

	for _, queue := range flow.flowStructure.Queues {
		var queueType core.QueueInterface
		queueType = queues.NewSimpleQueue()

		if queue.Type != "" {
			queueType = silo.GetRegisteredQueue(queue.Type)
		}

		silo.CreateQueue(&queueType, queue.ID)
	}

	for _, processor := range flow.flowStructure.Processors {
		// TODO check if it exists :o
		procType := silo.GetRegisteredProcessor(processor.Type)
		silo.CreateProcessor(procType, processor.ID)
		silo.AddInputQueue(processor.ID, processor.InputQueue)
		silo.AddOutputQueue(processor.ID, processor.OutputQueue)
	}

	return nil
}

func (flow *Flow) validateComponentTypes(silo *silo.Silo) error {

	for _, queue := range flow.flowStructure.Queues {
		if queue.Type != "" {
			if silo.GetRegisteredQueue(queue.Type) == nil {
				return fmt.Errorf("queue Type %s for %v is not registered", queue.Type, queue.ID)
			}
		}
	}

	for _, proc := range flow.flowStructure.Processors {
		if proc.Type != "" {
			if silo.GetRegisteredProcessor(proc.Type) == nil {
				return fmt.Errorf("processor type %s for %v is not registered", proc.Type, proc.ID)
			}
		}
	}

	return nil
}
