package flowchart

import (
	"encoding/json"
	"fmt"

	"github.com/otherview/gambaru/core/models/flow_model"

	"github.com/pkg/errors"

	"github.com/otherview/gambaru/lib/queues"

	"github.com/otherview/gambaru/core/silo"
)

type FlowChart struct {
	flowStructure flow_model.Structure
	silo          *silo.Silo
}

// NewFlow creates a FlowChart model from a string
func NewFlow(silo *silo.Silo) *FlowChart {

	return &FlowChart{silo: silo}
}

// LoadJson Loads a new Json into the flow format
func (flow *FlowChart) LoadJson(flowJson string) error {
	var flowStruct flow_model.Structure
	if err := json.Unmarshal([]byte(flowJson), &flowStruct); err != nil {
		return err
	}
	flow.flowStructure = flowStruct

	return nil
}

// CreateFlow creates the flow in the silo
func (flow *FlowChart) CreateFlow() error {

	for _, queue := range flow.flowStructure.Queues {
		if err := flow.validateQueueType(&queue); err != nil {
			return errors.Wrap(err, "failed to validate the flow")
		}
	}

	for _, proc := range flow.flowStructure.Processors {
		if err := flow.validateProcessorType(&proc); err != nil {
			return errors.Wrap(err, "failed to validate the flow")
		}
	}

	for _, queue := range flow.flowStructure.Queues {
		if err := flow.createQueue(&queue); err != nil {
			return errors.Wrap(err, "failed to validate the flow")
		}
	}

	for _, processor := range flow.flowStructure.Processors {
		if err := flow.createProcessor(&processor); err != nil {
			return errors.Wrap(err, "failed to validate the flow")
		}
	}

	return nil
}

// CreateProcessor validates and creates a processor in the silo
func (flow *FlowChart) CreateProcessor(processor *flow_model.Processor) error {
	if err := flow.validateProcessorType(processor); err != nil {
		return errors.Wrap(err, "failed to validate the flow")
	}

	if err := flow.createProcessor(processor); err != nil {
		return errors.Wrap(err, "failed to validate the flow")
	}

	return nil
}

// CreateQueue validates and creates a queue in the silo
func (flow *FlowChart) CreateQueue(queue *flow_model.Queue) error {
	if err := flow.validateQueueType(queue); err != nil {
		return errors.Wrap(err, "failed to validate the flow")
	}

	if err := flow.createQueue(queue); err != nil {
		return errors.Wrap(err, "failed to validate the flow")
	}

	return nil
}

// SaveJson exports the silo flow to a json format
func (flow *FlowChart) SaveJson() (*flow_model.Structure, error) {

	processors, err := flow.silo.ExportProcessors()
	if err != nil {
		return nil, err
	}

	//queues, err := flow.silo.ExportQueues()
	//if err != nil {
	//	return nil, err
	//}

	return &flow_model.Structure{
		//Queues:     queues,
		Processors: processors,
	}, nil
}

func (flow *FlowChart) validateProcessorType(proc *flow_model.Processor) error {
	if proc.Type != "" {
		if flow.silo.GetRegisteredProcessor(proc.Type) == nil {
			return fmt.Errorf("processor type %s for %v is not registered", proc.Type, proc.ID)
		}
	}
	return nil
}

func (flow *FlowChart) validateQueueType(queue *flow_model.Queue) error {
	if queue.Type != "" {
		if flow.silo.GetRegisteredQueue(queue.Type) == nil {
			return fmt.Errorf("queue Type %s for %v is not registered", queue.Type, queue.ID)
		}
	}

	if err := flow.createQueue(queue); err != nil {
		return errors.Wrap(err, "failed to validate the flow")
	}

	return nil
}

func (flow *FlowChart) createQueue(queue *flow_model.Queue) error {
	queueType := flow.silo.GetRegisteredQueue(queue.Type)
	if queueType == nil {
		// TODO probably alert here
		queueType = queues.NewBufferQueue()
	}

	flow.silo.CreateQueue(queueType, queue.ID)

	return nil
}

func (flow *FlowChart) createProcessor(processor *flow_model.Processor) error {

	// TODO check if it exists :o
	procType := flow.silo.GetRegisteredProcessor(processor.Type)
	flow.silo.CreateProcessor(procType, processor.ID, processor.Config)
	for _, queue := range processor.InputQueues {
		flow.silo.AddInputQueue(processor.ID, queue)
	}
	for _, queue := range processor.OutputQueues {
		flow.silo.AddOutputQueue(processor.ID, queue)
	}

	return nil
}
