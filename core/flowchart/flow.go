package flowchart

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/otherview/gambaru/lib/queues"

	"github.com/otherview/gambaru/core/silo"
)

type FlowChart struct {
	flowStructure Structure
}

// NewFlow creates a FlowChart model from a string
func NewFlow(flowJson string) *FlowChart {

	var flow Structure
	if err := json.Unmarshal([]byte(flowJson), &flow); err != nil {
		// TODO yep
		panic(err)
	}

	return &FlowChart{flowStructure: flow}
}

// CreateFlow creates the flow in the silo
func (flow *FlowChart) CreateFlow(silo *silo.Silo) error {

	if err := flow.validateComponentTypes(silo); err != nil {
		return errors.Wrap(err, "failed to validate the flow")
	}

	for _, queue := range flow.flowStructure.Queues {

		queueType := silo.GetRegisteredQueue(queue.Type)
		if queueType == nil {
			// TODO probably alert here
			queueType = queues.NewSimpleQueue()
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

func (flow *FlowChart) validateComponentTypes(silo *silo.Silo) error {

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
